package database

type MysqlDatabaseConfig struct {
	DbType        string
	DbUser        string
	DbPassword    string
	DbHost        string
	DbPort        string
	DbDataName    string
	DbTablePrefix string
}
