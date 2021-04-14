# 3Network

> 数据库

## 主视图相关集合

### BasicInfo

> 集合
>
> 系统基本信息，该集合中只有一个文档。

```json
{
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
    "band": {
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

### NodeInfo
> 集合
>
> 节点及其信息。

```json
{
    "id": "Bj001",
    "name": "北京站",
    "city": "北京",
    "state": "正常", // 正常、繁忙、拥塞、故障    
    "location": {
        "Ing": [
            116,
            23,
            17
        ],
        "lat": [
            39,
            54,
            27
        ]
    },
    "type": "业务上下站", // 光中继站、电中继站、城域汇聚站    
    "Precision": "最高时钟基准", // 超高时钟基准、次高时钟基准、时钟网关、非时频网络节点    
    "error": "5ms",
    "throughput": 1000,
    "forwardingRate": 0.9
    "Topology": "img/1.jpg"
}
```

### NodeMessage
> 集合
>
> 节点收到的报文。
```json
{
    "id": "Bj001",
    "message": "正常",
    "rTime": "1613724232.828657"
}
```

### FlowChange

> 集合
>
> 网络流量变化。
```json
{
    "time": "1613724232.828657",
    "flowData": "80",
    //PB
    "rate": 0.9
}
```

### SituationHandle
> 集合
>
> 事件处理。
```json
{
    "id": "312",
    "type": "",
    "handler": "",
    "handleTime": "",
    "handleState": "",
    "origin": "",
    "message": ""
}
```





## 光网络相关集合

### LightNetworkLinkInfo

> 集合
>
> 光网络链路信息。
```json
{
    "id": "",
    "name": "",
    "node1ID": "",
    "node1Name": "",
    "node2ID": "",
    "node2Name": "",
    "contain": "",
    "state": ""
}

```





## 时频网络相关集合

### TimeNetworkLinkInfo

> 集合
>
> 时频网络链路信息。
```json
{
    "id": "",
    "name": "",
    "node1ID": "",
    "node1Name": "",
    "node2ID": "",
    "node2Name": "",
    "band": "",
}
```





## 数据网络相关集合

### DataNetworkBasicInfo

> 集合
>
> 数据网络基本信息。
```json
{
    "bandUsed": "",
    "bandTotal": "",
    "linkNum": 9121,
    "nodeNum": 1000,
    "Timestamp": 1753,
    "Location": 13541,
}
```

### DataNetworkLinkBasicInfo
> 集合
>
> 数据网络链路信息。
```json
{
    "id": "",
    "name": "",
    "node1ID": "",
    "node1Name": "",
    "node2ID": "",
    "node2Name": "",
    "band": ""
}
```

### DataNetworkLinkDetail

>集合
>
>数据网络链路参数变化。

```json
{
    "lId": "",
    "time": "1613724232.828657",
    "precisionError": 12,
    "loss": 0.01,
    "used": 0.59
}
```

### DataNetworkFlowChange

> 集合
>
> 数据网络流量变化。
```json
{
    "time": "1613724232.828657",
    "flowData": "80",
    //PB
}
```

### DataNetworkErrorAlarm
> 集合
>
> 数据网络接收报文。
```json
{
    "id": "",
    "rTime": "1613724232.828657",
    "type" : "",
    "level": "",
    "msg": "",
    "state": "",
    "beginNode": "",
    "endNode": ""
}
```

