## /auth

```text
首次验证
```

#### 接口状态

> 需修改

#### 接口URL

> /auth

#### 请求方式

> POST

#### Content-Type

> form-data

#### 请求Body参数

| 参数名   | 示例值   | 参数类型 | 是否必填 | 参数描述     |
| -------- | -------- | -------- | -------- | ------------ |
| username | 3network | Text     | 是       | 用于验证身份 |
| password | 3network | Text     | 是       | 用于验证身份 |

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"code": 0
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| ------ | ------ | -------- | -------- |
| code   | 0      | Text     | 成功     |

#### 成功响应示例

```javascript
{
	"code": 1
}
```

| 参数名 | 示例值 | 参数类型 | 参数描述 |
| ------ | ------ | -------- | -------- |
| code   | 1      | Text     | 失败     |

## /home

```text
主视图首屏
```

#### 接口状态

> 需修改

#### 接口URL

> /

#### 请求方式

> GET

#### Content-Type

> form-data

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
    "response": 0,
    "version": "v2.1",
    "code": "code71233",
    "state": {
        "database": 0,
        "network": 0,
        "cluster": 0
    },
    "people": {
        "principal": "牛猛",
        "duty": "娓奇"
    },
    "situation": {
        "handled": 100,
        "handling": 210,
        "unhandled": 100
    },
    "nodeList": [
        {
            "id": "Bj001",
            "name": "北京站",
            "city": "北京",
            "state": "正常",
            // 正常、繁忙、拥塞、故障
            "type": "业务上下站",
            // 光中继站、电中继站、城域汇聚站
            "Precision": "最高时钟基准",
            // 超高时钟基准、次高时钟基准、时钟网关、非时频网络节点
        },
        //...
    ],
    "nodeMessages": [
        {
            "id": "Bj001",
            "name": "北京",
            "message": "正常",
        },
        //...
    ],
    "flowChange": [
        {
            "time": "1613724232.828657",
            "flowData": "80",
            //PB
            "rate": 0.9
        },
        //...
    ]
}
```

| 参数名               | 示例值            | 参数类型 | 参数描述 |
| -------------------- | ----------------- | -------- | -------- |
| response             | 0                 | Text     | 请求成功 |
| version              | v2.1              | Text     |          |
| code                 | code71233         | Text     | 失败     |
| state                |                   | Text     |          |
| state.database       | 0                 | Text     |          |
| state.network        | 0                 | Text     |          |
| state.cluster        | 0                 | Text     |          |
| people               |                   | Text     |          |
| people.principal     | 牛猛              | Text     |          |
| people.duty          | 娓奇              | Text     |          |
| situation            |                   | Text     |          |
| situation.handled    | 100               | Text     |          |
| situation.handling   | 210               | Text     |          |
| situation.unhandled  | 100               | Text     |          |
| nodeList             |                   | Text     |          |
| nodeList.id          | Bj001             | Text     |          |
| nodeList.name        | 北京站            | Text     |          |
| nodeList.city        | 北京              | Text     |          |
| nodeList.state       | 正常              | Text     |          |
| nodeList.type        | 业务上下站        | Text     |          |
| nodeList.Precision   | 最高时钟基准      | Text     |          |
| nodeMessages         |                   | Text     |          |
| nodeMessages.id      | Bj001             | Text     |          |
| nodeMessages.name    | 北京              | Text     |          |
| nodeMessages.message | 正常              | Text     |          |
| flowChange           |                   | Text     |          |
| flowChange.time      | 1613724232.828657 | Text     |          |
| flowChange.flowData  | 80                | Text     |          |
| flowChange.rate      | 0.9               | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1,
}

```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /situationHandle

```text
主视图情况处理
```

#### 接口状态

> 需修改

#### 接口URL

> /situationHandle

#### 请求方式

> GET

#### Content-Type

