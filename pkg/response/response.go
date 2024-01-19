package response

import (
	"go_gin_framework/pkg/err"
	"time"
)

type Response struct {
	Status  int         `json:"status"`
	ErrCode int         `json:"err_code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data, omitempty"`
	Time    int64       `json:"time"`
}

func SuccessResponse(result interface{}) Response {
	return Response{
		Status:  200,
		Message: "success",
		Data:    result,
		Time:    time.Now().Unix(),
	}
}

func ErrorResponse(message string ) Response {
	return Response{
		Status:  200,
		ErrCode: 1,
		Message: message,
		Time:    time.Now().Unix(),
	}
}
