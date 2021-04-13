## 3Network

> 集合

### BasicInfo

> 文档

```json
{
    "version": "v2.1",
    "code": "code71233",
    "state": {
        "database": 0,
        "network": 0,
        "cluster": 0
    },
    "people":{
        "principal":"牛猛",
        "duty":"娓奇"
    }
}
```

### NodeInfo
> 文档

```json
{
    "id": "Bj001",
    "name": "北京站",
    "city": "北京",
    "state": "正常",
    // 正常、繁忙、拥塞、故障
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
    "type": "业务上下站",
    // 光中继站、电中继站、城域汇聚站
    "Precision": "最高时钟基准",
    // 超高时钟基准、次高时钟基准、时钟网关、非时频网络节点
    "error": "5ms",
    "Topology": "img/1.jpg"
}
```

### NodeMessage
> 文档
```json
{
    "id": "Bj001",
    "message": "正常",
    "rTime": "1613724232.828657"
}
```

### FlowChange

> 文档
```json
{
    "time": "1613724232.828657",
    "flowData": "80",
    //PB
    "rate": 0.9
}
```

### SituationHandle
> 文档
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

---

### LightNetworkLinkInfo
> 文档
```json
{
    "id": "",
    "name": "",
    "node1ID": "",
    "node1Name": "",
    "node2ID": "",
    "node2Name": "",
    "contain": ""
}

```



---

### TimeNetworkLinkInfo
> 文档
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

---

### DataNetworkBasicInfo
> 文档
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

### DataNetworkLinkInfo
> 文档
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



### DataNetworkFlowChange
> 文档
```json
{
    "time": "1613724232.828657",
    "flowData": "80",
    //PB
}
```

### DataNetworkErrorAlarm
> 文档
```json
{
    "id": "",
    "level": "",
    "msg": "",
    "state": "",
    "beginNode": "",
    "endNode": ""
}
```

