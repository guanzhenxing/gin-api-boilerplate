package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"gin-api-boilerplate/pkg/log"
	"gin-api-boilerplate/util"
)

// 配置
type Config struct {
	DbConfig     dbConfig     //数据库配置
	ServerConfig serverConfig //服务配置
}

var C = Config{}

//初始化
func Init(cfpath string) error {

	// 初始化配置文件
	if err := C.initConfig(cfpath); err != nil {
		return err
	}

	//初始化服务配置
	if err := C.initServerConfig(); err != nil {
		return err
	}

	// 初始化日志包
	if err := C.initLog(); err != nil {
		return err
	}

	//初始化DB
	if err := C.initDbConfig(); err != nil {
		return err
	}

	return nil
}

// 初始化Viper
func (c *Config) initConfig(cfgFile string) error {

	//设置并解析配置文件。
	//多个配置文件可以参考：https://stackoverrun.com/cn/q/12929186
	if len(cfgFile) > 0 {
		viper.SetConfigFile(cfgFile) // 如果指定了配置文件，则解析指定的配置文件
	} else { // 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("./conf")      // optionally look for config in the working directory
		viper.SetConfigName("application") // name of config file (without extension)
	}

	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Viper file not found; ignore error if desired
		} else {
			// Viper file was found but another error was produced
		}
		return err
	}

	//从环境变量读取配置
	viper.AutomaticEnv()                                   // 读取匹配的环境变量
	viper.SetEnvPrefix(viper.GetString("server.name"))     // 读取环境变量的前缀
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) //将 _ 换成 .

	// 监控配置文件变化并热加载程序
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Sugar().Infof("Viper file %s changed .", e.Name)
	})
	return nil
}

//初始化日志
func (c *Config) initLog() error {

	logCfg := log.Cfg{}

	err := viper.UnmarshalKey("log", &logCfg)
	if err != nil {
		return err
	}
	logCfg.Development = util.IsDebugMode(viper.GetString("env"))

	log.Init(logCfg)

	defer log.Sync()

	return nil
}
