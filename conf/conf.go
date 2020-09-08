package conf

import (
	"github.com/hesen/blog/middleware"
)

// IgnoredAuthPath 不需要验证的路由
var IgnoredAuthPath = &middleware.PathRegexp{ "login$", }
// [...]string{
// 	"login$",
// }
