package usecase

import (
	"context"
	"github.com/haodam/user-backend-golang/internal/modules/user/model"
)

type (
	IUserAuthed interface {
		Register(ctx context.Context, arg *model.RegisterEntity) (codeResult int, err error)
		VerifyOTP(ctx context.Context, req *model.VerifyOTPInput) (out *model.VerifyOTPOutput, err error)
		UpdatePasswordRegister(ctx context.Context, token string, Password string) (userId int, err error)
		Login(ctx context.Context, req *model.LoginInput) (codeResult int, out *model.LoginOutput, err error)

		// IsTwoFactorEnabled two-factor authentication
		IsTwoFactorEnabled(ctx context.Context, userId int) (codeResult int, rs bool, err error)
		// SetupTwoFactorAuth setup authentication
		SetupTwoFactorAuth(ctx context.Context, req *model.SetupTwoFactorAuthInput) (codeResult int, err error)
		// VerifyTwoFactorAuth Verify Two-Factor Authentication
		VerifyTwoFactorAuth(ctx context.Context, req *model.TwoFactorVerificationInput) (codeResult int, err error)
	}

	IUserInfo interface {
		// TODO
	}

	IUserAdmin interface {
		// TODO
	}
)

var (
	localUserAuthed IUserAuthed
)

func UserAuthed() IUserAuthed {
	if localUserAuthed == nil {
		panic("implement localUserAuthed not found for interface IUserAuthed")
	}
	return localUserAuthed
}

func InitUserAuthed(i IUserAuthed) {
	localUserAuthed = i
}
