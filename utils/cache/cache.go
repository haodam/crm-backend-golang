package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {

	rs, err := global.Rdb.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return fmt.Errorf("key %s not found", key)
	} else if err != nil {
		return err
	}
	// Convert rs to object
	if err := json.Unmarshal([]byte(rs), obj); err != nil {
		return fmt.Errorf("failed to unmarshal")
	}

	return nil
}
