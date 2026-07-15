package middleware

import (
	reponse "api/comment/response"
	"api/resp"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strconv"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

const (
	AdminSecretKey = "adminSecretkeySheYangWenHuaJuHuHaoYaoQiuAdmin"
	UserSecretKey  = "UserSecretkeySheYangWenHuaJuHuHaoYaoQiuUser"
)

func (m *UserMiddleware) AdminToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var xToken string = r.Header.Get("X-Token")
		if xToken == "" {
			returnData := &resp.CommonReply{
				Code:    300,
				Message: "请求参数Token获取错误",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, nil)
			return
		}
		// 解析token字符串

		signedToken, tokenErr := ParseToken(xToken, AdminSecretKey)
		if tokenErr != nil {
			returnData := &resp.CommonReply{
				Code:    301,
				Message: "请求参数解析错误",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, tokenErr)
			return
		}

		r.Header.Add("uid", strconv.Itoa(int(signedToken["userId"].(float64))))
		next(w, r)
	}
}

func (m *UserMiddleware) UserToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var xToken string = r.Header.Get("X-Token")
		if xToken == "" {
			returnData := &resp.CommonReply{
				Code:    300,
				Message: "请求参数Token获取错误",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, nil)
			return
		}
		signedToken, tokenErr := ParseToken(xToken, UserSecretKey)
		if tokenErr != nil {
			returnData := &resp.CommonReply{
				Code:    301,
				Message: "请求参数解析错误",
				Data:    []interface{}{},
			}
			reponse.NewResponse(w, returnData, r, tokenErr)
			return
		}
		r.Header.Add("uid", strconv.Itoa(int(signedToken["userId"].(float64))))
		next(w, r)
	}
}

func ParseToken(tokenString string, secretKey string) (map[string]interface{}, error) {
	// 解析token字符串
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证token的签名方法是否是我们期望的算法（HS256）
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥用于签名验证
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return map[string]interface{}(claims), nil // 将claims转换为map[string]interface{}类型返回
	} else {
		return nil, fmt.Errorf("确保用户在有效登入状态中") // 如果token无效，返回错误信息
	}
}
