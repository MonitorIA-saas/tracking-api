package database

import (
	"context"

	"github.com/monitoria-saas/tracking-api/logger"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoConnection struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewMongoConnection(ctx context.Context, logger logger.Logger, database string) (*mongoConnection, error) {
	logger.Info("Connecting to mongo database")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("MONGO_URL")))

	if err != nil {
		logger.Critic("failed to connect to mongo db", err)
		return nil, err
	}

	logger.Info("Connected to mongo")
	return &mongoConnection{
		Client:   client,
		Database: client.Database(database),
	}, nil
}

func (c *mongoConnection) Disconnect(ctx context.Context) {
	c.Client.Disconnect(ctx)
}
