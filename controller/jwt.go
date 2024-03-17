package controller

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const TokenExpireDuration = time.Hour * 24 //过期时间

var CustomSecret = []byte("夭璟就是最配的！") //签名字符串

type CustomClaim struct {
	Id       int    `json:"id"`
	Username string `json:"username"`

	jwt.RegisteredClaims
}

// 生成token
func GenToken(id int, username string) (string, error) {
	claim := CustomClaim{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "满宝呀",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) //创建签名对象
	return token.SignedString(CustomSecret)
}

// 解析token
func ParseToken(tokenString string) (*CustomClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaim{}, func(token *jwt.Token) (i interface{}, err error) {
		return CustomSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claim, ok := token.Claims.(*CustomClaim); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
