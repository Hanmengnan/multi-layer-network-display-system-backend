package webSocket

import (
	database "3network-backend/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type dataNetworkFlowChangeResponse struct {
	Response   string
	FlowChange []database.DataNetworkFlow
}

func DataNetworkFlowChange(c *gin.Context) {
	rawData := database.GetDataNetworkFlow()
	res := dataNetworkFlowChangeResponse{Response: "0", FlowChange: rawData}
	c.JSON(http.StatusOK, res)
}
