package config

// 该文件用于初始化viper
import (
	"github.com/spf13/viper"
	"go-talk-talk/global"
)

func ViperInit() {
	// 设置配置文件名称
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("toml")
	// 设置配置文件所在目录，可以设置多个
	viper.AddConfigPath(".")
	// 设置默认值
	setDefault()
	// 设置mysql默认值
	loadMysqlConfig()
}

func setDefault() {
	isSetAndDefaultValue("server.HTTP_PORT", 8088)
	isSetAndDefaultValue("server.READ_TIMEOUT", 60)
	isSetAndDefaultValue("server.WRITE_TIMEOUT", 60)

	isSetAndDefaultValue("database.TYPE", "mysql")
	isSetAndDefaultValue("database.USER", "root")
	isSetAndDefaultValue("database.HOST", "127.0.0.1")
	isSetAndDefaultValue("database.PORT", "3306")
	isSetAndDefaultValue("database.DATABASE_NAME", "go_talk_talk")
	isSetAndDefaultValue("database.TABLE_PREFIX", "talk")

	isSetAndDefaultValue("logger.LOG_PATH", "runtime/logs/")
	isSetAndDefaultValue("logger.LOG_SAVE_NAME", "log")
	isSetAndDefaultValue("logger.LOG_FILE_EXT", "log")
}

// 没有设置就给默认值
func isSetAndDefaultValue(key string, defaultValue any) {
	if !viper.IsSet(key) {
		viper.SetDefault(key, defaultValue)
	}
}

// 加载mysql配置
func loadMysqlConfig() {
	global.MysqlDateBaseConfig.DbDataName = viper.GetString("database.DATABASE_NAME")
	global.MysqlDateBaseConfig.DbUser = viper.GetString("database.USER")
	global.MysqlDateBaseConfig.DbHost = viper.GetString("database.HOST")
	global.MysqlDateBaseConfig.DbPort = viper.GetString("database.PORT")
	global.MysqlDateBaseConfig.DbTablePrefix = viper.GetString("database.TABLE_PREFIX")
	global.MysqlDateBaseConfig.DbType = viper.GetString("database.TYPE")
	global.MysqlDateBaseConfig.DbPassword = viper.GetString("database.PASSWORD")
}
