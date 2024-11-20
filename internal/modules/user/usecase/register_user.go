package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/modules/user"
	"github.com/haodam/user-backend-golang/internal/modules/user/model"
	database "github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/pkg/response"
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
	Register(ctx context.Context, arg *model.RegisterEntity) (codeResult int, err error)
}

type registerUserUseCase struct {
	d *database.Queries
}

func NewRegisterUserUseCase(d *database.Queries) IUserRegister {
	return &registerUserUseCase{d: d}
}

var _ IUserRegister = (*registerUserUseCase)(nil)

// VerifyKey     string `json:"verify_key"` la email
// VerifyType    int    `json:"verify_type"` 1 la dang ky bang email , 2 la dang ky bang so dt
// VerifyPurpose string `json:"verify_purpose"` TEST_USER

func (r *registerUserUseCase) Register(ctx context.Context, arg *model.RegisterEntity) (codeResult int, err error) {

	// Step1: Hash Email
	fmt.Printf("VerifyKey: %s\n", arg.VerifyKey)
	fmt.Printf("VerifyType: %d\n", arg.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(arg.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	// Step2: Check user exists in uer base
	userFound, err := r.d.CheckUserBaseExists(ctx, arg.VerifyKey)
	if err != nil {
		return response.ErrCodeUserHasExists, err
	}
	// check email already registered (example@gmail.com)
	if userFound > 0 {
		return response.ErrCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	// Step3: Create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Rdb.Get(ctx, userKey).Result()

	switch {
	case err == nil:
		fmt.Println("Key doesn't exist")
	case err != nil:
		fmt.Println("get otp for user failed", arg.VerifyKey)
	case otpFound != "":
		return response.ErrCodeOtpNotExists, fmt.Errorf("")
	}

	// Step4: Generate OTP
	otpNew := random.GenerateSixDigOtp()
	if arg.VerifyPurpose == "TEST_USER" {
		otpNew = 123456 // Hard code
	}
	fmt.Printf("OTP is :::%d\n", otpNew)

	// Step5: Save OTP in Redis with expiration time
	err = global.Rdb.SetEx(ctx, userKey, strconv.Itoa(otpNew), time.Duration(2)*time.Minute).Err()
	if err != nil {
		return response.ErrInvalidOTP, err
	}
	// Step6: Send OTP
	switch arg.VerifyType {
	case user.EMAIL:
		// Hard code to email (example@gmail.com)
		err := sendto.SendTextEmailOtp([]string{arg.VerifyKey}, user.HOST_EMAIL, strconv.Itoa(otpNew))
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		// Step7: Save OTP to MYSQL
		result, err := r.d.InsertOTPVerify(ctx, database.InsertOTPVerifyParams{
			VerifyOtp:     strconv.Itoa(otpNew),
			VerifyKey:     arg.VerifyKey,
			VerifyKeyHash: hashKey,
			VerifyType:    sql.NullInt32{Int32: 1, Valid: true},
		})
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		// step8: gelatosID
		lastIdVerifyUser, err := result.LastInsertId()
		if err != nil {
			return response.ErrSendEmailOtp, err
		}
		log.Println("lastIdVerifyUser", lastIdVerifyUser)
		return response.ErrCodeSuccess, nil
	case user.MOBILE:
		//TO DO
		return response.ErrCodeSuccess, nil
	default:
		fmt.Println("unhandled default case")
	}

	return response.ErrCodeSuccess, nil
}
