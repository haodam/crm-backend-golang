package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type GinOption func(*gin.Engine)

// NewGin Khởi tạo Gin Framework
func NewGin(options ...GinOption) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())

	// Thêm các options cho Gin
	for _, option := range options {
		option(r)
	}

	return r
}

func AddMiddlewares(ms ...func(c *gin.Context)) GinOption {
	return func(g *gin.Engine) {
		for _, m := range ms {
			g.Use(m)
		}
	}
}

func AddGroupRoutes(gr []GroupRoute) GinOption {
	return func(g *gin.Engine) {
		for _, r := range gr {
			r.AddGroupRoute(g)
		}
	}
}

func AddRoutes(rs []Route) GinOption {
	return func(g *gin.Engine) {
		for _, r := range rs {
			r.AddRoute(g)
		}
	}
}

func StrictSlash(strict bool) GinOption {
	return func(g *gin.Engine) {
		g.RemoveExtraSlash = strict
	}
}

func SetMaximumMultipartSize(size int64) GinOption {
	return func(g *gin.Engine) {
		g.MaxMultipartMemory = size
	}
}

func AddGinOptions(options ...GinOption) GinOption {
	return func(e *gin.Engine) {
		for _, o := range options {
			o(e)
		}
	}
}

func AddHealthCheckRoute() GinOption {
	return func(g *gin.Engine) {
		g.GET("/health", func(c *gin.Context) {
			c.String(http.StatusOK, "OK")
		})
	}
}
