package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sort"
	"time"
)

type DataNetworkBasicInfo struct {
	_id         primitive.ObjectID `json:"id" bson:"_id"`
	BandUsed    int                `json:"bandUsed"`
	BandTotal   int                `json:"bandTotal"`
	LinkNum     int                `json:"linkNum"`
	NodeNum     int                `json:"nodeNum"`
	Timestamp   int                `json:"timestampNum"`
	Location    int                `json:"locationNum"`
	DaliFlow    float64            `json:"daliFlow"`
	MonthlyFlow float64            `json:"monthlyFlow"`
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

type ParameterChange struct {
	_id            primitive.ObjectID `json:"id" bson:"_id"`
	Id             string             `json:"id"`
	Time           time.Time          `json:"time"`
	PrecisionError float64            `json:"precisionError"`
	Loss           float64            `json:"loss"`
	Used           float64            `json:"used"`
}

type NodeDetail struct {
	_id            primitive.ObjectID `json:"_id" bson:"_id"`
	Name           string             `json:"name"`
	Precision      string             `json:"precision"`
	Error          int                `json:"error"`
	Throughput     int                `json:"throughput"`
	ForwardingRate float64            `json:"forwardingRate"`
}

type MessageDateStatistics struct {
	ReceiveDate string
	SiteId      string
	//SiteName         string
	NormalMessageNum int32
	BusyMessageNum   int32
	ErrorMessageNum  int32
}

type MessageLinkStatistics struct {
	SiteId           string
	LinkName         string
	NormalMessageNum int32
	BusyMessageNum   int32
	ErrorMessageNum  int32
}

func GetDataNetworkBasicInfo() *DataNetworkBasicInfo {
	var info DataNetworkBasicInfo
	res := staticDatabase.Collection("dataNetworkBasicInfo").FindOne(context.TODO(), bson.M{})
	err = res.Decode(&info)
	return &info
}

func GetLinkBasicInfo(linkId string) LinkInfo {
	var info LinkInfo
	err = staticDatabase.Collection("linkInfo").FindOne(context.TODO(), bson.M{"id": linkId}).Decode(&info)
	return info
}

func GetLinkParameterChange(linkId string) []ParameterChange {
	var cursor *mongo.Cursor
	var infoList []ParameterChange
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"time", -1}})
	findOptions.SetLimit(30 * 24)

	cursor, err = dynamicDatabase.Collection("linkParameter").Find(context.TODO(), bson.M{"id": linkId}, findOptions)
	find(&infoList, cursor)
	err = cursor.Close(context.TODO())
	return infoList
}

func GetLinkMessageOriginStatistics(linkId string) map[string]int32 {
	infoList := make(map[string]int32)
	var info LinkInfo
	err = staticDatabase.Collection("linkInfo").FindOne(context.TODO(), bson.M{"id": linkId}).Decode(&info)

	statistics1 := GetNodeMessageOriginStatistics(info.Node1Id)
	statistics2 := GetNodeMessageOriginStatistics(info.Node2Id)

	for key, value := range statistics1 {
		if _, ok := infoList[key]; !ok {
			infoList[key] = infoList[key] + value
		} else {
			infoList[key] = value
		}
	}
	for key, value := range statistics2 {
		if _, ok := infoList[key]; !ok {
			infoList[key] = infoList[key] + value
		} else {
			infoList[key] = value
		}
	}
	return infoList
}

func GetNodeDetail(nodeId string) NodeDetail {
	var info NodeDetail
	err = staticDatabase.Collection("nodeInfo").FindOne(context.TODO(), bson.M{"id": nodeId}).Decode(&info)
	return info
}

