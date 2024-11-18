package main

import (
	"github.com/haodam/user-backend-golang/internal/initialize"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           API Documentation Ecommerce Backend SHOP
// @version         1.0.0
// @description     This is a sample server seller server.
// @termsOfService  https://github.com/haodam/crm-backend-golang

// @contact.name   DAM ANH HAO
// @contact.url    https://github.com/haodam/crm-backend-golang
// @contact.email  example@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8002
// @BasePath  /v1/2024
// @schema http
func main() {
	r := initialize.Initialize()
	// http://localhost:8002/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err := r.Run(":8002")
	if err != nil {
		return
	}
}
