package home

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/serv"
	"api/internal/svc"
	"api/reqs/activityConfigReq"
	"api/resp"
	"api/resp/activityConfigResp"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ActivityConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewActivityConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ActivityConfigLogic {
	return &ActivityConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 活动配置列表
func (l *ActivityConfigLogic) ListLogic(req activityConfigReq.ActivityConfigListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereStr := "is_delete = ?"
	whereArgs := []interface{}{0}
	if req.ConfigName != "" {
		whereStr += " AND config_name LIKE ?"
		whereArgs = append(whereArgs, "%"+req.ConfigName+"%")
	}
	if req.ActivityType > 0 {
		whereStr += " AND activity_type = ?"
		whereArgs = append(whereArgs, req.ActivityType)
	}
	whereSlice := append([]interface{}{whereStr}, whereArgs...)
	count, list, err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.GetList(whereSlice, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置列表失败"
		return returnData, err
	}
	data := activityConfigResp.ActivityConfigList{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    count,
		List:     []activityConfigResp.ActivityConfigInfo{},
	}
	if count == 0 {
		returnData.Data = data
		return returnData, nil
	}
	addUids := make([]int, 0, len(list))
	for _, c := range list {
		addUids = append(addUids, c.AddUid)
	}
	adminUserService := serv.NewAdminUserService(l.ctx, l.svcCtx)
	userMap, _ := adminUserService.GetAdminIdUserMap(addUids)

	for _, c := range list {
		addUser := ""
		if name, ok := userMap[c.AddUid]; ok {
			addUser = name
		}
		data.List = append(data.List, activityConfigResp.ActivityConfigInfo{
			ConfigId:     c.Id,
			ConfigName:   c.ConfigName,
			ConfigImage:  c.ConfigImage,
			StartTime:    helper.TimeEnumFuncObject.StringTime(c.StartTime),
			EndTime:      helper.TimeEnumFuncObject.StringTime(c.EndTime),
			IsDefault:    c.IsDefault,
			ActivityType: c.ActivityType,
			Status:       c.Status,
			AddTime:      helper.TimeEnumFuncObject.StringTime(c.AddTime),
			AddUser:      addUser,
		})
	}
	returnData.Data = data
	return returnData, nil
}

// 添加活动配置
func (l *ActivityConfigLogic) AddLogic(req activityConfigReq.ActivityConfigAddReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	if req.ActivityType < 1 || req.ActivityType > 4 {
		returnData.Code = 100001
		returnData.Message = "活动类型错误，1:签到活动 2:秒杀活动 3:商品活动 4:景点活动"
		return returnData, nil
	}
	startTime, err := time.ParseInLocation(helper.TimeEnumObject.DateTimeLayout, req.StartTime, time.Local)
	if err != nil {
		returnData.Code = 100002
		returnData.Message = "开始时间格式错误，格式：2006-01-02 15:04:05"
		return returnData, nil
	}
	endTime, err := time.ParseInLocation(helper.TimeEnumObject.DateTimeLayout, req.EndTime, time.Local)
	if err != nil {
		returnData.Code = 100003
		returnData.Message = "结束时间格式错误，格式：2006-01-02 15:04:05"
		return returnData, nil
	}
	if !endTime.After(startTime) {
		returnData.Code = 100004
		returnData.Message = "结束时间必须大于开始时间"
		return returnData, nil
	}
	insertMap := map[string]interface{}{
		"config_name":   req.ConfigName,
		"config_image":  req.ConfigImage,
		"start_time":    startTime.Format(helper.TimeEnumObject.DateTimeLayout),
		"end_time":      endTime.Format(helper.TimeEnumObject.DateTimeLayout),
		"activity_type": req.ActivityType,
		"is_default":    0,
		"status":        0,
		"add_uid":       uid,
		"update_uid":    uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100005
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.Insert([]map[string]interface{}{insertMap}, session); err != nil {
		session.Rollback()
		returnData.Code = 100006
		returnData.Message = "添加活动配置失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 修改活动配置基础信息
func (l *ActivityConfigLogic) UpdateLogic(req activityConfigReq.ActivityConfigUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	if req.ActivityType < 1 || req.ActivityType > 4 {
		returnData.Code = 100001
		returnData.Message = "活动类型错误，1:签到活动 2:秒杀活动 3:商品活动 4:景点活动"
		return returnData, nil
	}
	whereConfig := []interface{}{"id = ? AND is_delete = ?", req.ConfigId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.GetInfo(whereConfig)
	if err != nil {
		returnData.Code = 100002
		returnData.Message = "查询活动配置信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100003
		returnData.Message = "未查询到活动配置信息"
		return returnData, nil
	}
	startTime, err := time.ParseInLocation(helper.TimeEnumObject.DateTimeLayout, req.StartTime, time.Local)
	if err != nil {
		returnData.Code = 100004
		returnData.Message = "开始时间格式错误，格式：2006-01-02 15:04:05"
		return returnData, nil
	}
	endTime, err := time.ParseInLocation(helper.TimeEnumObject.DateTimeLayout, req.EndTime, time.Local)
	if err != nil {
		returnData.Code = 100005
		returnData.Message = "结束时间格式错误，格式：2006-01-02 15:04:05"
		return returnData, nil
	}
	if !endTime.After(startTime) {
		returnData.Code = 100006
		returnData.Message = "结束时间必须大于开始时间"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"config_name":   req.ConfigName,
		"config_image":  req.ConfigImage,
		"start_time":    startTime.Format(helper.TimeEnumObject.DateTimeLayout),
		"end_time":      endTime.Format(helper.TimeEnumObject.DateTimeLayout),
		"activity_type": req.ActivityType,
		"update_uid":    uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100007
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.Update(whereConfig, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100008
		returnData.Message = "修改活动配置失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 删除活动配置
func (l *ActivityConfigLogic) DeleteLogic(req activityConfigReq.ActivityConfigDeleteReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereConfig := []interface{}{"id = ? AND is_delete = ?", req.ConfigId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.GetInfo(whereConfig)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到活动配置信息"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"is_delete":  1,
		"update_uid": uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.Update(whereConfig, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "删除活动配置失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 设置为默认活动配置（同时清除其他默认）
func (l *ActivityConfigLogic) SetDefaultLogic(req activityConfigReq.ActivityConfigSetDefaultReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereConfig := []interface{}{"id = ? AND is_delete = ?", req.ConfigId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.GetInfo(whereConfig)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到活动配置信息"
		return returnData, nil
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	// 先清除所有默认
	clearWhere := []interface{}{"is_default = ? AND is_delete = ?", 1, 0}
	clearMap := map[string]interface{}{"is_default": 0, "update_uid": uid}
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.UpdateAll(clearWhere, clearMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "清除默认标记失败"
		return returnData, err
	}
	// 再设置当前为默认
	setMap := map[string]interface{}{"is_default": 1, "update_uid": uid}
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.Update(whereConfig, setMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100005
		returnData.Message = "设置默认活动配置失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 上下线活动配置
func (l *ActivityConfigLogic) UpdateStatusLogic(req activityConfigReq.ActivityConfigUpdateStatusReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	if req.Status != 0 && req.Status != 1 {
		returnData.Code = 100001
		returnData.Message = "状态值错误，0:下线 1:上线"
		return returnData, nil
	}
	whereConfig := []interface{}{"id = ? AND is_delete = ?", req.ConfigId, 0}
	has, config, err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.GetInfo(whereConfig)
	if err != nil {
		returnData.Code = 100002
		returnData.Message = "查询活动配置信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100003
		returnData.Message = "未查询到活动配置信息"
		return returnData, nil
	}
	if config.Status == req.Status {
		if req.Status == 1 {
			returnData.Message = "活动配置已上线，无需重复操作"
		} else {
			returnData.Message = "活动配置已下线，无需重复操作"
		}
		returnData.Code = 100004
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"status":     req.Status,
		"update_uid": uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100005
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.Update(whereConfig, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100006
		returnData.Message = "操作失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 活动配置项列表
func (l *ActivityConfigLogic) ItemListLogic(req activityConfigReq.ActivityConfigItemListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereSlice := []interface{}{"config_id = ? AND is_delete = ?", req.ConfigId, 0}
	count, list, err := l.svcCtx.CultrueDbStruct.SyActivityConfigItemDao.GetList(whereSlice, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置项列表失败"
		return returnData, err
	}
	data := activityConfigResp.ActivityConfigItemList{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    count,
		List:     []activityConfigResp.ActivityConfigItemInfo{},
	}
	if count == 0 {
		returnData.Data = data
		return returnData, nil
	}
	addUids := make([]int, 0, len(list))
	for _, item := range list {
		addUids = append(addUids, item.AddUid)
	}
	adminUserService := serv.NewAdminUserService(l.ctx, l.svcCtx)
	userMap, _ := adminUserService.GetAdminIdUserMap(addUids)

	for _, item := range list {
		addUser := ""
		if name, ok := userMap[item.AddUid]; ok {
			addUser = name
		}
		data.List = append(data.List, activityConfigResp.ActivityConfigItemInfo{
			ItemId:     item.Id,
			ConfigId:   item.ConfigId,
			ActivityId: item.ActivityId,
			Sort:       item.Sort,
			AddTime:    helper.TimeEnumFuncObject.StringTime(item.AddTime),
			AddUser:    addUser,
		})
	}
	returnData.Data = data
	return returnData, nil
}

// 添加活动配置项
func (l *ActivityConfigLogic) ItemAddLogic(req activityConfigReq.ActivityConfigItemAddReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereConfig := []interface{}{"id = ? AND is_delete = ?", req.ConfigId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyActivityConfigDao.GetInfo(whereConfig)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到活动配置信息"
		return returnData, nil
	}
	insertMap := map[string]interface{}{
		"config_id":   req.ConfigId,
		"activity_id": req.ActivityId,
		"sort":        req.Sort,
		"add_uid":     uid,
		"update_uid":  uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigItemDao.Insert([]map[string]interface{}{insertMap}, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "添加活动配置项失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 修改活动配置项
func (l *ActivityConfigLogic) ItemUpdateLogic(req activityConfigReq.ActivityConfigItemUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereItem := []interface{}{"id = ? AND is_delete = ?", req.ItemId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyActivityConfigItemDao.GetInfo(whereItem)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置项信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到活动配置项信息"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"activity_id": req.ActivityId,
		"sort":        req.Sort,
		"update_uid":  uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigItemDao.Update(whereItem, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改活动配置项失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 删除活动配置项
func (l *ActivityConfigLogic) ItemDeleteLogic(req activityConfigReq.ActivityConfigItemDeleteReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereItem := []interface{}{"id = ? AND is_delete = ?", req.ItemId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyActivityConfigItemDao.GetInfo(whereItem)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置项信息失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到活动配置项信息"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"is_delete":  1,
		"update_uid": uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100003
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyActivityConfigItemDao.Update(whereItem, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "删除活动配置项失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}
