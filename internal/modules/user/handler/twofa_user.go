package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/modules/user/model"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
	"github.com/haodam/user-backend-golang/pkg/response"
	"github.com/haodam/user-backend-golang/utils/context"
	"log"
)

var TwoFA = new(sUser2FA)

type sUser2FA struct {
}

// User Setup Two-Factor Authentication
// @Summary      ser Setup Two-Factor Authentication
// @Description  ser Setup Two-Factor Authentication
// @Tags         account 2fa
// @Accept       json
// @Produce      json
// @param Authorization header string true "Authorization token"
// @Param        payload body model.SetupTwoFactorAuthInput true "payload"
// @Success      200  {object}  response.ResponseData
// @Failure      500  {object}  response.ErrorResponseData
// @Router       /user/two-factor/setup [post]

func (c *sUser2FA) SetupTwoFactorAuth(ctx *gin.Context) {
	var params model.SetupTwoFactorAuthInput
	if err := ctx.ShouldBindJSON(&params); err != nil {
		// Handle error
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "Missing or invalid setupTwoFactorAuth parameter")
		return
	}
	userId, err := context.GetUserIdFromUUID(ctx.Request.Context())
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, "UserId is not valid")
		return
	}

	log.Println("userId:::", userId)
	params.UserId = uint32(userId)
	codeResult, err := usecase.UserAuthed().SetupTwoFactorAuth(ctx, &params)
	if err != nil {
		response.ErrorResponse(ctx, response.ErrCodeTwoFactorAuthSetupFailed, err.Error())
		return
	}
	response.SuccessResponse(ctx, codeResult, nil)
}
