package course

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/model"
	"api/reqs/courseMediaReq"
	"api/resp"
	"api/resp/courseMediaResp"
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type CourseMediaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCourseMediaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseMediaLogic {
	return &CourseMediaLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *CourseMediaLogic) List(req courseMediaReq.CourseMediaListReq) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 20
	}
	if req.PageSize > 200 {
		req.PageSize = 200
	}
	whereSQL := "is_delete = ?"
	args := []interface{}{0}
	if req.CourseId > 0 {
		whereSQL += " AND course_id = ?"
		args = append(args, req.CourseId)
	}
	if req.MediaType != nil {
		if *req.MediaType != 0 && *req.MediaType != 1 && *req.MediaType != 3 {
			return invalidCourseMediaReply("媒体类型参数错误"), nil
		}
		whereSQL += " AND media_type = ?"
		args = append(args, *req.MediaType)
	}
	where := append([]interface{}{whereSQL}, args...)
	count, mediaList, err := l.svcCtx.CultrueDbStruct.CourseMediaDao.GetList(where, req.Page, req.PageSize, "id desc")
	if err != nil {
		result.Code = 100001
		result.Message = "查询媒体资源列表失败"
		return result, err
	}
	data := courseMediaResp.CourseMediaListResp{
		Page: req.Page, PageSize: req.PageSize, Count: count, List: make([]courseMediaResp.CourseMediaInfo, 0),
	}
	for _, m := range mediaList {
		data.List = append(data.List, mediaToInfo(m))
	}
	result.Data = data
	return result, nil
}

func (l *CourseMediaLogic) Detail(req courseMediaReq.CourseMediaDetailReq) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseMediaReply("资源ID不能为空"), nil
	}
	has, media, err := l.svcCtx.CultrueDbStruct.CourseMediaDao.GetInfo([]interface{}{"id = ? AND is_delete = ?", req.Id, 0})
	if err != nil {
		result.Code = 100001
		result.Message = "查询媒体资源详情失败"
		return result, err
	}
	if !has {
		return courseMediaNotFoundReply(), nil
	}
	result.Data = mediaToInfo(media)
	return result, nil
}

func (l *CourseMediaLogic) Add(req courseMediaReq.CourseMediaAddReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.CourseId <= 0 {
		return invalidCourseMediaReply("请选择课程"), nil
	}
	if req.MediaType != 0 && req.MediaType != 1 && req.MediaType != 3 {
		return invalidCourseMediaReply("媒体类型参数错误"), nil
	}
	if strings.TrimSpace(req.MediaUrl) == "" {
		return invalidCourseMediaReply("请填写资源地址"), nil
	}
	insert := map[string]interface{}{
		"course_id":  req.CourseId,
		"media_type": req.MediaType,
		"media_url":  strings.TrimSpace(req.MediaUrl),
		"add_uid":    operatorID,
		"update_uid": operatorID,
	}
	if err := l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseMediaDao.Insert([]map[string]interface{}{insert}, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "新增媒体资源失败"
		return result, err
	}
	result.Message = "新增成功"
	return result, nil
}

func (l *CourseMediaLogic) Update(req courseMediaReq.CourseMediaUpdateReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseMediaReply("资源ID不能为空"), nil
	}
	if req.CourseId <= 0 {
		return invalidCourseMediaReply("请选择课程"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseMediaDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询媒体资源失败"
		return result, err
	}
	if !has {
		return courseMediaNotFoundReply(), nil
	}
	update := map[string]interface{}{
		"course_id":   req.CourseId,
		"media_type":  req.MediaType,
		"media_url":   strings.TrimSpace(req.MediaUrl),
		"update_uid":  operatorID,
		"update_time": time.Now(),
	}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseMediaDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "修改媒体资源失败"
		return result, err
	}
	result.Message = "修改成功"
	return result, nil
}

func (l *CourseMediaLogic) Delete(req courseMediaReq.CourseMediaDeleteReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseMediaReply("资源ID不能为空"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseMediaDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询媒体资源失败"
		return result, err
	}
	if !has {
		return courseMediaNotFoundReply(), nil
	}
	update := map[string]interface{}{"is_delete": 1, "update_uid": operatorID, "update_time": time.Now()}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseMediaDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "删除媒体资源失败"
		return result, err
	}
	result.Message = "删除成功"
	return result, nil
}

func (l *CourseMediaLogic) withTransaction(operation func(*xorm.Session) error) error {
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		return err
	}
	defer func() {
		if recovered := recover(); recovered != nil {
			session.Rollback()
			panic(recovered)
		}
	}()
	if err := operation(session); err != nil {
		session.Rollback()
		return err
	}
	return session.Commit()
}

func mediaToInfo(m *model.CourseMedia) courseMediaResp.CourseMediaInfo {
	return courseMediaResp.CourseMediaInfo{
		Id:         m.Id,
		CourseId:   m.CourseId,
		MediaType:  m.MediaType,
		MediaUrl:   m.MediaUrl,
		CreateTime: helper.TimeEnumFuncObject.StringTime(m.CreateTime),
		UpdateTime: helper.TimeEnumFuncObject.StringTime(m.UpdateTime),
	}
}

func invalidCourseMediaReply(message string) *resp.CommonReply {
	return &resp.CommonReply{Code: 400, Message: message, Data: []interface{}{}}
}

func courseMediaNotFoundReply() *resp.CommonReply {
	return &resp.CommonReply{Code: 404, Message: "媒体资源不存在", Data: []interface{}{}}
}
