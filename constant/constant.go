package constant

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SuccessResp 成功响应
func SuccessResp(data interface{}) *Response {
	return &Response{
		Code:    0,
		Message: "success",
		Data:    data,
	}
}

// FailResp 失败响应
func FailResp(msg string) *Response {
	return &Response{
		Code:    ErrGeneral,
		Message: msg,
	}
}