> form-data

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"handled": [
		{
			"id": "312",
			"type": "",
			"handler": "",
			"handleTime": "",
			"handleState": "",
			"origin": "",
			"message": ""
		}
	],
	"handling": [
		{
			"id": "312",
			"type": "",
			"handler": "",
			"handleTime": "",
			"handleState": "",
			"origin": "",
			"message": ""
		}
	],
	"unhandled": [
		{
			"id": "312",
			"type": "",
			"handler": "",
			"handleTime": "",
			"handleState": "",
			"origin": "",
			"message": ""
		}
	]
}
```

| 参数名                | 示例值 | 参数类型 | 参数描述 |
| --------------------- | ------ | -------- | -------- |
| handled               |        | Text     |          |
| handled.id            | 312    | Text     |          |
| handled.type          |        | Text     |          |
| handled.handler       |        | Text     |          |
| handled.handleTime    |        | Text     |          |
| handled.handleState   |        | Text     |          |
| handled.origin        |        | Text     |          |
| handled.message       |        | Text     |          |
| handling              |        | Text     |          |
| handling.id           | 312    | Text     |          |
| handling.type         |        | Text     |          |
| handling.handler      |        | Text     |          |
| handling.handleTime   |        | Text     |          |
| handling.handleState  |        | Text     |          |
| handling.origin       |        | Text     |          |
| handling.message      |        | Text     |          |
| unhandled             |        | Text     |          |
| unhandled.id          | 312    | Text     |          |
| unhandled.type        |        | Text     |          |
| unhandled.handler     |        | Text     |          |
| unhandled.handleTime  |        | Text     |          |
| unhandled.handleState |        | Text     |          |
| unhandled.origin      |        | Text     |          |
| unhandled.message     |        | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /light

```text
光网络首屏
```

#### 接口状态

> 需修改

#### 接口URL

> /light

#### 请求方式

> GET

#### Content-Type

> form-data

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
    "response": 0,
    "version": "v2.1",
    "code": "code71233",
    "state": {
        "database": 0,
        "network": 0,
        "cluster": 0
    },
    "people": {
        "principal": "牛猛",
        "duty": "娓奇"
    },
    "situation": {
        "handled": 100,
        "handling": 210,
        "unhandled": 100
    },
    "nodeList": [
        {
            "id": "Bj001",
            "name": "北京站",
            "city": "北京",
            "state": "正常",
            // 正常、繁忙、拥塞、故障
            "type": "业务上下站",
            // 光中继站、电中继站、城域汇聚站
            "Precision": "最高时钟基准",
            // 超高时钟基准、次高时钟基准、时钟网关、非时频网络节点
        },
        //...
    ],
    "nodeMessages": [
        {
            "id": "Bj001",
            "name": "北京",
            "message": "正常",
        },
        //...
    ],
    "flowChange": [
        {
            "time": "1613724232.828657",
            "flowData": "80",
            //PB
            "rate": 0.9
        },
        //...
    ]
}
```

| 参数名                | 示例值 | 参数类型 | 参数描述 |
| --------------------- | ------ | -------- | -------- |
| response              | 0      | Text     | 请求成功 |
| siteNum               |        | Text     |          |
| siteNum.business      | 100    | Text     |          |
| siteNum.electric      | 50     | Text     |          |
| siteNum.light         | 40     | Text     |          |
| siteNum.error         | 1      | Text     |          |
| links                 |        | Text     |          |
| links.id              |        | Text     |          |
| links.name            |        | Text     |          |
| links.node1ID         |        | Text     |          |
| links.node1Name       |        | Text     |          |
| links.node2ID         |        | Text     |          |
| links.node2Name       |        | Text     |          |
| links.state           |        | Text     |          |
| bandMonitor           | 10     | Text     |          |
| bandSetting           |        | Text     |          |
| bandSetting.light     | MD01   | Text     |          |
| bandSetting.data      | MD03   | Text     |          |
| bandSetting.emergency | MD09   | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1,
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /lightNodesAndLinks

```text
光网络节点与链路
```

#### 接口状态

> 需修改

#### 接口URL

> /light/nodesAndLinks

#### 请求方式

> GET

#### Content-Type

