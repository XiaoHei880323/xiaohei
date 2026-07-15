package scenic

import (
	reponse "api/comment/response"
	helper "api/comment/help"
	"api/internal/svc"
	"api/reqs/scenicSpotReq"
	"api/resp"
	"api/resp/scenicSpotResp"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type ScenicSpotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewScenicSpotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ScenicSpotLogic {
	return &ScenicSpotLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}

func (l *ScenicSpotLogic) ListLogic(req scenicSpotReq.ScenicSpotListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereStr := "is_delete = ?"
	whereArgs := []interface{}{0}
	if req.SpotName != "" {
		whereStr += " AND spot_name LIKE ?"
		whereArgs = append(whereArgs, "%"+req.SpotName+"%")
	}
	whereSlice := append([]interface{}{whereStr}, whereArgs...)
	count, list, err := l.svcCtx.CultrueDbStruct.SyScenicSpotDao.GetList(whereSlice, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询景点列表失败"
		return returnData, err
	}
	data := scenicSpotResp.GetScenicSpotList{
		Page: req.Page, PageSize: req.PageSize, Count: count,
		List: []scenicSpotResp.ScenicSpotInfo{},
	}
	for _, s := range list {
		data.List = append(data.List, scenicSpotResp.ScenicSpotInfo{
			Id:          s.Id,
			SpotName:    s.SpotName,
			Longitude:   s.Longitude,
			Latitude:    s.Latitude,
			TicketPrice: s.TicketPrice,
			Description: s.Description,
			AddTime:     helper.TimeEnumFuncObject.StringTime(s.AddTime),
			Status:      s.Status,
		})
	}
	returnData.Data = data
	return returnData, nil
}

func (l *ScenicSpotLogic) AddLogic(req scenicSpotReq.ScenicSpotAddReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	insertMap := map[string]interface{}{
		"spot_name":    req.SpotName,
		"longitude":    req.Longitude,
		"latitude":     req.Latitude,
		"ticket_price": req.TicketPrice,
		"description":  req.Description,
		"add_uid":      uid,
		"update_uid":   uid,
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100001
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SyScenicSpotDao.Insert([]map[string]interface{}{insertMap}, session); err != nil {
		session.Rollback()
		returnData.Code = 100002
		returnData.Message = "新增景点失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

func (l *ScenicSpotLogic) UpdateLogic(req scenicSpotReq.ScenicSpotUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyScenicSpotDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询景点失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "景点不存在"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"spot_name":    req.SpotName,
		"longitude":    req.Longitude,
		"latitude":     req.Latitude,
		"ticket_price": req.TicketPrice,
		"description":  req.Description,
		"update_uid":   uid,
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
	if err = l.svcCtx.CultrueDbStruct.SyScenicSpotDao.Update(where, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改景点失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

func (l *ScenicSpotLogic) DeleteLogic(req scenicSpotReq.ScenicSpotDeleteReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyScenicSpotDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询景点失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "景点不存在"
		return returnData, nil
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
	if err = l.svcCtx.CultrueDbStruct.SyScenicSpotDao.Update(
		where, map[string]interface{}{"is_delete": 1, "update_uid": uid}, session,
	); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "删除景点失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 上下架景点
func (l *ScenicSpotLogic) UpdateStatusLogic(req scenicSpotReq.ScenicSpotUpdateStatusReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyScenicSpotDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询景点失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "景点不存在"
		return returnData, nil
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
	if err = l.svcCtx.CultrueDbStruct.SyScenicSpotDao.Update(
		where, map[string]interface{}{"status": req.Status, "update_uid": uid}, session,
	); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改景点状态失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}
