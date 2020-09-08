package middleware

import (
	"github.com/gin-gonic/gin"
	"regexp"
	"strings"
	"net/http"

	"github.com/hesen/blog/pkg/res"
	"github.com/hesen/blog/pkg/util"
)

// ParseToken 解析token
func ParseToken(ignoredAuthPaths *PathRegexp) gin.HandlerFunc {
	return func (c *gin.Context) {
		for _, ignoredAuthPath := range *ignoredAuthPaths {
			match, _ := regexp.MatchString(ignoredAuthPath, c.FullPath())
			if match {
				c.Next()
				return
			}
		}
		token := c.GetHeader("Authorization")
		
		if token == "" {
			c.JSON(http.StatusOK, res.Error(res.InvalidTokenCode))
			// 从功能表现来看，无论是c.Next() 还是 return，只要没有调用c.Abort()
			// 该中间件 return后，还是会去调用后续的中间件或者路由处理器
			// 猜测，gin底层中间件是一个队列，只要当前执行的中间件完毕（return 或者 c.Next()）都会去执行下一个
			// 所以c.Next()只是控制提前调用下一个中间件，不调用c.Next()(因为c.Next()内部有for 循环 执行中间handler，所以执行的是for循环里的中间handler)也会到下一个中间件执行
			// 除非调用c.Abort()
			c.Abort()
			return
		}
		user, err := util.ParseJwt(strings.Split(token, " ")[1])

		if err != nil {
			c.JSON(http.StatusOK, res.Error(res.InvalidTokenCode))
			c.Abort()
			return
		}

		c.Set("userInfo", user)

		c.Next()
	}
}
