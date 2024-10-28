package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler/model/req"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
	"net/http"
)

type UserRegisterHandler struct {
	userCase usecase.IUserRegisterService
}

func NewUserRegisterHandler(userCase usecase.IUserRegisterService) *UserRegisterHandler {
	return &UserRegisterHandler{userCase: userCase}
}

func (ur *UserRegisterHandler) UserRegisterHandler(ctx *gin.Context) {

	var params req.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Email params: %s", params.Email)
	result := ur.userCase.Execute(ctx, params.Email, params.Purpose)
	ctx.JSON(http.StatusOK, result)

}
