package serv

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

type TokenService struct {
}

func NewTokenService() TokenService {
	return TokenService{}
}

func (s TokenService) GenerateToken(secretKey string, claims jwt.Claims) (string, error) {
	// 创建token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 生成签名的token字符串
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s TokenService) ParseToken(tokenString string, secretKey string) (map[string]interface{}, error) {
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
		return nil, fmt.Errorf("invalid token") // 如果token无效，返回错误信息
	}
}
func main() {
	//	secretKey := "your_secret_key" // 你的密钥应该足够复杂且保密
	//	token, err := GenerateToken(secretKey)
	//	if err != nil {
	//		fmt.Println("Error generating token:", err)
	//		return
	//	}
	//	fmt.Println("Generated Token:", token)
}
