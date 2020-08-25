package config

import "github.com/spf13/viper"

type serverConfig struct {
	Name string
	Env  string
	Port int
	Ping pingConfig
}
type pingConfig struct {
	Url          string
	MaxTryNumber int
}

func (c *Config) initServerConfig() error {
	return viper.UnmarshalKey("server", &C.ServerConfig) //初始化Server配置
}
