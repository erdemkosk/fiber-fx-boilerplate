package mongodb

import (
	"context"
	"log"

	"github.com/erdemkosk/fiber-fx-boilerplate/bootstrap/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func NewMongoDB(config *config.Config) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(config.MongoDbUrl)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	log.Println("MongoDB connected")

	return &MongoDB{Client: client}, nil
}
