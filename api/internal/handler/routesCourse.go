package handler

import (
	"api/internal/handler/courseHandler"
	"api/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func CourseHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	// 学生管理
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "list", Handler: courseHandler.StudentListHandler(serverCtx)},
				{Method: http.MethodPost, Path: "detail", Handler: courseHandler.StudentDetailHandler(serverCtx)},
				{Method: http.MethodPost, Path: "add", Handler: courseHandler.StudentAddHandler(serverCtx)},
				{Method: http.MethodPost, Path: "update", Handler: courseHandler.StudentUpdateHandler(serverCtx)},
				{Method: http.MethodPost, Path: "delete", Handler: courseHandler.StudentDeleteHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/course/personnel/student/"),
	)

	// 教师管理
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "list", Handler: courseHandler.TeacherListHandler(serverCtx)},
				{Method: http.MethodPost, Path: "detail", Handler: courseHandler.TeacherDetailHandler(serverCtx)},
				{Method: http.MethodPost, Path: "add", Handler: courseHandler.TeacherAddHandler(serverCtx)},
				{Method: http.MethodPost, Path: "update", Handler: courseHandler.TeacherUpdateHandler(serverCtx)},
				{Method: http.MethodPost, Path: "delete", Handler: courseHandler.TeacherDeleteHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/course/personnel/teacher/"),
	)
}