> form-data

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
    "response": 0,
    "nodes": [
        {
            "id": "Bj001",
            "name": "北京站",
        },
        {
            "id": "Bj001",
            "name": "北京站",
        }
    ],
    "links": [
        {
            "id": "",
            "name": "",
            "node1ID": "",
            "node1Name": "",
            "node2ID": "",
            "node2Name": "",
            "contain": ""
        }
    ]
}
```

| 参数名          | 示例值 | 参数类型 | 参数描述 |
| --------------- | ------ | -------- | -------- |
| response        | 0      | Text     | 请求成功 |
| nodes           |        | Text     |          |
| nodes.id        | Bj001  | Text     |          |
| nodes.name      | 北京站 | Text     |          |
| links           |        | Text     |          |
| links.id        |        | Text     |          |
| links.name      |        | Text     |          |
| links.node1ID   |        | Text     |          |
| links.node1Name |        | Text     |          |
| links.node2ID   |        | Text     |          |
| links.node2Name |        | Text     |          |
| links.contain   |        | Text     |          |

#### 成功响应示例

```javascript
{
	"response": 1
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /bandSet

```text
光网络波段设置
```

#### 接口状态

> 需修改

#### 接口URL

> /bandSet

#### 请求方式

> POST

#### Content-Type

> form-data

#### 请求Body参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| ------ | ------ | -------- | -------- | -------- |
| type   | light  | Text     | 是       |          |
| begin  | MD01   | Text     | 是       |          |
| end    | MD05   | Text     | 是       |          |

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"response": 0
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 0      | Text     | 请求成功 |

#### 成功响应示例

```javascript
{
    "response": 1
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /time

```text
时频网络首屏
```

#### 接口状态

> 需修改

#### 接口URL

> /time

#### 请求方式

> GET

#### Content-Type

> form-data

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
    "response": 0,
    "siteNum": [
        100,
        120,
        150,
        200
    ],
    "sites": [
        {
            "id": "Bj001",
            "name": "北京站",
            "city": "北京",
            "state": "正常",
            // 正常、繁忙、拥塞、故障
            "type": "业务上下站",
            // 光中继站、电中继站、城域汇聚站
            "Precision": "最高时钟基准",
            // 超高时钟基准、次高时钟基准、时钟网关、非时频网络节点
            "error": "5ms",
        },
    ],
    "links": [
        {
            "id": "",
            "name": "",
            "node1ID": "",
            "node1Name": "",
            "node2ID": "",
            "node2Name": "",
            "band": "",
        }
    ]
}
```

| 参数名          | 示例值       | 参数类型 | 参数描述 |
| --------------- | ------------ | -------- | -------- |
| response        | 0            | Text     | 请求成功 |
| siteNum         | 100          | Text     |          |
| sites           |              | Text     |          |
| sites.id        | Bj001        | Text     |          |
| sites.name      | 北京站       | Text     |          |
| sites.city      | 北京         | Text     |          |
| sites.state     | 正常         | Text     |          |
| sites.type      | 业务上下站   | Text     |          |
| sites.Precision | 最高时钟基准 | Text     |          |
| sites.error     | 5ms          | Text     |          |
| links           |              | Text     |          |
| links.id        |              | Text     |          |
| links.name      |              | Text     |          |
| links.node1ID   |              | Text     |          |
| links.node1Name |              | Text     |          |
| links.node2ID   |              | Text     |          |
| links.node2Name |              | Text     |          |
| links.band      |              | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1,
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /data

```text
数据网络首屏
```

#### 接口状态

> 需修改

#### 接口URL

> /data

#### 请求方式

> GET

#### Content-Type

> form-data

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
    "response": 0,
    "bandUsed": "",
    "bandTotal": "",
    "linkNum": 9121,
    "nodeNum": 1000,
    "Timestamp": 1753,
    "Location": 13541,
    "nodes": [
        {
            "id": "SH001",
            "name": "上海站",
            "city": "上海",
            "state": "正常",
            "location": {
                "type": "Ponit",
                "coordinates": [
                    121.4648,
                    31.2891
                ]
            },
            "type": "业务上下站",
            "Precision": "最高时钟基准",
            "error": "5",
            "throughput": "900",
            "forwardingRate": "0.8",
            "Topology": "img/3.jpg"
        }
    ],
    "links": [
        {
            "id": "",
            "name": "",
            "node1ID": "",
            "node1Name": "",
            "node2ID": "",
            "node2Name": "",
        }
    ],
    "flowChange": [
        {
            "time": "1613724232.828657",
            "flowData": "80",
            //PB
        }
    ],
    "errorAlarm": [
        {
            "id": "",
            "rTime": "1613724232.828657",
            "type": "",
            "level": "",
            "msg": "",
            "state": "",
            "beginNode": "",
            "endNode": ""
        }
    ]
}
```

| 参数名               | 示例值            | 参数类型 | 参数描述 |
| -------------------- | ----------------- | -------- | -------- |
| response             | 0                 | Text     | 请求成功 |
| bandUsed             |                   | Text     |          |
| bandTotal            |                   | Text     |          |
| linkNum              | 9121              | Text     |          |
| nodeNum              | 1000              | Text     |          |
| Timestamp            | 1753              | Text     |          |
| Location             | 13541             | Text     |          |
| links                |                   | Text     |          |
| links.id             |                   | Text     |          |
| links.name           |                   | Text     |          |
| links.node1ID        |                   | Text     |          |
| links.node1Name      |                   | Text     |          |
| links.node2ID        |                   | Text     |          |
| links.node2Name      |                   | Text     |          |
| flowChange           |                   | Text     |          |
| flowChange.time      | 1613724232.828657 | Text     |          |
| flowChange.flowData  | 80                | Text     |          |
| errorAlarm           |                   | Text     |          |
| errorAlarm.id        |                   | Text     |          |
| errorAlarm.rTime     | 1613724232.828657 | Text     |          |
| errorAlarm.type      |                   | Text     |          |
| errorAlarm.level     |                   | Text     |          |
| errorAlarm.msg       |                   | Text     |          |
| errorAlarm.state     |                   | Text     |          |
| errorAlarm.beginNode |                   | Text     |          |
| errorAlarm.endNode   |                   | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /data/links

```text
光网络查询链路
```

#### 接口状态

> 需修改

#### 接口URL

> /data/links?linkId

#### 请求方式

> GET

#### Content-Type

> form-data

#### 请求Query参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| ------ | ------ | -------- | -------- | -------- |
| linkId |        | Text     | 是       |          |

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
    "response": 0,
    "id": "",
    "name": "",
    "node1Id": "",
    "node1Name": "",
    "node2Id": "",
    "node2Name": "",
    "band": "",
    "detail": [
        {
            "time": "1613724232.828657",
            "precisionError": 12,
            "loss": 0.01,
            "used": 0.59
        }
    ],
}
```

