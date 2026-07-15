package auth

import (
	helper "api/comment/help"
	reponse "api/comment/response"
	"api/internal/serv"
	"api/internal/svc"
	"api/reqs/userReq"
	"api/resp"
	"api/resp/userResp"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UserAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAuthLogic {
	return &UserAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 用户登入
func (l *UserAuthLogic) UserLogin(req userReq.UserLoginReq) (*resp.CommonReply, error) {
	returnData := reponse.ReturnStruct()
	whereSlice := []interface{}{
		"status = ? and is_delete = ? and phone = ? and pwd = ?",
		0, 0, req.Phone, helper.Md5HelperObject.Sha256ToString(req.Pwd),
	}
	has, userInfo, err := l.svcCtx.CultrueDbStruct.SyUserDao.GetInfo(whereSlice)
	if err != nil {
		returnData.Code = 100001
		returnData.Message = "查询用户信息错误"
		return returnData, err
	}
	if !has {
		returnData.Code = 100002
		returnData.Message = "手机号或密码不正确"
		return returnData, nil
	}

	claims := jwt.MapClaims{
		"userId": userInfo.Id,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(), // 7天有效期
		"iat":    time.Now().Unix(),
	}
	tokenService := serv.NewTokenService()
	signedToken, tokenErr := tokenService.GenerateToken(l.svcCtx.Config.Auth.UserSecretKey, claims)
	if tokenErr != nil {
		returnData.Code = 100003
		returnData.Message = "生成token错误"
		return returnData, tokenErr
	}

	returnData.Data = userResp.UserLoginResp{
		UserId:   userInfo.Id,
		UserName: userInfo.UserName,
		NickName: userInfo.RelName,
		Phone:    userInfo.Phone,
		HeadImg:  userInfo.HeadImg,
		Points:   userInfo.Points,
		Token:    signedToken,
	}
	return returnData, nil
}
