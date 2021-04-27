package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type bandSetting struct {
	Begin string
	End   string
}

type basicInfo struct {
	_id     primitive.ObjectID `json:"id" bson:"_id"`
	Version string             `bson:"version"`
	Code    string             `bson:"code"`
	State   struct {
		Database int64
		Network  int64
		Cluster  int64
	} `bson:"state"`
	People struct {
		Principal string
		Duty      string
	} `bson:"people"`
	Band struct {
		Light     bandSetting
		Data      bandSetting
		Emergency bandSetting
	} `bson:"band"`
}

type location struct {
	Type        string
	Coordinates []float64
}

type NodeInfo struct {
	_id      primitive.ObjectID `json:"id" bson:"_id"`
	Id       string
	Name     string
	City     string
	state    string
	Location location
	Type     string
	Topology string `json:"topology"`
}

type LinkInfo struct {
	_id         primitive.ObjectID `json:"id" bson:"_id"`
	Id          string
	Name        string
	Node1Id     string
	Node1Name   string
	Node2Id     string
	Node2Name   string
	Contain     int64
	Band        string
	Builded     bool
	UsedForTime bool
	UsedForData bool
}

type situationHandleInfo struct {
	_id         primitive.ObjectID `json:"_id" bson:"_id"`
	Id          string             `json:"id"`
	Type        int64              `json:"type"`
	Handler     string             `json:"handler"`
	HandleTime  string             `json:"handleTime"`
	HandleState int64              `json:"handleState"`
	Origin      string             `json:"origin"`
	Message     string             `json:"message"`
}

type NetworkFlow struct {
	_id         primitive.ObjectID `json:"id" bson:"_id"`
	Time        string
	FlowData    string
	DailyGrowth float64
}

func GetSystemBasicInfo() *basicInfo {
	var systemBasicInfo basicInfo
	err = staticDatabase.Collection("basicInfo").FindOne(context.TODO(), bson.D{}).Decode(&systemBasicInfo)
	if err != nil {
		log.Fatal(err)
	}
	return &systemBasicInfo
}

func GetSituationHandleInfo() map[string][]situationHandleInfo {
	var situationHandleInfoItem situationHandleInfo
	var situationHandle = make(map[string][]situationHandleInfo)
	var cursor *mongo.Cursor

	closure := func(aimSituaion int64) []situationHandleInfo {
		var situaionList []situationHandleInfo
		cursor, err = staticDatabase.Collection("nodeMessage").Find(context.TODO(), bson.M{"type": aimSituaion})
		find(&situationHandleInfoItem, &situaionList, cursor)
		err = cursor.Close(context.TODO())
		return situaionList
	}

	situationHandle["handled"] = closure(2)
	situationHandle["handling"] = closure(1)
	situationHandle["unhandled"] = closure(0)

	err = cursor.Close(context.TODO())

	return situationHandle
}

func GetNetworkFlow() []NetworkFlow {
	var cursor *mongo.Cursor
	var infoList []NetworkFlow
	var info NetworkFlow
	cursor, err = staticDatabase.Collection("flowChange").Find(context.TODO(), bson.M{})
	find(&info, &infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}
