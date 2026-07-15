package courseHandler

import (
	reponse "api/comment/response"
	courseLogic "api/internal/logic/admin/course"
	"api/internal/svc"
	"api/reqs/courseEvaluationReq"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CourseEvaluationListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseEvaluationReq.CourseEvaluationListReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseEvaluationLogic(r.Context(), svcCtx).List(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseEvaluationDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseEvaluationReq.CourseEvaluationDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseEvaluationLogic(r.Context(), svcCtx).Detail(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseEvaluationAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseEvaluationReq.CourseEvaluationAddReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseEvaluationLogic(r.Context(), svcCtx).Add(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseEvaluationUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseEvaluationReq.CourseEvaluationUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseEvaluationLogic(r.Context(), svcCtx).Update(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func CourseEvaluationDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req courseEvaluationReq.CourseEvaluationDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := courseLogic.NewCourseEvaluationLogic(r.Context(), svcCtx).Delete(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}
