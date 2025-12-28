package settings

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/monitoria-saas/tracking-api/logger"
	"github.com/spf13/viper"
)

func LoadEnv(logger logger.Logger) {
	if _, err := os.Stat(".env"); err == nil {
		if err = godotenv.Load(); err != nil {
			logger.Critic("erro ao carregar .env", errors.New("Failed to load .env file"))
		}
	} else {
		logger.Info("arquivo .env não existe")
	}

	viper.AutomaticEnv()
	logger.Info("environmental variables loaded")
}
