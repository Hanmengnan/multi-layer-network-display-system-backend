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
	Response int64               `json:"response"`
	NodeList []database.NodeInfo `json:"nodeList"`
}

type linksResponse struct {
	Response int64               `json:"response"`
	LinkList []database.LinkInfo `json:"linkList"`
}

func SysInfo(c *gin.Context) {
	res := database.GetSystemBasicInfo()
	c.JSON(http.StatusOK, *res)
}

func NodeInfo(c *gin.Context) {
	c.JSON(http.StatusOK, nodesResponse{0, database.GetNodeInfo()})
}

func LinkInfo(c *gin.Context) {
	c.JSON(http.StatusOK, linksResponse{0, database.GetLinkInfo()})
}
