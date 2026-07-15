package logic

import (
	reponse "api/comment/response"
	"api/internal/svc"
	"api/reqs/admin"
	"api/resp"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApiLogic {
	return &ApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApiLogic) Api(req admin.UserList) (resp *resp.CommonReply, err error) {
	// todo: add your logic here and delete this line
	returnData := reponse.ReturnStruct()
	claims := jwt.MapClaims{
		"userId": 1,
		"exp":    time.Now().Add(time.Hour * 1).Unix(), // token 1小时后过期
		"iat":    time.Now().Unix(),
	}
	token, tokenErr := GenerateToken(l.svcCtx.Config.Auth.AdminSecretKey, claims)
	if tokenErr != nil {
		return nil, tokenErr
	}
	fmt.Println("                                      ")
	fmt.Printf("token=> %v", token)
	fmt.Println("                                      ")
	tokenMap, tokenMapErr := ParseToken(token, l.svcCtx.Config.Auth.AdminSecretKey)
	if tokenMapErr != nil {
		return nil, tokenMapErr
	}
	fmt.Printf("tokenMap => %v", tokenMap)
	return returnData, err
}

func GenerateToken(secretKey string, claims jwt.Claims) (string, error) {
	// 创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名的token字符串
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string, secretKey string) (map[string]interface{}, error) {
	fmt.Println("tokenString => ", tokenString)
	// 解析token字符串
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证token的签名方法是否是我们期望的算法（HS256）
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥用于签名验证
		return []byte(secretKey), nil
	})
	fmt.Printf("token.claims=> %v", token.Claims.(jwt.MapClaims))
	fmt.Println()
	fmt.Printf("token.Valid=> %v", token.Valid)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]interface{}(claims), nil // 将claims转换为map[string]interface{}类型返回
	} else {
		return nil, fmt.Errorf("invalid token") // 如果token无效，返回错误信息
	}
}
