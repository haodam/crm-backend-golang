package postgres

import (
	"context"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"log"
)

func checkErrPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitPostgresConnection() {

	config := global.Config.Postgres
	// "postgres://%s:%s@localhost:%d/%s?timezone=Local"
	dsn := "postgres://%s:%s@%s:%d/%s?timezone=Local"
	var s = fmt.Sprintf(dsn, config.UserName, config.Password, config.Host, config.Port, config.Database)
	conn, err := pgx.Connect(context.Background(), s)
	if err != nil {
		log.Fatal("cannot connect to db")
	} else {
		fmt.Print("Successfully connected to db")
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {
			return
		}
	}(conn, context.Background())
}
