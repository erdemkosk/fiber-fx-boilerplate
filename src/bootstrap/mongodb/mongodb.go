package mongodb

import (
	"context"

	"github.com/erdemkosk/fiber-fx-boilerplate/src/bootstrap/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type MongoDB struct {
	Client *mongo.Client
	Logger *zap.Logger
}

func NewMongoDB(config *config.Config, logger *zap.Logger) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(config.Mongo.URL)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return &MongoDB{Client: client, Logger: logger}, nil
}

func (m *MongoDB) onStart(ctx context.Context) error {
	m.Logger.Info("📦 MongoDB started successfully.")
	return nil
}

func (m *MongoDB) onStop(ctx context.Context) error {
	m.Logger.Info("🔌 MongoDB stopping...")

	if err := m.Client.Disconnect(ctx); err != nil {
		m.Logger.Error("Error disconnecting MongoDB: %v", zap.Error(err))
		return err
	} else {
		m.Logger.Info("✅ MongoDB disconnected successfully.")
	}

	return nil
}
