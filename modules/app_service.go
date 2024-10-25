package modules

import (
	"go-login-server/routes"
	"go-login-server/validators"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	"gorm.io/gorm"

	"log"
	"os"
)

type AppService struct {
	DB *gorm.DB
}

func NewAppService(db *gorm.DB) *AppService {
    return &AppService{
        DB: db,
    }
}
 
/**
 * (Public) StartServer
 * 서버를 시작합니다.
 */
func (as *AppService) StartServer() {

    // Gin의 기존 Validator 엔진 가져오기
    // https://gin-gonic.com/docs/examples/custom-validators/
    if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
        validators.RegisterCustomValidators(v)
    } else {
        log.Fatal("Validator 엔진을 가져오지 못했습니다.")
    }

    // 라우트 설정 및 서버 시작
    router := routes.SetupRoutes(as.DB)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    if err := router.Run(":" + port); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}

/**
 * (Public) OnModuleStart
 * 모듈이 시작될 때 호출되는 함수입니다.
 */
func (as *AppService) OnModuleStart(redisService *RedisService) {
    redisService.Init();

    log.Println("Redis Service Initialized")
}