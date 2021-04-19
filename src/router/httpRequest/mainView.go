package httpRequest

import (
	"github.com/gin-gonic/gin"
	"net/http"

	database "3network-backend/src/model"
)

var (
	err error
)

func HomeView(c *gin.Context) {
	rawData := database.GetSystemBasicInfo()
	c.JSON(http.StatusOK, *rawData)
}

func HandleSituation(c *gin.Context) {
}
