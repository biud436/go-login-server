// common/errors.go
package common

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AppError 구조체 정의
type AppError struct {
    Code       string `json:"code"`
    Message    string `json:"message"`
    StatusCode int    `json:"-"`
}

// Error 메서드 구현 (error 인터페이스 만족)
func (e *AppError) Error() string {
    return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// 공통 에러 변수 정의
var (
    ErrUsernameExists       = &AppError{Code: "USERNAME_EXISTS", Message: "Username already exists", StatusCode: http.StatusBadRequest}
    ErrInvalidCredentials   = &AppError{Code: "INVALID_CREDENTIALS", Message: "Invalid username or password", StatusCode: http.StatusUnauthorized}
    ErrAuthorizationHeader = &AppError{Code: "AUTH_HEADER_MISSING", Message: "Authorization header missing", StatusCode: http.StatusUnauthorized}
    ErrInvalidAuthHeader    = &AppError{Code: "INVALID_AUTH_HEADER", Message: "Invalid authorization header", StatusCode: http.StatusUnauthorized}
    ErrInvalidToken         = &AppError{Code: "INVALID_TOKEN", Message: "Invalid token", StatusCode: http.StatusUnauthorized}
    ErrJWTSecretNotSet      = &AppError{Code: "JWT_SECRET_NOT_SET", Message: "JWT_SECRET not set in environment variables", StatusCode: http.StatusInternalServerError}
    ErrHashingPassword      = &AppError{Code: "HASHING_PASSWORD", Message: "Error hashing password", StatusCode: http.StatusInternalServerError}
    ErrGeneratingToken      = &AppError{Code: "GENERATING_TOKEN", Message: "Could not generate token", StatusCode: http.StatusInternalServerError}
    ErrPasswordLength       = &AppError{Code: "PASSWORD_LENGTH", Message: "Password must be at least 8 characters long", StatusCode: http.StatusBadRequest}
    ErrInvalidInput         = &AppError{Code: "INVALID_INPUT", Message: "Invalid input", StatusCode: http.StatusBadRequest}
)

// RespondWithError 함수 정의
func RespondWithError(c *gin.Context, err error) {
    if appErr, ok := err.(*AppError); ok {
        // AppError 타입인 경우
        log.Printf("Error [%s]: %s", appErr.Code, appErr.Message)
        c.JSON(appErr.StatusCode, gin.H{
            "error":  appErr.Message,
            "code":   appErr.Code,
        })
    } else {
        // 일반 에러인 경우
        log.Printf("Unexpected Error: %s", err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Internal Server Error",
            "code":  "INTERNAL_ERROR",
        })
    }
}
