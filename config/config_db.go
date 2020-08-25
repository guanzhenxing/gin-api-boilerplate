package config

import "github.com/spf13/viper"

type dbConfig struct {
	Host string
	//Port      uint16
	Username string
	Password string
	Database string
	//Charset   string
	//Collation string
	//Loc       string
	//ParseTime bool
	Gorm gormConfig
}
type gormConfig struct {
	LogMode bool
}

func (c *Config) initDbConfig() error {
	return viper.UnmarshalKey("mysql", &C.DbConfig)
}
