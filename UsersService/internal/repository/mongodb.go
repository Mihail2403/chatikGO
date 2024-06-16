package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	IMG_COLLECTION = "img"
)

type MongoConfig struct {
	Host string
	Port string
	DB   string
}

func NewMongoDB(ctx context.Context, config MongoConfig) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(config.DB)
	return db, nil
}
