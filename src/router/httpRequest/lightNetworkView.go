package httpRequest

import (
	"github.com/gin-gonic/gin"
	database "multi-layer-network-display-system-backend/src/model"
	"net/http"
)

type NodeHealthStatistics struct {
	NodeNum           map[string]int32 `json:"nodeNum"`
	NodeHealthPercent map[string]int32 `json:"nodeHealthPercent"`
	ErrorNum          struct {
		Num     int32   `json:"num"`
		Percent float64 `json:"percent"`
	} `json:"errorNum"`
}

func BandSet(c *gin.Context) {
}

func LightNetNodeStatistics(c *gin.Context) {
	var res NodeHealthStatistics
	var sum int32 = 0
	var healthSum int32 = 0

	nodeStatistics := database.GetNodeTypeStatistics()
	nodeStatisticsHealth := database.GetNodeTypeStatisticsHealth()

	for _, value := range nodeStatistics {
		sum += value
	}

	for _, value := range nodeStatisticsHealth {
		healthSum += value
	}

	res.NodeNum = nodeStatistics
	res.NodeHealthPercent = nodeStatisticsHealth
	res.ErrorNum.Num = sum - healthSum
	res.ErrorNum.Percent = float64(res.ErrorNum.Num) / float64(sum) * 100

	c.JSON(http.StatusOK, res)
}

func LightNetNodeInfo(c *gin.Context) {
	json := make(map[string]string)
	err = c.BindJSON(&json)
	if err != nil {
		return
	}
	nodeId := json["nodeId"]
	c.JSON(http.StatusOK, database.GetNodeInfo(nodeId))
}

func LightNetLinkInfo(c *gin.Context) {
	json := make(map[string]string)
	err = c.BindJSON(&json)
	if err != nil {
		return
	}
	linkId := json["linkId"]
	linkInfo := database.GetLinkBasicInfo(linkId)
	node1Info := database.GetNodeInfo(linkInfo.Node1Id)
	node2Info := database.GetNodeInfo(linkInfo.Node2Id)
	c.JSON(http.StatusOK, gin.H{"node1Name": node1Info.Name, "node1State": node1Info.State, "node2Name": node2Info.Name, "node2State": node2Info.State})
}
