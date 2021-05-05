## 通用接口

### 全部节点

#### URL

> /nodeList

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json
{
    "nodeList": [
        {
            "id": "SH01",
            "name": "上海站",
            "city": "上海",
            "state": "正常",
            "location": [
                121.47824,
                31.236423
            ],
            "type": "业务上下站",
            "topology": "img/1.jpg"
        },
        {
            "id": "CC01",
            "name": "长春站",
            "city": "长春",
            "state": "正常",
            "location": [
                125.330606,
                43.820913
            ],
            "type": "业务上下站",
            "topology": "img/1.jpg"
        },
        {
            "id": "SY01",
            "name": "三亚站",
            "city": "三亚",
            "state": "繁忙",
            "location": [
                109.515685,
                18.262581
            ],
            "type": "业务上下站",
            "topology": "img/1.jpg"
        },
        {
            "id": "SJZ01",
            "name": "石家庄站",
            "city": "石家庄",
            "state": "故障",
            "location": [
                114.51463,
                38.046721
            ],
            "type": "业务上下站",
            "topology": "img/1.jpg"
        },
        {
            "id": "BJ01",
            "name": "北京站",
            "city": "北京",
            "state": "故障",
            "location": [
                116.403613,
                39.913802
            ],
            "type": "业务上下站",
            "topology": "img/1.jpg"
        }
    ],
    "response": 0
}
```

### 全部链路

#### URL

> /linkList

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json
{
    "linkList": [
        {
            "id": "BJ-SJZ-01",
            "name": "北京-石家庄-01",
            "node1Id": "BJ01",
            "node1Name": "北京-01",
            "node2Id": "SJZ01",
            "node2Name": "石家庄-01",
            "contain": 69,
            "band": "",
            "builded": true,
            "usedForTime": true,
            "usedForData": true
        }
    ],
    "response": 0
}
```

### WebSocket





## 主视图

### 系统信息

#### URL

> /nodeList

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json
{
    "systemName": "网络运行管理仿真系统",
    "version": "v2.8",
    "code": "code71233",
    "state": {
        "database": "health",
        "network": "health",
        "cluster": "health"
    },
    "people": {
        "Principal": {
            "name": "张三",
            "photo": "./"
        },
        "Duty": {
            "name": "李四",
            "photo": "./"
        }
    },
    "Band": {
        "Light": {
            "Begin": "MD01",
            "End": "MD02"
        },
        "Data": {
            "Begin": "MD03",
            "End": "MD05"
        },
        "Emergency": {
            "Begin": "MD06",
            "End": "MD10"
        }
    },
    "startTime": "2020-06-17T00:00:00Z"
}
```

### 

### 情况处理

#### URL

> /eventList

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```

### 

### 节点列表

> 同通用接口中的 “ 全部节点 ”

#### URL

> /nodeList

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```javascript

```

### 

### 基本信息

#### URL

> /netInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json
{
    "recordTime": "2021-05-04T09:00:02Z",
    "connectivity": 98.2,
    "throughput": 52001,
    "utilization": 67.2,
    "responseTime": 80.2
}
```

### 

### 流量信息

#### URL

> /flowInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```javascript
{
    "day": [
        {
            "time": "2021-05-04T01:41:33+08:00",
            "flowData": 80,
            "dailyGrowth": 0.12
        },
        {
            "time": "2021-05-03T01:43:00+08:00",
            "flowData": 90,
            "dailyGrowth": 0.2
        },
        {
            "time": "2021-05-02T01:44:26+08:00",
            "flowData": 90,
            "dailyGrowth": 0.1
        },
        {
            "time": "2021-05-01T01:45:53+08:00",
            "flowData": 110,
            "dailyGrowth": 0.23
        },
        {
            "time": "2021-04-30T01:47:19+08:00",
            "flowData": 100,
            "dailyGrowth": -0.231
        },
        {
            "time": "2021-04-29T01:48:45+08:00",
            "flowData": 70,
            "dailyGrowth": -0.2
        },
        {
            "time": "2021-04-28T01:47:19+08:00",
            "flowData": 120,
            "dailyGrowth": -0.231
        },
        {
            "time": "2021-04-27T01:47:19+08:00",
            "flowData": 110,
            "dailyGrowth": -0.231
        },
        {
            "time": "2021-04-26T01:47:19+08:00",
            "flowData": 100,
            "dailyGrowth": -0.231
        },
        {
            "time": "2021-04-25T01:47:19+08:00",
            "flowData": 110,
            "dailyGrowth": -0.231
        },
        {
            "time": "2021-04-24T01:47:19+08:00",
            "flowData": 102,
            "dailyGrowth": -0.231
        },
        {
            "time": "2021-04-23T01:47:19+08:00",
            "flowData": 120,
            "dailyGrowth": -0.231
        }
    ],
    "month": [
        {
            "time": "2021-05",
            "flowData": 260,
            "dailyGrowth": 0
        },
        {
            "time": "2021-04",
            "flowData": 942,
            "dailyGrowth": 0
        }
    ],
    "week": [
        {
            "time": "2021-05-04T01:41:33+08:00",
            "flowData": 80,
            "dailyGrowth": 0.12
        },
        {
            "time": "2021-05-03T01:43:00+08:00",
            "flowData": 90,
            "dailyGrowth": 0.2
        },
        {
            "time": "2021-05-02T01:44:26+08:00",
            "flowData": 90,
            "dailyGrowth": 0.1
        },
        {
            "time": "2021-05-01T01:45:53+08:00",
            "flowData": 110,
            "dailyGrowth": 0.23
        },
        {
            "time": "2021-04-30T01:47:19+08:00",
            "flowData": 100,
            "dailyGrowth": -0.231
        },
        {
            "time": "2021-04-29T01:48:45+08:00",
            "flowData": 70,
            "dailyGrowth": -0.2
        },
        {
            "time": "2021-04-28T01:47:19+08:00",
            "flowData": 120,
            "dailyGrowth": -0.231
        }
    ]
}
```



## 光网络视图

### 节点信息概览

#### URL

> /lightNetInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



### 波段报文数量查询

#### URL

> /lightNetBandMsg

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



### 运行配置设置

#### URL

> /lightNetBandSet

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



### 链路状态查询

#### URL

> /lightNetLinkInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



### 节点设备状态查询

#### URL

> /lightNetNodeInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



### 波段利用率查询

#### URL

> /lightNetBandUse

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



## 时频网络视图

### 网络概览

#### URL

> /timeNetInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



### 各基站精度查询

#### URL

> /timeNetNodeInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



### 拓扑图

#### URL

> /timeNetLinkInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json

```



