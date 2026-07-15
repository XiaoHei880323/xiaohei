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

type AdminActivityLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminActivityLogic {
	return &AdminActivityLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取活动的列表
func (l AdminActivityLogic) ActivityListLogic(req activityReq.ActivityListReq) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	sql := "is_delete = ?"
	whereSlice := []interface{}{1}
	if req.ActivityName != "" {
		sql = sql + " and activity_name like \"% ? %\""
		whereSlice = append(whereSlice, req.ActivityName)
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
		returnData.Message = "查询活动数据错误"
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
	//获取用户的信息
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

// 添加活动信息
func (l AdminActivityLogic) ActivityAddLogic(req activityReq.ActivityAddInfoReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	insertMap := make(map[string]interface{})
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
	//开启事务
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
		returnData.Message = "插入活动数据信息失败"
		return returnData, insertErr
	}
	session.Commit()
	return returnData, nil
}

// 修改活动
func (l AdminActivityLogic) ActivityUpdateLogic(req activityReq.ActivityUpdateInfoReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	adminRoleServie := serv.NewAdminUserRoleService(l.ctx, l.svcCtx)
	userRole := adminRoleServie.AmdinRole(userId)
	if userRole != 1 {
		returnData.Code = 100001
		returnData.Message = "当前用户角色不可以进行修改活动"
		return returnData, nil
	}
	//获取活动信息
	whereActivity := []interface{}{"id = ? and is_delete = ?", req.ActivityId, 0}
	activityHas, _, activityErr := l.svcCtx.CultrueDbStruct.SyActivityDao.GetInfo(whereActivity)
	if activityErr != nil {
		returnData.Code = 100002
		returnData.Message = "查询活动信息错误"
		return returnData, activityErr
	}
	if !activityHas {
		returnData.Code = 100003
		returnData.Message = "未查询到当前的活动信息"
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
	updateErr := l.svcCtx.CultrueDbStruct.SyActivityDao.Update(whereActivity, updateMap, session)
	if updateErr != nil {
		returnData.Code = 100005
		returnData.Message = "修改活动信息失败"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}

// 删除活动
func (l AdminActivityLogic) ActivityDeleteLogic(req activityReq.ActivityDeleteReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	adminRoleServie := serv.NewAdminUserRoleService(l.ctx, l.svcCtx)
	userRole := adminRoleServie.AmdinRole(userId)
	if userRole != 1 {
		returnData.Code = 100001
		returnData.Message = "当前用户角色不可以进行删除活动"
		return returnData, nil
	}
	//获取活动信息
	whereActivity := []interface{}{"id = ? and is_delete = ?", req.ActivityId, 0}
	activityHas, _, activityErr := l.svcCtx.CultrueDbStruct.SyActivityDao.GetInfo(whereActivity)
	if activityErr != nil {
		returnData.Code = 100002
		returnData.Message = "查询活动信息错误"
		return returnData, activityErr
	}
	if !activityHas {
		returnData.Code = 100003
		returnData.Message = "未查询到当前的活动信息"
		return returnData, nil
	}
	updateMap := make(map[string]interface{})
	updateMap["is_delete"] = 1
	updateMap["update_uid"] = userId
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
	updateErr := l.svcCtx.CultrueDbStruct.SyActivityDao.Update(whereActivity, updateMap, session)
	if updateErr != nil {
		returnData.Code = 100005
		returnData.Message = "修改活动信息失败"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}
