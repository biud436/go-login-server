// routes/register/register.go
package register

import (
	"go-login-server/handlers"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes는 /register 엔드포인트에 대한 라우트를 설정합니다.
func RegisterRoutes(router *gin.RouterGroup, db *gorm.DB) {
    router.POST("/register", handlers.Register(db))
}
