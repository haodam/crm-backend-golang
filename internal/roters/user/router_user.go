package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/wire"
)

type RouterUser struct{}

func (us *RouterUser) InitUserRouter(Router *gin.RouterGroup) {

	//ur := repository.NewUserRepository()
	//userService := usecase.NewUserService(ur, nil)
	//userHanderNonDenpency := handler.NewUserRegisterHandler(userService)
	userHandler, _ := wire.InitUserRouterHandler()
	fmt.Println(userHandler)

	// Public user
	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.POST("/register", nil)
	}

	//Private user
	userRouterPrivate := Router.Group("user")
	{
		userRouterPrivate.GET("/get_info")
	}

}
