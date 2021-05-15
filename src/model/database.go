package model

import (
	"context"
	"fmt"
	"reflect"
	"time"

	config "multi-layer-network-display-system-backend/src/config"

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

func find(infoListPtr interface{}, cursor *mongo.Cursor) {
	infoListValue := reflect.ValueOf(infoListPtr).Elem()
	infoListType := infoListValue.Type()
	infoELemType := infoListType.Elem()
	infoSlice := reflect.MakeSlice(reflect.SliceOf(infoELemType), 0, 0)

	for cursor.Next(context.TODO()) {
		newInfo := reflect.New(infoELemType)
		err = cursor.Decode(newInfo.Interface())
		if err != nil {
			fmt.Printf("%v", err)
			return
		}
		infoSlice = reflect.Append(infoSlice, newInfo.Elem())
	}

	infoListValue.Set(infoSlice)
}
