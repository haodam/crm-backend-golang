//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler"
	database "github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
)

func InitUserRouterHandler() (*handler.userHandlerImpl, error) {
	wire.Build(
		database.New,
		usecase.NewUserService,
		handler.NewUserRegisterHandler,
	)
	return new(handler.userHandlerImpl), nil
}
