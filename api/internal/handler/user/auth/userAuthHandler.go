package auth

import (
	reponse "api/comment/response"
	authLogic "api/internal/logic/user/auth"
	"api/internal/svc"
	"api/reqs/userReq"
	"api/resp"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 用户登入
func UserLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req userReq.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			reponse.NewResponse(w, &resp.CommonReply{
				Code:    400,
				Message: "请求参数错误！",
				Data:    []interface{}{},
			}, r, err)
			return
		}
		l := authLogic.NewUserAuthLogic(r.Context(), svcCtx)
		resp, err := l.UserLogin(req)
		reponse.NewResponse(w, resp, r, err)
	}
}
