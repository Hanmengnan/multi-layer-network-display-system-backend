package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetNodeTypeStatisticsHealth() map[string]int32 {
	var cursor *mongo.Cursor
	var res []bson.M
	var infoList = make(map[string]int32)
	matchStage := bson.D{
		{"$match", bson.D{{"state", "正常"}}},
	}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$type"},
			{"count", bson.M{"$sum": 1}}},
		},
	}
	opt := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err = staticDatabase.Collection("nodeInfo").Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage}, opt)

	if err != nil {
		log.Println(err)
		return nil
	}

	err = cursor.All(context.TODO(), &res)

	if err != nil {
		log.Println(err)
		return nil
	}

	for _, item := range res {
		infoList[item["_id"].(string)] = item["count"].(int32)
	}
	return infoList
}

func GetLinkOverload(nodeId string) map[string]float64 {
	var cursor *mongo.Cursor
	var infoList []graphNode
	var parameter ParameterChange
	var res = make(map[string]float64)
	cursor, err = staticDatabase.Collection("linkInfo").Find(context.TODO(), bson.D{{"$or", []interface{}{bson.M{"node1ID": nodeId}, bson.M{"node2ID": nodeId}}}})
	if err != nil {
		return nil
	}
	find(&infoList, cursor)

	var option = options.FindOne()
	option.SetSort(bson.M{"time": -1})

	for _, item := range infoList {
		err = dynamicDatabase.Collection("linkParameter").FindOne(context.TODO(), bson.M{"id": item.Id}, option).Decode(&parameter)
		if err != nil {
			return nil
		}
		res[item.Name] = parameter.Used
	}
	return res
}

func GetNodeInfo(nodeId string) NodeInfo {
	var info NodeInfo
	err = staticDatabase.Collection("nodeInfo").FindOne(context.TODO(), bson.M{"id": nodeId}).Decode(&info)
	if err != nil {
		return NodeInfo{}
	}
	return info
}
