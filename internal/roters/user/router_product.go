package user

import "github.com/gin-gonic/gin"

type RouterProduct struct{}

func (pr *RouterProduct) CreateRouterGroup(Router *gin.RouterGroup) {

	// public router
	productRouterPublic := Router.Group("/product")
	{
		productRouterPublic.GET("/search")
		productRouterPublic.GET("/detail/:id")
	}
	// private router
}
