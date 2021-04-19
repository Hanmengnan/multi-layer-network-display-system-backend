package webSocket

import (
	database "3network-backend/src/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var (
	connection *websocket.Conn
	err        error
)

func SituationHandle(c *gin.Context) {
	connection, err = upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		info := database.GetSituationHandleInfo()
		err = connection.WriteJSON(info)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
