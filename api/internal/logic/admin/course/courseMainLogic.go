package course

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/model"
	"api/reqs/courseMainReq"
	"api/resp"
	"api/resp/courseMainResp"
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type CourseMainLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type CourseMainRow struct {
	model.CourseMain `xorm:"extends"`
	StudentName      string `xorm:"student_name"`
	TeacherName      string `xorm:"teacher_name"`
}

func NewCourseMainLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseMainLogic {
	return &CourseMainLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *CourseMainLogic) List(req courseMainReq.CourseMainListReq) (*resp.CommonReply, error) {
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
	whereSQL := "cm.is_delete = ?"
	args := []interface{}{0}

	if studentName := strings.TrimSpace(req.StudentName); studentName != "" {
		studentIds, err := l.svcCtx.CultrueDbStruct.StudentDao.GetIdsByName(studentName)
		if err != nil {
			result.Code = 100001
			result.Message = "查询学生失败"
			return result, err
		}
		if len(studentIds) == 0 {
			result.Data = courseMainResp.CourseMainListResp{
				Page: req.Page, PageSize: req.PageSize, Count: 0, List: make([]courseMainResp.CourseMainInfo, 0),
			}
			return result, nil
		}
		placeholders := make([]string, len(studentIds))
		for i, id := range studentIds {
			placeholders[i] = "?"
			args = append(args, id)
		}
		whereSQL += " AND cm.student_id IN (" + strings.Join(placeholders, ",") + ")"
	}

	if teacherName := strings.TrimSpace(req.TeacherName); teacherName != "" {
		teacherIds, err := l.svcCtx.CultrueDbStruct.TeacherDao.GetIdsByName(teacherName)
		if err != nil {
			result.Code = 100001
			result.Message = "查询老师失败"
			return result, err
		}
		if len(teacherIds) == 0 {
			result.Data = courseMainResp.CourseMainListResp{
				Page: req.Page, PageSize: req.PageSize, Count: 0, List: make([]courseMainResp.CourseMainInfo, 0),
			}
			return result, nil
		}
		placeholders := make([]string, len(teacherIds))
		for i, id := range teacherIds {
			placeholders[i] = "?"
			args = append(args, id)
		}
		whereSQL += " AND cm.teacher_id IN (" + strings.Join(placeholders, ",") + ")"
	}

	if req.Status != nil {
		if *req.Status < 0 || *req.Status > 3 {
			return invalidCourseMainReply("课程状态参数错误"), nil
		}
		whereSQL += " AND cm.status = ?"
		args = append(args, *req.Status)
	}
	if req.StudentId > 0 {
		whereSQL += " AND cm.student_id = ?"
		args = append(args, req.StudentId)
	}
	if req.TeacherId > 0 {
		whereSQL += " AND cm.teacher_id = ?"
		args = append(args, req.TeacherId)
	}
	if dateBegin := strings.TrimSpace(req.DateBegin); dateBegin != "" {
		whereSQL += " AND cm.course_start_time >= ?"
		args = append(args, dateBegin)
	}
	if dateEnd := strings.TrimSpace(req.DateEnd); dateEnd != "" {
		whereSQL += " AND cm.course_end_time <= ?"
		args = append(args, dateEnd+" 23:59:59")
	}

	countArgs := make([]interface{}, len(args))
	copy(countArgs, args)
	offset := req.PageSize * (req.Page - 1)
	queryArgs := append(args, req.PageSize, offset)

	sqlStr := "SELECT cm.*, cs.name AS student_name, ct.name AS teacher_name FROM course_main cm " +
		"LEFT JOIN course_student cs ON cm.student_id = cs.id " +
		"LEFT JOIN course_teacher ct ON cm.teacher_id = ct.id " +
		"WHERE " + whereSQL + " ORDER BY cm.id DESC LIMIT ? OFFSET ?"
	countSQL := "SELECT COUNT(*) FROM course_main cm WHERE " + whereSQL

	total, err := l.svcCtx.Engine.SQL(countSQL, countArgs...).Count()
	if err != nil {
		result.Code = 100001
		result.Message = "查询课程总数失败"
		return result, err
	}

	rows := make([]*CourseMainRow, 0)
	if err := l.svcCtx.Engine.SQL(sqlStr, queryArgs...).Find(&rows); err != nil {
		result.Code = 100001
		result.Message = "查询课程列表失败"
		return result, err
	}

	data := courseMainResp.CourseMainListResp{
		Page: req.Page, PageSize: req.PageSize, Count: int(total), List: make([]courseMainResp.CourseMainInfo, 0),
	}
	for _, row := range rows {
		data.List = append(data.List, courseMainRowToInfo(row))
	}
	result.Data = data
	return result, nil
}

