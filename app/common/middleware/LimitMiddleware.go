package middleware

import (
	"SimpleDouYin/app/common"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

const (
	// Seconds 窗口大小
	Seconds = 1
	// Quota 请求上限
	Quota = 100
	// KeyPrefix key前缀
	KeyPrefix = "periodlimit"

	UserApiLoginRegister = "user_api_login_register"
	UserApiGetInfo       = "user_api_getinfo"
)

// OPTION模式设置密码
func op(r *redis.Redis) {
	r.Pass = common.RedisPass
}

func mid(key string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := limit.NewPeriodLimit(Seconds, Quota, redis.New(common.RedisAddr, op), KeyPrefix)
		// 0：表示错误，比如可能是redis故障、过载
		// 1：允许
		// 2：允许但是当前窗口内已到达上限
		// 3：拒绝
		code, err := l.Take(key)
		if err != nil {
			logx.Error(err)
			next(w, r)
		}
		// switch val => process request
		switch code {
		case limit.OverQuota:
			logx.Errorf("OverQuota key: %v", key)
			resp := common.Resp{
				Status: common.ErrRejectedRequest,
				Info:   "请求被拒绝，请稍后再试",
			}
			httpx.OkJson(w, resp)
			next(w, r)
		case limit.Allowed:
			logx.Infof("AllowedQuota key: %v", key)
			next(w, r)
		case limit.HitQuota:
			logx.Errorf("HitQuota key: %v", key)
			resp := common.Resp{
				Status: common.ErrLimitedRequest,
				Info:   "请求被限流，请稍后再试",
			}
			httpx.OkJson(w, resp)
			next(w, r)
		default:
			logx.Errorf("DefaultQuota key: %v", key)
			resp := common.Resp{
				Status: common.ErrOfServer,
				Info:   "服务器错误",
			}
			httpx.OkJson(w, resp)
			next(w, r)
		}
		next(w, r)
	}
}
