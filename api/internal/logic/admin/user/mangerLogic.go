package user

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/serv"
	"api/internal/svc"
	"api/reqs/admin"
	"api/resp"
	"api/resp/adminResp"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) AdminUserLogin(req admin.UserLogin) (resp *resp.CommonReply, err error) {
	// todo: add your logic here and delete this line
	returnData := reponse.ReturnStruct()
	whereSlice := []interface{}{"status = ? and is_delete = ? and phone = ? and pwd = ?", 0, 0, req.Phone, helper.Md5HelperObject.Sha256ToString(req.Pwd)}
	userInfoHas, userInfo, userInfoErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetInfo(whereSlice)
	if userInfoErr != nil {
		returnData.Code = 100001
		returnData.Message = "没有查询到用户信息"
		return returnData, userInfoErr
	}
	if !userInfoHas {
		returnData.Code = 100002
		returnData.Message = "用户信息不正确，请确认用户信息"
		return returnData, nil
	}
	user := adminResp.AdminInfoResp{
		UserId:   userInfo.Id,
		UserName: userInfo.UserName,
		NickName: userInfo.RelName,
	}
	// 设置claims
	claims := jwt.MapClaims{
		"userId": userInfo.Id,
		"exp":    time.Now().Add(time.Hour * 8).Unix(), // token 1小时后过期
		"iat":    time.Now().Unix(),
	}
	// 生成签名的token字符串
	tokenService := serv.NewTokenService()
	signedToken, tokenErr := tokenService.GenerateToken(l.svcCtx.Config.Auth.AdminSecretKey, claims)

	if tokenErr != nil {
		returnData.Code = 100003
		returnData.Message = "生成token错误"
		return returnData, tokenErr
	}
	user.Token = signedToken
	returnData.Data = user
	return returnData, err
}

// 修改自己个人信息
func (l *AdminLogic) UpdateInfoLogic(req admin.UpdateInfoReq) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	whereSlice := []interface{}{"id = ? and status = ? and is_delete = ?", req.UserId, 0, 0}
	userInfoHas, userInfo, userInfoErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetInfo(whereSlice)
	if userInfoErr != nil {
		returnData.Code = 100001
		returnData.Message = "没有查询到用户信息"
		return returnData, userInfoErr
	}
	if !userInfoHas {
		returnData.Code = 100002
		returnData.Message = "用户信息不正确，请确认用户信息"
		return returnData, nil
	}
	update := make(map[string]interface{})
	update["update_aid"] = req.UserId

	if req.OldPwd != "" { //如果存在历史的密钥
		if helper.Md5HelperObject.Sha256ToString(req.OldPwd) != userInfo.Pwd {
			returnData.Code = 100005
			returnData.Message = "用户的原始密码和先密码不一致"
			return returnData, nil
		}
		if req.NewOldPwd != "" {
			update["pwd"] = helper.Md5HelperObject.Sha256ToString(req.NewOldPwd)
		}
	}

	if req.Phone != 0 && req.Phone != int(userInfo.Phone) {
		whereUpdate := []interface{}{"phone = ? and  status = ? and is_delete = ? and id != ?", req.Phone, 0, 0, userInfo.Id}
		phoneCount, phoneErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetCount(whereUpdate)
		if phoneErr != nil {
			returnData.Code = 100003
			returnData.Message = "查询手机号码用户数据错误"
			return returnData, phoneErr
		}
		if phoneCount > 1 {
			returnData.Code = 100004
			returnData.Message = "手机号码已经被其他用户使用"
			return returnData, phoneErr
		}
		update["phone"] = req.Phone
	}
	if req.UserName != "" {
		update["user_name"] = req.UserName
	}
	if req.RelName != "" {
		update["rel_name"] = req.RelName
	}

	//开启事务
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if errSession := session.Begin(); nil != errSession {
		returnData.Code = 100006
		returnData.Message = "开启事务失败"
		return returnData, errSession
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	updateError := l.svcCtx.CultrueDbStruct.SyAdminDao.Update(whereSlice, update, session)
	if updateError != nil {
		session.Rollback()
		returnData.Code = 100007
		returnData.Message = "开启事务失败"
		return returnData, updateError
	}
	session.Commit()
	return returnData, nil
}

// 添加管理员用户
func (l *AdminLogic) GetAdminUserListLogic(req admin.GetAdminUserListReq) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	sql := "is_delete = ?"
	whereSlice := []interface{}{0}
	if req.Phone != "" {
		sql += " and phone = ?"
		whereSlice = append(whereSlice, req.Phone)
	}
	if req.UserName != "" {
		sql += " and user_name = ?"
		whereSlice = append(whereSlice, req.UserName)
	}
	if req.RelName != "" {
		sql += " and rel_name = ?"
		whereSlice = append(whereSlice, req.RelName)
	}
	where := []interface{}{}
	where = append(where, sql)
	where = append(where, whereSlice...)
	userCount, userList, userErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetCorList(where, req.Page, req.PageSize, "id desc")
	if userErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询管理员列表错误"
		return returnData, userErr
	}
	data := adminResp.GetAdminUserListResp{
		PageSize: req.PageSize,
		Page:     req.Page,
		Count:    userCount,
	}
	if userCount < 1 {
		returnData.Data = data
		return returnData, nil
	}
	adminList := make([]adminResp.AdminInfoResp, 0)
	for _, v := range userList {
		adminInfo := adminResp.AdminInfoResp{
			UserId:   v.Id,
			UserName: v.UserName,
			NickName: v.RelName,
			Status:   v.Status,
			CreateAt: helper.TimeEnumFuncObject.StringTime(v.CreateTime),
		}
		adminList = append(adminList, adminInfo)
	}
	data.AdminList = adminList
	returnData.Data = data
	return returnData, nil
}

