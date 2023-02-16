package key

const (
	// RedisUserNameCacheKey redis用户名缓存Key
	RedisUserNameCacheKey = "username"
	// RedisUserIdCacheKey redis用户id缓存Key
	RedisUserIdCacheKey = "user_id"
	// RedisUserIDLastTimeStamp 上次生成id时间戳
	RedisUserIDLastTimeStamp = "userid_last_timestamp"
	// RedisVideoIDLastTimeStamp 上次生成video_id时间戳
	RedisVideoIDLastTimeStamp = "video_id_last_timestamp"
	// RedisHashMod redis big key拆分 哈希模数
	RedisHashMod uint64 = 9
)

const (
	// LimitKeyPrefix redis限流key前缀
	LimitKeyPrefix = "periodlimit"
	// LimitKeyUserApi redis限流key user的api服务
	LimitKeyUserApi = "user_api"
	// LimitKeyVideoApi redis限流key video的api服务
	LimitKeyVideoApi = "video_api"
	// LimitKeyActionApi redis限流key action的api服务
	LimitKeyActionApi = "action_api"
	// LimitKeyChatApi redis限流key chat的api服务
	LimitKeyChatApi = "chat_api"
)

const (
	// MAXBytes video服务限制文件最大为50M
	MAXBytes = 50 * 1024 * 1024
	// FeedNum Feed视频流一次性返回的视频数
	FeedNum = 10
	// LogReadBuffer 日志buffer
	LogReadBuffer = 100
	// LogMsgQueueName 日志队列名
	LogMsgQueueName = "douyin_log"
)
