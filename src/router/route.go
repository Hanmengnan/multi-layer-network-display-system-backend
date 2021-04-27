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
	router.GET("/flow", webSocket.FlowChange)
	router.GET("/nodes", httpRequest.NodeInfo)
	router.GET("/links", httpRequest.LinkInfo)

	//光网络路由
	router.GET("/light", httpRequest.LightNetworkView)
	router.POST("/link/bandSet", httpRequest.BandSet)
	//数据网络路由
	router.GET("/data", httpRequest.DataNetworkView)
	router.GET("/data/link", httpRequest.DataNetworkLinkDetail)
	router.GET("/data/node", httpRequest.DataNetworkNodeDetail)

	//时频网络路由
	router.GET("/time", httpRequest.TimeNetworkView)

	server = &http.Server{
		Addr:           ":8070",
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
