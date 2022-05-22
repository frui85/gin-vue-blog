package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// code = 1000... 表示用户模块错误
	ERROR_USERNAME_USED       = 1001
	ERROR_USER_NOT_EXIST      = 1002
	ERROR_PASSWORD_WRONG      = 1101
	ERROR_PASSWORD_TYPE_WRONG = 1102
	ERROR_TOKEN_EXIST         = 1201
	ERROR_TOKEN_EXPIRED       = 1202
	ERROR_TOKEN_WRONG         = 1203
	ERROR_TOKEN_TYPE_WRONG    = 1204

	// code = 2000... 表示文章模块错误
	ERROR_CATENAME_USED = 2001

	// code = 3000... 表示分类模块错误

	// code = xxx... 待补充
)

var CodeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	ERROR_USERNAME_USED:       "用户名已存在",
	ERROR_USER_NOT_EXIST:      "用户不存在",
	ERROR_PASSWORD_WRONG:      "密码不正确",
	ERROR_PASSWORD_TYPE_WRONG: "密码格式错误",
	ERROR_TOKEN_EXIST:         "TOKEN不存在",
	ERROR_TOKEN_EXPIRED:       "TOKEN已过期",
	ERROR_TOKEN_WRONG:         "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG:    "TOKEN格式错误",

	ERROR_CATENAME_USED: "分类已存在",
}

func GetErrMsg(code int) string {
	return CodeMsg[code]
}
