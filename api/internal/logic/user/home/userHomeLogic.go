package home

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/serv"
	"api/internal/svc"
	"api/model"
	"api/reqs/userReq"
	"api/resp"
	"api/resp/userResp"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// hotThreshold 前 N 条数据打上热销标签
const hotThreshold = 3

type UserHomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomeLogic {
	return &UserHomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取首页配置数据（banner 等）
func (l *UserHomeLogic) GetHomeConfig(req userReq.HomeConfigReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"is_delete = ?", 0}
	count, list, err := l.svcCtx.CultrueDbStruct.SyHomeConfigDao.GetList(where, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询首页配置数据错误"
		return returnData, err
	}
	items := make([]userResp.HomeConfigItem, 0, len(list))
	for _, v := range list {
		items = append(items, userResp.HomeConfigItem{
			ConfigId:     v.Id,
			ConfigName:   v.ConfigName,
			ConfigImage:  v.ConfigImage,
			ActivityId:   v.ActivityId,
			ActivityType: v.ActivityType,
			Sort:         v.Sort,
		})
	}
	returnData.Data = userResp.HomeConfigResp{Count: count, List: items}
	return returnData, nil
}

// 获取首页数据：活动、商品、景点（倒序 + 热销标签）
func (l *UserHomeLogic) GetHomeData(req userReq.HomeDataReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()

	// 活动（签到 + 秒杀）id 倒序
	var activityList []*model.SyActivity
	activityErr := l.svcCtx.Engine.
		Where("is_delete = ?", 0).
		OrderBy("id desc").
		Limit(req.PageSize, req.PageSize*(req.Page-1)).
		Find(&activityList)
	if activityErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动数据错误"
		return returnData, activityErr
	}
	activities := make([]userResp.ActivityItem, 0, len(activityList))
	for i, v := range activityList {
		activities = append(activities, userResp.ActivityItem{
			ActivityId:    v.Id,
			ActivityName:  v.ActivityName,
			ActivityImage: v.ActivityImg,
			ActivityType:  v.ActivityType,
			StartTime:     helper.TimeEnumFuncObject.StringTime(v.ActivityStartingTime),
			EndTime:       helper.TimeEnumFuncObject.StringTime(v.ActivityEndTime),
			Points:        v.ActivityPoints,
			IsHot:         i < hotThreshold,
		})
	}

	// 商品 id 倒序
	var goodList []*model.SyGood
	goodErr := l.svcCtx.Engine.
		OrderBy("id desc").
		Limit(req.PageSize, req.PageSize*(req.Page-1)).
		Find(&goodList)
	if goodErr != nil {
		returnData.Code = 100002
		returnData.Message = "查询商品数据错误"
		return returnData, goodErr
	}
	goods := make([]userResp.GoodItem, 0, len(goodList))
	for i, v := range goodList {
		goods = append(goods, userResp.GoodItem{
			GoodId:    v.Id,
			GoodName:  v.GoodName,
			GoodImage: v.GoodImg,
			GoodPrice: v.GoodPrice,
			IsHot:     i < hotThreshold,
		})
	}

	// 景点 id 倒序
	var scenicList []*model.SyScenicSpot
	scenicErr := l.svcCtx.Engine.
		Where("is_delete = ?", 0).
		OrderBy("id desc").
		Limit(req.PageSize, req.PageSize*(req.Page-1)).
		Find(&scenicList)
	if scenicErr != nil {
		returnData.Code = 100003
		returnData.Message = "查询景点数据错误"
		return returnData, scenicErr
	}
	scenics := make([]userResp.ScenicItem, 0, len(scenicList))
	for i, v := range scenicList {
		scenics = append(scenics, userResp.ScenicItem{
			SpotId:      v.Id,
			SpotName:    v.SpotName,
			Longitude:   v.Longitude,
			Latitude:    v.Latitude,
			TicketPrice: v.TicketPrice,
			IsHot:       i < hotThreshold,
		})
	}

	returnData.Data = userResp.HomeDataResp{
		Activities: activities,
		Goods:      goods,
		Scenics:    scenics,
	}
	return returnData, nil
}

