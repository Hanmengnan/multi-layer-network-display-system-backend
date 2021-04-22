package model

import (
	"context"
	"fmt"
	"log"
	"reflect"
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
)

func init() {
	connect()
}

func connect() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
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
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Disconnect(ctx)
	if err != nil {
		panic(err)
	}
}

func find(info interface{}, infoList interface{}, cursor *mongo.Cursor) {
	infoListValue := reflect.ValueOf(infoList).Elem()
	resEleArray := make([]reflect.Value, 0)

	for cursor.Next(context.TODO()) {
		err = cursor.Decode(info)
		if err != nil {
			log.Fatal(err)
		}
		resEleArray = append(resEleArray, reflect.ValueOf(info).Elem())
	}
	infoListValue.Set(reflect.Append(infoListValue, resEleArray...))
}
