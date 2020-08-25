# gin-api-boilerplate

基于 Gin 的 API 项目模板。

该项目根据[基于 Go 语言构建企业级的 RESTful API 服务](https://juejin.im/book/5b0778756fb9a07aa632301e)改写。

## 功能

- 基于 [gin](https://github.com/gin-gonic/gin) Web 框架
- 使用 [pflag](https://github.com/spf13/pflag) 替换 Go 的 flag 包
- 使用 [Viper](https://github.com/spf13/viper) 读取配置文件
- 使用 [uber-go/zap](https://github.com/uber-go/zap) 记录和管理 API 日志
- 使用 [gorm](https://github.com/jinzhu/gorm) 持久化数据
- 使用 [redigo](https://github.com/gomodule/redigo) 操作redis
- 使用 [go-playground/validator](https://github.com/go-playground/validator) 做参数校验
- 定义通用的“自定义业务错误信息” [pkg/ecode](./pkg/ecode)
- 使用 pprof 性能分析
- 生成 Swagger 在线文档
- 用 Makefile 管理 API 项目，用脚本运行项目
- 给 API 命令增加版本功能
- 自定义中间件
    - 每个请求都加上 X-Request-Id
    

### TODO 

- [接口签名（JWT or TOKEN or ...）](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b1862375188257d3b39f939)
- [用HTTPS加密API请求](https://juejin.im/book/5b0778756fb9a07aa632301e/section/5b186251e51d4506dc0abb5c)
- [将应用部署到Docker](https://segmentfault.com/a/1190000013960558)
- [预防CSRF攻击](https://github.com/utrack/gin-csrf)



## 可以使用的中间件

- [官方自带的中间件](https://github.com/gin-contrib)
- [Collection of middlewares created by the community](https://github.com/gin-gonic/contrib)

## 项目目录
```text
.
├── Makefile                # Makefile文件，一般大型软件系统都是采用make来作为编译工具
├── README.md               # API目录README
├── admin.sh                # 进程的start|stop|status|restart控制文件
├── conf                    # 配置文件统一存放目录
│   ├── application.yml     # 配置文件
│   ├── server.crt          # TLS配置文件
│   └── server.key
├── config                  # 专门用来处理配置和配置文件的Go package
│   ├── config.go           # 配置
│   ├── config_db.go        # 数据库配置
│   ├── config_log.go       # 日志配置
│   └── config_rds.go       # Redis配置
├── docs
│   └── docs.go
├── handler                 # 类似MVC架构中的C，用来读取输入，并将处理流程转发给实际的处理函数，最后返回结果
│   ├── actuator            # 健康检查handler
│   └── handler.go
├── models                  # 数据库相关的操作统一放在这里，包括数据库初始化和对表的增删改查
│   └── init.go             # 初始化和连接数据库
├── main.go                 # Go程序唯一入口
├── pkg                     # 初始化和连接数据库
│   ├── constvar            # 常量统一存放位置
│   ├── ecode               # 错误码存放位置
│   ├── log                 # 日志包
│   ├── redis               # 封装Redis
│   └── version             # 版本包
├── router                  # 路由相关处理
│   ├── middleware          # API服务器用的是Gin Web框架，Gin中间件存放位置
│   └── router.go
├── service                 # 实际业务处理函数存放位置
│   └── service.go
└── util                    # 工具类函数存放目录
    └── util.go
```

