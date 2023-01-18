package middleware

import (
	"SimpleDouYin/app/common"
	"SimpleDouYin/app/common/key"
	"SimpleDouYin/app/common/status"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type LimitMiddleware struct {
	key     string
	Seconds int
	Quota   int
}

func NewLimitMiddleware(key string, seconds, quota int) *LimitMiddleware {
	return &LimitMiddleware{key: key, Seconds: seconds, Quota: quota}
}

// OPTION模式设置密码
func op(r *redis.Redis) {
	r.Pass = common.RedisPass
}

func (lm *LimitMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := limit.NewPeriodLimit(lm.Seconds, lm.Quota, redis.New(common.RedisAddr, op), key.LimitKeyPrefix)
		// 0：表示错误，比如可能是redis故障、过载
		// 1：允许
		// 2：允许但是当前窗口内已到达上限
		// 3：拒绝
		code, err := l.Take(lm.key)
		if err != nil {
			logx.Error(err)
			return
		}
		// switch val => process request
		switch code {
		case limit.OverQuota:
			logx.Errorf("OverQuota key: %v", lm.key)
			resp := status.Resp{
				Status: status.ErrRejectedRequest,
				Info:   "请求被拒绝，请稍后再试",
			}
			httpx.OkJson(w, resp)
			return
		case limit.Allowed:
			logx.Infof("AllowedQuota key: %v", lm.key)
			next(w, r)
		case limit.HitQuota:
			logx.Errorf("HitQuota key: %v", lm.key)
			resp := status.Resp{
				Status: status.ErrLimitedRequest,
				Info:   "请求被限流，请稍后再试",
			}
			httpx.OkJson(w, resp)
			return
		default:
			logx.Errorf("DefaultQuota key: %v", lm.key)
			resp := status.Resp{
				Status: status.ErrOfServer,
				Info:   "服务器错误",
			}
			httpx.OkJson(w, resp)
			return
		}
	}
}
