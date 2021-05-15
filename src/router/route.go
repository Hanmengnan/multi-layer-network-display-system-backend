package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
	"time"

	"multi-layer-network-display-system-backend/src/router/httpRequest"
	"multi-layer-network-display-system-backend/src/router/webSocket"
)

var (
	router *gin.Engine
	server *http.Server
	err    error
)

func init() {
	var wg sync.WaitGroup
	wg.Add(2)

	router = gin.Default()

	// WebSocket
	router.GET("/ws", webSocket.WSHandler)
	//主视图路由
	router.GET("/sysInfo", httpRequest.SysInfo)
	router.GET("/nodeList", httpRequest.NodeList)
	router.GET("/linkList", httpRequest.LinkList)
	router.GET("/flowInfo", httpRequest.FlowChange)
	router.GET("/netInfo", httpRequest.NetInfo)

	router.GET("/nodeStatistics", httpRequest.NodeStatistics)

	//光网络路由
	router.GET("/lightNetInfo", httpRequest.LightNetworkView)
	router.GET("/lightNetBandMsg", httpRequest.LightNetworkView)
	router.GET("/lightNetBandSet", httpRequest.LightNetworkView)
	router.GET("/lightNetBandUse", httpRequest.LightNetworkView)
	router.GET("/lightNetNodeInfo", httpRequest.LightNetworkView)
	router.GET("/lightNetLinkInfo", httpRequest.LightNetworkView)

	//数据网络路由
	router.GET("/dataNetInfo", httpRequest.DataNetworkView)
	router.POST("/dataNetNodeInfo", httpRequest.DataNetworkNodeDetail)
	router.POST("/dataNetLinkInfo", httpRequest.DataNetworkLinkDetail)

	//时频网络路由
	router.GET("/timeNetInfo", httpRequest.TimeNetworkView)
	router.GET("/timeNetNodeStatistics", httpRequest.NodePrecisionStatistics)
	router.POST("/timeNetRoute", httpRequest.RouteTopology)

	server = &http.Server{
		Addr:           ":8070",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		err = server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()
	go webSocket.Test()

	wg.Wait()
}
