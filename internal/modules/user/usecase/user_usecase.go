package usecase

import (
	"context"
	"github.com/haodam/user-backend-golang/internal/modules/user/entity"
)

type (
	IUserLogin interface {
		Register(ctx context.Context, arg *entity.RegisterInput) (codeResult int, err error)
		Login(ctx context.Context, arg *entity.LoginInput) (codeResult int, out *entity.LoginOutput, err error)
	}

	IUserInfo interface {
		// TODO
	}

	IUserAdmin interface {
		// TODO
	}
)

var (
	localUserLogin IUserLogin
)

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin not found for interface IUserLogin")
	}
	return localUserLogin
}
