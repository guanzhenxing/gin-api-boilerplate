package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gin-api-boilerplate/util"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gin-api-boilerplate/config"
	"gin-api-boilerplate/model"
	"gin-api-boilerplate/pkg/log"
	v "gin-api-boilerplate/pkg/version"
	"gin-api-boilerplate/router"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfpath  = pflag.StringP("config", "c", "", "gin-api-boilerplate config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

var httpServer *http.Server

func main() {

	if loadOption() {
		return
	}

	// init db
	model.DB.Init()
	defer model.DB.Close()

	runServer()
}

func loadOption() bool {
	pflag.Parse()
	if *version {
		v := v.Get()
		marshalled, err := json.MarshalIndent(&v, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return true
	}

	// init config
	if err := config.Init(*cfpath); err != nil {
		panic(err)
	}
	return false
}

func runServer() {

	// Create the Gin engine.
	g := gin.New()
	g.Use(gin.Recovery())

	// Set gin mode.
	if util.IsDebugMode(config.C.ServerConfig.Env) {
		gin.SetMode(gin.DebugMode)
	}
	if util.IsTestMode(config.C.ServerConfig.Env) {
		gin.SetMode(gin.TestMode)
	}
	if util.IsReleaseMode(config.C.ServerConfig.Env) {
		gin.SetMode(gin.ReleaseMode)
	}

	// Routes.
	router.Init(g)

	addr := fmt.Sprintf(":%d", config.C.ServerConfig.Port)
	httpServer = &http.Server{
		Addr:    addr,
		Handler: g,
	}

	// Start to listening the incoming requests.
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Info("Start to listening the incoming requests on https address:" + viper.GetString("tls.addr"))
			log.Info(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}

	go func() {
		log.Info("Start to listening the incoming requests on http address:" + addr)
		err := httpServer.ListenAndServe().Error()
		log.Sugar().Infof(err)
	}()

	signalHandler()
	pingServerHandler()
}

// 定时请求检查健康接口，确保路由可工作
func pingServerHandler() {

	var ping = func() error {
		for i := 0; i < config.C.ServerConfig.Ping.MaxTryNumber; i++ {
			// PingConfig the server by sending a GET request to `/health`.
			resp, err := http.Get(config.C.ServerConfig.Ping.Url + "/actuator/health")
			if err == nil && resp.StatusCode == 200 {
				return nil
			}
			// Sleep for a second to continue the next ping.
			log.Info("Waiting for the router, retry in 1 second.")
			time.Sleep(time.Second)
		}
		return errors.New("CannotConnectToTheRouter")
	}

	go func() {
		if err := ping(); err != nil {
			log.Sugar().Fatalf("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
}

//处理信号，为了优雅关闭服务
func signalHandler() {
	var (
		ch = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT:
			log.Sugar().Infof("Get a %s, stop the server process", si.String())

			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			if err := httpServer.Shutdown(ctx); err != nil {
				log.Sugar().Errorf("Http Server Shutdown Fail", err)
			}

			log.Logger.Debug("HttpServer Graceful Shutdown")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
