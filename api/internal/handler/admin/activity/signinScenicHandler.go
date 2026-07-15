package activity

import (
	reponse "api/comment/response"
	activityLogic "api/internal/logic/admin/activity"
	"api/internal/svc"
	"api/reqs/signinScenicReq"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SigninScenicListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req signinScenicReq.SigninScenicListReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSigninScenicLogic(r.Context(), svcCtx)
		resp, err := l.ListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

func SigninScenicAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req signinScenicReq.SigninScenicAddReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSigninScenicLogic(r.Context(), svcCtx)
		resp, err := l.AddLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

func SigninScenicUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req signinScenicReq.SigninScenicUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSigninScenicLogic(r.Context(), svcCtx)
		resp, err := l.UpdateLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}

func SigninScenicDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req signinScenicReq.SigninScenicDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			parseErr(w, r, err)
			return
		}
		l := activityLogic.NewSigninScenicLogic(r.Context(), svcCtx)
		resp, err := l.DeleteLogic(req, uidFromHeader(r))
		reponse.NewResponse(w, resp, r, err)
	}
}
