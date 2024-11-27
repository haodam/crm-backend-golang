package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/modules/user/model"
	database "github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/utils/crypto"
	string2 "github.com/haodam/user-backend-golang/utils/string"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type IVerifyUserRegister interface {
	VerifyOTP(ctx context.Context, req *model.VerifyOTPInput) (out *model.VerifyOTPOutput, err error)
}
type verifyUserUseCase struct {
	d *database.Queries
}

var _ IVerifyUserRegister = (*verifyUserUseCase)(nil)

func NewVerifyUserUseCase(d *database.Queries) IVerifyUserRegister {
	return &verifyUserUseCase{d: d}
}

func (v *verifyUserUseCase) VerifyOTP(ctx context.Context, req *model.VerifyOTPInput) (out *model.VerifyOTPOutput, err error) {

	// hash email
	hashKey := crypto.GetHash(strings.ToLower(req.VerifyKey))

	// VD: Neu otpKey la u:123abc:otp, thi attemptKey la u:123abc:otp:attempts.
	otpKey := string2.GetUserKey(hashKey)
	attemptKey := fmt.Sprintf("%s:attempts", otpKey) // u:<hashKey>:otp:attempts

	// Kiem tra gioi han so lan thu
	attempts, _ := global.Rdb.Get(ctx, attemptKey).Int()
	if attempts >= 5 {
		return out, err
	}

	// Get OTP
	otpFound, err := global.Rdb.Get(ctx, hashKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return out, err
		}
		return out, err
	}

	if req.VerifyCode != otpFound {
		_, _ = global.Rdb.Incr(ctx, attemptKey).Result()
		_ = global.Rdb.Expire(ctx, attemptKey, time.Minute)
		return out, err
	}

	// Neu OTP chinh xac , xoa bo dem so lan thu
	_ = global.Rdb.Del(ctx, attemptKey).Err()

	infoOTP, err := v.d.GetInfoOTP(ctx, req.VerifyKey)
	if err != nil {
		return out, err
	}

	// Update status verified
	err = v.d.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// Out put
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"

	return out, nil
}
