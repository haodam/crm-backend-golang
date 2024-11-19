package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler/model/req"
	"github.com/haodam/user-backend-golang/pkg/response"
)

func (u *userHandlerImpl) VerifyOTP(ctx *gin.Context) {
	var params req.UserReqVerifyOTPRequest
	if err := ctx.ShouldBindJSON(&params); err != nil {
		response.ErrorResponse(ctx, response.ErrCodeParamInvalid, err.Error())
		return
	//}
	//result, err := u.verifyUserUseCase.VerifyOTP(ctx, params.VerifyKey, params.VerifyCode)
	//if err != nil {
	//	response.ErrorResponse(ctx, response.ErrInvalidOTP, err.Error())
	//	return
	//}
	response.SuccessResponse(ctx, response.ErrCodeSuccess, nil)
}
