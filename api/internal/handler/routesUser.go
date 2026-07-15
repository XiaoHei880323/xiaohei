package handler

import (
	"api/internal/handler/user/auth"
	"api/internal/handler/user/home"
	"api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterUserHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 用户登入（无需 token）
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "login",
				Handler: auth.UserLoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/"),
	)

	server.AddRoutes(
		[]rest.Route{
			// 4-1: 首页配置数据（banner 等）
			{
				Method:  http.MethodPost,
				Path:    "config",
				Handler: home.HomeConfigHandler(serverCtx),
			},
			// 4-2: 首页数据（活动/商品/景点，倒序 + 热销标签）
			{
				Method:  http.MethodPost,
				Path:    "data",
				Handler: home.HomeDataHandler(serverCtx),
			},
			// 4-3: 有效公告列表（按公告时间倒序）
			{
				Method:  http.MethodPost,
				Path:    "notice",
				Handler: home.NoticeListHandler(serverCtx),
			},
			// 4-4: 公告详情
			{
				Method:  http.MethodPost,
				Path:    "noticeDetail",
				Handler: home.NoticeDetailHandler(serverCtx),
			},
			// 4-5: 首页活动配置（最多5条，优先当前有效，否则取默认）
			{
				Method:  http.MethodPost,
				Path:    "activityConfig",
				Handler: home.UserActivityConfigHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/home/"),
	)
}
