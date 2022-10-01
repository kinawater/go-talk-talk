package global

import (
	"go-talk-talk/database"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

var MysqlDateBaseConfig *database.MysqlDatabaseConfig
