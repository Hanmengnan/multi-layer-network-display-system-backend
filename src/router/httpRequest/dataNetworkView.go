package httpRequest

import (
	database "3network-backend/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type dataNetworkResponse struct {
	Response     int64
	BandUsed     int64
	BandTotal    int64
	LinkNum      int64
	NodeNum      int64
	TimestampNum int64
	LocationNum  int64
}

type DataNetworkLinkDetailResponse struct {
	Response int64 `json:"response"`
	Detail   []database.LinkDetail
}
type DataNetworkNodeDetailResponse struct {
	Response int64 `json:"response"`
	Detail   []database.NodeDetail
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
	res.Response = 0
	c.JSON(http.StatusOK, res)
}

func DataNetworkLinkDetail(c *gin.Context) {
	var res DataNetworkLinkDetailResponse
	linkId := c.Query("linkId")
	res.Response = 0
	res.Detail = database.GetDataNetworkLinkDetail(linkId)

	c.JSON(http.StatusOK, res)
}

func DataNetworkNodeDetail(c *gin.Context) {
	var res DataNetworkNodeDetailResponse
	nodeId := c.Query("nodeId")
	res.Response = 0
	res.Detail = database.GetDataNetworkNodeDetail(nodeId)

	c.JSON(http.StatusOK, res)
}
