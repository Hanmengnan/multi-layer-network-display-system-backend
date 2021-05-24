package webSocket

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	database "multi-layer-network-display-system-backend/src/model"
	"net/http"
)

type webSocketDataChannel struct {
	netParameter chan database.NetParameter
	nodeList     chan []database.NodeInfo
	situations   chan map[string]interface{}
	nodeOverLoad chan map[string]float64
	nodeDetail   chan DataNetworkNodeDetailResponse
	linkDetail   chan DataNetworkLinkDetailResponse
}

var (
	nodeBandOverloadIndicator = ""
	dataNetNodeInfoIndicator  = ""
	dataNetLinkInfoIndicator  = ""
)

var (
	DataChannel = webSocketDataChannel{
		make(chan database.NetParameter, 1),
		make(chan []database.NodeInfo, 1),
		make(chan map[string]interface{}, 1),
		make(chan map[string]float64, 1),
		make(chan DataNetworkNodeDetailResponse, 1),
		make(chan DataNetworkLinkDetailResponse, 1),
	}
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

	var rdataType, sdataType string
	var sdata interface{}
	var rdata string

	go func() {
		var message map[string]interface{}
		for {
			_, byteData, err := ws.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			err = json.Unmarshal(byteData, &message)
			if err != nil {
				log.Println(err)
				return
			}

			rdataType = message["dataType"].(string)
			rdata = message["data"].(string)
			switch rdataType {
			case "lightNetNodeOverload":
				if nodeBandOverloadIndicator = rdata; nodeBandOverloadIndicator != "" {
					DataChannel.nodeOverLoad <- database.GetLinkOverload(rdata)
				}
			case "dataNetNodeInfo":
				if dataNetNodeInfoIndicator = rdata; dataNetNodeInfoIndicator != "" {
					DataChannel.nodeDetail <- DataNetworkNodeDetail(rdata)
				}
			case "daraNetLinkInfo":
				if dataNetLinkInfoIndicator = rdata; dataNetLinkInfoIndicator != "" {
					DataChannel.linkDetail <- DataNetworkLinkDetail(rdata)
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case sdata = <-DataChannel.netParameter:
				sdataType = "parameterChange"
			case sdata = <-DataChannel.nodeList:
				sdataType = "nodeList"
			case sdata = <-DataChannel.situations:
				sdataType = "situation"
			case sdata = <-DataChannel.nodeOverLoad:
				sdataType = "nodeOverLoad"
			case sdata = <-DataChannel.nodeDetail:
				sdataType = "nodeDetail"
			case sdata = <-DataChannel.linkDetail:
				sdataType = "linkDetail"
			}
			err := ws.WriteJSON(gin.H{"dataType": sdataType, "data": sdata})
			if err != nil {
				return
			}
		}
	}()
}

func NewDataDetect() {
	for {
		messageType := <-database.NewDataChannel
		switch messageType {
		case "parameterChange":
			DataChannel.netParameter <- *database.GetNetParameter()
		case "nodeList":
			DataChannel.nodeList <- database.GetNodeList()
		case "situation":
			DataChannel.situations <- database.GetSituationHandleInfo()
		case "nodeOverLoad":
			if nodeBandOverloadIndicator != "" {
				DataChannel.nodeOverLoad <- database.GetLinkOverload(nodeBandOverloadIndicator)
			}
		case "dataNetNodeInfo":
			if dataNetNodeInfoIndicator != "" {
				DataChannel.nodeDetail <- DataNetworkNodeDetail(dataNetNodeInfoIndicator)
			}
		case "daraNetLinkInfo":
			if dataNetLinkInfoIndicator != "" {
				DataChannel.linkDetail <- DataNetworkLinkDetail(dataNetNodeInfoIndicator)
			}
		}
	}
}
