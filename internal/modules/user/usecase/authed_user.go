package usecase

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/haodam/user-backend-golang/utils/auth"
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

	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// add user_id to user info table
	newUserInfo, err := s.r.AddUserHaveUserId(ctx, repository.AddUserHaveUserIdParams{
		UserID:               uint64(user_id),
		UserAccount:          infoOTP.VerifyKey,
		UserNickname:         sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Now(), Valid: false},
		UserEmail:            sql.NullString{String: infoOTP.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	})
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err = newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}
	return int(user_id), nil
}

func (s *sUserAuthed) Login(ctx context.Context, req *model.LoginInput) (codeResult int, out *model.LoginOutput, err error) {

	// Step1. Check user account exits in database user base
	useBase, err := s.r.GetOneUserInfo(ctx, req.UserAccount)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// Step2. Check password
	if !crypto.MatchingPassword(useBase.UserPassword, req.UserPassword, useBase.UserSalt) {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("user password not match")
	}

	// Step3. Check two-factor authentication
	isTwoFactorEnable, err := s.r.IsTwoFactorEnabled(ctx, uint32(useBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("does not match password")
	}
	if isTwoFactorEnable > 0 {
		// Send OTP to req.TwoFactorEmail
		keyUserLoginTwoFactor := crypto.GetHash("2fa:otp" + strconv.Itoa(int(useBase.UserID)))
		err = global.Rdb.SetEx(ctx, keyUserLoginTwoFactor, "111111", time.Duration(user.TIME_OTP_REGISTER)*time.Minute).Err()
		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("set otp redis failed")
		}
		// Send OTP via two factor Email
		// get email 2FA
		infoUserTwoFactor, err := s.r.GetTwoFactorMethodByIDAndType(ctx, repository.GetTwoFactorMethodByIDAndTypeParams{
			UserID:            uint32(useBase.UserID),
			TwoFactorAuthType: repository.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		})
		if err != nil {
			return response.ErrCodeAuthFailed, out, fmt.Errorf("get two factor method failed")
		}
		// go send.SendEmailInJavaByAPI
		log.Println("infoUserTwoFactor:", infoUserTwoFactor)
		go func() {
			err := sendto.SendTextEmailOtp([]string{infoUserTwoFactor.TwoFactorEmail.String}, user.HOST_EMAIL, "111111")
			if err != nil {
				return
			}
		}()
		out.Message = "send OTP 2FA to Email, pls het OTP by Email..."
		return response.ErrCodeSuccess, out, nil
	}
	// Step4. update password time
	go func() {
		err := s.r.LoginUserBase(ctx, repository.LoginUserBaseParams{
			UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
			UserAccount:  req.UserAccount,
			UserPassword: req.UserPassword,
		})
		if err != nil {
			return
		}
	}()

	// Step5. Create UUID user
	subToken := string2.GenerateCliTokenUUID(int(useBase.UserID))
	log.Println(subToken)

	// Step6. Get user_info table
	infoUser, err := s.r.GetUser(ctx, uint64(useBase.UserID))
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// convert to json
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrCodeAuthFailed, out, fmt.Errorf("convert to json failed: %v", err)
	}

	// 7. give infoUserJson to redis with key = subToken
	err = global.Rdb.Set(ctx, subToken, infoUserJson, time.Duration(user.TIME_2FA_OTP_REGISTER)*time.Minute).Err()
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	// 8. create token
	out.Token, err = auth.CreateToken(subToken)
	if err != nil {
		return response.ErrCodeAuthFailed, out, err
	}

	return 200, out, err
}

// SetupTwoFactorAuth setup authentication
func (s *sUserAuthed) SetupTwoFactorAuth(ctx context.Context, req *model.SetupTwoFactorAuthInput) (codeResult int, err error) {

	// Step1. Check is Two FactorAuth Enabled --> true return
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, req.UserId)
	if err != nil {
		return response.ErrCodeAuthFailed, err
	}

	if isTwoFactorAuth > 0 {
		return response.ErrCodeAuthFailed, fmt.Errorf("user already has two-factor auth")
	}

	// Step2. create new type Auth
	err = s.r.EnableTwoFactorTypeEmail(ctx, repository.EnableTwoFactorTypeEmailParams{
		UserID:            req.UserId,
		TwoFactorAuthType: repository.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
		TwoFactorEmail:    sql.NullString{String: req.TwoFactorEmail, Valid: true},
	})
	if err != nil {
		return response.ErrCodeAuthFailed, err
	}

	// Step3. Send OTP to req.TwoFactorEmail
	keyUserTwoFactor := crypto.GetHash("2fa" + strconv.Itoa(int(req.UserId)))
	go func() {
		err := global.Rdb.Set(ctx, keyUserTwoFactor, "123456", time.Duration(user.TIME_2FA_OTP_REGISTER)*time.Minute).Err()
		if err != nil {
			return
		}
	}()
	return response.ErrCodeSuccess, err
}

// VerifyTwoFactorAuth Verify Two-Factor Authentication
func (s *sUserAuthed) VerifyTwoFactorAuth(ctx context.Context, req *model.TwoFactorVerificationInput) (codeResult int, err error) {

	//Step1. Check is two factor enable
	isTwoFactorAuth, err := s.r.IsTwoFactorEnabled(ctx, req.UserId)
	if err != nil {
		return response.ErrCodeAuthFailed, err
	}

	if isTwoFactorAuth > 0 {
		return response.ErrCodeAuthFailed, fmt.Errorf("user already has two-factor auth")
	}
	//Step2. Check OTP in redis available
	keyUserTwoFactor := crypto.GetHash("2fa" + strconv.Itoa(int(req.UserId)))
	otpVerifyAuth, err := global.Rdb.Get(ctx, keyUserTwoFactor).Result()
	if errors.Is(err, redis.Nil) {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("key %s does not exists", keyUserTwoFactor)
	} else if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	//Step3. Check OTP
	if otpVerifyAuth == req.TwoFactorCode {
		return response.ErrCodeTwoFactorAuthVerifyFailed, fmt.Errorf("OTP does not match")
	}
	//Step4. update status
	err = s.r.UpdateTwoFactorStatus(ctx, repository.UpdateTwoFactorStatusParams{
		UserID:            req.UserId,
		TwoFactorAuthType: repository.PreGoAccUserTwoFactor9999TwoFactorAuthTypeEMAIL,
	})
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	//Step5. Remove OTP
	_, err = global.Rdb.Del(ctx, keyUserTwoFactor).Result()
	if err != nil {
		return response.ErrCodeTwoFactorAuthVerifyFailed, err
	}
	return 200, nil
}

// IsTwoFactorEnabled two-factor authentication
func (s *sUserAuthed) IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error) {
	// TO DO
	return 0, true, err
}
