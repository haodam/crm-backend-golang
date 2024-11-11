package global

import (
	"database/sql"
	"github.com/haodam/user-backend-golang/pkg/logger"
	"github.com/haodam/user-backend-golang/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.ZapLogger
	Rdb    *redis.Client
	Mdb    *gorm.DB
	MdbC   *sql.DB
)
