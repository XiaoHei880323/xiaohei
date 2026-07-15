package goods

import (
	reponse "api/comment/response"
	"api/internal/svc"
	"api/reqs/goodReq"
	"api/resp"
	"api/resp/goodResp"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GoodAdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGoodAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GoodAdminLogic {
	return &GoodAdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 添加商品
func (l GoodAdminLogic) GoodAddLogic(req goodReq.GoodAddReq, userId int) (resp *resp.CommonReply, err error) {
	returnData := reponse.ReturnStruct()
	insertMap := make(map[string]interface{})
	insertMap["good_name"] = req.GoodName
	insertMap["good_img"] = req.GoodImage
	insertMap["good_price"] = req.GoodPrice
	insertMap["good_desc"] = req.GoodDesc
	insertMap["status"] = 1
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
	insertErr := l.svcCtx.CultrueDbStruct.SyGoodDao.Insert(addArray, session)
	if insertErr != nil {
		session.Rollback()
		returnData.Code = 100001
		returnData.Message = "商品添加失败"
		return returnData, insertErr
	}
	session.Commit()
	return returnData, nil
}

// 获取活动商品信息
func (l GoodAdminLogic) GetGoodListLogic(req goodReq.GoodListReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	sql := "is_delete = ?"
	whereSlice := []interface{}{0}
	if req.GoodName != "" {
		sql = sql + " AND good_name LIKE \"%?%\""
		whereSlice = append(whereSlice, req.GoodName)
	}
	where := []interface{}{}
	where = append(where, sql)
	where = append(where, whereSlice...)
	goodListCount, goodList, goodListErr := l.svcCtx.CultrueDbStruct.SyGoodDao.GetGoodList(where, req.Page, req.PageSize, "id desc")
	if goodListErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询商品信息错误"
		return returnData, goodListErr
	}
	data := goodResp.GetGoodListResp{
		Count:    goodListCount,
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	if goodListCount == 0 {
		returnData.Data = data
		return returnData, nil
	}
	for _, goodInfo := range goodList {
		info := goodResp.GetGoodInfoResp{
			GoodId:    goodInfo.Id,
			GoodName:  goodInfo.GoodName,
			GoodImg:   goodInfo.GoodImg,
			GoodPrice: goodInfo.GoodPrice,
			GoodDesc:  goodInfo.GoodDesc,
			Status:    goodInfo.Status,
		}
		data.GoodList = append(data.GoodList, info)
	}
	returnData.Data = data
	return returnData, nil
}

// 修改商品
func (l GoodAdminLogic) GoodUpdateLogic(req goodReq.GoodUpdateReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereGood := []interface{}{"id = ?", req.GoodId}
	goodInfohas, _, goodInfoErr := l.svcCtx.CultrueDbStruct.SyGoodDao.GetInfo(whereGood)
	if goodInfoErr != nil {
		returnData.Code = 100001
		returnData.Message = "查询商品详情错误"
		return returnData, goodInfoErr
	}
	if !goodInfohas {
		returnData.Code = 100002
		returnData.Message = "未查询到商品信息"
		return returnData, nil
	}
	updateGood := make(map[string]interface{})
	updateGood["good_name"] = req.GoodName
	updateGood["good_img"] = req.GoodImage
	updateGood["good_price"] = req.GoodPrice
	updateGood["good_desc"] = req.GoodDesc
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
	updateErr := l.svcCtx.CultrueDbStruct.SyGoodDao.Update(whereGood, updateGood, session)
	if updateErr != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改商品信息错误"
		return returnData, updateErr
	}
	session.Commit()
	return returnData, nil
}

// 删除商品（软删除）
func (l GoodAdminLogic) GoodDeleteLogic(req goodReq.GoodDeleteReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.GoodId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyGoodDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询商品错误"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "商品不存在"
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
	if err = l.svcCtx.CultrueDbStruct.SyGoodDao.Update(where, map[string]interface{}{"is_delete": 1}, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "删除商品失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}

// 上下架商品
func (l GoodAdminLogic) GoodUpdateStatusLogic(req goodReq.GoodUpdateStatusReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	where := []interface{}{"id = ? AND is_delete = ?", req.GoodId, 0}
	has, _, err := l.svcCtx.CultrueDbStruct.SyGoodDao.GetInfo(where)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询商品错误"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "商品不存在"
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
	if err = l.svcCtx.CultrueDbStruct.SyGoodDao.Update(where, map[string]interface{}{"status": req.Status}, session); err != nil {
		session.Rollback()
		returnData.Code = 100004
		returnData.Message = "修改商品状态失败"
		return returnData, err
	}
	session.Commit()
	return returnData, nil
}
