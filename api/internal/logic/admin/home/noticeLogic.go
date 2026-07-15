package home

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/serv"
	"api/internal/svc"
	"api/reqs/noticeReq"
	"api/resp"
	"api/resp/noticeResp"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type NoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NoticeLogic {
	return &NoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 公告列表
func (l *NoticeLogic) ListLogic(req noticeReq.NoticeListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereStr := "is_delete = ?"
	whereArgs := []interface{}{0}
	if req.NoticeName != "" {
		whereStr += " AND notice_name LIKE ?"
		whereArgs = append(whereArgs, "%"+req.NoticeName+"%")
	}
	whereSlice := append([]interface{}{whereStr}, whereArgs...)
	count, list, err := l.svcCtx.CultrueDbStruct.SyNoticeDao.GetList(whereSlice, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询公告列表失败"
		return returnData, err
	}
	data := noticeResp.NoticeList{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    count,
		List:     []noticeResp.NoticeInfo{},
	}
	if count == 0 {
		returnData.Data = data
		return returnData, nil
	}
	// 收集添加人和发布人ID
	userIds := make([]int, 0)
	for _, n := range list {
		userIds = append(userIds, n.AddUid)
		if n.PublishUid > 0 {
			userIds = append(userIds, n.PublishUid)
		}
	}
	adminUserService := serv.NewAdminUserService(l.ctx, l.svcCtx)
	userMap, _ := adminUserService.GetAdminIdUserMap(userIds)

	for _, n := range list {
		addUser := ""
		if name, ok := userMap[n.AddUid]; ok {
			addUser = name
		}
		publishUser := ""
		if name, ok := userMap[n.PublishUid]; ok {
			publishUser = name
		}
		data.List = append(data.List, noticeResp.NoticeInfo{
			NoticeId:      n.Id,
			NoticeName:    n.NoticeName,
			PublishTime:   helper.TimeEnumFuncObject.StringTime(n.PublishTime),
			NoticeContent: n.NoticeContent,
			NoticeStatus:  n.NoticeStatus,
			AddUser:       addUser,
			PublishUser:   publishUser,
			AddTime:       helper.TimeEnumFuncObject.StringTime(n.AddTime),
		})
	}
	returnData.Data = data
	return returnData, nil
}

// 添加公告
func (l *NoticeLogic) AddLogic(req noticeReq.NoticeAddReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	publishUid := req.PublishUid
	if publishUid == 0 {
		publishUid = uid
	}
	insertMap := map[string]interface{}{
		"notice_name":    req.NoticeName,
		"notice_content": req.NoticeContent,
		"notice_status":  req.NoticeStatus,
		"add_uid":        uid,
		"publish_uid":    publishUid,
		"update_uid":     uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100001
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyNoticeDao.Insert([]map[string]interface{}{insertMap}, session); err != nil {
		session.Rollback()
		returnData.Code = 100002
		returnData.Message = "添加公告失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 修改公告
func (l *NoticeLogic) UpdateLogic(req noticeReq.NoticeUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereNotice := []interface{}{"id = ? AND is_delete = ?", req.NoticeId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyNoticeDao.GetInfo(whereNotice)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询公告信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到公告信息"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"notice_name":    req.NoticeName,
		"notice_content": req.NoticeContent,
		"notice_status":  req.NoticeStatus,
		"update_uid":     uid,
	}
	if req.PublishUid > 0 {
		updateMap["publish_uid"] = req.PublishUid
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyNoticeDao.Update(whereNotice, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改公告失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 发布/下线公告
func (l *NoticeLogic) StatusLogic(req noticeReq.NoticeStatusReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereNotice := []interface{}{"id = ? AND is_delete = ?", req.NoticeId, 0}
	has, notice, err := l.svcCtx.CultrueDbStruct.SyNoticeDao.GetInfo(whereNotice)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询公告信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到公告信息"
		return returnData, nil
	}
	// 校验操作合法性：未发布才能发布，已发布才能下线
	if req.NoticeStatus == 1 && notice.NoticeStatus == 1 {
		returnData.Code = 100003
		returnData.Message = "公告已发布，无需重复发布"
		return returnData, nil
	}
	if req.NoticeStatus == 0 && notice.NoticeStatus == 0 {
		returnData.Code = 100004
		returnData.Message = "公告已下线，无需重复下线"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"notice_status": req.NoticeStatus,
		"update_uid":    uid,
	}
	if req.NoticeStatus == 1 {
		updateMap["publish_time"] = time.Now().Format(helper.TimeEnumObject.DateTimeLayout)
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100005
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyNoticeDao.Update(whereNotice, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100006
		returnData.Message = "操作失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

func (l *NoticeLogic) DeleteLogic(req noticeReq.NoticeDeleteReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereNotice := []interface{}{"id = ? AND is_delete = ?", req.NoticeId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyNoticeDao.GetInfo(whereNotice)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询公告信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到公告信息"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"is_delete":  1,
		"update_uid": uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyNoticeDao.Update(whereNotice, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "删除公告失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}
