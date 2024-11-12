package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/common"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler/model/req"
	"net/http"
)

func (u *userHandlerImpl) HandleUserRegister(ctx *gin.Context) {

	var params req.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.ResponseErr(ctx, http.StatusBadRequest)
		return
	}
	err := u.registerUserUseCase.Register(
		ctx.Request.Context(),
		params.VerifyKey,
		params.VerifyType,
		params.VerifyPurpose)

	if err != nil {
		common.ResponseErr(ctx, http.StatusInternalServerError)
		return
	}

	common.SimpleResponseOK(ctx, http.StatusOK, nil)

}
