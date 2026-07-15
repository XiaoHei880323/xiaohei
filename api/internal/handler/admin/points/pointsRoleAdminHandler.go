package points

import (
	reponse "api/comment/response"
	pointsLogic "api/internal/logic/admin/points"
	"api/internal/svc"
	"api/reqs/pointsRoleReq"
	"api/resp"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"strconv"
)

// 获取积分兑换规则列表
func PointsRoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req pointsRoleReq.PointsRoleListReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := pointsLogic.NewPointsRoleAdminLogic(r.Context(), svcCtx)
		resp, err := l.GetPointsRoleListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 新增积分规则
func PointsRoleAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req pointsRoleReq.PointsRoleAddReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := pointsLogic.NewPointsRoleAdminLogic(r.Context(), svcCtx)
		resp, err := l.PointsRoleAddLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改积分兑换规则
func PointsRoleUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req pointsRoleReq.PointsRoleUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		uid, _ := strconv.Atoi(r.Header.Get("uid"))
		l := pointsLogic.NewPointsRoleAdminLogic(r.Context(), svcCtx)
		resp, err := l.PointsRoleUpdateLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}
