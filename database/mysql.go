package database

import (
	"fmt"
	"go-talk-talk/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

func init() {
	// 设置db的日志
	DBLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 一秒以上的sql记入慢日志
			LogLevel:      logger.Info, //
			Colorful:      false,       // 禁止彩色打印
		},
	)

	// 链接数据库
	DBLinkConfigFormat := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		global.MysqlDateBaseConfig.DbUser,
		global.MysqlDateBaseConfig.DbPassword,
		global.MysqlDateBaseConfig.DbHost,
		global.MysqlDateBaseConfig.DbPort,
		global.MysqlDateBaseConfig.DbDataName)
	db, err := gorm.Open(mysql.Open(DBLinkConfigFormat), &gorm.Config{
		SkipDefaultTransaction: true, // 默认关闭事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			TablePrefix:   global.MysqlDateBaseConfig.DbTablePrefix,
		},
		//开启日志监控级别
		Logger: DBLogger,
	})
	if err != nil {
		log.Println(err)
	}
	sqlDb, err := db.DB()

	/*设置连接池*/

	// 最大空闲数
	sqlDb.SetMaxIdleConns(10)
	// 最大连接数
	sqlDb.SetMaxOpenConns(50)
	// 最长连接时间
	sqlDb.SetConnMaxLifetime(5 * time.Minute)

	global.MysqlDB = db

}
