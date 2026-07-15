package user

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/reqs/admin"
	"api/resp"
	"api/resp/adminResp"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取用户的数据
func (l *UserLogic) GetUserListLogic(req admin.GetUserListReq) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	sql := "is_delete = ?"
	whereSlice := []interface{}{0}
	if req.Phone != "" {
		sql += " and phone = ?"
		whereSlice = append(whereSlice, req.Phone)
	}
	if req.UserNick != "" {
		sql += " and rel_name like ?"
		whereSlice = append(whereSlice, "%"+req.UserNick+"%")
	}
	where := []interface{}{}
	where = append(where, sql)
	where = append(where, whereSlice...)
	userCount, userList, userErr := l.svcCtx.CultrueDbStruct.SyUserDao.GetUserListAndPage(where, req.Page, req.PageSize, "id desc")
	if userErr != nil {
		returnData.Code = 10001
		returnData.Message = "查询用户信息错误"
		return returnData, userErr
	}
	data := adminResp.GetUserListAndPageResp{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    userCount,
	}
	if userCount == 0 {
		returnData.Data = data
		return returnData, nil
	}
	for _, user := range userList {
		info := adminResp.GetUserInfoResp{
			UserId:       user.Id,
			UserName:     user.UserName,
			UserNickName: user.RelName,
			Phone:        user.Phone,
			Points:       user.Points,
			Status:       user.Status,
			AddTime:      helper.TimeEnumFuncObject.StringTime(user.CreateTime),
		}
		data.UserList = append(data.UserList, info)
	}
	returnData.Data = data
	return returnData, nil
}

// 重置密码
func (l *UserLogic) UpdatePwdUserLogic(req admin.UpdateUserPwd, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	newPwd := helper.Md5HelperObject.Sha256ToString("123456")
	updateUserWhere := []interface{}{"id = ?", req.UserId}
	updateData := make(map[string]interface{})
	updateData["pwd"] = newPwd
	updateData["update_aid"] = userId
	//开启事务
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if errSession := session.Begin(); nil != errSession {
		returnData.Code = 100001
		returnData.Message = "开启事务失败"
		return returnData, errSession
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	updateErr := l.svcCtx.CultrueDbStruct.SyUserDao.Update(updateUserWhere, updateData, session)
	if updateErr != nil {
		session.Rollback()
		returnData.Code = 100002
		returnData.Message = "修改用户信息不正确"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}

// 修改用户积分数据信息
func (l *UserLogic) UpdateUserPointsLogic(req admin.UpdateUserPointsReq, uid int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()

	//查询用户
	findUserWhere := []interface{}{"is_delete = ? and status =? and id = ? ", 0, 0, req.UserId}
	userHas, userInfo, userErr := l.svcCtx.CultrueDbStruct.SyUserDao.GetInfo(findUserWhere)
	if userErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询用户数据信息错误"
		return returnData, userErr
	}
	if !userHas {
		returnData.Code = 100002
		returnData.Message = "未查到用户数据信息"
		return returnData, userErr
	}

	//查询用户的所有的积分综合
	findUserPoints := []interface{}{"user_id = ? and is_delete = ?", userInfo.Id, 0}
	userPoint, userPointErr := l.svcCtx.CultrueDbStruct.SyUserPointsDao.GetUserPointsSum(findUserPoints)
	if userPointErr != nil {
		returnData.Code = 100003
		returnData.Message = "查询用户积分错误"
		return returnData, userErr
	}
	if req.Source == 1 {
		req.Source = 3
	} else {
		req.Source = 4
		req.Points = -req.Points
	}
	if req.Source == 4 && userPoint == 0 {
		returnData.Code = 100007
		returnData.Message = "用户当前积分为0，不可以再扣减"
		return returnData, nil
	}
	whereAdmin := []interface{}{"id = ?", uid}
	_, adminInfo, _ := l.svcCtx.CultrueDbStruct.SyAdminDao.GetInfo(whereAdmin)
	add := make(map[string]interface{})
	add["user_id"] = req.UserId
	add["points"] = req.Points
	add["source"] = req.Source
	add["notes"] = "管理员：" + adminInfo.RelName + ",在后台进行修改"
	insert := make([]map[string]interface{}, 0)
	insert = append(insert, add)

	updateUserPoint := make(map[string]interface{})
	updateUserPoint["points"] = req.Points + userPoint
	//开启事务
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if errSession := session.Begin(); nil != errSession {
		returnData.Code = 100004
		returnData.Message = "开启事务失败"
		return returnData, errSession
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	//添加记录
	insertErr := l.svcCtx.CultrueDbStruct.SyUserPointsDao.Insert(insert, session)
	if insertErr != nil {
		session.Rollback()
		returnData.Code = 100005
		returnData.Message = "添加用户的积分数据错误"
		return returnData, insertErr
	}
	//修改用户的积分
	updateUserPointError := l.svcCtx.CultrueDbStruct.SyUserDao.Update(findUserWhere, updateUserPoint, session)
	if updateUserPointError != nil {
		session.Rollback()
		returnData.Code = 100006
		returnData.Message = "修改用户的数据错误"
		return returnData, updateUserPointError
	}
	session.Commit()
	return returnData, nil
}

// 获取用户的积分
func (l *UserLogic) UserPointsListLogin(req admin.GetAdminToUserPointListReq) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	whereInterface := []interface{}{"is_delete = ? and user_id = ?", 0, req.UserId}
	userPointsCount, userPointsList, userPointsListErr := l.svcCtx.CultrueDbStruct.SyUserPointsDao.GetUserPointsListPage(whereInterface, req.Page, req.PageSize, "id desc")
	if userPointsListErr != nil {
		returnData.Code = 100001
		returnData.Message = "开启事务失败"
		return returnData, userPointsListErr
	}
	data := adminResp.GetUserPointsListResp{
		Count:    userPointsCount,
		PageSize: req.PageSize,
		Page:     req.Page,
	}
	if userPointsCount == 0 {
		returnData.Data = data
		return returnData, nil
	}
	for _, v := range userPointsList {
		info := adminResp.GetUserPointsInfoResp{
			SourceId: v.SourceId,
			Points:   v.Points,
			Source:   v.Source,
			CreateAt: helper.InterfaceHelperObject.ToString(v.CreateTime),
			Notes:    v.Notes,
		}
		data.List = append(data.List, info)
	}
	returnData.Data = data
	return returnData, nil
}

// 新增用户
func (l *UserLogic) AddUserLogic(req admin.AddUserReq, adminId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()

	// 校验手机号是否已存在
	checkWhere := []interface{}{"is_delete = ? and phone = ?", 0, req.Phone}
	has, _, checkErr := l.svcCtx.CultrueDbStruct.SyUserDao.GetInfo(checkWhere)
	if checkErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询用户信息失败"
		return returnData, checkErr
	}
	if has {
		returnData.Code = 100002
		returnData.Message = "该手机号已被注册"
		return returnData, nil
	}

	add := map[string]interface{}{
		"user_name":  req.UserName,
		"rel_name":   req.RelName,
		"phone":      req.Phone,
		"pwd":        helper.Md5HelperObject.Sha256ToString("123456"),
		"add_aid":    adminId,
		"update_aid": adminId,
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

	if insertErr := l.svcCtx.CultrueDbStruct.SyUserDao.Insert([]map[string]interface{}{add}, session); insertErr != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "新增用户失败"
		return returnData, insertErr
	}
	session.Commit()
	return returnData, nil
}

