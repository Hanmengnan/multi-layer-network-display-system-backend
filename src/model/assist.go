package model

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math/rand"
	"time"
)

type graphNode struct {
	_id  primitive.ObjectID `json:"__id" bson:"__id"`
	Id   string             `json:"id" bson:"id"`
	Name string             `json:"name" bson:"name"`
}

type graphLink struct {
	_id       primitive.ObjectID `json:"id" bson:"_id"`
	Id        string             `json:"id" bson:"id"`
	Name      string             `json:"name"`
	Node1Id   string             `json:"node1Id"`
	Node1Name string             `json:"node1Name"`
	Node2Id   string             `json:"node2Id"`
	Node2Name string             `json:"node2Name"`
}

var CP_KV = map[string]string{"合肥": "安徽", "芜湖": "安徽", "蚌埠": "安徽", "淮南": "安徽", "马鞍山": "安徽", "淮北": "安徽", "铜陵": "安徽", "安庆": "安徽", "黄山": "安徽", "滁州": "安徽", "阜阳": "安徽", "宿州": "安徽", "巢湖": "安徽", "六安": "安徽", "亳州": "安徽", "池州": "安徽", "宣城": "安徽", "北京": "北京", "东城区": "北京", "西城区": "北京", "崇文区": "北京", "宣武区": "北京", "朝阳区": "北京", "丰台区": "北京", "石景山区": "北京", "海淀区": "北京", "门头沟区": "北京", "房山区": "北京", "通州区": "北京", "顺义区": "北京", "昌平区": "北京", "大兴区": "北京", "怀柔区": "北京", "平谷区": "北京", "密云县": "北京", "延庆县": "北京", "重庆": "重庆", "万州区": "重庆", "涪陵区": "重庆", "渝中区": "重庆", "大渡口区": "重庆", "江北区": "重庆", "沙坪坝区": "重庆", "九龙坡区": "重庆", "南岸区": "重庆", "北碚区": "重庆", "万盛区": "重庆", "双桥区": "重庆", "渝北区": "重庆", "巴南区": "重庆", "黔江区": "重庆", "长寿区": "重庆", "綦江县": "重庆", "潼南县": "重庆", "铜梁县": "重庆", "大足县": "重庆", "荣昌县": "重庆", "璧山县": "重庆", "梁平县": "重庆", "城口县": "重庆", "丰都县": "重庆", "垫江县": "重庆", "武隆县": "重庆", "忠县": "重庆", "开县": "重庆", "云阳县": "重庆", "奉节县": "重庆", "巫山县": "重庆", "巫溪县": "重庆", "石柱土家族自治县": "重庆", "秀山土家族苗族自治县": "重庆", "酉阳土家族苗族自治县": "重庆", "彭水苗族土家族自治县": "重庆", "江津市": "重庆", "合川市": "重庆", "永川市": "重庆", "南川市": "重庆", "福州": "福建", "厦门": "福建", "莆田": "福建", "三明": "福建", "泉州": "福建", "漳州": "福建", "南平": "福建", "龙岩": "福建", "宁德": "福建", "兰州": "甘肃", "嘉峪关": "甘肃", "金昌": "甘肃", "白银": "甘肃", "天水": "甘肃", "武威": "甘肃", "张掖": "甘肃", "平凉": "甘肃", "酒泉": "甘肃", "庆阳": "甘肃", "定西": "甘肃", "陇南": "甘肃", "临夏": "甘肃", "甘南": "甘肃", "广州": "广东", "韶关": "广东", "深圳": "广东", "珠海": "广东", "汕头": "广东", "佛山": "广东", "江门": "广东", "湛江": "广东", "茂名": "广东", "肇庆": "广东", "惠州": "广东", "梅州": "广东", "汕尾": "广东", "河源": "广东", "阳江": "广东", "清远": "广东", "东莞": "广东", "中山": "广东",
	"潮州": "广东", "揭阳": "广东", "云浮": "广东", "南宁": "广西", "柳州": "广西", "桂林": "广西", "梧州": "广西", "北海": "广西", "防城港": "广西", "钦州": "广西", "贵港": "广西", "玉林": "广西", "百色": "广西", "贺州": "广西", "河池": "广西", "贵阳": "贵州", "六盘水": "贵州", "遵义": "贵州", "安顺": "贵州", "铜仁": "贵州", "黔西南": "贵州", "毕节": "贵州", "黔东南": "贵州", "黔南": "贵州", "海口": "海南", "三亚": "海南", "其他": "台湾", "石家庄": "河北", "唐山": "河北", "秦皇岛": "河北", "邯郸": "河北", "邢台": "河北", "保定": "河北", "张家口": "河北", "承德": "河北", "沧州": "河北", "廊坊": "河北", "衡水": "河北", "哈尔滨": "黑龙江", "齐齐哈尔": "黑龙江", "鸡西": "黑龙江", "鹤岗": "黑龙江", "双鸭山": "黑龙江", "大庆": "黑龙江", "伊春": "黑龙江", "佳木斯": "黑龙江", "七台河": "黑龙江", "牡丹江": "黑龙江", "黑河": "黑龙江", "绥化": "黑龙江", "大兴安岭": "黑龙江", "郑州": "河南", "开封": "河南", "洛阳": "河南", "平顶山": "河南", "安阳": "河南", "鹤壁": "河南", "新乡": "河南", "焦作": "河南", "濮阳": "河南", "许昌": "河南", "漯河": "河南", "三门峡": "河南", "南阳": "河南", "商丘": "河南", "信阳": "河南", "周口": "河南", "驻马店": "河南", "武汉": "湖北", "黄石": "湖北", "十堰": "湖北", "宜昌": "湖北", "襄樊": "湖北", "鄂州": "湖北", "荆门": "湖北", "孝感": "湖北", "荆州": "湖北", "黄冈": "湖北", "咸宁": "湖北", "随州": "湖北", "恩施土家族苗族自治州": "湖北", "长沙": "湖南", "株洲": "湖南", "湘潭": "湖南", "衡阳": "湖南", "邵阳": "湖南", "岳阳": "湖南", "常德": "湖南", "张家界": "湖南", "益阳": "湖南", "郴州": "湖南", "永州": "湖南", "怀化": "湖南", "娄底": "湖南", "湘西土家族苗族自治州": "湖南", "呼和浩特": "内蒙古", "包头": "内蒙古", "乌海": "内蒙古", "赤峰": "内蒙古", "通辽": "内蒙古", "鄂尔多斯": "内蒙古", "呼伦贝尔": "内蒙古", "兴安盟": "内蒙古", "锡林郭勒盟": "内蒙古", "乌兰察布盟": "内蒙古", "巴彦淖尔盟": "内蒙古", "阿拉善盟": "内蒙古", "南京": "江苏", "无锡": "江苏", "徐州": "江苏", "常州": "江苏", "苏州": "江苏", "南通": "江苏", "连云港": "江苏", "淮安": "江苏", "盐城": "江苏", "扬州": "江苏", "镇江": "江苏", "泰州": "江苏", "宿迁": "江苏", "南昌": "江西", "景德镇": "江西", "萍乡": "江西", "九江": "江西", "新余": "江西", "鹰潭": "江西", "赣州": "江西", "吉安": "江西", "宜春": "江西", "抚州": "江西", "上饶": "江西", "长春": "吉林", "吉林": "吉林", "四平": "吉林", "辽源": "吉林", "通化": "吉林", "白山": "吉林", "松原": "吉林", "白城": "吉林", "延边朝鲜族自治州": "吉林", "沈阳": "辽宁", "大连": "辽宁", "鞍山": "辽宁", "抚顺": "辽宁", "本溪": "辽宁", "丹东": "辽宁", "锦州": "辽宁", "营口": "辽宁", "阜新": "辽宁", "辽阳": "辽宁", "盘锦": "辽宁", "铁岭": "辽宁", "朝阳": "辽宁", "葫芦岛": "辽宁", "银川": "宁夏", "石嘴山": "宁夏", "吴忠": "宁夏", "固原": "宁夏", "西宁": "青海", "海东": "青海", "海北": "青海", "黄南": "青海", "海南": "青海", "果洛": "青海", "玉树": "青海", "海西": "青海", "太原": "山西", "大同": "山西", "阳泉": "山西", "长治": "山西", "晋城": "山西", "朔州": "山西", "晋中": "山西", "运城": "山西", "忻州": "山西", "临汾": "山西", "吕梁": "山西", "济南": "山东", "青岛": "山东", "淄博": "山东", "枣庄": "山东", "东营": "山东", "烟台": "山东", "潍坊": "山东", "济宁": "山东", "泰安": "山东", "威海": "山东", "日照": "山东", "莱芜": "山东", "临沂": "山东", "德州": "山东", "聊城": "山东", "滨州": "山东", "菏泽": "山东", "上海": "上海", "黄浦区": "上海", "卢湾区": "上海", "徐汇区": "上海", "长宁区": "上海", "静安区": "上海", "普陀区": "上海", "闸北区": "上海", "虹口区": "上海", "杨浦区": "上海", "闵行区": "上海", "宝山区": "上海", "嘉定区": "上海", "浦东新区": "上海", "金山区": "上海", "松江区": "上海", "青浦区": "上海", "南汇区": "上海", "奉贤区": "上海", "崇明县": "上海", "成都": "四川", "自贡": "四川", "攀枝花": "四川", "泸州": "四川", "德阳": "四川", "绵阳": "四川", "广元": "四川", "遂宁": "四川", "内江": "四川", "乐山": "四川", "南充": "四川", "眉山": "四川", "宜宾": "四川", "广安": "四川", "达州": "四川", "雅安": "四川", "巴中": "四川", "资阳": "四川", "阿坝": "四川", "甘孜": "四川", "凉山": "四川", "天津": "天津", "和平区": "天津", "河东区": "天津", "河西区": "天津", "南开区": "天津", "河北区": "天津", "红桥区": "天津", "塘沽区": "天津", "汉沽区": "天津", "大港区": "天津", "东丽区": "天津", "西青区": "天津", "津南区": "天津", "北辰区": "天津", "武清区": "天津", "宝坻区": "天津", "宁河县": "天津", "静海县": "天津", "蓟县": "天津", "拉萨": "西藏", "昌都": "西藏", "山南": "西藏", "日喀则": "西藏", "那曲": "西藏", "阿里": "西藏", "林芝": "西藏", "乌鲁木齐": "新疆", "克拉玛依": "新疆", "吐鲁番": "新疆", "哈密": "新疆", "昌吉": "新疆", "博尔塔拉": "新疆", "巴音郭楞": "新疆", "阿克苏": "新疆", "克孜勒苏": "新疆", "喀什": "新疆", "和田": "新疆", "伊犁": "新疆", "塔城": "新疆", "阿勒泰": "新疆", "昆明": "云南", "曲靖": "云南", "玉溪": "云南", "保山": "云南", "昭通": "云南", "楚雄": "云南", "红河": "云南", "文山": "云南", "思茅": "云南", "西双版纳": "云南", "大理": "云南", "德宏": "云南", "丽江": "云南", "怒江": "云南", "迪庆": "云南", "临沧": "云南", "杭州": "浙江", "宁波": "浙江", "温州": "浙江", "嘉兴": "浙江", "湖州": "浙江", "绍兴": "浙江", "金华": "浙江", "衢州": "浙江", "舟山": "浙江", "台州": "浙江", "丽水": "浙江", "西安": "陕西", "铜川": "陕西", "宝鸡": "陕西", "咸阳": "陕西", "渭南": "陕西", "延安": "陕西", "汉中": "陕西", "榆林": "陕西", "安康": "陕西", "商洛": "陕西", "台北": "台湾", "高雄": "台湾", "香港": "香港", "澳门": "澳门"}

