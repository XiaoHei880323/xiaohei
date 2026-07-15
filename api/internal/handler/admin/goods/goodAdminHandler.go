package goods

import (
	reponse "api/comment/response"
	goodsLogic "api/internal/logic/admin/goods"
	"api/internal/svc"
	"api/reqs/goodReq"
	"api/resp"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

// 添加商品
func GoodAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req goodReq.GoodAddReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		uidString := r.Header.Get("uid")
		uid, _ := strconv.Atoi(uidString)
		l := goodsLogic.NewGoodAdminLogic(r.Context(), svcCtx)
		resp, err := l.GoodAddLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 获取商品列表
func GoodListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req goodReq.GoodListReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := goodsLogic.NewGoodAdminLogic(r.Context(), svcCtx)
		resp, err := l.GetGoodListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改商品
func GoodUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req goodReq.GoodUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := goodsLogic.NewGoodAdminLogic(r.Context(), svcCtx)
		resp, err := l.GoodUpdateLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 删除商品
func GoodDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req goodReq.GoodDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			reponse.NewResponse(w, &resp.CommonReply{Code: 400, Message: "请求参数错误！", Data: []interface{}{}}, r, err)
			return
		}
		l := goodsLogic.NewGoodAdminLogic(r.Context(), svcCtx)
		resp, err := l.GoodDeleteLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 上下架商品
func GoodUpdateStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req goodReq.GoodUpdateStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			reponse.NewResponse(w, &resp.CommonReply{Code: 400, Message: "请求参数错误！", Data: []interface{}{}}, r, err)
			return
		}
		l := goodsLogic.NewGoodAdminLogic(r.Context(), svcCtx)
		resp, err := l.GoodUpdateStatusLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}
