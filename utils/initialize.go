package utils

import (
	"fmt"
	"github.com/haodam/user-backend-golang/configs"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/pkg/database/mysql"
	"github.com/haodam/user-backend-golang/pkg/logger"
	"github.com/haodam/user-backend-golang/pkg/transports/https/routers"
	"go.uber.org/zap"
)

func Initialize() {

	configs.MustLoadConfig()
	fmt.Println("Loading configs...", global.Config.Mysql.Username)

	global.Logger = logger.NewLogger(global.Config.Logger)
	global.Logger.Info("Config Log ok !!", zap.String("ok", "success"))

	mysql.InitMysql()

	r := routers.NewRouter()
	err := r.Run(":8082")
	if err != nil {
		return
	}
}