| 参数名                | 示例值            | 参数类型 | 参数描述 |
| --------------------- | ----------------- | -------- | -------- |
| response              | 0                 | Text     | 请求成功 |
| id                    |                   | Text     |          |
| name                  |                   | Text     |          |
| node1ID               |                   | Text     |          |
| node1Name             |                   | Text     |          |
| node2ID               |                   | Text     |          |
| node2Name             |                   | Text     |          |
| band                  |                   | Text     |          |
| precisionError        | 12                | Text     |          |
| loss                  | 0.01              | Text     |          |
| used                  | 0.59              | Text     |          |
| detail                |                   | Text     |          |
| detail.time           | 1613724232.828657 | Text     |          |
| detail.precisionError | 12                | Text     |          |
| detail.loss           | 0.01              | Text     |          |
| detail.used           | 0.59              | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /data/nodes

```text
光网络查询节点
```

#### 接口状态

> 需修改

#### 接口URL

> /data/nodes?nodeId=

#### 请求方式

> GET

#### Content-Type

> form-data

#### 请求Query参数

| 参数名 | 示例值 | 参数类型 | 是否必填 | 参数描述 |
| ------ | ------ | -------- | -------- | -------- |
| nodeId |        | Text     | 是       |          |

#### 预执行脚本

```javascript
暂无预执行脚本
```

#### 后执行脚本

```javascript
暂无后执行脚本
```

#### 成功响应示例

```javascript
{
	"response": 0,
	"siteNum": {
		"business": 100,
		"electric": 50,
		"light": 40,
		"error": 1
	},
	"links": [
		{
			"id": "",
			"name": "",
			"node1ID": "",
			"node1Name": "",
			"node2ID": "",
			"node2Name": "",
			"state": ""
		}
	],
	"bandMonitor": [
		10,
		12,
		13,
		14,
		50,
		11,
		12,
		13,
		14,
		19,
		20
	],
	"bandSetting": {
		"light": [
			"MD01",
			"MD02"
		],
		"data": [
			"MD03",
			"MD05"
		],
		"emergency": [
			"MD09",
			"MD10"
		]
	}
}
```

| 参数名         | 示例值 | 参数类型 | 参数描述 |
| -------------- | ------ | -------- | -------- |
| response       | 0      | Text     | 请求成功 |
| id             | Bj001  | Text     |          |
| name           | 北京站 | Text     |          |
| city           | 北京   | Text     |          |
| state          | 正常   | Text     |          |
| error          | 5ms    | Text     |          |
| throughput     | 1000   | Text     |          |
| forwardingRate | 0.9    | Text     |          |

#### 成功响应示例

```javascript
{
	"response": 1
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |