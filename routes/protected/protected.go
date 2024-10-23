// routes/protected/protected.go
package protected

import (
	"go-login-server/common"

	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDashboard godoc
// @Summary 사용자 대시보드 가져오기
// @Description 인증된 사용자의 대시보드를 가져옵니다
// @Tags protected
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} common.Response
// @Failure 500 {object} common.Response
// @Router /v1/protected/dashboard [get]
func GetDashboard(c *gin.Context) {
    username, exists := c.Get("username")
    if !exists {
        common.ErrorResponse(c, &common.AppError{
            Code:       "CONTEXT_ERROR",
            Message:    "Username not found in context",
            StatusCode: http.StatusInternalServerError,
        })
        return
    }
    common.SuccessResponse(c, http.StatusOK, "Welcome "+username.(string), nil)
}
