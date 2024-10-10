package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
)

type UserHandler struct {
	userhandler usecase.IUserService
}

func NewUserHandler(userhandler usecase.IUserService) *UserHandler {
	return &UserHandler{
		userhandler: userhandler,
	}
}

func (uh *UserHandler) Register(c *gin.Context) {

}
