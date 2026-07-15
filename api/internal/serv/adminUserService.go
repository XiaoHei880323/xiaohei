package serv

import (
	"api/internal/svc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AdminUserService struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminUserService(ctx context.Context, svcCtx *svc.ServiceContext) *AdminUserService {
	return &AdminUserService{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 获取用户map【userId】=userName
func (s AdminUserService) GetAdminIdUserMap(userIds []int) (map[int]string, error) {
	userList, userlistErr := s.svcCtx.CultrueDbStruct.SyAdminDao.GetAdminWhereInInt("id", userIds)
	if userlistErr != nil {
		return nil, userlistErr
	}
	userMap := make(map[int]string)
	if len(userList) == 0 {
		return userMap, nil
	}
	for _, userInfo := range userList {
		userMap[userInfo.Id] = userInfo.RelName
	}
	return userMap, nil
}
