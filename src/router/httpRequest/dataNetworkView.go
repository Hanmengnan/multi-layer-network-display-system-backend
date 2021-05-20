package httpRequest

import (
	"github.com/gin-gonic/gin"
	database "multi-layer-network-display-system-backend/src/model"
	"net/http"
)

func DataNetworkView(c *gin.Context) {
	basicInfo := database.GetDataNetworkBasicInfo()
	c.JSON(http.StatusOK, basicInfo)
}

func FlowChange(c *gin.Context) {
	flowStatistic := database.GetNetworkFlow()
	c.JSON(http.StatusOK, flowStatistic)
}
