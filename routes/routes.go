// routes/routes.go
package routes

import (
	"go-login-server/middleware"
	"go-login-server/routes/login"
	"go-login-server/routes/protected"
	"go-login-server/routes/register"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "go-login-server/docs" // swag init 명령어로 생성된 docs 패키지 임포트
)

// SetupRoutes는 Gin 엔진을 설정하고 모든 라우트를 등록합니다.
func SetupRoutes(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    // 요청 ID 미들웨어 추가
    // router.Use(middleware.RequestIDMiddleware())

    // 요청 로깅 및 복구 미들웨어 추가
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // Swagger 라우트 설정
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // API 버전 그룹화 (예: /v1)
    v1 := router.Group("/v1")
    {
        // /v1/register 라우트 설정
        register.RegisterRoutes(v1, db)

        // /v1/login 라우트 설정
        login.LoginRoutes(v1, db)

        // 보호된 라우트 그룹 설정
        authorized := v1.Group("/protected")
        authorized.Use(middleware.AuthMiddleware())
        {
            authorized.GET("/dashboard", protected.GetDashboard)
        }
    }

    // 기본 라우트 또는 헬스 체크 등 추가 라우트 설정 가능
    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":  "success",
            "message": "Go Login Server is running",
        })
    })

    return router
}
