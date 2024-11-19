package constant

// HTTP 相关通用错误码
const (
	ErrBadRequest         = 400 // 客户端请求错误
	ErrUnauthorized       = 401 // 未授权
	ErrForbidden          = 403 // 禁止访问
	ErrNotFound           = 404 // 资源未找到
	ErrInternalServer     = 500 // 服务器内部错误
	ErrServiceUnavailable = 503 // 服务不可用
)

// 业务逻辑相关通用错误码
const (
	ErrGeneral = int32(1000)
)
