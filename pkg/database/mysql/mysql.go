package mysql

import (
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func checkErrPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysql() {

	var m = global.Config.Mysql
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{})
	if err != nil {
		checkErrPanic(err, "InitMysql initiaLization error")
	}
	global.Logger.Info("InitMysql Successfully")
	global.Mdb = db

	// Set Pool
	SetPool()
}

// Mo nhom ket noi , giup cai thien hieu suat

func SetPool() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Printf("mysql error : %s ::", err)
		global.Logger.Error("InitMysql error", zap.Error(err))
	}
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime) * time.Second)

}

func migrateTables() {

}
