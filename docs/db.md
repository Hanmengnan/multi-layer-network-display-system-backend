# 	networkStaticData

> 静态数据数据库

## 主视图相关集合

### basicInfo

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
        "light": {
            "begin": "MD01",
            "end": "MD02"
        },
        "data": {
            "begin": "MD03",
            "end": "MD05"
        },
        "emergency": {
            "begin": "MD06",
            "end": "MD10"
        },
    }
}
```

### nodeInfo
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
        "type":"Ponit",
        "coordinates": [116.4551,40.2539]
    },
    "type": "业务上下站", // 光中继站、电中继站、城域汇聚站    
    "Precision": "最高时钟基准", // 超高时钟基准、次高时钟基准、时钟网关、非时频网络节点    
    "error": "5",
    "throughput": "1000",
    "forwardingRate": "0.9",
    "Topology": "img/1.jpg"
}
```

### linkInfo

> 集合
>
> 链路及其信息

```json
{
    "id": "",
    "name": "",
    "node1ID": "",
    "node1Name": "",
    "node2ID": "",
    "node2Name": "",
    "contain": "",
    "bulided": true,
    "band": "",
    "usedForTime": true,
    "usedForData": true
    
}
```

## 光网络相关集合

## 时频网络相关集合

## 数据网络相关集合

### dataNetworkBasicInfo

> 集合
>
> 数据网络基本信息。
```json
{
    "bandUsed": 100,
    "bandTotal": 300,
    "linkNum": 9121,
    "nodeNum": 1000,
    "timestampNum": 1753,
    "location": 13541,
}
```

# networkDynamicData

> 动态数据数据库


### nodeMessage
> 集合
>
> 节点收到的报文。
```json
{
    "id": "Bj001",
    "message": "正常",
    "rTime": "1613724232.828657",
    "type": "",
    "handler": "",
    "handleTime": "",
    "handleState": "",
    "origin": ""
}
```

### flowChange

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

### linkErrorAlarm
> 集合
>
> 数据网络接收报文。
```json
{
    "id": "",
    "rTime": "1613724232.828657",
    "type" : 0,
    "level": "",
    "message": "",
    "state": "",
    "fromNodeId": "",
    "toNodeId": ""
    
}
```


### linkDetail

>集合
>
>数据网络链路参数变化。

```json
{
    "id": "",
    "time": "1613724232.828657",
    "precisionError": 12,
    "loss": 0.01,
    "used": 0.59
}
```
