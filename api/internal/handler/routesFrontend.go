package handler

import (
	"api/internal/handler/viewHandler"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

// RegisterFrontendHandlers 注册前端展示页面路由（与后台 API 路由分离）
// 顶层页面:  GET /admin/page/:name         ->  view/admin/{name}.html
// 模块页面:  GET /admin/page/:module/:name  ->  view/admin/{module}/{name}.html
// 静态资源:  GET /admin/static/:name        ->  view/admin/{name}（JS/CSS）
func RegisterFrontendHandlers(server *rest.Server) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/admin/page/:name",
				Handler: viewHandler.AdminViewFileHandler(),
			},
			{
				Method:  http.MethodGet,
				Path:    "/admin/page/:module/:name",
				Handler: viewHandler.AdminModuleViewFileHandler(),
			},
			{
				Method:  http.MethodGet,
				Path:    "/admin/static/:name",
				Handler: viewHandler.AdminStaticFileHandler(),
			},
		},
	)
}
