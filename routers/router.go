package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/hesen/blog/routers/v1"
	"github.com/hesen/blog/middleware"
	"github.com/hesen/blog/pkg/validator"
	"github.com/hesen/blog/conf"
)

func healthy(c *gin.Context) {
	c.String(200, "OK")
}

func initConfig() {
	validator.RegisterValidation()
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	initConfig()

	r.Use(middleware.Recovery())
	r.Use(middleware.Cors())
	r.Use(middleware.ParseToken(conf.IgnoredAuthPath))

	r.GET("/healthy", healthy)

	v1Api := r.Group("/api/v1")
	v1.Group(v1Api)
	
	return r
}