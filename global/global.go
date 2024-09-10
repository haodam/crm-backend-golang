package global

import (
	"github.com/haodam/user-backend-golang/pkg/logger"
	"github.com/haodam/user-backend-golang/pkg/setting"
)

var (
	Config setting.Config
	Logger *logger.ZapLogger
)
