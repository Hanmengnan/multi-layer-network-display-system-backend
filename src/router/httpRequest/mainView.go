package httpRequest

import (
	"github.com/gin-gonic/gin"
	"net/http"

	database "3network-backend/src/model"
)

var (
	err error
)

type nodesResponse struct {
	Response     int64
	NodeInfoList []database.NodeInfo
}

type linksResponse struct {
	Response     int64
	LinkInfoList []database.LinkInfo
}

func HomeView(c *gin.Context) {
	res := database.GetSystemBasicInfo()
	c.JSON(http.StatusOK, *res)
}

func NodeInfo(c *gin.Context) {
	c.JSON(http.StatusOK, nodesResponse{0, database.GetDataNetworkNodeInfo()})
}

func LinkInfo(c *gin.Context) {
	c.JSON(http.StatusOK, linksResponse{0, database.GetDataNetworkLinkInfo()})
}
