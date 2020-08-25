package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	Logger *zap.Logger
)

type Cfg struct {
	Development      bool
	Encoding         string   //编码方式  console, json
	Level            string   //日志级别
	Writers          []string //日志输出到哪些地方 stdout,file 等
	File             string   //日志文件
	MaxSize          int      // 每个日志文件保存的最大尺寸 单位：M
	MaxAge           int      // 日志文件最多保存多少个备份
	MaxBackups       int      // 文件最多保存多少天
	Compress         bool     //是否压缩
	ErrorOutputPaths []string
}

var (
	cfg       Cfg
	zapConfig zap.Config
)

func Init(logConfig Cfg) {

	cfg = logConfig

	if cfg.Development { //根据是否是开发模式，获得对应的配置信息
		zapConfig = zap.NewDevelopmentConfig()
	} else {
		zapConfig = zap.NewProductionConfig()
	}

	Logger = zap.New(getCore()) //创建Logger

}

// 获得core
//todo 不同等级的输出到不同的文件
func getCore() zapcore.Core {

	enc := getEncoder()
	ws := getLogWriter()
	level := getLevel(cfg.Level)

	core := zapcore.NewCore(enc, zapcore.NewMultiWriteSyncer(ws...), level)
	return core
}

func getLogWriter() []zapcore.WriteSyncer {
	var ws []zapcore.WriteSyncer
	if len(cfg.Writers) < 1 {
		ws = append(ws, zapcore.AddSync(os.Stdout))
	} else {
		for _, v := range cfg.Writers {
			if v == "stdout" {
				ws = append(ws, zapcore.AddSync(os.Stdout))
			} else if v == "file" {
				lumberJackLogger := lumberjack.Logger{
					Filename:   cfg.File,
					MaxSize:    cfg.MaxSize,
					MaxBackups: cfg.MaxBackups,
					MaxAge:     cfg.MaxAge,
					Compress:   cfg.Compress,
				}
				ws = append(ws, zapcore.AddSync(&lumberJackLogger))
			}
		}
	}
	return ws
}

func getEncoder() zapcore.Encoder {
	var enc zapcore.Encoder
	if cfg.Encoding == "json" {
		enc = zapcore.NewJSONEncoder(zapConfig.EncoderConfig)
	} else if cfg.Encoding == "console" {
		enc = zapcore.NewConsoleEncoder(zapConfig.EncoderConfig) //编码
	}
	return enc
}

func getLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}

//http://vearne.cc/?s=zap
