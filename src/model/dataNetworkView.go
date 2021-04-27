package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataNetworkBasicInfo struct {
	_id          primitive.ObjectID `json:"id" bson:"_id"`
	BandUsed     int64
	BandTotal    int64
	LinkNum      int64
	NodeNum      int64
	TimestampNum int64
	LocationNum  int64
}

type DataNetworkErrorAlarm struct {
	_id        primitive.ObjectID `json:"id" bson:"_id"`
	Id         string
	RTime      string
	Type       int64
	Level      string
	Message    string
	State      string
	FromNodeId string
	ToNodeId   string
}

type LinkDetail struct {
	_id            primitive.ObjectID `json:"id" bson:"_id"`
	Id             string
	Time           string
	PrecisionError float64
	Loss           float64
	Used           float64
}

type NodeDetail struct {
	_id            primitive.ObjectID `json:"id" bson:"_id"`
	Precision      string             `json:"precision"`
	Error          int64              `json:"error"`
	Throughput     int64              `json:"throughput"`
	ForwardingRate float64            `json:"forwardingRate"`
}

func GetDataNetworkBasicInfo() *DataNetworkBasicInfo {
	var info DataNetworkBasicInfo
	res := staticDatabase.Collection("dataNetworkBasicInfo").FindOne(context.TODO(), bson.M{})
	err = res.Decode(&info)
	return &info
}

func GetDataNetworkLinkInfo() []LinkInfo {
	var cursor *mongo.Cursor
	var info LinkInfo
	var infoList []LinkInfo
	cursor, err = staticDatabase.Collection("linkInfo").Find(context.TODO(), bson.M{})
	find(&info, &infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

func GetDataNetworkLinkDetail(linkId string) []LinkDetail {
	var cursor *mongo.Cursor
	var infoList []LinkDetail
	var info LinkDetail
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"time", -1}})
	findOptions.SetLimit(30 * 24)

	cursor, err = staticDatabase.Collection("linkDetail").Find(context.TODO(), bson.M{"id": linkId}, findOptions)
	find(&info, &infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

func GetDataNetworkNodeInfo() []NodeInfo {
	var cursor *mongo.Cursor
	var infoList []NodeInfo
	var info NodeInfo
	cursor, err = staticDatabase.Collection("nodeInfo").Find(context.TODO(), bson.M{})
	find(&info, &infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

// 节点的详细信息大部分是与报文相关
// 未完成

func GetDataNetworkNodeDetail(nodeId string) []NodeDetail {
	var cursor *mongo.Cursor
	var infoList []NodeDetail
	var info NodeDetail
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"time", -1}})
	findOptions.SetLimit(30 * 24)

	cursor, err = staticDatabase.Collection("linkDetail").Find(context.TODO(), bson.M{"id": nodeId}, findOptions)
	find(&info, &infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

// 该预警是链路预警
// 未完成

func GetDataNetworkErrorALarms() []DataNetworkErrorAlarm {
	var cursor *mongo.Cursor
	var infoList []DataNetworkErrorAlarm
	var info DataNetworkErrorAlarm
	cursor, err = staticDatabase.Collection("dataNetworkErrorAlarm").Find(context.TODO(), bson.M{})
	find(&info, &infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}
