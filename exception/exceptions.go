package exception

import "errors"

type ErrCode struct {
	err  error
	code int
}

type UdsError interface {
	Err() error
	Code() int
	ToJson() map[string]any
}

func NewErr(code int, msg string) UdsError {
	return &ErrCode{errors.New(msg), code}
}

var (
	ErrPassword    = NewErr(20000, "用户名或密码错误")
	ErrNoUserFound = NewErr(20003, "用户不存在")

	ErrFindUserRole    = NewErr(21000, "查询用户角色错误")
	ErrFindUserAccount = NewErr(21001, "查询用户账号错误")
	ErrFindUserCompany = NewErr(21002, "查询用户公司信息错误")

	ErrTelValid   = NewErr(30003, "手机号验证错误")
	ErrEmailValid = NewErr(30007, "邮箱验证错误")

	ErrNoAuth               = NewErr(40001, "无权限访问")
	ErrRefreshTokenExpired  = NewErr(40002, "RefreshToken过期")
	ErrTokenExpired         = NewErr(40003, "Token过期")
	ErrUnknow               = NewErr(40004, "未知错误")
	ErrTokenExpiredByLogout = NewErr(40005, "Token过期由于登出")
	ErrForbitUserUse        = NewErr(40006, "该用户账号被禁止")
	ErrNoAuthCredential     = NewErr(40007, "没有提供token")
	ErrNoMatchedPath        = NewErr(40008, "没有匹配到路径")
	ErrTokenFormat          = NewErr(40009, "Token格式错误")
	ErrPemParse             = NewErr(40010, "Pem解析错误")

	ErrNoMatchedRoute  = NewErr(40011, "没有匹配到路由")
	ErrRouteConfig     = NewErr(40012, "路由内部配置错误")
	ErrNoHealthService = NewErr(40013, "没有健康实例")

	ErrHttpCreateRequest = NewErr(40100, "创建Http请求错误")
	ErrHttpRequest       = NewErr(40101, "Http请求错误")
	ErrReadHttpResponse  = NewErr(40102, "读取Http请求错误")
)

func (e *ErrCode) ToJson() map[string]any {
	return map[string]any{
		"code": e.code,
		"msg":  e.err.Error(),
	}
}

func (e *ErrCode) Err() error {
	return e.err
}

func (e *ErrCode) Code() int {
	return e.code
}
