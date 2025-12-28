package main

import (
	"context"

	"github.com/monitoria-saas/tracking-api/internal/database"
	gingonic "github.com/monitoria-saas/tracking-api/internal/http/gin-gonic"
	"github.com/monitoria-saas/tracking-api/logger"
	"github.com/monitoria-saas/tracking-api/settings"
	"github.com/spf13/viper"
)

func main() {
	// general settings
	myLogger := logger.NewLogger()
	ctx := context.Background()

	myLogger.Info("STARTING APPLICATION")

	// preparing enviroments variables
	settings.LoadEnv(myLogger)

	//getting connection with mongodb
	connection, err := database.NewMongoConnection(ctx, myLogger, viper.GetString("MONGO_DATABASE"))

	if err != nil {
		panic("Failed to connect to mongo db")
	}

	defer connection.Disconnect(ctx)

	// starting api
	gingonic.StartAPI()
}
