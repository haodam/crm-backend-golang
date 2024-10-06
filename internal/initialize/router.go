package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/global"
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
