package model

// redis相关
const (
	// RedisUserPhoneNumbers 手机号是否注册的集合 长期有效
	RedisUserPhoneNumbers = "user_pnum"

	//验证码键值对 有效期两分钟
	//PoneNumber+VCode 值为VCode

	//今日发送次数
	//TodayTimes="pkey"+PoneNumber
)
