package user

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/middleware"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler"
)

type RouterUser struct{}

func (us *RouterUser) InitUserRouter(Router *gin.RouterGroup) {

	//ur := repository.NewUserRepository()
	//userService := usecase.NewUserService(ur, nil)
	//userHanderNonDenpency := handler.NewUserRegisterHandler(userService)
	//userHandler, _ := wire.InitUserRouterHandler()
	//fmt.Println(userHandler)

	// Public user
	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.POST("/register", handler.Authed.HandleUserRegister)
		userRouterPublic.POST("/verify_user", handler.Authed.HandleUserVerifyOTP)
		userRouterPublic.POST("/update_password_register", handler.Authed.HandleUserUpdatePasswordRegister)
		userRouterPublic.POST("/login", handler.Authed.Login)

	}

	//Private user
	userRouterPrivate := Router.Group("user")
	userRouterPrivate.Use(middleware.AuthedMiddleware())
	{
		userRouterPrivate.GET("/get_info")
		userRouterPrivate.GET("/two-factor/setup", handler.TwoFA.SetupTwoFactorAuth)
		userRouterPrivate.POST("/two-factor/verify", handler.TwoFA.VerifyTwoFactorAuth)

	}
}
