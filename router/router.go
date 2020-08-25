package router

import (
	"gin-api-boilerplate/config"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "gin-api-boilerplate/docs" // docs is generated by Swag CLI, you have to import it.
	"gin-api-boilerplate/handler/actuator"
	"gin-api-boilerplate/handler/user"
	"gin-api-boilerplate/router/middleware"
	"gin-api-boilerplate/util"
)

func Init(g *gin.Engine) {

	//初始化静态资源
	initStatic(g)

	//初始化全局中间件
	initGlobalMiddleware(g)

	//初始化handler
	initHandler(g)

}

func initStatic(g *gin.Engine) {

}

func initGlobalMiddleware(g *gin.Engine) {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)

	g.Use(middleware.Logging())
	g.Use(middleware.RequestId())
}

func initHandler(g *gin.Engine) {

	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	if util.IsDebugMode(config.C.ServerConfig.Env) || util.IsTestMode(config.C.ServerConfig.Env) {
		pprof.Register(g)                                                    // pprof router
		g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // swagger api docs
	}

	initActuatorHandlers(g)

	// api for authentication functionalities
	g.POST("/login", user.Login)

	// The user handlers, requiring authentication
	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.POST("", user.Create)
		u.DELETE("/:id", user.Delete)
		u.PUT("/:id", user.Update)
		u.GET("", user.List)
		u.GET("/:username", user.Get)
	}

}

// The health check handlers
func initActuatorHandlers(g *gin.Engine) {
	a := g.Group("/actuator")
	{
		a.GET("/health", actuator.HealthCheck)
		a.GET("/disk", actuator.DiskCheck)
		a.GET("/cpu", actuator.CPUCheck)
		a.GET("/ram", actuator.RAMCheck)
	}
}