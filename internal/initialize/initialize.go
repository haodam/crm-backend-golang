package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/pkg/database/mysql"
	"github.com/haodam/user-backend-golang/pkg/database/postgres"
	"github.com/haodam/user-backend-golang/pkg/logger"
	"github.com/haodam/user-backend-golang/pkg/redis"
	"github.com/haodam/user-backend-golang/utils/configs"
	"go.uber.org/zap"
)

func Initialize() *gin.Engine {

	configs.MustLoadConfig()
	fmt.Println("Loading configs...", global.Config.Mysql.Username)

	global.Logger = logger.NewLogger(global.Config.Logger)
	global.Logger.Info("Config Log ok !!", zap.String("ok", "success"))

	//mysql.InitMysql()
	mysql.InitMysqlC()
	postgres.InitPostgresConnection()
	redis.InitRedis()

	r := InitRouter()
	err := r.Run(":8002")
	if err != nil {
		return nil
	}
	return r
}
