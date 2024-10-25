// main.go
package main

import (
	"go-login-server/common"
	"go-login-server/models"
	"go-login-server/modules"
	"go-login-server/routes"
	"go-login-server/validators"

	"log"
	"os"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"go.uber.org/dig"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// @title Go 로그인 서버 API
// @version 1.0
// @description 이 API는 Go 기반 로그인 시스템을 위한 샘플 서버입니다.

// @contact.name API 지원
// @contact.url https://github.com/biud436
// @contact.email biud436@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

    log.Println("Dependency Injection Container Initializing...")
    container := dig.New()

    container.Provide(modules.NewConfigService)
    container.Provide(modules.NewRedisService)
    container.Provide(newAppService)

    container.Invoke(func(appService *AppService, redisService *modules.RedisService) {
        appService.onModuleStart(redisService)
        appService.startServer()
    })
}

type AppService struct {
	
}

func newAppService() *AppService {
    return &AppService{}
}
 

func (as *AppService) startServer() {
    // .env 파일 로드
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    // 환경 변수에서 DB 정보 로드
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // PostgreSQL DSN 구성
    dsn := "host=" + dbHost +
        " port=" + dbPort +
        " user=" + dbUser +
        " password=" + dbPassword +
        " dbname=" + dbName +
        " sslmode=disable"

    // 데이터베이스 연결
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        common.ErrorResponse(nil, err) // main.go에서는 gin.Context가 없으므로 별도의 에러 처리 필요
        log.Fatal("Failed to connect to database:", err)
    }

    // 마이그레이션
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        common.ErrorResponse(nil, err) // main.go에서는 gin.Context가 없으므로 별도의 에러 처리 필요
        log.Fatal("Failed to migrate database:", err)
    }

    // Gin의 기존 Validator 엔진 가져오기
    // https://gin-gonic.com/docs/examples/custom-validators/
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        validators.RegisterCustomValidators(v)
    } else {
        log.Fatal("Validator 엔진을 가져오지 못했습니다.")
    }

    // 라우트 설정 및 서버 시작
    router := routes.SetupRoutes(db)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    if err := router.Run(":" + port); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}

func (as *AppService) onModuleStart(redisService *modules.RedisService) {
    redisService.Init();

    log.Println("Redis Service Initialized")
}