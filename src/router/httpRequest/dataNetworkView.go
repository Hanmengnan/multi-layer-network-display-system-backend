package httpRequest

import (
	database "3network-backend/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type info struct {
	Title string      `json:"title"`
	Num   interface{} `json:"num"`
}

type DataNetworkLinkDetailResponse struct {
	Response         int                        `json:"response"`
	BasicInfo        database.LinkInfo          `json:"basicInfo"`
	ParameterChange  []database.ParameterChange `json:"parameterChange"`
	OriginStatistics map[string]int32           `json:"originStatistics"`
}
type DataNetworkNodeDetailResponse struct {
	Response         int `json:"response"`
	NodeDetail       database.NodeDetail
	LinkStatistics   []database.MessageLinkStatistics
	DateStatistics   []database.MessageDateStatistics
	OriginStatistics map[string]int32
}

func DataNetworkView(c *gin.Context) {
	basicInfo := database.GetDataNetworkBasicInfo()
	c.JSON(http.StatusOK, basicInfo)
}

func DataNetworkLinkDetail(c *gin.Context) {
	var res DataNetworkLinkDetailResponse
	json := make(map[string]string)
	err = c.BindJSON(&json)
	if err != nil {
		return
	}
	linkId := json["linkId"]

	res.Response = 0
	res.BasicInfo = database.GetLinkBasicInfo(linkId)
	res.ParameterChange = database.GetLinkParameterChange(linkId)
	res.OriginStatistics = database.GetLinkMessageOriginStatistics(linkId)

	c.JSON(http.StatusOK, res)
}

func DataNetworkNodeDetail(c *gin.Context) {
	var res DataNetworkNodeDetailResponse
	json := make(map[string]string)
	err = c.BindJSON(&json)
	if err != nil {
		return
	}
	nodeId := json["nodeId"]
	res.Response = 0
	res.LinkStatistics = database.GetNodeMessageLinkStatistics(nodeId)
	res.DateStatistics = database.GetNodeMessageDateStatistics(nodeId)
	res.OriginStatistics = database.GetNodeMessageOriginStatistics(nodeId)
	res.NodeDetail = database.GetNodeDetail(nodeId)
	c.JSON(http.StatusOK, res)
}
