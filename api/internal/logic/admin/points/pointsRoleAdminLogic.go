package points

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/svc"
	"api/reqs/pointsRoleReq"
	"api/resp"
	"api/resp/pointsRoleResp"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type PointsRoleAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPointsRoleAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PointsRoleAdminLogic {
	return &PointsRoleAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取时间
func (l PointsRoleAdminLogic) GetPointsRoleListLogic(req pointsRoleReq.PointsRoleListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	sql := "p.is_deleted = ?"
	whereSlice := []interface{}{0}
	if req.PointsStart > 0 {
		sql += " and p.exchange_points >= ?"
		whereSlice = append(whereSlice, req.PointsStart)
	}
	if req.PointsEnd > 0 {
		sql += " and p.exchange_points <= ?"
		whereSlice = append(whereSlice, req.PointsEnd)
	}
	if req.PointsRoleStartTime != "" {
		sql += " and p.exchange_start_time >=  ?"
		whereSlice = append(whereSlice, req.PointsRoleStartTime)
	}
	if req.PointsRoleEndTime != "" {
		sql += " and p.exchange_end_time <=  ?"
		whereSlice = append(whereSlice, req.PointsRoleEndTime)
	}
	if req.GoodName != "" {
		sql += " and g.good_name like \"?\""
		whereSlice = append(whereSlice, req.GoodName)
	}
	where := []interface{}{}
	where = append(where, sql)
	where = append(where, whereSlice...)
	pointsRoleCount, pointsRoleList, pointsRoleErr := l.svcCtx.CultrueDbStruct.SyPointsRoleDao.GetPointsRoleAndGood(where, req.Page, req.PageSize, "p.id desc")
	if pointsRoleErr != nil {
		returnData.Code = 10001
		returnData.Message = "获取积分兑换数据错误"
		return returnData, pointsRoleErr
	}
	data := pointsRoleResp.GetPointsRoleListResp{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    pointsRoleCount,
	}
	if pointsRoleCount == 0 {
		returnData.Data = data
		return returnData, nil
	}
	for _, pointsRoleInfo := range pointsRoleList {
		info := pointsRoleResp.GetPointsRoleInfoResp{
			PointId:        pointsRoleInfo.SyPointsRole.Id,
			Point:          pointsRoleInfo.SyPointsRole.ExchangePoints,
			GoodId:         pointsRoleInfo.SyGood.Id,
			GoodName:       pointsRoleInfo.SyGood.GoodName,
			GoodImage:      pointsRoleInfo.SyGood.GoodImg,
			PointStartTime: helper.TimeEnumFuncObject.StringTime(pointsRoleInfo.ExchangeStartTime),
			PointEndTime:   helper.TimeEnumFuncObject.StringTime(pointsRoleInfo.ExchangeEndTime),
			CreateTime:     helper.TimeEnumFuncObject.StringTime(pointsRoleInfo.SyPointsRole.CreateTime),
		}
		data.PointsRoleList = append(data.PointsRoleList, info)
	}
	returnData.Data = data
	return returnData, nil
}

// 新增积分兑换活动规则
func (l PointsRoleAdminLogic) PointsRoleAddLogic(req pointsRoleReq.PointsRoleAddReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	insert := make(map[string]interface{})
	insert["exchange_points"] = req.Points
	insert["good_id"] = req.GoodId
	insert["exchange_start_time"] = req.PointsStartTime
	insert["exchange_end_time"] = req.PointsEndTime
	insert["add_uid"] = uid
	insert["update_uid"] = uid
	add := make([]map[string]interface{}, 0)
	add = append(add, insert)
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

	insertErr := l.svcCtx.CultrueDbStruct.SyPointsRoleDao.Insert(add, session)
	if insertErr != nil {
		session.Rollback()
		returnData.Code = 100002
		returnData.Message = "添加积分兑换规则错误"
		return returnData, insertErr
	}
	session.Commit()
	return returnData, nil
}

// 修改积分兑换规则
func (l PointsRoleAdminLogic) PointsRoleUpdateLogic(req pointsRoleReq.PointsRoleUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"is_deleted = ? and id = ?", 0, req.PointId}
	pointsRoleHas, _, pointsRoleErr := l.svcCtx.CultrueDbStruct.SyPointsRoleDao.GetInfo(where)
	if pointsRoleErr != nil {
		returnData.Code = 100001
		returnData.Message = "获取积分兑换规则失败"
		return returnData, pointsRoleErr
	}
	if !pointsRoleHas {
		returnData.Code = 100002
		returnData.Message = "没有获取到对应的积分规则"
		return returnData, nil
	}
	update := make(map[string]interface{})
	update["exchange_points"] = req.Points
	update["good_id"] = req.GoodId
	update["exchange_start_time"] = req.PointsStartTime
	update["exchange_end_time"] = req.PointsEndTime
	update["update_uid"] = uid
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
	updateErr := l.svcCtx.CultrueDbStruct.SyPointsRoleDao.Update(where, update, session)
	if updateErr != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改积分兑换规则错误"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}
