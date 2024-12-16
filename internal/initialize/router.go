package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/middleware"
	"github.com/haodam/user-backend-golang/internal/roters"
)

func InitRouter() *gin.Engine {

	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middleware
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global

	r.Use(middleware.NewRateLimiter().GlobalRateLimiter())
	r.GET("/ping/100", func(c *gin.Context) { // 100 req/s
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.NewRateLimiter().PublicAPIRateLimiter())
	r.GET("/ping/80", func(c *gin.Context) { // 80 req/s
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.NewRateLimiter().UserAndPrivateAPIRateLimiter())
	r.GET("/ping/60", func(c *gin.Context) { // 60 req/s
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	managerRouter := roters.RouterGroupApp.Manager
	userRouter := roters.RouterGroupApp.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/check_status")
	}
	{
		managerRouter.InitUserRouter(MainGroup)
		managerRouter.InitAdminRouter(MainGroup)
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}

	return r
}
