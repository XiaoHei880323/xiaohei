package course

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/model"
	"api/reqs/studentReq"
	"api/resp"
	"api/resp/studentResp"
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type StudentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudentLogic {
	return &StudentLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *StudentLogic) List(req studentReq.StudentListReq) (*resp.CommonReply, error) {
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
	keyword := strings.TrimSpace(req.Keyword)
	if keyword != "" {
		whereSQL += " AND (student_no LIKE ? OR name LIKE ? OR phone LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like, like)
	}
	if className := strings.TrimSpace(req.ClassName); className != "" {
		whereSQL += " AND class_name LIKE ?"
		args = append(args, "%"+className+"%")
	}
	if req.Status != nil {
		if *req.Status != 0 && *req.Status != 1 {
			return invalidStudentReply("学生状态参数错误"), nil
		}
		whereSQL += " AND status = ?"
		args = append(args, *req.Status)
	}
	where := append([]interface{}{whereSQL}, args...)
	count, students, err := l.svcCtx.CultrueDbStruct.StudentDao.GetList(where, req.Page, req.PageSize, "id desc")
	if err != nil {
		result.Code = 100001
		result.Message = "查询学生列表失败"
		return result, err
	}
	data := studentResp.StudentListResp{
		Page: req.Page, PageSize: req.PageSize, Count: count, List: make([]studentResp.StudentInfo, 0),
	}
	for _, student := range students {
		data.List = append(data.List, studentToInfo(student))
	}
	result.Data = data
	return result, nil
}

func (l *StudentLogic) Detail(req studentReq.StudentDetailReq) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidStudentReply("学生ID不能为空"), nil
	}
	has, student, err := l.svcCtx.CultrueDbStruct.StudentDao.GetInfo([]interface{}{"id = ? AND is_delete = ?", req.Id, 0})
	if err != nil {
		result.Code = 100001
		result.Message = "查询学生详情失败"
		return result, err
	}
	if !has {
		return studentNotFoundReply(), nil
	}
	result.Data = studentToInfo(student)
	return result, nil
}

func (l *StudentLogic) Add(req studentReq.StudentAddReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	studentNo := strings.TrimSpace(req.StudentNo)
	name := strings.TrimSpace(req.Name)
	phone := strings.TrimSpace(req.Phone)
	if message := validateStudent(studentNo, name, req.Gender, req.Status, req.Password, true); message != "" {
		return invalidStudentReply(message), nil
	}
	status := 1
	if req.Status != nil {
		status = *req.Status
	}
	if reply, err := l.ensureUnique(0, studentNo, phone); reply != nil || err != nil {
		return reply, err
	}
	insert := map[string]interface{}{
		"student_no": studentNo,
		"name":       name,
		"gender":     req.Gender,
		"phone":      phone,
		"email":      strings.TrimSpace(req.Email),
		"class_name": strings.TrimSpace(req.ClassName),
		"password":   helper.Md5HelperObject.Sha256ToString(req.Password),
		"status":     status,
		"add_uid":    operatorID,
		"update_uid": operatorID,
	}
	if err := l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.StudentDao.Insert([]map[string]interface{}{insert}, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "新增学生失败"
		return result, err
	}
	result.Message = "新增成功"
	return result, nil
}

func (l *StudentLogic) Update(req studentReq.StudentUpdateReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidStudentReply("学生ID不能为空"), nil
	}
	studentNo := strings.TrimSpace(req.StudentNo)
	name := strings.TrimSpace(req.Name)
	phone := strings.TrimSpace(req.Phone)
	status := req.Status
	if message := validateStudent(studentNo, name, req.Gender, &status, req.Password, false); message != "" {
		return invalidStudentReply(message), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.StudentDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询学生失败"
		return result, err
	}
	if !has {
		return studentNotFoundReply(), nil
	}
	if reply, err := l.ensureUnique(req.Id, studentNo, phone); reply != nil || err != nil {
		return reply, err
	}
	update := map[string]interface{}{
		"student_no":  studentNo,
		"name":        name,
		"gender":      req.Gender,
		"phone":       phone,
		"email":       strings.TrimSpace(req.Email),
		"class_name":  strings.TrimSpace(req.ClassName),
		"status":      req.Status,
		"update_uid":  operatorID,
		"update_time": time.Now(),
	}
	if req.Password != "" {
		update["password"] = helper.Md5HelperObject.Sha256ToString(req.Password)
	}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.StudentDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "修改学生失败"
		return result, err
	}
	result.Message = "修改成功"
	return result, nil
}

func (l *StudentLogic) Delete(req studentReq.StudentDeleteReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidStudentReply("学生ID不能为空"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.StudentDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询学生失败"
		return result, err
	}
	if !has {
		return studentNotFoundReply(), nil
	}
	update := map[string]interface{}{"is_delete": 1, "update_uid": operatorID, "update_time": time.Now()}
	if err = l.withTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.StudentDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "删除学生失败"
		return result, err
	}
	result.Message = "删除成功"
	return result, nil
}

func (l *StudentLogic) ensureUnique(id int, studentNo, phone string) (*resp.CommonReply, error) {
	studentNoWhere := []interface{}{"student_no = ? AND is_delete = ? AND id != ?", studentNo, 0, id}
	has, _, err := l.svcCtx.CultrueDbStruct.StudentDao.GetInfo(studentNoWhere)
	if err != nil {
		reply := reponse.ReturnStruct()
		reply.Code = 100003
		reply.Message = "校验学号失败"
		return reply, err
	}
	if has {
		return invalidStudentReply("学号已存在"), nil
	}
	if phone == "" {
		return nil, nil
	}
	phoneWhere := []interface{}{"phone = ? AND is_delete = ? AND id != ?", phone, 0, id}
	has, _, err = l.svcCtx.CultrueDbStruct.StudentDao.GetInfo(phoneWhere)
	if err != nil {
		reply := reponse.ReturnStruct()
		reply.Code = 100004
		reply.Message = "校验手机号失败"
		return reply, err
	}
	if has {
		return invalidStudentReply("手机号已存在"), nil
	}
	return nil, nil
}

func (l *StudentLogic) withTransaction(operation func(*xorm.Session) error) error {
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

func validateStudent(studentNo, name string, gender int, status *int, password string, passwordRequired bool) string {
	if studentNo == "" {
		return "请输入学号"
	}
	if name == "" {
		return "请输入学生姓名"
	}
	if gender < 0 || gender > 2 {
		return "性别参数错误"
	}
	if status != nil && *status != 0 && *status != 1 {
		return "学生状态参数错误"
	}
	if passwordRequired && len(password) < 6 {
		return "密码不能少于6位"
	}
	if !passwordRequired && password != "" && len(password) < 6 {
		return "密码不能少于6位"
	}
	return ""
}

func studentToInfo(student *model.Student) studentResp.StudentInfo {
	return studentResp.StudentInfo{
		Id: student.Id, StudentNo: student.StudentNo, Name: student.Name,
		Gender: student.Gender, Phone: student.Phone, Email: student.Email,
		ClassName: student.ClassName, Status: student.Status,
		CreateTime: helper.TimeEnumFuncObject.StringTime(student.CreateTime),
		UpdateTime: helper.TimeEnumFuncObject.StringTime(student.UpdateTime),
	}
}

func invalidStudentReply(message string) *resp.CommonReply {
	return &resp.CommonReply{Code: 400, Message: message, Data: []interface{}{}}
}

func studentNotFoundReply() *resp.CommonReply {
	return &resp.CommonReply{Code: 404, Message: "学生不存在", Data: []interface{}{}}
}
