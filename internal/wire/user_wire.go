//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler"
	"github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
)

func InitUserRouterHandler() (*handler.UserHandler, error) {
	wire.Build(
		repository.NewUserRepository,
		usecase.NewUserService,
		handler.NewUserHandler,
	)
	return new(handler.UserHandler), nil
}
