package courseHandler

import (
	reponse "api/comment/response"
	courseLogic "api/internal/logic/admin/course"
	"api/internal/svc"
	"api/reqs/courseErrorCollectionReq"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CourseErrorCollectionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseErrorCollectionReq.CourseErrorCollectionListReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseErrorCollectionLogic(r.Context(), svcCtx).List(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseErrorCollectionDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseErrorCollectionReq.CourseErrorCollectionDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseErrorCollectionLogic(r.Context(), svcCtx).Detail(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseErrorCollectionAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseErrorCollectionReq.CourseErrorCollectionAddReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseErrorCollectionLogic(r.Context(), svcCtx).Add(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseErrorCollectionUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseErrorCollectionReq.CourseErrorCollectionUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseErrorCollectionLogic(r.Context(), svcCtx).Update(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseErrorCollectionDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseErrorCollectionReq.CourseErrorCollectionDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseErrorCollectionLogic(r.Context(), svcCtx).Delete(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}
