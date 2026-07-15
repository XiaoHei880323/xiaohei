package course

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/model"
	"api/reqs/courseErrorCollectionReq"
	"api/resp"
	"api/resp/courseErrorCollectionResp"
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type CourseErrorCollectionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCourseErrorCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseErrorCollectionLogic {
	return &CourseErrorCollectionLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *CourseErrorCollectionLogic) List(req courseErrorCollectionReq.CourseErrorCollectionListReq) (*resp.CommonReply, error) {
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
	if keyword := strings.TrimSpace(req.Keyword); keyword != "" {
		whereSQL += " AND (question LIKE ? OR correct_answer LIKE ? OR knowledge_point LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like, like)
	}
	where := append([]interface{}{whereSQL}, args...)
	count, list, err := l.svcCtx.CultrueDbStruct.CourseErrorCollectionDao.GetList(where, req.Page, req.PageSize, "id desc")
	if err != nil {
		result.Code = 100001
		result.Message = "查询错题集列表失败"
		return result, err
	}
	data := courseErrorCollectionResp.CourseErrorCollectionListResp{
		Page: req.Page, PageSize: req.PageSize, Count: count, List: make([]courseErrorCollectionResp.CourseErrorCollectionInfo, 0),
	}
	for _, item := range list {
		data.List = append(data.List, errorCollectionToInfo(item))
	}
	result.Data = data
	return result, nil
}

func (l *CourseErrorCollectionLogic) Detail(req courseErrorCollectionReq.CourseErrorCollectionDetailReq) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseErrorCollectionReply("错题ID不能为空"), nil
	}
	has, item, err := l.svcCtx.CultrueDbStruct.CourseErrorCollectionDao.GetInfo([]interface{}{"id = ? AND is_delete = ?", req.Id, 0})
	if err != nil {
		result.Code = 100001
		result.Message = "查询错题详情失败"
		return result, err
	}
	if !has {
		return courseErrorCollectionNotFoundReply(), nil
	}
	result.Data = errorCollectionToInfo(item)
	return result, nil
}

func (l *CourseErrorCollectionLogic) Add(req courseErrorCollectionReq.CourseErrorCollectionAddReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.CourseId <= 0 {
		return invalidCourseErrorCollectionReply("请选择课程"), nil
	}
	if strings.TrimSpace(req.Question) == "" {
		return invalidCourseErrorCollectionReply("请填写错题内容"), nil
	}
	if strings.TrimSpace(req.CorrectAnswer) == "" {
		return invalidCourseErrorCollectionReply("请填写正确答案"), nil
	}
	insert := map[string]interface{}{
		"course_id":       req.CourseId,
		"question":        strings.TrimSpace(req.Question),
		"correct_answer":  strings.TrimSpace(req.CorrectAnswer),
		"student_answer":  strings.TrimSpace(req.StudentAnswer),
		"analysis":        strings.TrimSpace(req.Analysis),
		"knowledge_point": strings.TrimSpace(req.KnowledgePoint),
		"add_uid":         operatorID,
		"update_uid":      operatorID,
	}
	if err := l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseErrorCollectionDao.Insert([]map[string]interface{}{insert}, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "新增错题失败"
		return result, err
	}
	result.Message = "新增成功"
	return result, nil
}

func (l *CourseErrorCollectionLogic) Update(req courseErrorCollectionReq.CourseErrorCollectionUpdateReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseErrorCollectionReply("错题ID不能为空"), nil
	}
	if req.CourseId <= 0 {
		return invalidCourseErrorCollectionReply("请选择课程"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseErrorCollectionDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询错题失败"
		return result, err
	}
	if !has {
		return courseErrorCollectionNotFoundReply(), nil
	}
	update := map[string]interface{}{
		"course_id":       req.CourseId,
		"question":        strings.TrimSpace(req.Question),
		"correct_answer":  strings.TrimSpace(req.CorrectAnswer),
		"student_answer":  strings.TrimSpace(req.StudentAnswer),
		"analysis":        strings.TrimSpace(req.Analysis),
		"knowledge_point": strings.TrimSpace(req.KnowledgePoint),
		"update_uid":      operatorID,
		"update_time":     time.Now(),
	}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseErrorCollectionDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "修改错题失败"
		return result, err
	}
	result.Message = "修改成功"
	return result, nil
}

func (l *CourseErrorCollectionLogic) Delete(req courseErrorCollectionReq.CourseErrorCollectionDeleteReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseErrorCollectionReply("错题ID不能为空"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseErrorCollectionDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询错题失败"
		return result, err
	}
	if !has {
		return courseErrorCollectionNotFoundReply(), nil
	}
	update := map[string]interface{}{"is_delete": 1, "update_uid": operatorID, "update_time": time.Now()}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseErrorCollectionDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "删除错题失败"
		return result, err
	}
	result.Message = "删除成功"
	return result, nil
}

func (l *CourseErrorCollectionLogic) withTransaction(operation func(*xorm.Session) error) error {
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

func errorCollectionToInfo(item *model.CourseErrorCollection) courseErrorCollectionResp.CourseErrorCollectionInfo {
	return courseErrorCollectionResp.CourseErrorCollectionInfo{
		Id:             item.Id,
		CourseId:       item.CourseId,
		Question:       item.Question,
		CorrectAnswer:  item.CorrectAnswer,
		StudentAnswer:  item.StudentAnswer,
		Analysis:       item.Analysis,
		KnowledgePoint: item.KnowledgePoint,
		CreateTime:     helper.TimeEnumFuncObject.StringTime(item.CreateTime),
		UpdateTime:     helper.TimeEnumFuncObject.StringTime(item.UpdateTime),
	}
}

func invalidCourseErrorCollectionReply(message string) *resp.CommonReply {
	return &resp.CommonReply{Code: 400, Message: message, Data: []interface{}{}}
}

func courseErrorCollectionNotFoundReply() *resp.CommonReply {
	return &resp.CommonReply{Code: 404, Message: "错题不存在", Data: []interface{}{}}
}
