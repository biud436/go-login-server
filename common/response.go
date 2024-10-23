// common/response.go
package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response는 API 응답의 공통 구조체입니다.
type Response struct {
    Status  string      `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Code    string      `json:"code,omitempty"`
}

// SuccessResponse는 성공적인 응답을 클라이언트에 보냅니다.
// data는 응답에 포함할 데이터입니다.
// message는 응답 메시지입니다.
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
    response := Response{
        Status:  "success",
        Message: message,
        Data:    data,
    }
    c.JSON(statusCode, response)
}

// ErrorResponse는 에러 응답을 클라이언트에 보냅니다.
// err는 발생한 에러입니다.
func ErrorResponse(c *gin.Context, err error) {
    if appErr, ok := err.(*AppError); ok {
        // AppError 타입인 경우
        response := Response{
            Status:  "error",
            Message: appErr.Message,
            Code:    appErr.Code,
        }
        c.JSON(appErr.StatusCode, response)
    } else {
        // 일반 에러인 경우
        response := Response{
            Status:  "error",
            Message: "Internal Server Error",
            Code:    "INTERNAL_ERROR",
        }
        c.JSON(http.StatusInternalServerError, response)
    }
}
