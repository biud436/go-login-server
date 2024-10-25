package modules

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	// RedisClient 선언 (대문자면 외부에서 접근 가능, 소문자면 외부에서 접근 불가능)
	redisClient *redis.Client

	// RedisContext 선언
	redisContext = context.Background()
)

/**
 * RedisService 구조체 선언
 */
type RedisService struct {
	Client *redis.Client
	Context context.Context
	configService *ConfigService
}

/**
 * RedisService 생성자 함수	
 */
func NewRedisService(configService *ConfigService) *RedisService {
	return &RedisService{
		Client: redisClient,
		Context: redisContext,
		configService: configService,
	}
}

func (rs *RedisService) Init() {
	
	rs.Client = redis.NewClient(&redis.Options{
		Addr: rs.configService.getEnv("REDIS_HOST", "localhost") + ":" + rs.configService.getEnv("REDIS_PORT", "6379"),
	})

	_, err := rs.Client.Ping(rs.Context).Result()
	if err != nil {
		log.Fatalf(("Failed to connect to Redis: %v"), err)
	}

	log.Println("Connected to Redis")
}

func (rs *RedisService) Get(key string) (string, error) {
	return rs.Client.Get(rs.Context, key).Result()
}

func (rs *RedisService) Set(key string, value interface{}, expiration int) error {
	return rs.Client.Set(rs.Context, key, value, 0).Err()
}

func (rs *RedisService) Del(key string) error {
	return rs.Client.Del(rs.Context, key).Err()
}

func (rs *RedisService) Exists(key string) (int64, error) {
	return rs.Client.Exists(rs.Context, key).Result()
}