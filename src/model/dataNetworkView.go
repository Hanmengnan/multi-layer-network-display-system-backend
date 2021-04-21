package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type location struct {
	Type        string
	Coordinates []float64
}

type DataNetworkBasicInfo struct {
	_id          primitive.ObjectID `json:"id" bson:"_id"`
	BandUsed     int
	BandTotal    int
	LinkNum      int
	NodeNum      int
	TimestampNum int
	LocationNum  int
}

type DataNetworkLinkInfo struct {
	_id       primitive.ObjectID `json:"id" bson:"_id"`
	Id        string
	Name      string
	Node1Id   string
	Node1Name string
	Node2Id   string
	Node2Name string
	Band      string
}

type DataNetworkNodeInfo struct {
	_id            primitive.ObjectID `json:"id" bson:"_id"`
	Id             string
	Name           string
	City           string
	state          string
	Location       location
	Type           string
	Precision      string  `json:"precision"`
	Error          int     `json:"error"`
	Throughput     int     `json:"throughput"`
	ForwardingRate float64 `json:"forwardingRate"`
	Topology       string  `json:"topology"`
}

type DataNetworkErrorAlarm struct {
	_id        primitive.ObjectID `json:"id" bson:"_id"`
	Id         string
	RTime      string
	Type       int
	Level      string
	Message    string
	State      string
	FromNodeId string
	ToNodeId   string
}

type DataNetworkLinkDetail struct {
	_id            primitive.ObjectID `json:"id" bson:"_id"`
	Id             string
	Time           string
	PrecisionError float64
	Loss           float64
	Used           float64
}

type DataNetworkFlow struct {
	_id      primitive.ObjectID `json:"id" bson:"_id"`
	Time     string
	FlowData string
}

func GetDataNetworkBasicInfo() *DataNetworkBasicInfo {
	var info DataNetworkBasicInfo
	res := staticDatabase.Collection("dataNetworkBasicInfo").FindOne(context.TODO(), bson.M{})
	err = res.Decode(&info)
	if err != nil {
		log.Fatal(err)
	}
	return &info
}

func GetDataNetworkLinkInfo(linkId ...string) []DataNetworkLinkInfo {
	fmt.Printf("%v", linkId)
	var info DataNetworkLinkInfo
	var infoList []DataNetworkLinkInfo
	if len(linkId) == 0 {
		var cursor *mongo.Cursor
		cursor, err = staticDatabase.Collection("dataNetworkLinkInfo").Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		for cursor.Next(context.TODO()) {
			err = cursor.Decode(&info)
			if err != nil {
				log.Fatal(err)
			}
			infoList = append(infoList, info)
		}
		cursor.Close(context.TODO())
		return infoList
	} else {
		res := staticDatabase.Collection("dataNetworkLinkDetail").FindOne(context.TODO(), bson.M{"id": linkId[0]})
		err = res.Decode(&info)
		infoList = append(infoList, info)
		return infoList
	}
}

func GetDataNetworkNodeInfo() []DataNetworkNodeInfo {
	var info DataNetworkNodeInfo
	var infoList []DataNetworkNodeInfo
	var cursor *mongo.Cursor
	cursor, err = staticDatabase.Collection("dataNetworkLinkInfo").Find(context.TODO(), bson.M{"type": "业务上下站"})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&info)
		if err != nil {
			log.Fatal(err)
		}
		infoList = append(infoList, info)
	}
	cursor.Close(context.TODO())
	return infoList
}

func GetDataNetworkLinkDetail(linkId string) []DataNetworkLinkDetail {
	var cursor *mongo.Cursor
	var infoList []DataNetworkLinkDetail
	info := new(DataNetworkLinkDetail)

	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"time", -1}})
	findOptions.SetLimit(30 * 24)

	cursor, err = staticDatabase.Collection("dataNetworkLinkDetail").Find(context.TODO(), bson.M{"id": linkId}, findOptions)
	res := find(info, cursor)
	err = cursor.Close(context.TODO())

	for _, item := range res {
		infoList = append(infoList, *item.(*DataNetworkLinkDetail))
	}
	return infoList
}

func GetDataNetworkErrorALarms() []DataNetworkErrorAlarm {
	var info DataNetworkErrorAlarm
	var infoList []DataNetworkErrorAlarm
	var cursor *mongo.Cursor
	cursor, err = staticDatabase.Collection("dataNetworkLinkInfo").Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.TODO()) {
		err = cursor.Decode(&info)
		if err != nil {
			log.Fatal(err)
		}
		infoList = append(infoList, info)
	}
	cursor.Close(context.TODO())
	return infoList
}

func GetDataNetworkFlow() []DataNetworkFlow {
	var cursor *mongo.Cursor
	var infoList []DataNetworkFlow
	info := new(DataNetworkFlow)

	cursor, err = staticDatabase.Collection("dataNetworkFlowChange").Find(context.TODO(), bson.M{})
	res := find(info, cursor)
	err = cursor.Close(context.TODO())

	for _, item := range res {
		infoList = append(infoList, *item.(*DataNetworkFlow))
	}
	return infoList
}
