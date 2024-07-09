package errno

// 全局错误码定义
const (
	// OK 正确
	OK = 0
	// InternalServerError 服务器内部错误
	InternalServerError = 10001
	// ErrBind 绑定错误
	ErrBind = 10002
	// ErrValidation 参数验证错误
	ErrValidation = 10003
	// ErrDatabase 数据库错误
	ErrDatabase = 10004
	// ErrToken 错误的token
	ErrToken = 10005
	// ErrEncrypt 加密错误
	ErrEncrypt = 10006
	// ErrDecrypt 解密错误
	ErrDecrypt = 10007
	// ErrSign 签名错误
	ErrSign = 10008
	// ErrCheckSign 签名校验错误
	ErrCheckSign = 10009
	// ErrTokenInvalid token无效
	ErrTokenInvalid = 10010

	// ErrUserNotFound 用户不存在
	ErrUserNotFound = 20001
	// ErrUserExist 用户已经存在
	ErrUserExist = 20002
	// ErrPasswordIncorrect 密码错误
	ErrPasswordIncorrect = 20003
	// ErrEncryptPassword 加密密码错误
	ErrEncryptPassword = 20004
	// ErrEmailExist 邮箱已经存在
	ErrEmailExist = 20005
)

var ErrMsg = map[int32]string{
	OK:                  "成功",
	InternalServerError: "服务器内部错误",
	ErrBind:             "绑定参数错误",
	ErrValidation:       "参数验证错误",
	ErrDatabase:         "数据库错误",
	ErrToken:            "错误的token",
	ErrEncrypt:          "加密错误",
	ErrDecrypt:          "解密错误",
	ErrSign:             "签名错误",
	ErrCheckSign:        "签名校验错误",
	ErrTokenInvalid:     "token无效",

	ErrUserNotFound:      "用户不存在",
	ErrUserExist:         "用户已经存在",
	ErrPasswordIncorrect: "密码错误",
	ErrEncryptPassword:   "加密密码错误",
	ErrEmailExist:        "邮箱已经存在",
}

// Err2Msg 错误码转错误信息
func Err2Msg(code int32) string {
	msg, ok := ErrMsg[code]
	if ok {
		return msg
	}
	return ErrMsg[InternalServerError]
}
