package common

import (
	"SimpleDouYin/app/service/user/dao/model"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	UserId int64
	jwt.StandardClaims
}

// TSInfo 存储在JWTMap里供给Logic使用的信息
type TSInfo struct {
	UserId int64
}

const ExpTD = 7

// AccessTokenExpireDuration 定义AccessToken的过期时间
const AccessTokenExpireDuration = time.Hour * 24 * ExpTD

// MySecret 自定义签名字段
var MySecret = []byte("ayanami")

// GenAccessToken 生成AccessToken
func GenAccessToken(user model.User) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		user.UserID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(AccessTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "Lcuky",                                          // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
