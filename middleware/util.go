package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"

	"strings"
	"net/http"
	// "fmt"
	"runtime/debug"
	"log"
	"regexp"
)

// ParseIP IP解析中间件
/*
	结果挂载在 RealIP 

	example:
		r := gin.Default()

		r.Use(middleware.ParseIp())

		...

		realIp := c.MustGet("RealIP").(middleware.ParseIpResult)

		fmt.Println(realIp.RealIP + "  " + strings.Join(realIp.ProxyIps, ","))
*/
func ParseIP(options ...ParseIPOptions) gin.HandlerFunc {
	return func (c *gin.Context) {
		var option ParseIPOptions
		if len(options) > 0 {
			option = options[0]
		} else {
			option = defaultParseIPOptions()
		}

		result := NewParseIPResult()

		if option.Source == REAL {
			result.RealIP = c.GetHeader("X-Real-IP")
		} else {
			result.ProxyIps = strings.Split(c.GetHeader("X-Forwarded-For"), ",")
			if len(result.ProxyIps) > 0 {
				result.RealIP = result.ProxyIps[0]
				result.ProxyIps = result.ProxyIps[1:]
			}
		}

		c.Set("RealIP", result)

		c.Next()
	}
}

// Recovery Panic恢复中间件
func Recovery() gin.HandlerFunc {
	return func (c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				log.Println(string(debug.Stack()))
				// c.JSON(http.StatusInternalServerError, gin.H{ "code": 5001, "message": fmt.Sprintf("error: %s", err) })
				c.JSON(http.StatusInternalServerError, gin.H{ "code": 5001, "message": "system error" })
			}
		}()
		c.Next()
	}
}

// Cors 跨域设置
func Cors(options ...CorsOptions) gin.HandlerFunc {
	return func(c *gin.Context) {

		option := CorsOptions{
			Origin: []string{ "*" },
			Methods: []string{ "GET", "PUT", "DELETE", "POST", "HEAD", "PATCH" },
			Headers: []string{ "Content-Type", "Authorization", "Origin" },
		}

		if len(options) > 0 {
			mergo.Merge(&options[0], option, mergo.WithOverride)
			option = options[0]
		}

		if option.Credentials {
			c.Header("Access-Control-Allow-Credentials", "true")
		}
		if len(option.ExposeHeaders) > 0 {
			c.Header("Access-Control-Expose-Headers", strings.Join(option.ExposeHeaders, ","))
		}
		if len(option.Origin) > 0 {
			if option.Origin[0] == "*" {
				c.Header("Access-Control-Allow-Origin", "*")
			} else {
				var isContain = false
				var reqOrigin = c.GetHeader("Origin")
				// 验证origin是否满足限制的规则
				for _, origin := range option.Origin {
					match, _ := regexp.MatchString(origin, reqOrigin)
					if match {
						isContain = true
						break
					}
				}
				if isContain {
					c.Header("Access-Control-Allow-Origin", reqOrigin)
				}
			}
		}

		if c.Request.Method == http.MethodOptions || c.Request.Method == strings.ToLower(http.MethodOptions) {
			if len(option.Methods) > 0 {
				c.Header("Access-Control-Allow-Methods", strings.Join(option.Methods, ","))
			}
			if len(option.Headers) > 0 {
				c.Header("Access-Control-Allow-Headers", strings.Join(option.Headers, ","))
			}
			if option.MaxAge > 0 {
				c.Header("Access-Control-Max-Age", string(option.MaxAge))
			}

			c.Header("Content-Length", "0")
			c.AbortWithStatus(http.StatusNoContent)
		} else {
			c.Next()
		}
	}
}
