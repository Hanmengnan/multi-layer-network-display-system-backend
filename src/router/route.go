package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"3network-backend/src/router/httpRequest"
	"3network-backend/src/router/webSocket"
)

var (
	router *gin.Engine
	server *http.Server
	err    error
)

func init() {
	router = gin.Default()
	//主视图路由
	router.GET("/", httpRequest.HomeView)
	router.GET("/situaionHandle", webSocket.SituationHandle)

	//光网络路由
	router.GET("/light", httpRequest.LightNetworkView)
	router.GET("/light/nodesAndLinks", httpRequest.LightNodesAndLinks)
	router.POST("/link/bandSet", httpRequest.BandSet)
	//数据网络路由
	router.GET("/data", httpRequest.DataNetworkView)
	router.GET("/data/nodes", httpRequest.DataNetworkNodeDetail)
	router.GET("/data/links", httpRequest.DataNetworkLinksDetail)
	router.GET("data/flowChange", webSocket.DataNetworkFlowChange)

	//时频网络路由
	router.GET("/time", httpRequest.TimeNetworkView)
	router.GET("/time/links", httpRequest.TimeNetworkLinks)

	server = &http.Server{
		Addr:           ":8090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
