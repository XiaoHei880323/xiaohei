package course

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/model"
	"api/reqs/courseEvaluationReq"
	"api/resp"
	"api/resp/courseEvaluationResp"
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type CourseEvaluationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCourseEvaluationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseEvaluationLogic {
	return &CourseEvaluationLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *CourseEvaluationLogic) List(req courseEvaluationReq.CourseEvaluationListReq) (*resp.CommonReply, error) {
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
	if req.EvalType != nil {
		if *req.EvalType != 0 && *req.EvalType != 1 && *req.EvalType != 3 {
			return invalidCourseEvaluationReply("评价类型参数错误"), nil
		}
		whereSQL += " AND eval_type = ?"
		args = append(args, *req.EvalType)
	}
	if keyword := strings.TrimSpace(req.Keyword); keyword != "" {
		whereSQL += " AND content LIKE ?"
		args = append(args, "%"+keyword+"%")
	}
	where := append([]interface{}{whereSQL}, args...)
	count, evaluations, err := l.svcCtx.CultrueDbStruct.CourseEvaluationDao.GetList(where, req.Page, req.PageSize, "id desc")
	if err != nil {
		result.Code = 100001
		result.Message = "查询评价列表失败"
		return result, err
	}
	data := courseEvaluationResp.CourseEvaluationListResp{
		Page: req.Page, PageSize: req.PageSize, Count: count, List: make([]courseEvaluationResp.CourseEvaluationInfo, 0),
	}
	for _, evaluation := range evaluations {
		data.List = append(data.List, evaluationToInfo(evaluation))
	}
	result.Data = data
	return result, nil
}

func (l *CourseEvaluationLogic) Detail(req courseEvaluationReq.CourseEvaluationDetailReq) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseEvaluationReply("评价ID不能为空"), nil
	}
	has, evaluation, err := l.svcCtx.CultrueDbStruct.CourseEvaluationDao.GetInfo([]interface{}{"id = ? AND is_delete = ?", req.Id, 0})
	if err != nil {
		result.Code = 100001
		result.Message = "查询评价详情失败"
		return result, err
	}
	if !has {
		return courseEvaluationNotFoundReply(), nil
	}
	result.Data = evaluationToInfo(evaluation)
	return result, nil
}

func (l *CourseEvaluationLogic) Add(req courseEvaluationReq.CourseEvaluationAddReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if message := validateCourseEvaluation(req.CourseId, req.EvalType, req.Content); message != "" {
		return invalidCourseEvaluationReply(message), nil
	}
	rating := req.Rating
	if rating < 1 || rating > 5 {
		rating = 5
	}
	insert := map[string]interface{}{
		"course_id":   req.CourseId,
		"eval_type":   req.EvalType,
		"content":     strings.TrimSpace(req.Content),
		"rating":      rating,
		"add_uid":     operatorID,
		"update_uid":  operatorID,
	}
	if err := l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseEvaluationDao.Insert([]map[string]interface{}{insert}, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "新增评价失败"
		return result, err
	}
	result.Message = "新增成功"
	return result, nil
}

func (l *CourseEvaluationLogic) Update(req courseEvaluationReq.CourseEvaluationUpdateReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseEvaluationReply("评价ID不能为空"), nil
	}
	if message := validateCourseEvaluation(req.CourseId, req.EvalType, req.Content); message != "" {
		return invalidCourseEvaluationReply(message), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseEvaluationDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询评价失败"
		return result, err
	}
	if !has {
		return courseEvaluationNotFoundReply(), nil
	}
	rating := req.Rating
	if rating < 1 || rating > 5 {
		rating = 5
	}
	update := map[string]interface{}{
		"course_id":   req.CourseId,
		"eval_type":   req.EvalType,
		"content":     strings.TrimSpace(req.Content),
		"rating":      rating,
		"update_uid":  operatorID,
		"update_time": time.Now(),
	}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseEvaluationDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "修改评价失败"
		return result, err
	}
	result.Message = "修改成功"
	return result, nil
}

func (l *CourseEvaluationLogic) Delete(req courseEvaluationReq.CourseEvaluationDeleteReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseEvaluationReply("评价ID不能为空"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseEvaluationDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询评价失败"
		return result, err
	}
	if !has {
		return courseEvaluationNotFoundReply(), nil
	}
	update := map[string]interface{}{"is_delete": 1, "update_uid": operatorID, "update_time": time.Now()}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseEvaluationDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "删除评价失败"
		return result, err
	}
	result.Message = "删除成功"
	return result, nil
}

func (l *CourseEvaluationLogic) withTransaction(operation func(*xorm.Session) error) error {
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

func validateCourseEvaluation(courseId, evalType int, content string) string {
	if courseId <= 0 {
		return "请选择课程"
	}
	if evalType != 0 && evalType != 1 && evalType != 3 {
		return "评价类型参数错误"
	}
	if strings.TrimSpace(content) == "" {
		return "请填写评价内容"
	}
	return ""
}

func evaluationToInfo(evaluation *model.CourseEvaluation) courseEvaluationResp.CourseEvaluationInfo {
	return courseEvaluationResp.CourseEvaluationInfo{
		Id:         evaluation.Id,
		CourseId:   evaluation.CourseId,
		EvalType:   evaluation.EvalType,
		Content:    evaluation.Content,
		Rating:     evaluation.Rating,
		CreateTime: helper.TimeEnumFuncObject.StringTime(evaluation.CreateTime),
		UpdateTime: helper.TimeEnumFuncObject.StringTime(evaluation.UpdateTime),
	}
}

func invalidCourseEvaluationReply(message string) *resp.CommonReply {
	return &resp.CommonReply{Code: 400, Message: message, Data: []interface{}{}}
}

func courseEvaluationNotFoundReply() *resp.CommonReply {
	return &resp.CommonReply{Code: 404, Message: "评价不存在", Data: []interface{}{}}
}
