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
)

type LimitMiddleware struct {
	key string
}

func NewLimitMiddleware(key string) *LimitMiddleware {
	return &LimitMiddleware{key: key}
}

// OPTION模式设置密码
func op(r *redis.Redis) {
	r.Pass = common.RedisPass
}

func (lm *LimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := limit.NewPeriodLimit(Seconds, Quota, redis.New(common.RedisAddr, op), common.LimitKeyPrefix)
		// 0：表示错误，比如可能是redis故障、过载
		// 1：允许
		// 2：允许但是当前窗口内已到达上限
		// 3：拒绝
		code, err := l.Take(lm.key)
		if err != nil {
			logx.Error(err)
			next(w, r)
		}
		// switch val => process request
		switch code {
		case limit.OverQuota:
			logx.Errorf("OverQuota key: %v", lm.key)
			resp := common.Resp{
				Status: common.ErrRejectedRequest,
				Info:   "请求被拒绝，请稍后再试",
			}
			httpx.OkJson(w, resp)
			return
		case limit.Allowed:
			logx.Infof("AllowedQuota key: %v", lm.key)
			return
		case limit.HitQuota:
			logx.Errorf("HitQuota key: %v", lm.key)
			resp := common.Resp{
				Status: common.ErrLimitedRequest,
				Info:   "请求被限流，请稍后再试",
			}
			httpx.OkJson(w, resp)
			return
		default:
			logx.Errorf("DefaultQuota key: %v", lm.key)
			resp := common.Resp{
				Status: common.ErrOfServer,
				Info:   "服务器错误",
			}
			httpx.OkJson(w, resp)
			return
		}
		next(w, r)
	}
}
