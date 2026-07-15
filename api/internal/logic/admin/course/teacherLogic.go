package course

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/model"
	"api/reqs/teacherReq"
	"api/resp"
	"api/resp/teacherResp"
	"context"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"xorm.io/xorm"
)

type TeacherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TeacherLogic {
	return &TeacherLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *TeacherLogic) List(req teacherReq.TeacherListReq) (*resp.CommonReply, error) {
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
	if keyword := strings.TrimSpace(req.Keyword); keyword != "" {
		whereSQL += " AND (teacher_no LIKE ? OR name LIKE ? OR phone LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like, like)
	}
	if department := strings.TrimSpace(req.Department); department != "" {
		whereSQL += " AND department LIKE ?"
		args = append(args, "%"+department+"%")
	}
	if req.Status != nil {
		if *req.Status != 0 && *req.Status != 1 {
			return invalidTeacherReply("教师状态参数错误"), nil
		}
		whereSQL += " AND status = ?"
		args = append(args, *req.Status)
	}
	where := append([]interface{}{whereSQL}, args...)
	count, teachers, err := l.svcCtx.CultrueDbStruct.TeacherDao.GetList(where, req.Page, req.PageSize, "id desc")
	if err != nil {
		result.Code = 100001
		result.Message = "查询教师列表失败"
		return result, err
	}
	data := teacherResp.TeacherListResp{
		Page: req.Page, PageSize: req.PageSize, Count: count, List: make([]teacherResp.TeacherInfo, 0),
	}
	for _, teacher := range teachers {
		data.List = append(data.List, teacherToInfo(teacher))
	}
	result.Data = data
	return result, nil
}

func (l *TeacherLogic) Detail(req teacherReq.TeacherDetailReq) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidTeacherReply("教师ID不能为空"), nil
	}
	has, teacher, err := l.svcCtx.CultrueDbStruct.TeacherDao.GetInfo([]interface{}{"id = ? AND is_delete = ?", req.Id, 0})
	if err != nil {
		result.Code = 100001
		result.Message = "查询教师详情失败"
		return result, err
	}
	if !has {
		return teacherNotFoundReply(), nil
	}
	result.Data = teacherToInfo(teacher)
	return result, nil
}

func (l *TeacherLogic) Add(req teacherReq.TeacherAddReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	teacherNo := strings.TrimSpace(req.TeacherNo)
	name := strings.TrimSpace(req.Name)
	phone := strings.TrimSpace(req.Phone)
	if message := validateTeacher(teacherNo, name, req.Gender, req.Status, req.Password, true); message != "" {
		return invalidTeacherReply(message), nil
	}
	status := 1
	if req.Status != nil {
		status = *req.Status
	}
	if reply, err := l.ensureTeacherUnique(0, teacherNo, phone); reply != nil || err != nil {
		return reply, err
	}
	insert := map[string]interface{}{
		"teacher_no": teacherNo,
		"name":       name,
		"gender":     req.Gender,
		"phone":      phone,
		"email":      strings.TrimSpace(req.Email),
		"title":      strings.TrimSpace(req.Title),
		"department": strings.TrimSpace(req.Department),
		"password":   helper.Md5HelperObject.Sha256ToString(req.Password),
		"status":     status,
		"add_uid":    operatorID,
		"update_uid": operatorID,
	}
	if err := l.withTeacherTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.TeacherDao.Insert([]map[string]interface{}{insert}, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "新增教师失败"
		return result, err
	}
	result.Message = "新增成功"
	return result, nil
}

func (l *TeacherLogic) Update(req teacherReq.TeacherUpdateReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidTeacherReply("教师ID不能为空"), nil
	}
	teacherNo := strings.TrimSpace(req.TeacherNo)
	name := strings.TrimSpace(req.Name)
	phone := strings.TrimSpace(req.Phone)
	status := req.Status
	if message := validateTeacher(teacherNo, name, req.Gender, &status, req.Password, false); message != "" {
		return invalidTeacherReply(message), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.TeacherDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询教师失败"
		return result, err
	}
	if !has {
		return teacherNotFoundReply(), nil
	}
	if reply, err := l.ensureTeacherUnique(req.Id, teacherNo, phone); reply != nil || err != nil {
		return reply, err
	}
	update := map[string]interface{}{
		"teacher_no":  teacherNo,
		"name":        name,
		"gender":      req.Gender,
		"phone":       phone,
		"email":       strings.TrimSpace(req.Email),
		"title":       strings.TrimSpace(req.Title),
		"department":  strings.TrimSpace(req.Department),
		"status":      req.Status,
		"update_uid":  operatorID,
		"update_time": time.Now(),
	}
	if req.Password != "" {
		update["password"] = helper.Md5HelperObject.Sha256ToString(req.Password)
	}
	if err = l.withTeacherTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.TeacherDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "修改教师失败"
		return result, err
	}
	result.Message = "修改成功"
	return result, nil
}

