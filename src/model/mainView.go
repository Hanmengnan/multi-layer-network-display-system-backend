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
		Database int
		Network  int
		Cluster  int
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

type situationHandleInfo struct {
	_id         primitive.ObjectID `json:"_id" bson:"_id"`
	Id          string             `json:"id"`
	Type        string             `json:"type"`
	Handler     string             `json:"handler"`
	HandleTime  string             `json:"handleTime"`
	HandleState string             `json:"handleState"`
	Origin      string             `json:"origin"`
	Message     string             `json:"message"`
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

	closure := func(aimSituaion string) []situationHandleInfo {
		var situaionList []situationHandleInfo
		cursor, err = dynamicDatabase.Collection("situationHandle").Find(context.TODO(), bson.M{"type": aimSituaion})
		if err != nil {
			log.Fatal(err)
		}
		for cursor.Next(context.TODO()) {
			err = cursor.Decode(&situationHandleInfoItem)
			if err != nil {
				log.Fatal(err)
			}
			situaionList = append(situaionList, situationHandleInfoItem)
		}
		return situaionList
	}

	situationHandle["handled"] = closure("2")
	situationHandle["handling"] = closure("1")
	situationHandle["unhandled"] = closure("0")

	err = cursor.Close(context.TODO())

	return situationHandle
}
