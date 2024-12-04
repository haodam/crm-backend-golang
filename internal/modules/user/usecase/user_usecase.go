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
