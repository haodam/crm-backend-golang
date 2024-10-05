package user

import "github.com/gin-gonic/gin"

type RouterUser struct{}

func (us *RouterUser) InitUserRouter(Router *gin.RouterGroup) {

	// Public user
	userRouterPublic := Router.Group("user")
	{
		userRouterPublic.POST("/register")
	}

	//Private user
	userRouterPrivate := Router.Group("user")
	{
		userRouterPrivate.GET("/get_info")
	}

}
