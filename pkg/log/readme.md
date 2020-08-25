
# log

## 使用示例



## uber-go/zap 相关

GitHub地址：https://github.com/uber-go/zap

### 示例：

自定义示例

```go

package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "fmt"
    "time"
)

func main() {
    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
    }

    // 设置日志级别
    atom := zap.NewAtomicLevelAt(zap.DebugLevel)

    config := zap.Config{
        Level:            atom,                                                // 日志级别
        Development:      true,                                                // 开发模式，堆栈跟踪
        Encoding:         "json",                                              // 输出格式 console 或 json
        EncoderConfig:    encoderConfig,                                       // 编码器配置
        InitialFields:    map[string]interface{}{"serviceName": "spikeProxy"}, // 初始化字段，如：添加一个服务器名称
        OutputPaths:      []string{"stdout", "./logs/spikeProxy.log"},         // 输出到指定文件 stdout（标准输出，正常颜色） stderr（错误输出，红色）
        ErrorOutputPaths: []string{"stderr"},
    }

    // 构建日志
    logger, err := config.Build()
    if err != nil {
        panic(fmt.Sprintf("log 初始化失败: %v", err))
    }
    logger.Info("log 初始化成功")

    logger.Info("无法获取网址",
        zap.String("url", "http://www.baidu.com"),
        zap.Int("attempt", 3),
        zap.Duration("backoff", time.Second),
    )
}

```

写入归档文件示例

Lumberjack是一个Go包，用于将日志写入滚动文件。
zap 不支持文件归档，如果要支持文件按大小或者时间归档，需要使用lumberjack，lumberjack也是[zap官方推荐](https://github.com/uber-go/zap/blob/master/FAQ.md)的。

```go
package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "time"
    "gopkg.in/natefinch/lumberjack.v2"
    "os"
)

func main() {
    hook := lumberjack.Logger{
        Filename:   "./logs/spikeProxy1.log", // 日志文件路径
        MaxSize:    128,                      // 每个日志文件保存的最大尺寸 单位：M
        MaxBackups: 30,                       // 日志文件最多保存多少个备份
        MaxAge:     7,                        // 文件最多保存多少天
        Compress:   true,                     // 是否压缩
    }

    encoderConfig := zapcore.EncoderConfig{
        TimeKey:        "time",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "linenum",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
        EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
        EncodeDuration: zapcore.SecondsDurationEncoder, //
        EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
        EncodeName:     zapcore.FullNameEncoder,
    }

    // 设置日志级别
    atomicLevel := zap.NewAtomicLevel()
    atomicLevel.SetLevel(zap.InfoLevel)

    core := zapcore.NewCore(
        zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
        zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
        atomicLevel,                                                                     // 日志级别
    )

    // 开启开发模式，堆栈跟踪
    caller := zap.AddCaller()
    // 开启文件及行号
    development := zap.Development()
    // 设置初始化字段
    filed := zap.Fields(zap.String("serviceName", "serviceName"))
    // 构造日志
    logger := zap.New(core, caller, development, filed)

    logger.Info("log 初始化成功")
    logger.Info("无法获取网址",
        zap.String("url", "http://www.baidu.com"),
        zap.Int("attempt", 3),
        zap.Duration("backoff", time.Second))
}
```