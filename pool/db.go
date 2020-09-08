package pool

import (
	// mysql 驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"

	"fmt"
	"log"

	"github.com/hesen/blog/pkg/setting"
)

// DB 全局数据库实列
var DB *gorm.DB

// Setupdatabse 初始化数据库连接
func Setupdatabse() {
	var (
		err error
		databaseType = setting.DatabaseSetting.DbType
		dbUser = setting.DatabaseSetting.DbUser
		dbPass = setting.DatabaseSetting.DbPassword
		dbName = setting.DatabaseSetting.DbName
		dbHost = setting.DatabaseSetting.DbHost
		dbPort = setting.DatabaseSetting.DbPort
		enableDebug = setting.DatabaseSetting.EnableDbDebug
		maxIdleConns = setting.DatabaseSetting.DbMaxIdleTime
		maxOpenConns = setting.DatabaseSetting.DbMaxOpenTime
	)

	DB, err = gorm.Open(databaseType, fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True", dbUser, dbPass, dbHost, dbPort, dbName))

	if err != nil {
		log.Fatalln(err)
	}

	// db.SingularTable(true)  //设置禁用表名的复数形式
	if enableDebug {
		DB.LogMode(true)              //打印日志，本地调试的时候可以打开看执行的sql语句
	}
	if maxIdleConns == 0 {
		maxIdleConns = 5
	}
	if maxOpenConns == 0 {
		maxOpenConns = 15
	}

	DB.DB().SetMaxIdleConns(maxIdleConns)     //设置空闲时的最大连接数
	DB.DB().SetMaxOpenConns(maxOpenConns)        //设置数据库的最大打开连接数
}

// FliterError 过滤错误，比如无记录返回时，不算错误
func FliterError(err error) error {
	if err == nil || gorm.IsRecordNotFoundError(err) {
		return nil
	}
	log.Println(err)
	return err
}