func (l *TeacherLogic) Delete(req teacherReq.TeacherDeleteReq, operatorID int) (*resp.CommonReply, error) {
	result := reponse.ReturnStruct()
	if req.Id <= 0 {
		return invalidTeacherReply("教师ID不能为空"), nil
	}
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.TeacherDao.GetInfo(where)
	if err != nil {
		result.Code = 100001
		result.Message = "查询教师失败"
		return result, err
	}
	if !has {
		return teacherNotFoundReply(), nil
	}
	update := map[string]interface{}{"is_delete": 1, "update_uid": operatorID, "update_time": time.Now()}
	if err = l.withTeacherTransaction(func(session *xorm.Session) error {
		return l.svcCtx.CultrueDbStruct.TeacherDao.Update(where, update, session)
	}); err != nil {
		result.Code = 100002
		result.Message = "删除教师失败"
		return result, err
	}
	result.Message = "删除成功"
	return result, nil
}

func (l *TeacherLogic) ensureTeacherUnique(id int, teacherNo, phone string) (*resp.CommonReply, error) {
	teacherNoWhere := []interface{}{"teacher_no = ? AND id != ?", teacherNo, id}
	has, _, err := l.svcCtx.CultrueDbStruct.TeacherDao.GetInfo(teacherNoWhere)
	if err != nil {
		reply := reponse.ReturnStruct()
		reply.Code = 100003
		reply.Message = "校验教师工号失败"
		return reply, err
	}
	if has {
		return invalidTeacherReply("教师工号已存在"), nil
	}
	if phone == "" {
		return nil, nil
	}
	phoneWhere := []interface{}{"phone = ? AND is_delete = ? AND id != ?", phone, 0, id}
	has, _, err = l.svcCtx.CultrueDbStruct.TeacherDao.GetInfo(phoneWhere)
	if err != nil {
		reply := reponse.ReturnStruct()
		reply.Code = 100004
		reply.Message = "校验手机号失败"
		return reply, err
	}
	if has {
		return invalidTeacherReply("手机号已存在"), nil
	}
	return nil, nil
}

func (l *TeacherLogic) withTeacherTransaction(operation func(*xorm.Session) error) error {
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

func validateTeacher(teacherNo, name string, gender int, status *int, password string, passwordRequired bool) string {
	if teacherNo == "" {
		return "请输入教师工号"
	}
	if name == "" {
		return "请输入教师姓名"
	}
	if gender < 0 || gender > 2 {
		return "性别参数错误"
	}
	if status != nil && *status != 0 && *status != 1 {
		return "教师状态参数错误"
	}
	if passwordRequired && len(password) < 6 {
		return "密码不能少于6位"
	}
	if !passwordRequired && password != "" && len(password) < 6 {
		return "密码不能少于6位"
	}
	return ""
}

func teacherToInfo(teacher *model.Teacher) teacherResp.TeacherInfo {
	return teacherResp.TeacherInfo{
		Id: teacher.Id, TeacherNo: teacher.TeacherNo, Name: teacher.Name,
		Gender: teacher.Gender, Phone: teacher.Phone, Email: teacher.Email,
		Title: teacher.Title, Department: teacher.Department, Status: teacher.Status,
		CreateTime: helper.TimeEnumFuncObject.StringTime(teacher.CreateTime),
		UpdateTime: helper.TimeEnumFuncObject.StringTime(teacher.UpdateTime),
	}
}

func invalidTeacherReply(message string) *resp.CommonReply {
	return &resp.CommonReply{Code: 400, Message: message, Data: []interface{}{}}
}

func teacherNotFoundReply() *resp.CommonReply {
	return &resp.CommonReply{Code: 404, Message: "教师不存在", Data: []interface{}{}}
}
