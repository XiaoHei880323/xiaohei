package courseHandler

import (
	reponse "api/comment/response"
	studentLogic "api/internal/logic/admin/course"
	"api/internal/svc"
	"api/reqs/studentReq"
	"api/resp"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func studentParseError(w http.ResponseWriter, r *http.Request, err error) {
	reponse.NewResponse(w, &resp.CommonReply{
		Code: 400, Message: "请求参数错误！", Data: []interface{}{},
	}, r, err)
}

func courseOperatorID(r *http.Request) int {
	uid, _ := strconv.Atoi(r.Header.Get("uid"))
	return uid
}

func StudentListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req studentReq.StudentListReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := studentLogic.NewStudentLogic(r.Context(), svcCtx).List(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func StudentDetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req studentReq.StudentDetailReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := studentLogic.NewStudentLogic(r.Context(), svcCtx).Detail(req)
		reponse.NewResponse(w, result, r, err)
	}
}

func StudentAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req studentReq.StudentAddReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := studentLogic.NewStudentLogic(r.Context(), svcCtx).Add(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func StudentUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req studentReq.StudentUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := studentLogic.NewStudentLogic(r.Context(), svcCtx).Update(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}

func StudentDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req studentReq.StudentDeleteReq
		if err := httpx.Parse(r, &req); err != nil {
			studentParseError(w, r, err)
			return
		}
		result, err := studentLogic.NewStudentLogic(r.Context(), svcCtx).Delete(req, courseOperatorID(r))
		reponse.NewResponse(w, result, r, err)
	}
}
