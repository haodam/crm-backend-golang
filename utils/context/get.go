package context

import (
	"context"
	"errors"
	"github.com/haodam/user-backend-golang/utils/cache"
	"log"
)

type InfoUserUUID struct {
	UserId      uint64 `json:"user_id"`
	UserAccount string `json:"user_account"`
}

func GetSubjectUUID(ctx context.Context) (string, error) {
	sUUID, ok := ctx.Value("subjectUUID").(string)
	if !ok {
		return "", errors.New("subjectUUID not found")
	}
	return sUUID, nil
}

func GetUserIdFromUUID(ctx context.Context) (uint64, error) {
	sUUid, err := GetSubjectUUID(ctx)
	if err != nil {
		return 0, err
	}

	// get info User redis from uuid
	var infoUser InfoUserUUID
	if err := cache.GetCache(ctx, sUUid, &infoUser); err != nil {
		log.Println("err:::", err)
		return 0, err
	}
	return infoUser.UserId, nil

}
