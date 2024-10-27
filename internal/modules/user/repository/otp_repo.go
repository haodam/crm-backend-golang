package repository

import (
	"context"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"time"
)

var ctx = context.Background()

type IOtpRegisterRepository interface {
	GenOTP(email string, otp int, expirationTime int64) error
}

type otpRegisterRepository struct{}

var _ IOtpRegisterRepository = (*otpRegisterRepository)(nil)

func NewOtpRegisterRepository() IOtpRegisterRepository {
	return &otpRegisterRepository{}

}

func (or *otpRegisterRepository) GenOTP(email string, otp int, expirationTime int64) error {

	key := fmt.Sprintf("usr:%s:otp", email) // usr:email:otp

	return global.Rdb.SetEx(ctx, key, otp, time.Duration(expirationTime)*time.Second).Err()
}
