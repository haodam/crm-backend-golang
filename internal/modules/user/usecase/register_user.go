package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/haodam/user-backend-golang/common"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/modules/user"
	database "github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/utils/crypto"
	"github.com/haodam/user-backend-golang/utils/random"
	"github.com/haodam/user-backend-golang/utils/sendto"
	utils "github.com/haodam/user-backend-golang/utils/string"
	"log"
	"strconv"
	"strings"
	"time"
)

type IUserRegister interface {
	Register(ctx context.Context, VerifyKey string, VerifyType int, VerifyPurpose string) *common.Error
}

type registerUserUseCase struct {
	d *database.Queries
}

func NewRegisterUserUseCase(d *database.Queries) IUserRegister {
	return &registerUserUseCase{d: d}
}

var _ IUserRegister = (*registerUserUseCase)(nil)

// VerifyKey     string `json:"verify_key"`
// VerifyType    int    `json:"verify_type"`
// VerifyPurpose string `json:"verify_purpose"`

func (r registerUserUseCase) Register(ctx context.Context, VerifyKey string, VerifyType int, VerifyPurpose string) *common.Error {

	// Step1: Hash Email
	fmt.Printf("VerifyKey: %s\n", VerifyKey)
	fmt.Printf("VerifyType: %d\n", VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	// Step2: Check user exists in uer base
	userFound, err := r.d.CheckUserBaseExists(ctx, VerifyKey)
	if err != nil {
		return &common.Error{
			Message:      fmt.Sprintf("user %v already exists", VerifyKey),
			DebugMessage: err.Error(),
			Code:         ErrCodeUserHasExists,
		}
	}
	// check email already registered (example@gmail.com)
	if userFound > 0 {
		return &common.Error{
			Message: fmt.Sprintf("user %v already registered", VerifyKey),
			Code:    ErrCodeUserHasExists,
		}
	}

	// Step3: Create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == nil:
		fmt.Println("Key doesn't exist")
	case err != nil:
		fmt.Println("get otp for user failed", VerifyKey)
		return &common.Error{
			DebugMessage: err.Error(),
			Code:         ErrInvalidOTP,
		}
	case otpFound != "":
		return &common.Error{
			DebugMessage: otpFound,
			Code:         ErrCodeOtpNotExists,
		}
	}

	// Step4: Generate OTP
	otpNew := random.GenerateSixDigOtp()
	if VerifyPurpose == "TEST_USER" {
		otpNew = 123456 // Hard code
	}
	fmt.Printf("OTP is :::%d\n", otpNew)

	// Step5: Save OTP in Redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(2)*time.Minute).Err()
	if err != nil {
		return &common.Error{
			DebugMessage: err.Error(),
			Code:         ErrInvalidOTP,
		}
	}
	// Step6: Send OTP
	switch VerifyType {
	case user.EMAIL:
		// Hard code to email (example@gmail.com)
		err := sendto.SendTextEmailOtp([]string{VerifyKey}, user.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return &common.Error{
				DebugMessage: err.Error(),
				Code:         ErrSendEmailOtp,
			}
		}
		// Step7: Save OTP to MYSQL
		result, err := r.d.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyKey:     VerifyKey,
			VerifyKeyHash: hashKey,
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
		})
		if err != nil {
			return &common.Error{
				DebugMessage: err.Error(),
				Code:         ErrSendEmailOtp,
			}
		}
		// step8: gelatosID
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return &common.Error{
				DebugMessage: err.Error(),
				Code:         ErrSendEmailOtp,
			}
		}
		log.Println("lastIdVerifyUser", lastIdVerifyUser)
		return nil
	case user.MOBILE:
		//TO DO
		return nil
	default:
		fmt.Println("unhandled default case")
	}

	return nil
}
