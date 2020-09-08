package main

import (
	"net/http"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/hesen/blog/pool"
	"github.com/hesen/blog/pkg/setting"
	"github.com/hesen/blog/routers"
	"github.com/hesen/blog/migration"
)

func main() {
	setting.Setup()

	gin.SetMode(setting.ServerSetting.RunMode)

	pool.Setupdatabse()
	migration.Migrations()
	defer pool.DB.Close()

	routersInit := routers.InitRouter()
	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.AppPort)

	server := &http.Server{
		Addr:          endPoint,
		Handler:        routersInit,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()
}
