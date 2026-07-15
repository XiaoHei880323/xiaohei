package activity

import (
	reponse "api/comment/response"
	activityLogic "api/internal/logic/admin/activity"
	"api/internal/svc"
	"api/reqs/activityReq"
	"api/resp"
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 秒杀活动列表
func SeckillActivityListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityReq.ActivityListReq
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := activityLogic.NewSeckillActivityLogic(r.Context(), svcCtx)
		resp, err := l.ActivityListLogic(req)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 添加秒杀活动
func SeckillActivityAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityReq.ActivityAddInfoReq
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
		l := activityLogic.NewSeckillActivityLogic(r.Context(), svcCtx)
		resp, err := l.ActivityAddLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 修改秒杀活动
func SeckillActivityUpdateInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityReq.ActivityUpdateInfoReq
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
		l := activityLogic.NewSeckillActivityLogic(r.Context(), svcCtx)
		resp, err := l.ActivityUpdateLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}

// 删除秒杀活动
func SeckillActivityDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req activityReq.ActivityDeleteReq
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
		l := activityLogic.NewSeckillActivityLogic(r.Context(), svcCtx)
		resp, err := l.ActivityDeleteLogic(req, uid)
		reponse.NewResponse(w, resp, r, err)
	}
}
