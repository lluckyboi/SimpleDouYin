package status

const SuccessCode = 0
const (
	ErrOfServer        = 500  //服务器错误
	ErrNoSuchUser      = 2001 //用户名不存在
	ErrAlreadyHaveUser = 2002 //用户名已存在
	ErrWrongPassword   = 2003 //密码错误
	ErrFailParseToken  = 2004 //解析token错误
	ErrLimitedRequest  = 2005 //请求被限流
	ErrRejectedRequest = 2006 //请求被拒绝
	ErrLengthErr       = 2007 //长度错误
	ErrParseTime       = 2008 //时间解析错误
	ErrUnknownAcType   = 2009 //未知action type
)

const (
	InfoErrOfServer = "服务器错误"
)

type Resp struct {
	Status int    `json:"status_code"`
	Info   string `json:"status_msg"`
}
