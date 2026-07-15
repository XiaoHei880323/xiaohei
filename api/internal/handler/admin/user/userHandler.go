package user

import (
	reponse "api/comment/response"
	userLogic "api/internal/logic/admin/user"
	"api/internal/svc"
	"api/reqs/admin"
	"api/resp"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

// 获取用户的数据信息
func GetUserListHander(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.GetUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := userLogic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.GetUserListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 重置用户密码
func UpdatePwdUserHander(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.UpdateUserPwd
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		uidstring := r.Header.Get("uid")
		uid, _ := strconv.Atoi(uidstring)
		l := userLogic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdatePwdUserLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改用户的积分
func UpdateUserPointsHander(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.UpdateUserPointsReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		uidstring := r.Header.Get("uid")
		uid, _ := strconv.Atoi(uidstring)
		l := userLogic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserPointsLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

func UserPointsListHander(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.GetAdminToUserPointListReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := userLogic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.UserPointsListLogin(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 新增用户
func AddUserHander(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.AddUserReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		uidstring := r.Header.Get("uid")
		uid, _ := strconv.Atoi(uidstring)
		l := userLogic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.AddUserLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改用户基本信息
func UpdateUserInfoHander(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.UpdateUserInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		uidstring := r.Header.Get("uid")
		uid, _ := strconv.Atoi(uidstring)
		l := userLogic.NewUserLogic(r.Context(), svcCtx)
		resp, err := l.UpdateUserInfoLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}
