package common

const (
	// RedisUserNameCacheKey redis用户名缓存Key
	RedisUserNameCacheKey = "username"
	// RedisUserIdCacheKey redis用户id缓存Key
	RedisUserIdCacheKey = "user_id"

	// LimitKeyPrefix redis限流key前缀
	LimitKeyPrefix = "periodlimit"
	// LimitKeyUserApi redis限流key user的api服务
	LimitKeyUserApi = "user_api"
)
