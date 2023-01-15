package common

const (
	ErrOfServer        = 500  //服务器错误
	ErrNoSuchUser      = 2001 //用户名不存在
	ErrAlreadyHaveUser = 2002 //用户名已存在
	ErrWrongPassword   = 2003 //密码错误
	ErrFailParseToken  = 2004 //解析token错误
)

const (
	InfoErrOfServer = "服务器错误"
)

type OKResp struct {
	Status int    `json:"status_code"`
	Info   string `json:"status_msg"`
}