## 数据网络视图

### 数据网络总览

#### URL

> /dataNetInfo

#### Method

> GET

#### Content-Type

> json

#### Example-Response

```json
{
    "bandUsed": 42,
    "bandTotal": 51,
    "linkNum": 4575,
    "nodeNum": 4,
    "timestampNum": 1915,
    "locationNum": 8114,
    "daliFlow": 100.2,
    "monthlyFlow": 5029.5
}
```



### 流量变化

> 同主视图中的 “流量信息” 





### 节点详情查询

#### URL

> /dataNetNodeInfo

#### Method

> POST

#### Content-Type

> json

#### Example-Response

```json
{
    "response": 0,
    "NodeDetail": {
        "name": "上海站",
        "precision": "超高时钟基准",
        "error": 5,
        "throughput": 1000,
        "forwardingRate": 0.9
    },
    "LinkStatistics": [
        {
            "SiteId": "SH01",
            "LinkName": "上海站-北京站",
            "NormalMessageNum": 0,
            "BusyMessageNum": 0,
            "ErrorMessageNum": 1
        },
        {
            "SiteId": "SH01",
            "LinkName": "上海站-长春站",
            "NormalMessageNum": 1,
            "BusyMessageNum": 1,
            "ErrorMessageNum": 0
        }
    ],
    "DateStatistics": [
        {
            "ReceiveDate": "1970-01-13",
            "SiteId": "SH01",
            "NormalMessageNum": 0,
            "BusyMessageNum": 0,
            "ErrorMessageNum": 1
        },
        {
            "ReceiveDate": "1970-01-19",
            "SiteId": "SH01",
            "NormalMessageNum": 1,
            "BusyMessageNum": 1,
            "ErrorMessageNum": 0
        }
    ],
    "OriginStatistics": {
        "北京站": 1,
        "长春站": 2
    }
}
```



### 链路详情查询

#### URL

> /dataNetLinkInfo

#### Method

> POST

#### Content-Type

> json

#### Example-Response

```json
{
    "response": 0,
    "basicInfo": {
        "id": "BJ-SJZ-01",
        "name": "北京-石家庄-01",
        "node1Id": "BJ01",
        "node1Name": "北京-01",
        "node2Id": "SJZ01",
        "node2Name": "石家庄-01",
        "contain": 69,
        "band": "",
        "builded": true,
        "usedForTime": true,
        "usedForData": true
    },
    "parameterChange": [
        {
            "id": "BJ-SJZ-01",
            "time": "1970-01-22T16:15:24Z",
            "precisionError": 12,
            "loss": 0.01,
            "used": 0.59
        },
        {
            "id": "BJ-SJZ-01",
            "time": "1970-01-21T16:15:24Z",
            "precisionError": 12,
            "loss": 0.01,
            "used": 0.59
        },
        {
            "id": "BJ-SJZ-01",
            "time": "1970-01-20T16:15:24Z",
            "precisionError": 12,
            "loss": 0.01,
            "used": 0.59
        },
        {
            "id": "BJ-SJZ-01",
            "time": "1970-01-19T16:15:24Z",
            "precisionError": 12,
            "loss": 0.01,
            "used": 0.59
        }
    ],
    "originStatistics": {
        "三亚站": 2,
        "长春站": 1
    }
}
```

