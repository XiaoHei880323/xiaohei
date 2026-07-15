package home

import (
	reponse "api/comment/response"
	homeLogic "api/internal/logic/user/home"
	"api/internal/svc"
	"api/reqs/userReq"
	"api/resp"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func parseErr(w http.ResponseWriter, r *http.Request, err error) {
	reponse.NewResponse(w, &resp.CommonReply{Code: 400, Message: "请求参数错误！", Data: []interface{}{}}, r, err)
}

// 获取首页配置数据（banner 轮播等）
func HomeConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userReq.HomeConfigReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewUserHomeLogic(r.Context(), svcCtx)
		resp, err := l.GetHomeConfig(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 获取首页数据（活动/商品/景点，倒序 + 热销标签）
func HomeDataHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userReq.HomeDataReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewUserHomeLogic(r.Context(), svcCtx)
		resp, err := l.GetHomeData(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 获取首页活动配置
func UserActivityConfigHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userReq.UserActivityConfigReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewUserHomeLogic(r.Context(), svcCtx)
		resp, err := l.GetActivityConfig(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 获取有效公告列表（按公告时间倒序）
func NoticeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userReq.NoticeListReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewUserHomeLogic(r.Context(), svcCtx)
		resp, err := l.GetNoticeList(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 获取公告详情
func NoticeDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userReq.NoticeDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := homeLogic.NewUserHomeLogic(r.Context(), svcCtx)
		resp, err := l.GetNoticeDetail(req)
		reponse.NewResponse(w, resp, r, err)
	}
}
