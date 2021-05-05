package httpRequest

import (
	"github.com/gin-gonic/gin"
	"net/http"

	database "3network-backend/src/model"
)

var (
	err error
)

func SysInfo(c *gin.Context) {
	res := database.GetSystemBasicInfo()
	c.JSON(http.StatusOK, *res)
}

func NodeList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": 0, "nodeList": database.GetNodeList()})
}

func LinkList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"response": 0, "linkList": database.GetLinkList()})
}

func NetInfo(c *gin.Context) {
	c.JSON(http.StatusOK, *database.GetNetParameter())
}
