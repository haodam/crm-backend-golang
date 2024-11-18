package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/common"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler/model/req"
	"github.com/haodam/user-backend-golang/pkg/response"
	"go.uber.org/zap"
	"net/http"
)

func (u *userHandlerImpl) HandleUserRegister(ctx *gin.Context) {

	var params req.UserRegistrationRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.ResponseErr(ctx, http.StatusBadRequest)
		return
	}
	codeStatus, err := u.registerUserUseCase.Register(
		ctx.Request.Context(),
		params.VerifyKey,
		params.VerifyType,
		params.VerifyPurpose)

	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)

}
