package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)
	}
	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "default")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong..." + name,
		"name":    name,
		"uid":     uid,
		"users":   []string{"cr7", "m10", "Hao"},
	})
}
