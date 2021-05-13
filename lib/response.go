package lib

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"error_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const StatusOK int = 0

func SuccessResponse(data interface{}) *Response {
	return &Response{
		Success: true,
		Code:    StatusOK,
		Message: "成功",
		Data:    data,
	}
}

func ErrorResponse(error *Error) *Response {
	return &Response{
		Success: false,
		Code:    error.Code,
		Message: error.Message,
		Data:    nil,
	}
}
