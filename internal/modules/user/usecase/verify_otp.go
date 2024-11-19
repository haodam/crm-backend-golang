package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	database "github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/utils/crypto"
	string2 "github.com/haodam/user-backend-golang/utils/string"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type IVerifyUserRegister interface {
	VerifyOTP(ctx context.Context, verifyKey string, verifyCode string) (token string, UserId string, message string, err error)
}
type verifyUserUseCase struct {
	d *database.Queries
}

var _ IVerifyUserRegister = (*verifyUserUseCase)(nil)

func NewVerifyUserUseCase(d *database.Queries) IVerifyUserRegister {
	return &verifyUserUseCase{d: d}
}

func (v *verifyUserUseCase) VerifyOTP(ctx context.Context, verifyKey string, verifyCode string) (token string, UserId string, message string, err error) {

	// hash email
	hashKey := crypto.GetHash(strings.ToLower(verifyKey))

	// VD: Neu otpKey la u:123abc:otp, thi attemptKey la u:123abc:otp:attempts.
	otpKey := string2.GetUserKey(hashKey)
	attemptKey := fmt.Sprintf("%s:attempts", otpKey) // u:<hashKey>:otp:attempts

	// Kiem tra gioi han so lan thu
	attempts, _ := global.Rdb.Get(ctx, attemptKey).Int()
	if attempts >= 5 {
		return "", "", "", fmt.Errorf("too many failed attempts, please try again later")
	}

	// Get OTP
	otpFound, err := global.Rdb.Get(ctx, hashKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return "", "", "", fmt.Errorf("OTP expired or not found")
		}
		return "", "", "", err
	}

	if verifyCode != otpFound {
		_, _ = global.Rdb.Incr(ctx, attemptKey).Result()
		_ = global.Rdb.Expire(ctx, attemptKey, time.Minute)
		return "", "", "", fmt.Errorf("otp verification failed")
	}

	// Neu OTP chinh xac , xoa bo dem so lan thu
	_ = global.Rdb.Del(ctx, attemptKey).Err()

	infoOTP, err := v.d.GetInfoOTP(ctx, verifyKey)
	if err != nil {
		return "", "", "", err
	}

	// Update status verified
	err = v.d.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return "", "", "", err
	}

	// Out put
	token = infoOTP.VerifyKeyHash
	message = "success"

	return token, "", message, nil
}
