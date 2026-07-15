package activity

import (
	reponse "api/comment/response"
	helper "api/comment/help"
	"api/internal/svc"
	"api/reqs/seckillGoodReq"
	"api/resp"
	"api/resp/seckillGoodResp"
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillGoodLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillGoodLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillGoodLogic {
	return &SeckillGoodLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 列表
func (l *SeckillGoodLogic) ListLogic(req seckillGoodReq.SeckillGoodListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	if req.ActivityId == 0 {
		returnData.Code = 100001
		returnData.Message = "活动ID不能为空"
		return returnData, nil
	}

	// 若按商品名过滤，先从商品表获取符合条件的 goodId 列表
	goodIdFilter := make([]int, 0)
	if req.GoodName != "" {
		where := []interface{}{"good_name LIKE ?", "%" + req.GoodName + "%"}
		_, goodList, err := l.svcCtx.CultrueDbStruct.SyGoodDao.GetGoodList(where, 1, 9999, "id asc")
		if err != nil {
			returnData.Code = 100002
			returnData.Message = "查询商品列表失败"
			return returnData, err
		}
		for _, g := range goodList {
			goodIdFilter = append(goodIdFilter, g.Id)
		}
		if len(goodIdFilter) == 0 {
			// 无匹配商品，直接返回空
			returnData.Data = seckillGoodResp.GetSeckillGoodList{
				Page: req.Page, PageSize: req.PageSize, Count: 0, List: []seckillGoodResp.GetSeckillGoodInfo{},
			}
			return returnData, nil
		}
	}

	// 查秒杀活动商品
	whereStr := "activity_id = ? AND is_delete = ?"
	whereArgs := []interface{}{req.ActivityId, 0}
	if len(goodIdFilter) > 0 {
		placeholders := strings.TrimRight(strings.Repeat("?,", len(goodIdFilter)), ",")
		whereStr += fmt.Sprintf(" AND good_id IN (%s)", placeholders)
		for _, id := range goodIdFilter {
			whereArgs = append(whereArgs, id)
		}
	}
	whereSlice := append([]interface{}{whereStr}, whereArgs...)

	count, sagList, err := l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.GetListByWhere(whereSlice, req.Page, req.PageSize)
	if err != nil {
		returnData.Code = 100003
		returnData.Message = "查询秒杀活动商品失败"
		return returnData, err
	}

	data := seckillGoodResp.GetSeckillGoodList{
		Page:     req.Page,
		PageSize: req.PageSize,
		Count:    count,
		List:     []seckillGoodResp.GetSeckillGoodInfo{},
	}
	if count == 0 {
		returnData.Data = data
		return returnData, nil
	}

	// 批量查商品信息
	goodIds := make([]int, 0, len(sagList))
	for _, item := range sagList {
		goodIds = append(goodIds, item.GoodId)
	}
	goodMap := make(map[int]struct {
		Name  string
		Img   string
		Price string
	})
	for _, gid := range goodIds {
		where := []interface{}{"id = ?", gid}
		has, g, gerr := l.svcCtx.CultrueDbStruct.SyGoodDao.GetInfo(where)
		if gerr == nil && has {
			goodMap[gid] = struct {
				Name  string
				Img   string
				Price string
			}{Name: g.GoodName, Img: g.GoodImg, Price: g.GoodPrice}
		}
	}

	for _, item := range sagList {
		info := seckillGoodResp.GetSeckillGoodInfo{
			Id:           item.Id,
			ActivityId:   item.ActivityId,
			GoodId:       item.GoodId,
			SeckillPrice: item.SeckillPrice,
			AddTime:      helper.TimeEnumFuncObject.StringTime(item.AddTime),
		}
		if g, ok := goodMap[item.GoodId]; ok {
			info.GoodName  = g.Name
			info.GoodImg   = g.Img
			info.GoodPrice = g.Price
		}
		data.List = append(data.List, info)
	}
	returnData.Data = data
	return returnData, nil
}

// 新增
func (l *SeckillGoodLogic) AddLogic(req seckillGoodReq.SeckillGoodAddReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	// 防重复：同一活动同一商品只允许存在一条
	where := []interface{}{"activity_id = ? AND good_id = ? AND is_delete = ?", req.ActivityId, req.GoodId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询已有商品失败"
		return returnData, err
	}
	if has {
		returnData.Code = 100002
		returnData.Message = "该商品已在此活动中，请勿重复添加"
		return returnData, nil
	}

	insertMap := map[string]interface{}{
		"activity_id":   req.ActivityId,
		"good_id":       req.GoodId,
		"seckill_price": req.SeckillPrice,
		"add_uid":       uid,
		"update_uid":    uid,
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
	if err = l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.Insert([]map[string]interface{}{insertMap}, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "新增秒杀商品失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 修改秒杀价
func (l *SeckillGoodLogic) UpdateLogic(req seckillGoodReq.SeckillGoodUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询记录失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "记录不存在"
		return returnData, nil
	}
	updateMap := map[string]interface{}{
		"seckill_price": req.SeckillPrice,
		"update_uid":    uid,
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
	if err = l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.Update(where, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 删除（软删除）
func (l *SeckillGoodLogic) DeleteLogic(req seckillGoodReq.SeckillGoodDeleteReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.Id, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询记录失败"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "记录不存在"
		return returnData, nil
	}
	updateMap := map[string]interface{}{"is_delete": 1, "update_uid": uid}
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
	if err = l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.Update(where, updateMap, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "删除失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 批量修改秒杀价
func (l *SeckillGoodLogic) BatchUpdatePriceLogic(req seckillGoodReq.SeckillGoodBatchUpdateReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	if len(req.Ids) == 0 {
		returnData.Code = 100001
		returnData.Message = "请选择要修改的商品"
		return returnData, nil
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100002
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.BatchUpdatePrice(req.Ids, req.SeckillPrice, uid, session); err != nil {
		session.Rollback()
		returnData.Code = 100003
		returnData.Message = "批量修改失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 批量删除
func (l *SeckillGoodLogic) BatchDeleteLogic(req seckillGoodReq.SeckillGoodBatchDeleteReq, uid int) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	if len(req.Ids) == 0 {
		returnData.Code = 100001
		returnData.Message = "请选择要删除的商品"
		return returnData, nil
	}
	session := l.svcCtx.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		returnData.Code = 100002
		returnData.Message = "开启事务失败"
		return returnData, err
	}
	defer func() {
		if r := recover(); r != nil {
			session.Rollback()
		}
	}()
	if err := l.svcCtx.CultrueDbStruct.SySeckillActivityGoodDao.BatchDelete(req.Ids, uid, session); err != nil {
		session.Rollback()
		returnData.Code = 100003
		returnData.Message = "批量删除失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}
