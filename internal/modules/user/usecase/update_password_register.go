package usecase

import (
	"context"
	"fmt"
	"github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/pkg/response"
	"github.com/haodam/user-backend-golang/utils/crypto"
	"log"
)

type IUpdatePassWordRegister interface {
	UpdatePasswordRegister(ctx context.Context, token string, Password string) (userId int, err error)
}

type updatePassWordRegister struct {
	r *repository.Queries
}

func newUpdatePassWordRegister(r *repository.Queries) IUpdatePassWordRegister {
	return &updatePassWordRegister{
		r: r,
	}
}

func (u *updatePassWordRegister) UpdatePasswordRegister(ctx context.Context, token string, Password string) (userId int, err error) {

	// Step 1 token is already verified : user_verify table
	infoOTP, err := u.r.GetInfoOTP(ctx, token)
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
	newUserBase, err := u.r.AddUserBase(ctx, userBase)
	log.Println("newUserBase::", newUserBase, userBase)
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	user_id, err := newUserBase.LastInsertId()
	if err != nil {
		return response.ErrCodeUserOtpNotExists, err
	}

	// add user_id to user info table
	newUserInfo, err := u.r.AddUserBase()
	return 0, nil
}