func (l *CourseMainLogic) Detail(req courseMainReq.CourseMainDetailReq) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseMainReply("课程ID不能为空"), nil
	}
	var row CourseMainRow
	has, err := l.svcCtx.Engine.SQL(
		"SELECT cm.*, cs.name AS student_name, ct.name AS teacher_name FROM course_main cm "+
			"LEFT JOIN course_student cs ON cm.student_id = cs.id "+
			"LEFT JOIN course_teacher ct ON cm.teacher_id = ct.id "+
			"WHERE cm.id = ? AND cm.is_delete = ?", req.Id, 0).Get(&row)
	if err != nil {
		result.Code = 100001
		result.Message = "查询课程详情失败"
		return result, err
	}
	if !has {
		return courseMainNotFoundReply(), nil
	}
	result.Data = courseMainRowToInfo(&row)
	return result, nil
}

func (l *CourseMainLogic) Add(req courseMainReq.CourseMainAddReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if message := validateCourseMain(req.StudentId, req.TeacherId, req.CourseStartTime, req.CourseEndTime, req.Status); message != "" {
		return invalidCourseMainReply(message), nil
	}
	status := 0
	if req.Status != nil {
		status = *req.Status
	}

	insert := map[string]interface{}{
		"student_id":        req.StudentId,
		"teacher_id":        req.TeacherId,
		"course_start_time": req.CourseStartTime,
		"course_end_time":   req.CourseEndTime,
		"meeting_link":      strings.TrimSpace(req.MeetingLink),
		"subject":           strings.TrimSpace(req.Subject),
		"status":            status,
		"add_uid":           operatorID,
		"update_uid":        operatorID,
	}
	if err := l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseMainDao.Insert([]map[string]interface{}{insert}, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "新增课程失败"
		return result, err
	}
	result.Message = "新增成功"
	return result, nil
}

func (l *CourseMainLogic) Update(req courseMainReq.CourseMainUpdateReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseMainReply("课程ID不能为空"), nil
	}
	status := req.Status
	if message := validateCourseMain(req.StudentId, req.TeacherId, req.CourseStartTime, req.CourseEndTime, &status); message != "" {
		return invalidCourseMainReply(message), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseMainDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询课程失败"
		return result, err
	}
	if !has {
		return courseMainNotFoundReply(), nil
	}
	update := map[string]interface{}{
		"student_id":        req.StudentId,
		"teacher_id":        req.TeacherId,
		"course_start_time": req.CourseStartTime,
		"course_end_time":   req.CourseEndTime,
		"meeting_link":      strings.TrimSpace(req.MeetingLink),
		"subject":           strings.TrimSpace(req.Subject),
		"status":            req.Status,
		"update_uid":        operatorID,
		"update_time":       time.Now(),
	}

	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseMainDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "修改课程失败"
		return result, err
	}
	result.Message = "修改成功"
	return result, nil
}

func (l *CourseMainLogic) Delete(req courseMainReq.CourseMainDeleteReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidCourseMainReply("课程ID不能为空"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.CourseMainDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询课程失败"
		return result, err
	}
	if !has {
		return courseMainNotFoundReply(), nil
	}
	update := map[string]interface{}{"is_delete": 1, "update_uid": operatorID, "update_time": time.Now()}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.CourseMainDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "删除课程失败"
		return result, err
	}
	result.Message = "删除成功"
	return result, nil
}

func (l *CourseMainLogic) withTransaction(operation func(*xorm.Session) error) error {
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

func validateCourseMain(studentId, teacherId int, courseStartTime, courseEndTime string, status *int) string {
	if studentId <= 0 {
		return "请选择学生"
	}
	if teacherId <= 0 {
		return "请选择老师"
	}
	if courseStartTime == "" {
		return "请填写上课开始时间"
	}
	if courseEndTime == "" {
		return "请填写上课结束时间"
	}
	if status != nil && (*status < 0 || *status > 3) {
		return "课程状态参数错误"
	}
	return ""
}

func courseMainRowToInfo(row *CourseMainRow) courseMainResp.CourseMainInfo {
	return courseMainResp.CourseMainInfo{
		Id:              row.Id,
		StudentId:       row.StudentId,
		StudentName:     row.StudentName,
		TeacherId:       row.TeacherId,
		TeacherName:     row.TeacherName,
		CourseStartTime: helper.TimeEnumFuncObject.StringTime(row.CourseStartTime),
		CourseEndTime:   helper.TimeEnumFuncObject.StringTime(row.CourseEndTime),
		MeetingLink:     row.MeetingLink,
		Subject:         row.Subject,
		Status:          row.Status,
		StudentEntered:  row.StudentEntered,
		TeacherEntered:  row.TeacherEntered,
		CreateTime:      helper.TimeEnumFuncObject.StringTime(row.CreateTime),
		UpdateTime:      helper.TimeEnumFuncObject.StringTime(row.UpdateTime),
	}
}

func invalidCourseMainReply(message string) *resp.CommonReply {
	return &resp.CommonReply{Code: 400, Message: message, Data: []interface{}{}}
}

func courseMainNotFoundReply() *resp.CommonReply {
	return &resp.CommonReply{Code: 404, Message: "课程不存在", Data: []interface{}{}}
}
