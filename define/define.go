package define

var MailPassword = "pzmaljknscpsebeg"

const TokenPrefix = "TOKEN_"          // 存储token的前缀
const CodePrefix = "Code_"            // 存储code的前缀
var TokenExpireTime = 30 * 60         // token过期时间
var EachCodeForEmailWaitTime = 2 * 60 // 验证码发送一次需要间隔多长时间
var CodeExpireTime = 2 * 60           // 验证码过期时间
var CacheExpireTime = 5 * 60          // 缓存过期时间

type ResponseResult struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func NewResponseResult(code int64, msg string, data any) *ResponseResult {
	return &ResponseResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
