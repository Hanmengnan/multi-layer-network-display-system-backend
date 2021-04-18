package model

import (
	"context"
	"fmt"
	"time"

	config "3network-backend/src/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client          *mongo.Client
	staticDatabase  *mongo.Database
	dynamicDatabase *mongo.Database
	err             error
	ctx, _          = context.WithTimeout(context.Background(), 5*time.Second)
)

func init() {
	connect()
}

func connect() {
	databaseURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/", config.DATABASE_USERNAME, config.DATABASE_PASSWORD, config.DATABASE_IP, config.DATABASE_PORT)
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(databaseURI))
	if err != nil {
		panic(err)
	} else {
		staticDatabase = client.Database(config.DATABASE_STATIC)
		dynamicDatabase = client.Database(config.DATABASE_DYNAMIC)
	}
}

func Disconnect() {
	err = client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}
