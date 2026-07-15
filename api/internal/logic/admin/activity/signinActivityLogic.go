package activity

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/serv"
	"api/internal/svc"
	"api/reqs/activityReq"
	"api/resp"
	"api/resp/activityResp"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

// ActivityType: 1=签到活动
const signinActivityType = 1

type SigninActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSigninActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SigninActivityLogic {
	return &SigninActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取签到活动列表
func (l SigninActivityLogic) ActivityListLogic(req activityReq.ActivityListReq) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	sql := "activity_type = ?"
	whereSlice := []interface{}{signinActivityType}
	if req.ActivityName != "" {
		sql += " and activity_name like ?"
		whereSlice = append(whereSlice, "%"+req.ActivityName+"%")
	}
	if req.ActivityStartTime != "" {
		sql += " and activity_starting_time = ?"
		whereSlice = append(whereSlice, req.ActivityStartTime)
	}
	if req.ActivityEndTime != "" {
		sql += " and activity_ending_time = ?"
		whereSlice = append(whereSlice, req.ActivityEndTime)
	}
	where := []interface{}{}
	where = append(where, sql)
	where = append(where, whereSlice...)
	activityCount, activityList, activityErr := l.svcCtx.CultrueDbStruct.SyActivityDao.GetCorList(where, req.Page, req.PageSize, "id desc")
	if activityErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询签到活动数据错误"
		return returnData, activityErr
	}
	data := activityResp.GetActivityList{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    activityCount,
	}
	if activityCount == 0 {
		returnData.Data = data
		return returnData, nil
	}
	userIds := make([]int, 0)
	for _, activity := range activityList {
		userIds = append(userIds, activity.AddUid)
	}
	adminUserService := serv.NewAdminUserService(l.ctx, l.svcCtx)
	userMap, userMapErr := adminUserService.GetAdminIdUserMap(userIds)
	if userMapErr != nil {
		returnData.Code = 100002
		returnData.Message = "查询添加活动人员错误"
		return returnData, userMapErr
	}
	for _, activity := range activityList {
		userName := ""
		if user, userOk := userMap[activity.AddUid]; userOk {
			userName = user
		}
		activityInfo := activityResp.GetActivityListInfoResp{
			ActivityId:          activity.Id,
			ActivityName:        activity.ActivityName,
			ActivityImage:       activity.ActivityImg,
			ActivityText:        activity.ActivityText,
			ActivityStartTime:   helper.TimeEnumFuncObject.StringTime(activity.ActivityStartingTime),
			ActivityEndTime:     helper.TimeEnumFuncObject.StringTime(activity.ActivityEndTime),
			ActivityPreviewTime: helper.TimeEnumFuncObject.StringTime(activity.ActivityPreviewTime),
			ActivityPoints:      activity.ActivityPoints,
			ActivityAddTime:     helper.TimeEnumFuncObject.StringTime(activity.AddTime),
			ActivityAddUser:     userName,
		}
		data.ActivityList = append(data.ActivityList, activityInfo)
	}
	returnData.Data = data
	return returnData, nil
}

// 添加签到活动
func (l SigninActivityLogic) ActivityAddLogic(req activityReq.ActivityAddInfoReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	insertMap := make(map[string]interface{})
	insertMap["activity_type"] = signinActivityType
	insertMap["activity_name"] = req.ActivityName
	insertMap["activity_img"] = req.ActivityImage
	insertMap["activity_text"] = req.ActivityText
	insertMap["activity_starting_time"] = req.ActivityStartTime
	insertMap["activity_end_time"] = req.ActivityEndTime
	insertMap["activity_preview_time"] = req.ActivityPreviewTime
	insertMap["activity_points"] = req.ActivityPoints
	insertMap["add_uid"] = userId
	insertMap["update_uid"] = userId
	addArray := make([]map[string]interface{}, 0)
	addArray = append(addArray, insertMap)
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if errSession := session.Begin(); nil != errSession {
		returnData.Code = 100002
		returnData.Message = "开启事务失败"
		return returnData, errSession
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	insertErr := l.svcCtx.CultrueDbStruct.SyActivityDao.Insert(addArray, session)
	if insertErr != nil {
		session.Rollback()
		returnData.Code = 100003
		returnData.Message = "插入签到活动数据失败"
		return returnData, insertErr
	}
	session.Commit()
	return returnData, nil
}

// 修改签到活动
func (l SigninActivityLogic) ActivityUpdateLogic(req activityReq.ActivityUpdateInfoReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	adminRoleServie := serv.NewAdminUserRoleService(l.ctx, l.svcCtx)
	userRole := adminRoleServie.AmdinRole(userId)
	if userRole != 1 {
		returnData.Code = 100001
		returnData.Message = "当前用户角色不可以进行修改签到活动"
		return returnData, nil
	}
	whereActivity := []interface{}{"id = ? and is_delete = ? and activity_type = ?", req.ActivityId, 1, signinActivityType}
	activityHas, _, activityErr := l.svcCtx.CultrueDbStruct.SyActivityDao.GetInfo(whereActivity)
	if activityErr != nil {
		returnData.Code = 100002
		returnData.Message = "查询签到活动信息错误"
		return returnData, activityErr
	}
	if !activityHas {
		returnData.Code = 100003
		returnData.Message = "未查询到当前的签到活动信息"
		return returnData, nil
	}
	updateMap := make(map[string]interface{})
	updateMap["activity_name"] = req.ActivityName
	updateMap["activity_img"] = req.ActivityImage
	updateMap["activity_text"] = req.ActivityText
	updateMap["activity_starting_time"] = req.ActivityStartTime
	updateMap["activity_end_time"] = req.ActivityEndTime
	updateMap["activity_preview_time"] = req.ActivityPreviewTime
	updateMap["activity_points"] = req.ActivityPoints
	updateMap["update_uid"] = userId
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
	updateErr := l.svcCtx.CultrueDbStruct.SyActivityDao.Update(whereActivity, updateMap, session)
	if updateErr != nil {
		returnData.Code = 100005
		returnData.Message = "修改签到活动信息失败"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}

// 删除签到活动
func (l SigninActivityLogic) ActivityDeleteLogic(req activityReq.ActivityDeleteReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	adminRoleServie := serv.NewAdminUserRoleService(l.ctx, l.svcCtx)
	userRole := adminRoleServie.AmdinRole(userId)
	if userRole != 1 {
		returnData.Code = 100001
		returnData.Message = "当前用户角色不可以进行删除签到活动"
		return returnData, nil
	}
	whereActivity := []interface{}{"id = ? and is_delete = ? and activity_type = ?", req.ActivityId, 0, signinActivityType}
	activityHas, _, activityErr := l.svcCtx.CultrueDbStruct.SyActivityDao.GetInfo(whereActivity)
	if activityErr != nil {
		returnData.Code = 100002
		returnData.Message = "查询签到活动信息错误"
		return returnData, activityErr
	}
	if !activityHas {
		returnData.Code = 100003
		returnData.Message = "未查询到当前的签到活动信息"
		return returnData, nil
	}
	updateMap := make(map[string]interface{})
	updateMap["is_delete"] = 1
	updateMap["update_uid"] = userId
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
	updateErr := l.svcCtx.CultrueDbStruct.SyActivityDao.Update(whereActivity, updateMap, session)
	if updateErr != nil {
		returnData.Code = 100005
		returnData.Message = "删除签到活动失败"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}
