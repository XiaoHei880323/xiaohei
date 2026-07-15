package home

import (
	reponse "api/comment/response"
	homeLogic "api/internal/logic/admin/home"
	"api/internal/svc"
	"api/reqs/activityConfigReq"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 活动配置列表
func ActivityConfigListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigListReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.ListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 添加活动配置
func ActivityConfigAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigAddReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.AddLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改活动配置
func ActivityConfigUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 删除活动配置
func ActivityConfigDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.DeleteLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 设置为默认活动配置
func ActivityConfigSetDefaultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigSetDefaultReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.SetDefaultLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 上下线活动配置
func ActivityConfigUpdateStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigUpdateStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStatusLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 活动配置项列表
func ActivityConfigItemListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigItemListReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.ItemListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 添加活动配置项
func ActivityConfigItemAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigItemAddReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.ItemAddLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改活动配置项
func ActivityConfigItemUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigItemUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.ItemUpdateLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 删除活动配置项
func ActivityConfigItemDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityConfigReq.ActivityConfigItemDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewActivityConfigLogic(r.Context(), svcCtx)
		resp, err := l.ItemDeleteLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}
