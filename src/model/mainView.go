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

type sysInfo struct {
	_id        primitive.ObjectID `json:"id" bson:"_id"`
	SystemName string             `json:"systemName"`
	Version    string             `json:"version"`
	Code       string             `json:"code"`
	State      struct {
		Database string `json:"database"`
		Network  string `json:"network"`
		Cluster  string `json:"cluster"`
	} `json:"state"`
	People struct {
		Principal string
		Duty      string
	} `json:"people"`
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
	_id      primitive.ObjectID `json:"_id" bson:"_id"`
	Id       string             `json:"id"`
	Name     string             `json:"name"`
	City     string             `json:"city"`
	State    string             `json:"state"`
	Location location           `json:"location"`
	Type     string             `json:"type"`
	Topology string             `json:"topology"`
}

type LinkInfo struct {
	_id         primitive.ObjectID `json:"id" bson:"_id"`
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Node1Id     string             `json:"startId"`
	Node1Name   string             `json:"start"`
	Node2Id     string             `json:"endId"`
	Node2Name   string             `json:"end"`
	Contain     int                `json:"contain"`
	Band        string             `json:"band"`
	Builded     bool               `json:"builded"`
	UsedForTime bool               `json:"usedForTime"`
	UsedForData bool               `json:"usedForData"`
}

type situationHandleInfo struct {
	_id         primitive.ObjectID `json:"_id" bson:"_id"`
	Id          string             `json:"id"`
	Type        int                `json:"type"`
	Handler     string             `json:"handler"`
	HandleTime  string             `json:"handleTime"`
	HandleState int                `json:"handleState"`
	Origin      string             `json:"origin"`
	Message     string             `json:"message"`
}

type NetworkFlow struct {
	_id         primitive.ObjectID `json:"id" bson:"_id"`
	Time        string
	FlowData    string
	DailyGrowth float64
}

func GetSystemBasicInfo() *sysInfo {
	var systemBasicInfo sysInfo
	err = staticDatabase.Collection("basicInfo").FindOne(context.TODO(), bson.D{}).Decode(&systemBasicInfo)
	if err != nil {
		log.Fatal(err)
	}
	return &systemBasicInfo
}

func GetSituationHandleInfo() map[string][]situationHandleInfo {
	var situationHandle = make(map[string][]situationHandleInfo)
	var cursor *mongo.Cursor

	closure := func(aimSituaion int) []situationHandleInfo {
		var situaionList []situationHandleInfo
		cursor, err = staticDatabase.Collection("nodeMessage").Find(context.TODO(), bson.M{"type": aimSituaion, "handleState": 0})
		find(&situaionList, cursor)
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
	cursor, err = staticDatabase.Collection("flowChange").Find(context.TODO(), bson.M{})
	find(&infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

func GetLinkInfo() []LinkInfo {
	var cursor *mongo.Cursor
	var infoList []LinkInfo
	cursor, err = staticDatabase.Collection("linkInfo").Find(context.TODO(), bson.M{})
	find(&infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

func GetNodeInfo() []NodeInfo {
	var cursor *mongo.Cursor
	var infoList []NodeInfo
	cursor, err = staticDatabase.Collection("nodeInfo").Find(context.TODO(), bson.M{})
	find(&infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}
