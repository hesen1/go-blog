package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/hesen/blog/entry/user"
)

// Group 路由组
func Group(r *gin.RouterGroup) {
	// 登录
	r.POST("/login", user.Login)
	// 获取用户列表
	r.GET("/users", user.GetUsers)
	// 用户注册
	r.POST("/users", user.Register)
}
