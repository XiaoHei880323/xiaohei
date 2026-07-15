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

	// 课程管理
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "list", Handler: courseHandler.CourseMainListHandler(serverCtx)},
				{Method: http.MethodPost, Path: "detail", Handler: courseHandler.CourseMainDetailHandler(serverCtx)},
				{Method: http.MethodPost, Path: "add", Handler: courseHandler.CourseMainAddHandler(serverCtx)},
				{Method: http.MethodPost, Path: "update", Handler: courseHandler.CourseMainUpdateHandler(serverCtx)},
				{Method: http.MethodPost, Path: "delete", Handler: courseHandler.CourseMainDeleteHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/course/manage/main/"),
	)

	// 课程评价管理
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "list", Handler: courseHandler.CourseEvaluationListHandler(serverCtx)},
				{Method: http.MethodPost, Path: "detail", Handler: courseHandler.CourseEvaluationDetailHandler(serverCtx)},
				{Method: http.MethodPost, Path: "add", Handler: courseHandler.CourseEvaluationAddHandler(serverCtx)},
				{Method: http.MethodPost, Path: "update", Handler: courseHandler.CourseEvaluationUpdateHandler(serverCtx)},
				{Method: http.MethodPost, Path: "delete", Handler: courseHandler.CourseEvaluationDeleteHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/course/manage/evaluation/"),
	)

	// 课程媒体资源管理
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "list", Handler: courseHandler.CourseMediaListHandler(serverCtx)},
				{Method: http.MethodPost, Path: "detail", Handler: courseHandler.CourseMediaDetailHandler(serverCtx)},
				{Method: http.MethodPost, Path: "add", Handler: courseHandler.CourseMediaAddHandler(serverCtx)},
				{Method: http.MethodPost, Path: "update", Handler: courseHandler.CourseMediaUpdateHandler(serverCtx)},
				{Method: http.MethodPost, Path: "delete", Handler: courseHandler.CourseMediaDeleteHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/course/manage/media/"),
	)

	// 课程错题集管理
	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.UserMiddle},
			[]rest.Route{
				{Method: http.MethodPost, Path: "list", Handler: courseHandler.CourseErrorCollectionListHandler(serverCtx)},
				{Method: http.MethodPost, Path: "detail", Handler: courseHandler.CourseErrorCollectionDetailHandler(serverCtx)},
				{Method: http.MethodPost, Path: "add", Handler: courseHandler.CourseErrorCollectionAddHandler(serverCtx)},
				{Method: http.MethodPost, Path: "update", Handler: courseHandler.CourseErrorCollectionUpdateHandler(serverCtx)},
				{Method: http.MethodPost, Path: "delete", Handler: courseHandler.CourseErrorCollectionDeleteHandler(serverCtx)},
			}...,
		),
		rest.WithPrefix("/course/manage/error-collection/"),
	)
}
