package webSocket

import (
	database "3network-backend/src/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type flowChangeResponse struct {
	Response   int64
	FlowChange []database.NetworkFlow
}

func FlowChange(c *gin.Context) {
	rawData := database.GetNetworkFlow()
	res := flowChangeResponse{Response: 0, FlowChange: rawData}
	c.JSON(http.StatusOK, res)
}
