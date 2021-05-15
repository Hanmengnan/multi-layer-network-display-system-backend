package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Route struct {
	_id   primitive.ObjectID `bson:"__id"`
	Start string             `bson:"start"`
	End   string             `bson:"end"`
	Route [][]string
}

func GetNodePrecisionStatistics() map[string]int32 {
	var res []bson.M
	var cursor *mongo.Cursor
	var infoList = make(map[string]int32)

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$precision"},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}

	sortStage := bson.D{{"$sort", bson.D{{"count", 1}}}}

	opts := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err = staticDatabase.Collection("nodeInfo").Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage}, opts)
	err = cursor.All(context.TODO(), &res)
	for _, item := range res {

		infoList[item["_id"].(string)] = item["count"].(int32)
	}
	return infoList

}

func GetRouteTopology(start, end string) Route {
	var res Route
	err = staticDatabase.Collection("graph").FindOne(context.TODO(), bson.M{"start": start, "end": end}).Decode(&res)
	if err != nil {
		return Route{}
	}
	return res
}
