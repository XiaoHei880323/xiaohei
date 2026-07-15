package courseHandler

import (
	reponse "api/comment/response"
	teacherLogic "api/internal/logic/admin/course"
	"api/internal/svc"
	"api/reqs/teacherReq"
	"api/resp"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func teacherParseError(w http.ResponseWriter, r *http.Request, err error) {
	reponse.NewResponse(w, &resp.CommonReply{
		Code: 400, Message: "请求参数错误！", Data: []interface{}{},
	}, r, err)
}

func TeacherListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req teacherReq.TeacherListReq
		if err := httpx.Parse(r, &req); err != nil {
			teacherParseError(w, r, err)
			return
		}
		result, err := teacherLogic.NewTeacherLogic(r.Context(), svcCtx).List(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func TeacherDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req teacherReq.TeacherDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			teacherParseError(w, r, err)
			return
		}
		result, err := teacherLogic.NewTeacherLogic(r.Context(), svcCtx).Detail(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func TeacherAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req teacherReq.TeacherAddReq
		if err := httpx.Parse(r, &req); err != nil {
			teacherParseError(w, r, err)
			return
		}
		result, err := teacherLogic.NewTeacherLogic(r.Context(), svcCtx).Add(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func TeacherUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req teacherReq.TeacherUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			teacherParseError(w, r, err)
			return
		}
		result, err := teacherLogic.NewTeacherLogic(r.Context(), svcCtx).Update(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func TeacherDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req teacherReq.TeacherDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			teacherParseError(w, r, err)
			return
		}
		result, err := teacherLogic.NewTeacherLogic(r.Context(), svcCtx).Delete(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}
