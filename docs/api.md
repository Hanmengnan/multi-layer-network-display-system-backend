## /

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
	}
}
```

| 参数名              | 示例值    | 参数类型 | 参数描述 |
| ------------------- | --------- | -------- | -------- |
| response            | 0         | Text     | 请求成功 |
| version             | v2.1      | Text     |          |
| code                | code71233 | Text     |          |
| state               |           | Text     |          |
| state.database      | 0         | Text     |          |
| state.network       | 0         | Text     |          |
| state.cluster       | 0         | Text     |          |
| people              |           | Text     |          |
| people.principal    | 牛猛      | Text     |          |
| people.duty         | 娓奇      | Text     |          |
| situation           |           | Text     |          |
| situation.handled   | 100       | Text     |          |
| situation.handling  | 210       | Text     |          |
| situation.unhandled | 100       | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1,
}

```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /flow

```text
暂无描述
```

#### 接口状态

> 需修改

#### 接口URL

> /flow

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
	"flowChange": [
		{
			"time": "1613724232.828657",
			"flowData": "80",
			"rate": 0.9
		}
	]
}
```

| 参数名              | 示例值            | 参数类型 | 参数描述 |
| ------------------- | ----------------- | -------- | -------- |
| response            | 0                 | Text     | 请求成功 |
| flowChange          |                   | Text     |          |
| flowChange.time     | 1613724232.828657 | Text     |          |
| flowChange.flowData | 80                | Text     |          |
| flowChange.rate     | 0.9               | Text     |          |

## /nodes

```text
获取全部节点
```

#### 接口状态

> 需修改

#### 接口URL

> /nodes

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
	]
}
```

| 参数名                     | 示例值       | 参数类型 | 参数描述 |
| -------------------------- | ------------ | -------- | -------- |
| response                   | 0            | Text     | 请求成功 |
| nodes                      |              | Text     |          |
| nodes.id                   | SH001        | Text     |          |
| nodes.name                 | 上海站       | Text     |          |
| nodes.city                 | 上海         | Text     |          |
| nodes.state                | 正常         | Text     |          |
| nodes.location             |              | Text     |          |
| nodes.location.type        | Ponit        | Text     |          |
| nodes.location.coordinates | 121.4648     | Text     |          |
| nodes.type                 | 业务上下站   | Text     |          |
| nodes.Precision            | 最高时钟基准 | Text     |          |
| nodes.error                | 5            | Text     |          |
| nodes.throughput           | 900          | Text     |          |
| nodes.forwardingRate       | 0.8          | Text     |          |
| nodes.Topology             | img/3.jpg    | Text     |          |

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
	"siteNum": {
		"business": 100,
		"electric": 50,
		"light": 40,
		"error": 1
	},
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

| 参数名                | 示例值 | 参数类型 | 参数描述 |
| --------------------- | ------ | -------- | -------- |
| response              | 0      | Text     | 请求成功 |
| siteNum               |        | Text     |          |
| siteNum.business      | 100    | Text     |          |
| siteNum.electric      | 50     | Text     |          |
| siteNum.light         | 40     | Text     |          |
| siteNum.error         | 1      | Text     |          |
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

## /light/links

```text
光网络节点与链路
```

#### 接口状态

> 需修改

#### 接口URL

> /light/links

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

## /light/bandNum

```text
暂无描述
```

#### 接口状态

> 需修改

#### 接口URL

> light/bandNum

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
	]
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 0      | Text     | 请求成功 |
| siteNum  | 100    | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1,
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /time/links

```text
暂无描述
```

#### 接口状态

> 需修改

#### 接口URL

> 未填写

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
	"links": [
		{
			"id": "",
			"name": "",
			"node1ID": "",
			"node1Name": "",
			"node2ID": "",
			"node2Name": "",
			"band": ""
		}
	]
}
```

| 参数名          | 示例值 | 参数类型 | 参数描述 |
| --------------- | ------ | -------- | -------- |
| response        | 0      | Text     | 请求成功 |
| links           |        | Text     |          |
| links.id        |        | Text     |          |
| links.name      |        | Text     |          |
| links.node1ID   |        | Text     |          |
| links.node1Name |        | Text     |          |
| links.node2ID   |        | Text     |          |
| links.node2Name |        | Text     |          |
| links.band      |        | Text     |          |

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
	"Location": 13541
}
```

| 参数名    | 示例值 | 参数类型 | 参数描述 |
| --------- | ------ | -------- | -------- |
| response  | 0      | Text     | 请求成功 |
| bandUsed  |        | Text     |          |
| bandTotal |        | Text     |          |
| linkNum   | 9121   | Text     |          |
| nodeNum   | 1000   | Text     |          |
| Timestamp | 1753   | Text     |          |
| Location  | 13541  | Text     |          |

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

> /data/links

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
	"links": [
		{
			"id": "",
			"name": "",
			"node1ID": "",
			"node1Name": "",
			"node2ID": "",
			"node2Name": "",
			"band": "",
			"state": "",
			"error": "",
			"loss": ""
		}
	]
}
```

| 参数名          | 示例值 | 参数类型 | 参数描述 |
| --------------- | ------ | -------- | -------- |
| response        | 0      | Text     | 请求成功 |
| links           |        | Text     |          |
| links.id        |        | Text     |          |
| links.name      |        | Text     |          |
| links.node1ID   |        | Text     |          |
| links.node1Name |        | Text     |          |
| links.node2ID   |        | Text     |          |
| links.node2Name |        | Text     |          |
| links.band      |        | Text     |          |
| links.state     |        | Text     |          |
| links.error     |        | Text     |          |
| links.loss      |        | Text     |          |

#### 成功响应示例

```javascript
{
    "response": 1
}
```

| 参数名   | 示例值 | 参数类型 | 参数描述 |
| -------- | ------ | -------- | -------- |
| response | 1      | Text     | 请求失败 |

## /data/flow

```text
暂无描述
```

#### 接口状态

> 需修改

#### 接口URL

> data/flow

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
	"flowChange": [
		{
			"time": "1613724232.828657",
			"flowData": "80",
			"rate": 0.9
		}
	]
}
```

| 参数名              | 示例值            | 参数类型 | 参数描述 |
| ------------------- | ----------------- | -------- | -------- |
| response            | 0                 | Text     | 请求成功 |
| flowChange          |                   | Text     |          |
| flowChange.time     | 1613724232.828657 | Text     |          |
| flowChange.flowData | 80                | Text     |          |
| flowChange.rate     | 0.9               | Text     |          |

## /data/node

```text
暂无描述
```

#### 接口状态

> 需修改

#### 接口URL

> /data/node?nodeId=

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

## /data/link

```text
暂无描述
```

#### 接口状态

> 需修改

#### 接口URL

> data/link?linkId=

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

## /dataErrorAlarm

```text
暂无描述
```

#### 接口状态

> 需修改

#### 接口URL

> data/errorAlarm

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
| errorAlarm           |                   | Text     |          |
| errorAlarm.id        |                   | Text     |          |
| errorAlarm.rTime     | 1613724232.828657 | Text     |          |
| errorAlarm.type      |                   | Text     |          |
| errorAlarm.level     |                   | Text     |          |
| errorAlarm.msg       |                   | Text     |          |
| errorAlarm.state     |                   | Text     |          |
| errorAlarm.beginNode |                   | Text     |          |
| errorAlarm.endNode   |                   | Text     |          |