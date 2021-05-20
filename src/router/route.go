package router

import (
	"github.com/gin-gonic/gin"
	"log"
	database "multi-layer-network-display-system-backend/src/model"
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
	wg.Add(3)

	gin.ForceConsoleColor()
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
	router.GET("/eventList", httpRequest.SituationHandle)

	//光网络路由
	router.GET("/lightNetNodeStatistics", httpRequest.LightNetNodeStatistics)
	router.POST("/lightNetNodeInfo", httpRequest.LightNetNodeInfo)
	router.POST("/lightNetLinkInfo", httpRequest.LightNetLinkInfo)

	//数据网络路由
	router.GET("/dataNetInfo", httpRequest.DataNetworkView)

	//时频网络路由
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
	// 服务监听

	go webSocket.NewDataDetect()
	// 监测是否有新数据写入

	time.Sleep(time.Second * 2)

	go database.InjectNewData()
	// 模拟数据库写入新数据

	wg.Wait()
}
