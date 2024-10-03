package redis

import (
	"context"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"go.uber.org/zap"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func InitRedis() {
	cfg := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", cfg.Host, cfg.Port),
		Password: cfg.Password, // no password set
		DB:       cfg.Database, // use default DB
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Redis init error", zap.Error(err))
	}

	fmt.Println("Redis init success")
	global.Rdb = rdb

	//err := rdb.Set(ctx, "key", "value", 0).Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//val, err := rdb.Get(ctx, "key").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key", val)
	//
	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
	// Output: key value
	// key2 does not exist
}
