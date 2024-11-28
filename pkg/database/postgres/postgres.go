package postgres

import (
	"context"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"github.com/jackc/pgx/v5"
	"log"
)

func InitPostgresConnection() {

	config := global.Config.Postgres
	// "postgres://%s:%s@localhost:%d/%s?timezone=Local"
	dsn := "postgres://%s:%s@localhost:%d/%s?timezone=Local"
	var s = fmt.Sprintf(dsn, config.UserName, config.Password, config.Host, config.Port)
	conn, err := pgx.Connect(context.Background(), s)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	defer conn.Close(context.Background())

}
