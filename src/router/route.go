package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"3network-backend/src/router/api"
)

var (
	router *gin.Engine
	server *http.Server
	err    error
)

func init() {
	router = gin.Default()
	//主视图路由
	router.GET("/", api.HomeView)
	//光网络路由
	router.GET("/light", api.LightNetworkView)
	router.GET("/light/nodesAndLinks", api.LightNodesAndLinks)
	router.POST("/link/bandSet", api.BandSet)
	//数据网络路由
	router.GET("/data", api.DataNetworkView)
	router.GET("/data/nodes", api.DataNetworkNodes)
	router.GET("/data/links", api.DataNetworkLinks)
	//时频网络路由
	router.GET("/time", api.TimeNetworkView)
	router.GET("/time/links", api.TimeNetworkLinks)

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