func GetNodeMessageLinkStatistics(nodeId string) []MessageLinkStatistics {
	var infoList []MessageLinkStatistics
	var info NodeInfo
	var res []bson.M
	var cursor *mongo.Cursor

	matchStage := bson.D{
		{"$match", bson.D{{"nowSite", nodeId}}},
	}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", bson.D{{"from", "$fromSite"}, {"type", "$type"}}},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}

	sortStage := bson.D{{"$sort", bson.D{{"count", -1}}}}
	limitStage := bson.D{{"$limit", 5}}

	opts := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err = dynamicDatabase.Collection("nodeMessage").Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage, sortStage, limitStage}, opts)
	err = cursor.All(context.TODO(), &res)

	err = staticDatabase.Collection("nodeInfo").FindOne(context.TODO(), bson.M{"id": nodeId}).Decode(&info)
	nowNodeName := info.Name

	tmp := make(map[string]map[string]int32)

	for _, item := range res {
		from := item["_id"].(primitive.M)["from"].(string)
		msgType := item["_id"].(primitive.M)["type"].(string)
		count := item["count"].(int32)
		if _, ok := tmp[from]; !ok {
			tmp[from] = map[string]int32{"正常": 0, "繁忙": 0, "故障": 0}
		}
		tmp[from][msgType] = count
	}
	for key, value := range tmp {
		err = staticDatabase.Collection("nodeInfo").FindOne(context.TODO(), bson.M{"id": key}).Decode(&info)
		infoList = append(infoList, MessageLinkStatistics{nodeId, nowNodeName + "-" + info.Name, value["正常"], value["繁忙"], value["故障"]})
	}
	return infoList
}

func GetNodeMessageDateStatistics(nodeId string) []MessageDateStatistics {
	var infoList []MessageDateStatistics
	var res []bson.M
	var cursor *mongo.Cursor

	matchStage := bson.D{
		{"$match", bson.D{{"nowSite", nodeId}}},
	}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", bson.D{
				{"date", bson.D{{"$dateToString", bson.D{{"format", "%Y-%m-%d"}, {"date", "$rTime"}}}}},
				{"type", "$type"},
			}},
			{"count", bson.D{
				{"$sum", 1},
			}},
		}},
	}

	sortStage := bson.D{{"$sort", bson.D{{"_id.date", -1}}}}
	limitStage := bson.D{{"$limit", 30}}

	opts := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err = dynamicDatabase.Collection("nodeMessage").Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage, sortStage, limitStage}, opts)
	err = cursor.All(context.TODO(), &res)

	tmp := make(map[string]map[string]int32)

	for _, item := range res {
		date := item["_id"].(primitive.M)["date"].(string)
		msgType := item["_id"].(primitive.M)["type"].(string)
		count := item["count"].(int32)
		if _, ok := tmp[date]; !ok {
			tmp[date] = map[string]int32{"正常": 0, "繁忙": 0, "故障": 0}
		}
		tmp[date][msgType] = count
	}
	for key, value := range tmp {
		infoList = append(infoList, MessageDateStatistics{key, nodeId, value["正常"], value["繁忙"], value["故障"]})
	}
	sort.Slice(infoList, func(i, j int) bool {
		return infoList[i].ReceiveDate < infoList[j].ReceiveDate
	})
	return infoList
}

func GetNodeMessageOriginStatistics(nodeId string) map[string]int32 {
	var cursor *mongo.Cursor
	var info NodeInfo
	infoList := make(map[string]int32)
	var res []bson.M

	matchStage := bson.D{
		{"$match", bson.D{{"nowSite", nodeId}}},
	}
	groupStage := bson.D{
		{"$group", bson.D{
			{"_id", "$originSite"},
			{"count", bson.D{
				{"$sum", 1}},
			},
		}},
	}
	sortStage := bson.D{{"$sort", bson.D{{"count", -1}}}}
	limitStage := bson.D{{"$limit", 7}}

	opts := options.Aggregate().SetMaxTime(5 * time.Second)
	cursor, err = dynamicDatabase.Collection("nodeMessage").Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage, sortStage, limitStage}, opts)
	err = cursor.All(context.TODO(), &res)

	for _, item := range res {
		err = staticDatabase.Collection("nodeInfo").FindOne(context.TODO(), bson.M{"id": item["_id"]}).Decode(&info)
		infoList[info.Name] = item["count"].(int32)
	}
	return infoList
}