// 修改用户基本信息（用户名、昵称、手机号、状态）
func (l *UserLogic) UpdateUserInfoLogic(req admin.UpdateUserInfoReq, adminId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()

	// 确认用户存在
	findWhere := []interface{}{"is_delete = ? and id = ?", 0, req.UserId}
	has, _, findErr := l.svcCtx.CultrueDbStruct.SyUserDao.GetInfo(findWhere)
	if findErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询用户信息失败"
		return returnData, findErr
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "用户不存在"
		return returnData, nil
	}

	// 若修改手机号，检查新号码是否已被占用
	if req.Phone != "" {
		phoneWhere := []interface{}{"is_delete = ? and phone = ? and id != ?", 0, req.Phone, req.UserId}
		phoneHas, _, phoneErr := l.svcCtx.CultrueDbStruct.SyUserDao.GetInfo(phoneWhere)
		if phoneErr != nil {
			returnData.Code = 100003
			returnData.Message = "查询手机号失败"
			return returnData, phoneErr
		}
		if phoneHas {
			returnData.Code = 100004
			returnData.Message = "该手机号已被其他用户使用"
			return returnData, nil
		}
	}

	update := map[string]interface{}{
		"update_aid": adminId,
		"status":     req.Status,
	}
	if req.UserName != "" {
		update["user_name"] = req.UserName
	}
	if req.RelName != "" {
		update["rel_name"] = req.RelName
	}
	if req.Phone != "" {
		update["phone"] = req.Phone
	}

	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err = session.Begin(); err != nil {
		returnData.Code = 100005
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()

	updateWhere := []interface{}{"id = ?", req.UserId}
	if updateErr := l.svcCtx.CultrueDbStruct.SyUserDao.Update(updateWhere, update, session); updateErr != nil {
		session.Rollback()
		returnData.Code = 100006
		returnData.Message = "修改用户信息失败"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}