var (
	graphNodeList      []graphNode
	nodeHashMap        map[string]bool
	backtrack          []string
	startNode, endNode string
	NewDataChannel     = make(chan string, 1)
)

func dfs(nodeName string, aimNode string) {
	if nodeName == aimNode {
		_, err = staticDatabase.Collection("graph").UpdateOne(context.TODO(), bson.M{"start": startNode, "end": endNode}, bson.M{"$push": bson.M{"route": backtrack}})
		if err != nil {
			log.Printf("%v", err)
		}
		return
	}

	var cursor *mongo.Cursor
	tmp := make([]graphLink, 0)
	cursor, err = staticDatabase.Collection("linkInfo").Find(context.TODO(), bson.D{{"$or", []interface{}{bson.M{"node1Name": nodeName}, bson.M{"node2Name": nodeName}}}})

	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	find(&tmp, cursor)
	for _, nextNode := range tmp {
		nextNodeName := map[bool]string{true: nextNode.Node2Name, false: nextNode.Node1Name}[nextNode.Node1Name == nodeName]
		if nodeHashMap[nextNodeName] == false {

			nodeHashMap[nextNodeName] = true
			backtrack = append(backtrack, nextNodeName)

			dfs(nextNodeName, aimNode)

			nodeHashMap[nextNodeName] = false
			backtrack = backtrack[:len(backtrack)-1]

		}

	}
	err = cursor.Close(context.TODO())
	return
}

