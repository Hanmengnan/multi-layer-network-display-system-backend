package httpRequest

import (
	database "3network-backend/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type dataNetworkResponse struct {
	Response     string
	BandUsed     int
	BandTotal    int
	LinkNum      int
	NodeNum      int
	TimestampNum int
	LocationNum  int
	Nodes        []database.DataNetworkNodeInfo
	Links        []database.DataNetworkLinkInfo
	FlowChange   []database.DataNetworkFlow
	ErrorAlarm   []database.DataNetworkErrorAlarm
}

type DataNetworkLinkDetailResponse struct {
	Response  int    `json:"response"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Node1Id   string `json:"node1ID"`
	Node1Name string `json:"node1Name"`
	Node2Id   string `json:"node2ID"`
	Node2Name string `json:"node2Name"`
	Band      string `json:"band"`
	Detail    []database.DataNetworkLinkDetail
}

func DataNetworkView(c *gin.Context) {
	var res = dataNetworkResponse{}
	basicInfo := database.GetDataNetworkBasicInfo()
	res.BandUsed = basicInfo.BandUsed
	res.BandTotal = basicInfo.BandTotal
	res.LinkNum = basicInfo.LinkNum
	res.NodeNum = basicInfo.NodeNum
	res.TimestampNum = basicInfo.TimestampNum
	res.LocationNum = basicInfo.LocationNum
	res.Links = database.GetDataNetworkLinkInfo()
	res.Nodes = database.GetDataNetworkNodeInfo()
	res.FlowChange = database.GetDataNetworkFlow()
	res.ErrorAlarm = database.GetDataNetworkErrorALarms()
	res.Response = "0"
	c.JSON(http.StatusOK, res)
}

func DataNetworkNodeDetail(c *gin.Context) {
	database.GetDataNetworkNodeInfo()
}

func DataNetworkLinksDetail(c *gin.Context) {
	var res DataNetworkLinkDetailResponse
	linkId := c.Query("linkId")
	basicInfo := database.GetDataNetworkLinkInfo(linkId)[0]
	detail := database.GetDataNetworkLinkDetail(linkId)

	res.Id = basicInfo.Id
	res.Name = basicInfo.Id
	res.Node1Id = basicInfo.Node1Id
	res.Node1Name = basicInfo.Node1Name
	res.Node2Id = basicInfo.Node2Id
	res.Node2Name = basicInfo.Node2Name
	res.Band = basicInfo.Node2Name
	res.Detail = detail

	c.JSON(http.StatusOK, res)
}
