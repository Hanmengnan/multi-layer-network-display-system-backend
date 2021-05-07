package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
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
		Principal struct {
			Name  string `json:"name"`
			Photo string `json:"photo"`
		}
		Duty struct {
			Name  string `json:"name"`
			Photo string `json:"photo"`
		}
	} `json:"people"`
	Band struct {
		Light     bandSetting
		Data      bandSetting
		Emergency bandSetting
	} `bson:"band"`
	StartTime time.Time `json:"startTime"`
}

type NetParameter struct {
	_id          primitive.ObjectID `bson:"__id" json:"_id,omitempty"`
	RecordTime   time.Time          `json:"recordTime"`
	Connectivity float64            `json:"connectivity"`
	Throughput   float64            `json:"throughput"`
	Utilization  float64            `json:"utilization"`
	ResponseTime float64            `json:"responseTime"`
}

type NodeInfo struct {
	_id      primitive.ObjectID `json:"_id" bson:"_id"`
	Id       string             `json:"id"`
	Name     string             `json:"name"`
	City     string             `json:"city"`
	State    string             `json:"state"`
	Location []float64          `json:"location"`
	Type     string             `json:"type"`
	Topology string             `json:"topology"`
}

type LinkInfo struct {
	_id         primitive.ObjectID `json:"id" bson:"_id"`
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Node1Id     string             `json:"node1Id"`
	Node1Name   string             `json:"node1Name"`
	Node2Id     string             `json:"node2Id"`
	Node2Name   string             `json:"node2Name"`
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
	Time        interface{}        `json:"time"`
	FlowData    float64            `json:"flowData"`
	DailyGrowth float64            `json:"dailyGrowth"`
}

func GetSystemBasicInfo() *sysInfo {
	var systemBasicInfo sysInfo
	err = staticDatabase.Collection("basicInfo").FindOne(context.TODO(), bson.D{}).Decode(&systemBasicInfo)
	if err != nil {
		log.Fatal(err)
	}
	return &systemBasicInfo
}

func GetNetParameter() *NetParameter {
	var info NetParameter
	var option = options.FindOne()
	option.SetSort(bson.M{"recordTime": -1})
	err = dynamicDatabase.Collection("netParameter").FindOne(context.TODO(), bson.M{}).Decode(&info)
	if err != nil {
		return nil
	}
	return &info
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

func GetLinkList() []LinkInfo {
	var cursor *mongo.Cursor
	var infoList []LinkInfo
	cursor, err = staticDatabase.Collection("linkInfo").Find(context.TODO(), bson.M{})
	find(&infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

func GetNodeList() []NodeInfo {
	var cursor *mongo.Cursor
	var infoList []NodeInfo
	cursor, err = staticDatabase.Collection("nodeInfo").Find(context.TODO(), bson.M{})
	find(&infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

func GetNetworkFlow() map[string][]NetworkFlow {
	var cursor *mongo.Cursor
	var res []bson.M
	var flowStatistics = make(map[string][]NetworkFlow)
	var dayFlow, weekFlow, monthFlow []NetworkFlow

	var option = options.Find()
	option.SetSort(bson.M{"time": -1})
	option.SetLimit(30)
	cursor, err = dynamicDatabase.Collection("flowChange").Find(context.TODO(), bson.M{}, option)
	find(&dayFlow, cursor)
	flowStatistics["day"] = dayFlow

	option.SetLimit(7)
	cursor, err = dynamicDatabase.Collection("flowChange").Find(context.TODO(), bson.M{}, option)
	find(&weekFlow, cursor)
	flowStatistics["week"] = weekFlow

	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", bson.M{"$dateToString": bson.M{"format": "%Y-%m", "date": "$time"}}},
			{"count", bson.D{{"$sum", "$flowData"}}}},
		},
	}
	sortStage := bson.D{{"$sort", bson.D{{"_id", -1}}}}
	limitStage := bson.D{{"$limit", 12}}

	cursor, err = dynamicDatabase.Collection("flowChange").Aggregate(context.TODO(), mongo.Pipeline{groupStage, sortStage, limitStage})
	err = cursor.All(context.TODO(), &res)
	for _, item := range res {
		monthFlow = append(monthFlow, NetworkFlow{Time: item["_id"].(string), FlowData: item["count"].(float64), DailyGrowth: 0})
	}
	flowStatistics["month"] = monthFlow
	err = cursor.Close(context.TODO())
	return flowStatistics
}

func GetNodeTypeStatistics() map[string]int32 {
	var cursor *mongo.Cursor
	var res []bson.M
	var infoList = make(map[string]int32)
	groupStage := bson.D{
		{
			"$group", bson.D{
				{"_id", "$type"},
				{"count", bson.M{"$sum": 1}}},
		},
	}
	opt := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err = staticDatabase.Collection("nodeInfo").Aggregate(context.TODO(), mongo.Pipeline{groupStage}, opt)
	if err != nil {
		return nil
	}

	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return nil
	}

	for _, item := range res {
		infoList[item["_id"].(string)] = item["count"].(int32)
	}
	return infoList
}

func GetNodeAreaStatistics() map[string]int32 {
	var cursor *mongo.Cursor
	var res []bson.M
	var infoList = make(map[string]int32)
	groupStage := bson.D{
		{
			"$group", bson.D{{"_id", "$city"},
				{"count", bson.M{"$sum": 1}}},
		},
	}
	opt := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err = staticDatabase.Collection("nodeInfo").Aggregate(context.TODO(), mongo.Pipeline{groupStage}, opt)
	if err != nil {
		return nil
	}

	err = cursor.All(context.TODO(), &res)
	if err != nil {
		return nil
	}

	for _, item := range res {
		infoList[item["_id"].(string)] = item["count"].(int32)
	}
	return infoList
}
