package config

// 该文件用于初始化viper
import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

type ServerConfig struct {
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MysqlDatabaseConfig struct {
	DbType        string
	DbUser        string
	DbPassword    string
	DbHost        string
	DbPort        string
	DbDataName    string
	DbTablePrefix string
}

type EmailConfig struct {
	From         string
	SmtpAddr     string
	SmtpUsername string
	SmtpPassword string
	SmtpHost     string
}

type RedisConfig struct {
	Host    string
	Port    string
	Network string
}

type LoggerConfig struct {
	LogPath     string
	LogSaveName string
	LogFileExt  string
}

var ServerConf ServerConfig
var MysqlDateBaseConfig MysqlDatabaseConfig
var EmailConf EmailConfig
var RedisConf RedisConfig
var LoggerConf LoggerConfig

var RunMode string

func init() {
	// 设置配置文件
	viper.SetConfigFile("./config")
	// 设置配置文件名称
	viper.SetConfigName("config")
	// 设置配置文件类型
	viper.SetConfigType("toml")
	// 设置配置文件所在目录，可以设置多个
	viper.AddConfigPath("./config")

	viper.WatchConfig()
	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {
		// 处理读取配置文件的错误
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// gin运行模式
	RunMode = viper.GetString("RUN_MODE")
	// 设置默认值
	setDefault()
	// 设置mysql默认值
	loadMysqlConfig()
	// 服务器配置
	loadServerConfig()
	// email配置
	loadEmailConfig()
	// redis配置
	loadRedisConfig()
	// 日志配置
	loadLoggerConfig()
}

func setDefault() {
	isSetAndDefaultValue("RUN_MODE", "debug")

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

	MysqlDateBaseConfig.DbDataName = viper.GetString("database.DATABASE_NAME")
	MysqlDateBaseConfig.DbUser = viper.GetString("database.USER")
	MysqlDateBaseConfig.DbHost = viper.GetString("database.HOST")
	MysqlDateBaseConfig.DbPort = viper.GetString("database.PORT")
	MysqlDateBaseConfig.DbTablePrefix = viper.GetString("database.TABLE_PREFIX")
	MysqlDateBaseConfig.DbType = viper.GetString("database.TYPE")
	MysqlDateBaseConfig.DbPassword = viper.GetString("database.PASSWORD")
}

// 加载服务器配置
func loadServerConfig() {
	ServerConf.HTTPPort = viper.GetInt("server.HTTP_PORT")
	ServerConf.ReadTimeout = time.Duration(viper.GetInt("server.READ_TIMEOUT")) * time.Second
	ServerConf.WriteTimeout = time.Duration(viper.GetInt("server.WRITE_TIMEOUT")) * time.Second
}

// 加载email配置
func loadEmailConfig() {
	EmailConf.From = viper.GetString("email.FROM")
	EmailConf.SmtpHost = viper.GetString("email.SMTP_HOST")
	EmailConf.SmtpPassword = viper.GetString("email.SMTP_PASSWORD")
	EmailConf.SmtpAddr = viper.GetString("email.SMTP_ADDR")
	EmailConf.SmtpUsername = viper.GetString("email.SMTP_USERNAME")
}

// 加载redis配置
func loadRedisConfig() {
	RedisConf.Host = viper.GetString("redis.REDIS_HOST")
	RedisConf.Port = viper.GetString("redis.REDIS_PORT")
	RedisConf.Network = viper.GetString("redis.REDIS_NETWORK")
}

// 加载logger日志配置
func loadLoggerConfig() {
	LoggerConf.LogPath = viper.GetString("logger.LOG_PATH")
	LoggerConf.LogSaveName = viper.GetString("logger.LOG_SAVE_NAME")
	LoggerConf.LogFileExt = viper.GetString("logger.LOG_FILE_EXT")
}
