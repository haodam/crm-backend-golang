package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/common"
	"github.com/haodam/user-backend-golang/global"
	"github.com/haodam/user-backend-golang/internal/modules/user/model"
	"github.com/haodam/user-backend-golang/pkg/response"
	"go.uber.org/zap"
	"net/http"
)

// User Registration documentation
// @Summary      User Registration
// @Description  When user is registered send otp to email
// @Tags         account management
// @Accept       json
// @Produce      json
// @Param        payload body model.RegisterInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/register [post]

func (u *userHandlerImpl) HandleUserRegister(ctx *gin.Context) {

	var params *model.RegisterEntity
	if err := ctx.ShouldBindJSON(&params); err != nil {
		common.ResponseErr(ctx, http.StatusBadRequest)
		return
	}
	codeStatus, err := u.registerUserUseCase.Register(
		ctx.Request.Context(),
		params,
	)

	if err != nil {
		global.Logger.Error("Error registering user OTP", zap.Error(err))
		response.ErrorResponse(ctx, codeStatus, err.Error())
		return
	}

	response.SuccessResponse(ctx, codeStatus, nil)
}
