package user

import (
	"github.com/gin-gonic/gin"
	"github.com/haodam/user-backend-golang/internal/modules/user/handler"
	"github.com/haodam/user-backend-golang/internal/modules/user/repository"
	"github.com/haodam/user-backend-golang/internal/modules/user/usecase"
)

type RouterUser struct{}

func (us *RouterUser) InitUserRouter(Router *gin.RouterGroup) {

	ur := repository.NewUserRepository()
	userService := usecase.NewUserService(ur)
	userHanderNonDenpency := handler.NewUserHandler(userService)

	// Public user
	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.POST("/register", userHanderNonDenpency.Register)
	}

	//Private user
	userRouterPrivate := Router.Group("user")
	{
		userRouterPrivate.GET("/get_info")
	}

}