// 获取首页活动配置（最多5条，优先当前时间有效的，无则取默认）
func (l *UserHomeLogic) GetActivityConfig(req userReq.UserActivityConfigReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	now := time.Now()
	// 查询当前时间有效且上线的活动配置，最多5条
	var list []*model.SyActivityConfig
	err := l.svcCtx.Engine.
		Where("is_delete = ? AND status = ? AND start_time <= ? AND end_time >= ?", 0, 1, now, now).
		OrderBy("id desc").
		Limit(5, 0).
		Find(&list)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询活动配置数据错误"
		return returnData, err
	}
	// 若无有效数据则获取默认配置
	if len(list) == 0 {
		var defaultList []*model.SyActivityConfig
		defaultErr := l.svcCtx.Engine.
			Where("is_delete = ? AND is_default = ?", 0, 1).
			OrderBy("id desc").
			Limit(5, 0).
			Find(&defaultList)
		if defaultErr != nil {
			returnData.Code = 100002
			returnData.Message = "查询默认活动配置数据错误"
			return returnData, defaultErr
		}
		list = defaultList
	}
	result := make([]userResp.UserActivityConfigItem, 0, len(list))
	for _, c := range list {
		items, _ := l.svcCtx.CultrueDbStruct.SyActivityConfigItemDao.GetByConfigId(c.Id)
		itemData := make([]userResp.UserActivityConfigItemData, 0, len(items))
		for _, it := range items {
			itemData = append(itemData, userResp.UserActivityConfigItemData{
				ItemId:     it.Id,
				ActivityId: it.ActivityId,
				Sort:       it.Sort,
			})
		}
		result = append(result, userResp.UserActivityConfigItem{
			ConfigId:     c.Id,
			ConfigName:   c.ConfigName,
			ConfigImage:  c.ConfigImage,
			StartTime:    helper.TimeEnumFuncObject.StringTime(c.StartTime),
			EndTime:      helper.TimeEnumFuncObject.StringTime(c.EndTime),
			IsDefault:    c.IsDefault,
			ActivityType: c.ActivityType,
			Items:        itemData,
		})
	}
	returnData.Data = userResp.UserActivityConfigResp{List: result}
	return returnData, nil
}

// 获取有效公告列表（按公告时间倒序）
func (l *UserHomeLogic) GetNoticeList(req userReq.NoticeListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"is_delete = ? AND notice_status = ?", 0, 1}
	count, list, err := l.svcCtx.CultrueDbStruct.SyNoticeDao.GetList(where, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询公告数据错误"
		return returnData, err
	}
	items := make([]userResp.NoticeItem, 0, len(list))
	for _, v := range list {
		items = append(items, userResp.NoticeItem{
			NoticeId:   v.Id,
			NoticeName: v.NoticeName,
		})
	}
	returnData.Data = userResp.NoticeListResp{Count: count, List: items}
	return returnData, nil
}

// 获取公告详情
func (l *UserHomeLogic) GetNoticeDetail(req userReq.NoticeDetailReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ? AND notice_status = ?", req.NoticeId, 0, 1}
	has, notice, err := l.svcCtx.CultrueDbStruct.SyNoticeDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询公告详情错误"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "未查询到公告信息"
		return returnData, nil
	}
	publishUser := ""
	if notice.PublishUid > 0 {
		adminUserService := serv.NewAdminUserService(l.ctx, l.svcCtx)
		userMap, _ := adminUserService.GetAdminIdUserMap([]int{notice.PublishUid})
		if name, ok := userMap[notice.PublishUid]; ok {
			publishUser = name
		}
	}
	returnData.Data = userResp.NoticeDetailResp{
		NoticeId:      notice.Id,
		NoticeName:    notice.NoticeName,
		PublishTime:   helper.TimeEnumFuncObject.StringTime(notice.PublishTime),
		PublishUser:   publishUser,
		NoticeContent: notice.NoticeContent,
	}
	return returnData, nil
}
