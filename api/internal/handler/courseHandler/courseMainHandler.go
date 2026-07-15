package courseHandler

import (
	reponse "api/comment/response"
	courseLogic "api/internal/logic/admin/course"
	"api/internal/svc"
	"api/reqs/courseMainReq"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CourseMainListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMainReq.CourseMainListReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMainLogic(r.Context(), svcCtx).List(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMainDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMainReq.CourseMainDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMainLogic(r.Context(), svcCtx).Detail(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMainAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMainReq.CourseMainAddReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMainLogic(r.Context(), svcCtx).Add(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMainUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMainReq.CourseMainUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMainLogic(r.Context(), svcCtx).Update(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMainDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMainReq.CourseMainDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMainLogic(r.Context(), svcCtx).Delete(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}