// 添加管理员用户
func (l *AdminLogic) AddAdminUserLogic(req admin.AddAdminUserReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	//添加用户的权限
	adminRoleServie := serv.NewAdminUserRoleService(l.ctx, l.svcCtx)
	userRole := adminRoleServie.AmdinRole(userId)
	if userRole != 1 {
		returnData.Code = 100001
		returnData.Message = "当前用户角色不可以进行添加活动"
		return returnData, nil
	}
	whereSlice := []interface{}{"id = ?", req.UserId}
	userInfoHas, userInfo, userInfoErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetInfo(whereSlice)
	if userInfoErr != nil {
		returnData.Code = 100001
		returnData.Message = "没有查询到用户信息"
		return returnData, userInfoErr
	}
	if !userInfoHas {
		returnData.Code = 100002
		returnData.Message = "用户信息不正确，请确认用户信息"
		return returnData, nil
	}
	//判断用户的手机号码是否正确
	findWhereSlice := []interface{}{"status = ? and is_delete = ? and phone = ? ", 0, 0, req.Phone}
	findUserInfoHas, _, findUserInfoErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetInfo(findWhereSlice)
	if findUserInfoErr != nil {
		returnData.Code = 100005
		returnData.Message = "查询需要新增的用户错误"
		return returnData, findUserInfoErr
	}
	if findUserInfoHas {
		returnData.Code = 100006
		returnData.Message = "查询需要新增用户手机号码存在"
		return returnData, nil
	}
	addUser := map[string]interface{}{}
	addUser["user_name"] = req.UserName
	addUser["rel_name"] = req.RelName
	addUser["phone"] = req.Phone
	addUser["pwd"] = helper.Md5HelperObject.Sha256ToString("123456")
	addUser["add_aid"] = userInfo.Id
	insertArr := make([]map[string]interface{}, 0)
	insertArr = append(insertArr, addUser)
	//开启事务
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if errSession := session.Begin(); nil != errSession {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, errSession
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	addAdminUserErr := l.svcCtx.CultrueDbStruct.SyAdminDao.Insert(insertArr, session)
	if addAdminUserErr != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "添加用户数据错误"
		return returnData, addAdminUserErr
	}
	session.Commit()
	return returnData, nil
}

// 修改其他管理员的状态
func (l *AdminLogic) UpdateAdminUserLogic(req admin.UpdateAdminUserReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	//添加用户的权限
	adminRoleServie := serv.NewAdminUserRoleService(l.ctx, l.svcCtx)
	userRole := adminRoleServie.AmdinRole(userId)
	if userRole != 1 {
		returnData.Code = 100001
		returnData.Message = "当前用户角色不可以进行添加活动"
		return returnData, nil
	}

	whereSlice := []interface{}{"id = ? and status = ? and is_delete = ?", req.UserId, 0, 0}
	userInfoHas, userInfo, userInfoErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetInfo(whereSlice)
	if userInfoErr != nil {
		returnData.Code = 100001
		returnData.Message = "没有查询到用户信息"
		return returnData, userInfoErr
	}
	if !userInfoHas {
		returnData.Code = 100002
		returnData.Message = "用户信息不正确，请确认用户信息"
		return returnData, nil
	}
	updateUserStatusWhere := []interface{}{"id = ?", req.UpdateUserId}
	updateUserHas, updateUserInfo, updateUserErr := l.svcCtx.CultrueDbStruct.SyAdminDao.GetInfo(updateUserStatusWhere)
	if updateUserErr != nil {
		returnData.Code = 100003
		returnData.Message = "用户信息不正确，请确认用户信息"
		return returnData, updateUserErr
	}
	if !updateUserHas {
		returnData.Code = 100004
		returnData.Message = "未查询到用户数据信息"
		return returnData, nil
	}

	update := make(map[string]interface{})
	update["update_aid"] = userInfo.Id
	if req.Type == 1 { //重新设置密码
		update["pwd"] = helper.Md5HelperObject.Sha256ToString("123456")
	} else if req.Type == 2 { //用户状态的修改
		if updateUserInfo.Status == 0 {
			update["status"] = 1
		} else if updateUserInfo.Status == 1 {
			update["status"] = 0
		}
	}
	//开启事务
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if errSession := session.Begin(); nil != errSession {
		returnData.Code = 100005
		returnData.Message = "开启事务失败"
		return returnData, errSession
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	updateAdminErr := l.svcCtx.CultrueDbStruct.SyAdminDao.Update(updateUserStatusWhere, update, session)
	if updateAdminErr != nil {
		session.Rollback()
		returnData.Code = 100006
		returnData.Message = "修改用户信息错误"
		return returnData, updateAdminErr
	}
	session.Commit()
	return returnData, nil
}
