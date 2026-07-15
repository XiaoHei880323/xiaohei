package home

import (
	reponse "api/comment/response"
	homeLogic "api/internal/logic/admin/home"
	"api/internal/svc"
	"api/reqs/noticeReq"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 公告列表
func NoticeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req noticeReq.NoticeListReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewNoticeLogic(r.Context(), svcCtx)
		resp, err := l.ListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 添加公告
func NoticeAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req noticeReq.NoticeAddReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewNoticeLogic(r.Context(), svcCtx)
		resp, err := l.AddLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改公告
func NoticeUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req noticeReq.NoticeUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewNoticeLogic(r.Context(), svcCtx)
		resp, err := l.UpdateLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 删除公告
func NoticeDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req noticeReq.NoticeDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewNoticeLogic(r.Context(), svcCtx)
		resp, err := l.DeleteLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 发布/下线公告
func NoticeStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req noticeReq.NoticeStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := homeLogic.NewNoticeLogic(r.Context(), svcCtx)
		resp, err := l.StatusLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}
