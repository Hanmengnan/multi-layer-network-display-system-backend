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
	res := database.GetSystemBasicInfo()
	c.JSON(http.StatusOK, *res)
}

func HandleSituation(c *gin.Context) {
}
