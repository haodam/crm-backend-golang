package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler/model/req"
	"net/http"
)

func (u *userHandlerImpl) HandleUserRegister(ctx *gin.Context) {

	var params req.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Email params: %s", params.Email)

}
