package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/global"
)

func InitRouter() gin.Engine {

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

	managerRouter := router.ManagerRouter{}

}
