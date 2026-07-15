package scenic

import (
	reponse "api/comment/response"
	scenicLogic "api/internal/logic/admin/scenic"
	"api/internal/svc"
	"api/reqs/scenicSpotReq"
	"api/resp"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func parseErr(w http.ResponseWriter, r *http.Request, err error) {
	reponse.NewResponse(w, &resp.CommonReply{Code: 400, Message: "请求参数错误！", Data: []interface{}{}}, r, err)
}

func uidFromHeader(r *http.Request) int {
	v, _ := strconv.Atoi(r.Header.Get("uid"))
	return v
}

func ScenicSpotListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req scenicSpotReq.ScenicSpotListReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := scenicLogic.NewScenicSpotLogic(r.Context(), svcCtx)
		resp, err := l.ListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

func ScenicSpotAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req scenicSpotReq.ScenicSpotAddReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := scenicLogic.NewScenicSpotLogic(r.Context(), svcCtx)
		resp, err := l.AddLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

func ScenicSpotUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req scenicSpotReq.ScenicSpotUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := scenicLogic.NewScenicSpotLogic(r.Context(), svcCtx)
		resp, err := l.UpdateLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

func ScenicSpotDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req scenicSpotReq.ScenicSpotDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := scenicLogic.NewScenicSpotLogic(r.Context(), svcCtx)
		resp, err := l.DeleteLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

func ScenicSpotUpdateStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req scenicSpotReq.ScenicSpotUpdateStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := scenicLogic.NewScenicSpotLogic(r.Context(), svcCtx)
		resp, err := l.UpdateStatusLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}
