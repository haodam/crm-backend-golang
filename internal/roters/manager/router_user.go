package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (us *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

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
