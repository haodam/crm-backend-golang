package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/modules/user"
	"github.com/haodam/user-backend-golang/internal/modules/user/model"
	"github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/pkg/response"
	"github.com/haodam/user-backend-golang/utils/crypto"
	"github.com/haodam/user-backend-golang/utils/random"
	"github.com/haodam/user-backend-golang/utils/sendto"
	string2 "github.com/haodam/user-backend-golang/utils/string"
	utils "github.com/haodam/user-backend-golang/utils/string"
	"github.com/redis/go-redis/v9"
)

type sUserAuthed struct {
	r *repository.Queries
}

func NewAuthedUserUseCase(r *repository.Queries) *sUserAuthed {
	return &sUserAuthed{r: r}
}

// VerifyKey     string `json:"verify_key"` la email
// VerifyType    int    `json:"verify_type"` 1 la dang ky bang email , 2 la dang ky bang so dt
// VerifyPurpose string `json:"verify_purpose"` TEST_USER

func (s *sUserAuthed) Register(ctx context.Context, arg *model.RegisterEntity) (codeResult int, err error) {

	// Step1: Hash Email
	fmt.Printf("VerifyKey: %s\n", arg.VerifyKey)
	fmt.Printf("VerifyType: %d\n", arg.VerifyType)
	hashKey := crypto.GetHash(strings.ToLower(arg.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	// Step2: Check user exists in uer base
	userFound, err := s.r.CheckUserBaseExists(ctx, arg.VerifyKey)
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
		result, err := s.r.InsertOTPVerify(ctx, repository.InsertOTPVerifyParams{
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

func (s *sUserAuthed) VerifyOTP(ctx context.Context, req *model.VerifyOTPInput) (out *model.VerifyOTPOutput, err error) {

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

	infoOTP, err := s.r.GetInfoOTP(ctx, req.VerifyKey)
	if err != nil {
		return out, err
	}

	// Update status verified
	err = s.r.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// Out put
	out.Token = infoOTP.VerifyKeyHash
	out.Message = "success"

	return out, nil
}

func (s *sUserAuthed) UpdatePasswordRegister(ctx context.Context, token string, Password string) (userId int, err error) {

	// Step 1 token is already verified : user_verify table
	infoOTP, err := s.r.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// step 1.1 check is verified OK
	if infoOTP.IsVerified.Int32 == 0 {
		return response.ErrCodeUserOtpNotExists, fmt.Errorf("user OTP not verified")
	}

	// step 2 check token ins exists in user_base
	// update user_base table
	log.Println("infoOTP:", infoOTP)
	userBase := repository.AddUserBaseParams{}
	userBase.UserAccount = infoOTP.VerifyKey
	userSalt, err := crypto.GenerateSalt(16)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	userBase.UserSalt = userSalt
	userBase.UserPassword = crypto.HashPassword(Password, userSalt)

	// add userBase to user_base table
	newUserBase, err := s.r.AddUserBase(ctx, userBase)
	log.Println("newUserBase::", newUserBase, userBase)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	//user_id, err := newUserBase.LastInsertId()
	//if err != nil {
	//	return response.ErrCodeUserOtpNotExists, err
	//}
	//
	//// add user_id to user info table
	//newUserInfo, err := u.r.AddUserBase()
	return 0, nil
}
