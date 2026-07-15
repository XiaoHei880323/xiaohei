package courseHandler

import (
	reponse "api/comment/response"
	courseLogic "api/internal/logic/admin/course"
	"api/internal/svc"
	"api/reqs/courseMediaReq"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CourseMediaListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMediaReq.CourseMediaListReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMediaLogic(r.Context(), svcCtx).List(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMediaDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMediaReq.CourseMediaDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMediaLogic(r.Context(), svcCtx).Detail(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMediaAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMediaReq.CourseMediaAddReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMediaLogic(r.Context(), svcCtx).Add(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMediaUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMediaReq.CourseMediaUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMediaLogic(r.Context(), svcCtx).Update(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseMediaDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseMediaReq.CourseMediaDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseMediaLogic(r.Context(), svcCtx).Delete(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}
