package webSocket

import (
	database "3network-backend/src/model"
	"github.com/gin-gonic/gin"
	"log"
)

type flowChangeResponse struct {
	Response   int64
	FlowChange map[string][]database.NetworkFlow
}

func FlowChange(c *gin.Context) {
	connection, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		res := flowChangeResponse{Response: 0, FlowChange: database.GetNetworkFlow()}
		err = connection.WriteJSON(res)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
