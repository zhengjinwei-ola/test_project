package server

import (
	"fmt"
	"log"
	"net/http"
	"test_project/app/pb"
	"test_project/app/server/route"
	"test_project/util/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
)

func Run() {
	// 初始化 gin
	serverConfig := config.ConfigServer.Get("server")
	serverConfigS := &pb.ServerConfig{}
	err := gconv.Struct(serverConfig, serverConfigS)
	if err != nil {
		panic(err)
	}
	gin.SetMode(serverConfigS.GetRunMode())
	engine := gin.New()
	engine.Use(cors.New(cors.Config{
		AllowMethods:    []string{"PUT", "POST", "PATCH", "GET", "DELETE"},
		AllowAllOrigins: true,
	}))
	//engine.Use(middleware.CatchError())
	route.InitRoute(engine)

	httpServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", gconv.Int(serverConfigS.GetPort())),
		Handler:        engine,
		ReadTimeout:    time.Duration(gconv.Int64(serverConfigS.GetReadTimeOut())) * time.Second,
		WriteTimeout:   time.Duration(gconv.Int64(serverConfigS.GetWriteTimeOut())) * time.Second,
		MaxHeaderBytes: gconv.Int(serverConfigS.GetMaxHeaderBytes()),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe err:", err)
	}
}
