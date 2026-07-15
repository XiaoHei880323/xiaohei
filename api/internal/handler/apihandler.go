package handler

import (
	"api/reqs/admin"
	"api/resp"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"api/internal/logic"
	"api/internal/svc"

	//"github.com/zeromicro/go-zero/rest/httpx"
	reponse "api/comment/response"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req admin.UserList
		if err := httpx.Parse(r, &req); err != nil {
			returnData := &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, err)
			return
		}
		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(req)
		reponse.NewResponse(w, resp, r, err)
	}
}
