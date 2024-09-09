package utils

import (
	"fmt"
	"github.com/haodam/user-backend-golang/configs"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/pkg/transports/https/routers"
)

func Initialize() {

	configs.MustLoadConfig()
	fmt.Println("Loading configs...", global.Config.Mysql.Username)

	r := routers.NewRouter()
	err := r.Run(":8082")
	if err != nil {
		return
	}
}
