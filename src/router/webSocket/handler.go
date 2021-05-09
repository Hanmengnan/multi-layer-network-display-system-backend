package webSocket

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	database "multi-layer-network-display-system-backend/src/model"
	"net/http"
	"time"
)

type webSocketDataChannel struct {
	netParameter chan database.NetParameter
	nodeList     chan []database.NodeInfo
}

var (
	DataChannel = webSocketDataChannel{make(chan database.NetParameter, 1), make(chan []database.NodeInfo, 1)}
)

func WSHandler(c *gin.Context) {
	upGrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	ws, _ := upGrader.Upgrade(c.Writer, c.Request, nil)

	var dataType string
	var data interface{}

	for {
		select {
		case data = <-DataChannel.netParameter:
			dataType = "parameterChange"
		case data = <-DataChannel.nodeList:
			dataType = "nodeList"
		}
		err := ws.WriteJSON(gin.H{"dataType": dataType, "data": data})
		if err != nil {
			return
		}
	}
}

func Test() {
	for {
		info1 := database.GetNodeList()
		DataChannel.nodeList <- info1
		info2 := database.GetNetParameter()
		DataChannel.netParameter <- *info2

		time.Sleep(3 * time.Second)
	}
}
