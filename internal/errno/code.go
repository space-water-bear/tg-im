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

	/*
	 用户相关
	*/
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

	/*
		好友相关
	*/
	// ErrFriendNotFound 好友不存在
	ErrFriendNotFound = 30001
	// ErrFriendExist 好友已经存在
	ErrFriendExist = 30002
	// ErrFriendRequestNotFound 好友请求不存在
	ErrFriendRequestNotFound = 30003
	// ErrFriendRequestExist 好友请求已经存在
	ErrFriendRequestExist = 30004
	// ErrFriendRequestRejected 好友请求被拒绝
	ErrFriendRequestRejected = 30005
	// ErrFriendBlackList 好友在黑名单中
	ErrFriendBlackList = 30006
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

	ErrFriendNotFound:        "好友不存在",
	ErrFriendExist:           "好友已经存在",
	ErrFriendRequestNotFound: "好友请求不存在",
	ErrFriendRequestExist:    "好友请求已经存在",
	ErrFriendRequestRejected: "好友请求被拒绝",
	ErrFriendBlackList:       "好友在黑名单中",
}

// Err2Msg 错误码转错误信息
func Err2Msg(code int32) string {
	msg, ok := ErrMsg[code]
	if ok {
		return msg
	}
	return ErrMsg[InternalServerError]
}
