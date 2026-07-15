package serv

import (
	helper "api/comment/help"
	"api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUserRoleService struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserRoleService(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserRoleService {
	return &AdminUserRoleService{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

/*
判断用户角色 1:超级管理员
*/
func (s AdminUserRoleService) AmdinRole(uid int) int {
	role := 0
	superUserId := []int{1}
	if helper.ArrayHelpFuncObject.InArray(uid, superUserId) {
		role = 1
		return role
	}

	return role
}
