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
// TODO: 1、跑以下 == 比较原理文章里面的列子。2、切片和数组的区别（主要是声明，初始化方面，怎么样声明、初始化算切片，怎样算数组），问题来源== 比较原理文章