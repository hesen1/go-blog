package middleware

const (
	// FORWARDED 解析用户Ip方式，从X-Forwarded-For解析
	FORWARDED = "forwarded"
	// REAL 解析用户Ip方式，从X-Real-IP解析
	REAL = "real"
)

// ParseIPOptions IP解析中间件配置
type ParseIPOptions struct {
	/*
		解析来源：
			forwarded: X-Forwarded-For
			real: X-Real-IP

		default: forwarded
	*/
	Source string
	/*
		真实ip索引
		当source参数为 forwarded 时，此参数有效

		default: 0
	*/
	Index int8
}

// defaultParseIPOptions 获取 ParseIpOptions 默认配置
func defaultParseIPOptions() ParseIPOptions {
	return ParseIPOptions {
		Source: "forwarded", Index: 0,
	}
}

// ParseIPResult IP解析信结果
type ParseIPResult struct {
	// 真实IP
	RealIP string
	/*
		所有代理IP
		当解析方式 source 是 real时，该字段值为 nil
	*/
	ProxyIps []string
}

// PathRegexp 路由正则表达式
type PathRegexp []string

// CorsOptions 跨域配置
type CorsOptions struct {
	// Access-Control-Allow-Origin,默认 *
	Origin []string
	// Access-Control-Allow-Methods，默认 GET,PUT,DELETE,POST,HEAD,PATCH
	Methods []string
	// Access-Control-Allow-Headers, 默认 Content-Type,Authorization,Origin
	Headers []string
	// Access-Control-Allow-Credentials, 默认不设置
	Credentials bool
	// Access-Control-Max-Age,默认不设置
	MaxAge int64
	// Access-Control-Expose-Headers, 需要暴露给外部的请求头，默认不设置
	ExposeHeaders []string
}

// NewParseIPResult 获取新实列
func NewParseIPResult() ParseIPResult {
	return ParseIPResult{}
}
