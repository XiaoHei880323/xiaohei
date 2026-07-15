package svc

import (
	"api/comment/dblink"
	"api/dblinkMysql"
	"api/internal/config"
	"api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config           config.Config
	Engine           *xorm.Engine
	CultrueDbStruct  *dblinkMysql.CultrueDbStruct
	UserMiddle       rest.Middleware // 管理员 token 中间件
	UserTokenMiddle  rest.Middleware // 用户端 token 中间件
}

func NewServiceContext(c config.Config) *ServiceContext {
	engine := dblink.Database(c.Mysql.DataSource)
	engine.ShowSQL(true)
	m := middleware.NewUserMiddleware()
	return &ServiceContext{
		Config:          c,
		Engine:          engine,
		CultrueDbStruct: dblinkMysql.NewCultrueDbStruct(engine),
		UserMiddle:      m.AdminToken,
		UserTokenMiddle: m.UserToken,
	}
}
