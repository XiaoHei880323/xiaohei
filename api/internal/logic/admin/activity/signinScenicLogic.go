package activity

import (
	reponse "api/comment/response"
	helper "api/comment/help"
	"api/internal/svc"
	"api/reqs/signinScenicReq"
	"api/resp"
	"api/resp/signinScenicResp"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SigninScenicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSigninScenicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigninScenicLogic {
	return &SigninScenicLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

// ListLogic 获取某活动的景点列表，并计算积分合计
func (l *SigninScenicLogic) ListLogic(req signinScenicReq.SigninScenicListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	if req.ActivityId == 0 {
		returnData.Code = 100001
		returnData.Message = "活动ID不能为空"
		return returnData, nil
	}
	whereSlice := []interface{}{"activity_id = ? AND is_delete = ?", req.ActivityId, 0}
	count, list, err := l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.GetList(whereSlice, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100002
		returnData.Message = "查询关联景点失败"
		return returnData, err
	}
	// 计算启用景点的积分合计
	totalPoints, _ := l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.SumPointsByActivity(req.ActivityId)

	data := signinScenicResp.GetSigninScenicList{
		Page:        req.Page,
		PageSize:    req.PageSize,
		Count:       count,
		TotalPoints: totalPoints,
		List:        []signinScenicResp.SigninScenicInfo{},
	}
	// 批量查景点名
	for _, item := range list {
		spotName := ""
		where := []interface{}{"id = ? AND is_delete = ?", item.ScenicId, 0}
		has, spot, serr := l.svcCtx.CultrueDbStruct.SyScenicSpotDao.GetInfo(where)
		if serr == nil && has {
			spotName = spot.SpotName
		}
		data.List = append(data.List, signinScenicResp.SigninScenicInfo{
			Id:         item.Id,
			ActivityId: item.ActivityId,
			ScenicId:   item.ScenicId,
			SpotName:   spotName,
			SignPoints: item.SignPoints,
			QrCodeUrl:  item.QrCodeUrl,
			Status:     item.Status,
			AddTime:    helper.TimeEnumFuncObject.StringTime(item.AddTime),
		})
	}
	returnData.Data = data
	return returnData, nil
}

// AddLogic 新增活动-景点关联
func (l *SigninScenicLogic) AddLogic(req signinScenicReq.SigninScenicAddReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	// 防重复
	where := []interface{}{"activity_id = ? AND scenic_id = ? AND is_delete = ?", req.ActivityId, req.ScenicId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询失败"
		return returnData, err
	}
	if has {
		returnData.Code = 100002
		returnData.Message = "该景点已添加到此活动，请勿重复"
		return returnData, nil
	}
	insertMap := map[string]interface{}{
		"activity_id": req.ActivityId,
		"scenic_id":   req.ScenicId,
		"sign_points": req.SignPoints,
		"qr_code_url": req.QrCodeUrl,
		"status":      req.Status,
		"add_uid":     uid,
		"update_uid":  uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err = l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.Insert([]map[string]interface{}{insertMap}, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "添加景点关联失败"
		return returnData, err
	}
	// 同步更新活动积分
	l.syncActivityPoints(req.ActivityId)
	session.Commit()
	return returnData, nil
}

// UpdateLogic 修改积分/二维码/状态
func (l *SigninScenicLogic) UpdateLogic(req signinScenicReq.SigninScenicUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, item, err := l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询记录失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "记录不存在"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"sign_points": req.SignPoints,
		"qr_code_url": req.QrCodeUrl,
		"status":      req.Status,
		"update_uid":  uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err = l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.Update(where, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改失败"
		return returnData, err
	}
	l.syncActivityPoints(item.ActivityId)
	session.Commit()
	return returnData, nil
}

// DeleteLogic 删除关联（软删除）
func (l *SigninScenicLogic) DeleteLogic(req signinScenicReq.SigninScenicDeleteReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, item, err := l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询记录失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "记录不存在"
		return returnData, nil
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err = l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.Update(
		where, map[string]interface{}{"is_delete": 1, "update_uid": uid}, session,
	); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "删除失败"
		return returnData, err
	}
	l.syncActivityPoints(item.ActivityId)
	session.Commit()
	return returnData, nil
}

// syncActivityPoints 重新计算并写回活动积分
func (l *SigninScenicLogic) syncActivityPoints(activityId int) {
	total, err := l.svcCtx.CultrueDbStruct.SySigninActivityScenicDao.SumPointsByActivity(activityId)
	if err != nil {
		return
	}
	updateMap := map[string]interface{}{"activity_points": total}
	where := []interface{}{"id = ?", activityId}
	_ = l.svcCtx.CultrueDbStruct.SyActivityDao.Update(where, updateMap, nil)
}
