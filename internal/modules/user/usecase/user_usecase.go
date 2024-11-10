package usecase

type (
	IUserLogin interface {
		//Register(ctx context.Context, req *entity.RegisterInput) *common.Error
		//Login(ctx context.Context, arg entity.LoginInput) (codeResult int, out *entity.LoginOutput, err error)
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
