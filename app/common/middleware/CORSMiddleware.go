package middleware

import "net/http"

const (
	AllowOrigin      = "Access-Control-Allow-Origin"
	AllowHeaders     = "Access-Control-Allow-Headers"
	AllowMethods     = "Access-Control-Allow-Methods"
	ExposeHeaders    = "Access-Control-Expose-Headers"
	AllowCredentials = "Access-Control-Allow-Credentials"
)

type CORSMiddleware struct{}

func NewCORSMiddleware() *CORSMiddleware {
	return &CORSMiddleware{}
}

func (m *CORSMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		// 放行全部请求
		w.Header().Set(AllowOrigin, origin)
		w.Header().Set(AllowHeaders, "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-Sign-Id")
		w.Header().Set(AllowMethods, "POST, GET, PATCH, OPTIONS, DELETE, PUT")
		w.Header().Set(ExposeHeaders, "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		w.Header().Set(AllowCredentials, "true")

		next(w, r)
	}
}
