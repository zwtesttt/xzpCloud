package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/zwtesttt/xzpCloud/pkg/api"
	"time"
)

const (
	SecretKey = "xzpCloud"
	Expire    = 24 * time.Hour
)

// 生成JWT-token
func GenerateJWT(claims jwt.MapClaims) (string, error) {
	claims["exp"] = time.Now().Add(Expire).Unix()
	claims["iat"] = time.Now().Unix()

	// 创建一个新的 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名生成 token 字符串
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", api.FindCodeError(api.InvalidToken)
	}

	return tokenString, nil
}

// 验证JWT-token
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	// 解析token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, api.FindCodeError(api.InvalidToken)
	}

	// 验证token是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, api.FindCodeError(api.InvalidToken)
}
