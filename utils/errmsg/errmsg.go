package errmsg

const (
	SUCCSE = 200
	ERROR  = 500

	//code = 1000... 用户模块的错误
	ERROR_USERNAME_USED     = 1001
	ERROR_PASSWORD_WROING   = 1002
	ERROR_USER_NOT_EXIST    = 1003
	ERROR_TOCKEN_EXIST      = 1004
	ERROR_TOCKEN_RUNTIME    = 1005
	ERROR_TOCKEN_WRONG      = 1006
	ERROR_TOCKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT     = 1008
	ERROR_USERNAME_NOT_NULL = 1009
	ERROR_USER_LOGINING     = 1010
	//code = 2000... 文章模块的错误
	ERROR_ART_NOT_EXIST = 2001
	//code = 3000... 分类模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002

	//code = 4000... redis相关问题
	ERROR_REIDS = 4000
)

var codemsg = map[int]string{
	SUCCSE:                  "OK",
	ERROR:                   "FAIL",
	ERROR_USERNAME_USED:     "用户名已存在",
	ERROR_USERNAME_NOT_NULL: "用户名不能为空",
	ERROR_PASSWORD_WROING:   "密码错误",
	ERROR_USER_NOT_EXIST:    "用户不存在",
	ERROR_TOCKEN_EXIST:      "TOCKEN不存在",
	ERROR_TOCKEN_RUNTIME:    "TOCKEN已过期",
	ERROR_TOCKEN_WRONG:      "TOCKEN不正确",
	ERROR_TOCKEN_TYPE_WRONG: "TOCKEN格式不正确",
	ERROR_CATENAME_USED:     "分类名已存在",
	ERROR_ART_NOT_EXIST:     "文章不存在",
	ERROR_CATE_NOT_EXIST:    "分类不存在",
	ERROR_USER_NO_RIGHT:     "该用户无权限",
	ERROR_REIDS:             "Redis错误",
	ERROR_USER_LOGINING:     "用户已登录",
}

func GetErrMsg(code int) string {
	return codemsg[code]
}
