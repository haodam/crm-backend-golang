package handler

import (
	"github.com/gin-gonic/gin"
	database "github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
)

type IUserHandler interface {
	HandleUserRegister(ctx *gin.Context)
	VerifyOTP(ctx *gin.Context)
}

type userHandlerImpl struct {
	registerUserUseCase usecase.IUserRegister
	verifyUserUseCase   usecase.IVerifyUserRegister
}

func NewUserHandler(d *database.Queries) IUserHandler {
	return &userHandlerImpl{
		registerUserUseCase: usecase.NewRegisterUserUseCase(d),
		verifyUserUseCase:   usecase.NewVerifyUserUseCase(d),
	}
}
