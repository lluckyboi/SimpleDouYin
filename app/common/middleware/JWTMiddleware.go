package middleware

import (
	"SimpleDouYin/app/common/jwt"
	"SimpleDouYin/app/common/status"
	"github.com/zeromicro/go-zero/rest/httpx"
	"log"
	"net/http"
	"strings"
)

type JWTMiddleware struct {
	JWTMap *jwt.JWTMap
}

func NewJWTMiddleware(j *jwt.JWTMap) *JWTMiddleware {
	return &JWTMiddleware{
		JWTMap: j,
	}
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start JWT Auth")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			resp := status.Resp{
				Status: 401,
				Info:   "请求头中auth为空",
			}
			httpx.OkJson(w, resp)
			return
		}
		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			resp := status.Resp{
				Status: 401,
				Info:   "请求头中auth格式有误",
			}
			httpx.OkJson(w, resp)
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			resp := status.Resp{
				Status: 401,
				Info:   "无效的Token",
			}
			httpx.OkJson(w, resp)
			return
		}

		//写入上下文
		userInfo := jwt.TSInfo{
			UserId: mc.UserId,
		}
		m.JWTMap.Set("UserInfo", userInfo)
		log.Println("JWT Auth Success")
		next(w, r)
	}
}
