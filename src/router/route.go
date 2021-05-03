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
	//router.GET("/home", httpRequest.HomeView)
	router.GET("/sysInfo", httpRequest.SysInfo)
	router.GET("/situaionHandle", webSocket.SituationHandle)
	router.GET("/flow", webSocket.FlowChange)
	router.GET("/nodeList", httpRequest.NodeInfo)
	router.GET("/linkList", httpRequest.LinkInfo)

	//光网络路由
	router.GET("/light", httpRequest.LightNetworkView)
	router.POST("/link/bandSet", httpRequest.BandSet)
	//数据网络路由
	router.GET("/DataNetInfo", httpRequest.DataNetworkView)
	router.POST("/DataNetNodeInfo", httpRequest.DataNetworkNodeDetail)
	router.POST("/DataNetLinkInfo", httpRequest.DataNetworkLinkDetail)

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
