// routes/login/login.go
package login

import (
	"go-login-server/handlers"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// LoginRoutes는 /login 엔드포인트에 대한 라우트를 설정합니다.
func LoginRoutes(router *gin.RouterGroup, db *gorm.DB) {
    router.POST("/login", handlers.Login(db))
}
