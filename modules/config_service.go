package modules

import (
	"os"
	"strconv"
)

type ConfigService struct {

}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (cs *ConfigService) getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value;
	}
	return defaultVal;
}

func (cs *ConfigService) getEnvAsInt(name string, defaultVal int) int {
	valueStr := cs.getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value;
	}
	return defaultVal;
}