package setting


import (
	"log"
	"time"
	// "os"

	"github.com/go-ini/ini"
)

// Server 后端应用配置
type Server struct {
	RunMode      string
	AppPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	JwtSecret string
}

// ServerSetting 后端应用配置实例
var ServerSetting = &Server{}

// Database 数据库配置
type Database struct {
	DbType string
	DbUser string
	DbPassword string
	DbHost string
	DbPort int32
	DbName string
	EnableDbDebug bool
	DbMaxIdleTime int
	DbMaxOpenTime int
}

// DatabaseSetting 数据库配置实例
var DatabaseSetting = &Database{}

var cfg *ini.File

// Setup 初始化环境变量
func Setup() {
	var err error
	var envPath = "conf/env.ini"
	// if (os.Getenv("GIN_MODE") == "release") {
	// 	envPath = "env.ini"
	// } else {
	// 	envPath = "conf/env.ini"
	// }
	cfg, err = ini.Load(envPath)
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	// mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
	// mapTo("redis", RedisSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

// mapTo 读取内存里的环境配置项
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
