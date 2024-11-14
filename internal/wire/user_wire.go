//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
)

func InitUserRouterHandler() (*handler.userHandlerImpl, error) {
	wire.Build(
		repository.NewUserRepository,
		repository.NewOtpRegisterRepository,
		usecase.NewUserService,
		handler.NewUserRegisterHandler,
	)
	return new(handler.UserRegisterHandler), nil
}
