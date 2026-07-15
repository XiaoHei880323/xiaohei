package user

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	userLogic "api/internal/logic/admin/user"
	"api/internal/svc"
	"api/reqs/admin"
	"api/resp"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

func AdminLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.UserLogin
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		fmt.Println("req => %v", helper.ConvertStrutToJson(req))
		l := userLogic.NewAdminLogic(r.Context(), svcCtx)
		resp, err := l.AdminUserLogin(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

func UpdateInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.UpdateInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := userLogic.NewAdminLogic(r.Context(), svcCtx)
		resp, err := l.UpdateInfoLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 获取管理员列表
func GetAdminUserListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.GetAdminUserListReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := userLogic.NewAdminLogic(r.Context(), ctx)
		resp, err := l.GetAdminUserListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 添加管理员用户
func AddAdminUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.AddAdminUserReq
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
		l := userLogic.NewAdminLogic(r.Context(), svcCtx)
		resp, err := l.AddAdminUserLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改其他管理员的状态
func UpdateAdminUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.UpdateAdminUserReq
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
		l := userLogic.NewAdminLogic(r.Context(), svcCtx)
		resp, err := l.UpdateAdminUserLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}
