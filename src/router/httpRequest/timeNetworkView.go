package httpRequest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	database "multi-layer-network-display-system-backend/src/model"
	"net/http"
)

func TimeNetworkView(c *gin.Context) {
}

func TimeNetworkLinks(c *gin.Context) {
}

func NodePrecisionStatistics(c *gin.Context) {
	res := database.GetNodePrecisionStatistics()
	c.JSON(http.StatusOK, res)
}

func RouteTopology(c *gin.Context) {
	json := make(map[string]string)
	err = c.BindJSON(&json)
	if err != nil {
		return
	}
	start := json["start"]
	end := json["end"]
	fmt.Printf("%v%v", start, end)

	res := database.GetRouteTopology(start, end)

	c.JSON(http.StatusOK, gin.H{"route": res.Route})
}