func GraphGenerate() {
	var cursor *mongo.Cursor
	staticDatabase.Collection("graph").Drop(context.TODO())
	cursor, err = staticDatabase.Collection("nodeInfo").Find(context.TODO(), bson.M{})
	find(&graphNodeList, cursor)

	nodeHashMap = make(map[string]bool)
	for _, item := range graphNodeList {
		nodeHashMap[item.Name] = false
	}

	backtrack = make([]string, 0)

	for _, start := range graphNodeList {
		for _, end := range graphNodeList {
			_, err = staticDatabase.Collection("graph").InsertOne(context.TODO(), bson.M{"start": start.Id, "end": end.Id, "route": []interface{}{}})
			if err != nil {
				log.Println(err)
			}
			if start.Name != end.Name {
				nodeHashMap[start.Name] = true
				backtrack = append(backtrack, start.Name)

				startNode, endNode = start.Id, end.Id
				dfs(start.Name, end.Name)

				nodeHashMap[start.Name] = false
				backtrack = backtrack[:len(backtrack)-1]
			}
		}
		err = cursor.Close(context.TODO())
	}
	return
}

func InjectNewData() {
	var cursor *mongo.Cursor
	var nodes []struct {
		_id primitive.ObjectID
		Id  string
	}
	var links []struct {
		_id primitive.ObjectID
		Id  string
	}
	cursor, _ = staticDatabase.Collection("nodeInfo").Find(context.TODO(), bson.M{})
	find(&nodes, cursor)
	cursor, _ = staticDatabase.Collection("linkInfo").Find(context.TODO(), bson.M{})
	find(&links, cursor)
	for {
		choice := rand.Intn(5)
		//choice := 1
		switch choice {
		case 0:
			dynamicDatabase.Collection("netParameter").InsertOne(context.TODO(), bson.M{
				"recordTime":   time.Now(),
				"connectivity": rand.Float64()*10 + 90,
				"throughput":   52001,
				"utilization":  rand.Float64()*50 + 40,
				"responseTime": rand.Float64() * 200,
			})
			NewDataChannel <- "parameterChange"
		case 1:
			var tmp NodeInfo
			nodeNum, _ := staticDatabase.Collection("nodeInfo").CountDocuments(context.TODO(), bson.M{})
			skipNum := rand.Intn(int(nodeNum))
			option := options.FindOne()
			option.SetSkip(int64(skipNum))
			err = staticDatabase.Collection("nodeInfo").FindOne(context.TODO(), bson.M{}, option).Decode(&tmp)
			staticDatabase.Collection("nodeInfo").FindOneAndUpdate(context.TODO(), bson.M{"id": tmp.Id}, bson.M{"$set": bson.M{
				"state":          ([]string{"正常", "繁忙", "故障"})[rand.Intn(3)],
				"error":          rand.Intn(20),
				"forwardingRate": rand.Float64()*30 + 70,
			}})
			NewDataChannel <- "nodeList"
			NewDataChannel <- "nodeDetail"
		case 2:
			NewDataChannel <- "situation"
		case 3:
			dynamicDatabase.Collection("linkParameter").InsertOne(context.TODO(), bson.M{
				"id":             links[rand.Intn(len(links))].Id,
				"time":           time.Now(),
				"precisionError": rand.Intn(20),
				"loss":           rand.Float64() * 10,
				"used":           rand.Float64()*50 + 50,
			})
			NewDataChannel <- "nodeOverLoad"
		default:
			NewDataChannel <- "nodeDetail"
		}
		time.Sleep(time.Second * 1)
	}
}
