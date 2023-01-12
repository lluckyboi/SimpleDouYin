package common

import (
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"go-zero-demo/app/service/user/dao/model"
	"time"
)

type MyClaims struct {
	UserPnum string
	UserRole string
	IsRef    bool //是否refresh_token
	jwt.StandardClaims
}

// TSInfo 存储在JWTMap里供给Logic使用的信息
type TSInfo struct {
	UserPnum string
	UserRole string
	IsRef    bool //是否refresh_token
}

const ExpTD = 7

// AccessTokenExpireDuration 定义AccessToken的过期时间
const AccessTokenExpireDuration = time.Hour * 24 * ExpTD

// RefreshTokenExpireDuration 定义RefreshToken的过期时间
const RefreshTokenExpireDuration = AccessTokenExpireDuration*2 + time.Hour*24

// MySecret 自定义签名字段
var MySecret = []byte("ayanami")

// GenAccessToken 生成AccessToken
func GenAccessToken(user model.User) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		user.UserPnum,
		user.UserRole,
		false,
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

// GenRefreshToken 生成RefreshToken
func GenRefreshToken(user model.User) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		user.UserPnum,
		user.UserRole,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(RefreshTokenExpireDuration).Unix(), // 过期时间
			Issuer:    "Lcuky",                                           // 签发人
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
