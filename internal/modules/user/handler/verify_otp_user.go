package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/modules/user/model"
	"github.com/haodam/user-backend-golang/pkg/response"
)

func (u *userHandlerImpl) VerifyOTP(ctx *gin.Context) {
	var params model.VerifyOTPInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	}
	result, err := u.verifyUserUseCase.VerifyOTP(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrInvalidOTP, err.Error())
		return
	}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, result)
}
