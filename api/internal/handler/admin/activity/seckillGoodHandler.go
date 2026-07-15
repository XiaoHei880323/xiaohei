package activity

import (
	reponse "api/comment/response"
	activityLogic "api/internal/logic/admin/activity"
	"api/internal/svc"
	"api/reqs/seckillGoodReq"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 列表
func SeckillGoodListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req seckillGoodReq.SeckillGoodListReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSeckillGoodLogic(r.Context(), svcCtx)
		resp, err := l.ListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 新增
func SeckillGoodAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req seckillGoodReq.SeckillGoodAddReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSeckillGoodLogic(r.Context(), svcCtx)
		resp, err := l.AddLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改秒杀价
func SeckillGoodUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req seckillGoodReq.SeckillGoodUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSeckillGoodLogic(r.Context(), svcCtx)
		resp, err := l.UpdateLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

// 删除
func SeckillGoodDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req seckillGoodReq.SeckillGoodDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSeckillGoodLogic(r.Context(), svcCtx)
		resp, err := l.DeleteLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

// 批量修改秒杀价
func SeckillGoodBatchUpdatePriceHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req seckillGoodReq.SeckillGoodBatchUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSeckillGoodLogic(r.Context(), svcCtx)
		resp, err := l.BatchUpdatePriceLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

// 批量删除
func SeckillGoodBatchDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req seckillGoodReq.SeckillGoodBatchDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSeckillGoodLogic(r.Context(), svcCtx)
		resp, err := l.BatchDeleteLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}
