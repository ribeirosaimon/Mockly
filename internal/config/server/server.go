package server

import (
	"github.com/gin-gonic/gin"
	"github.com/ribeirosaimon/Mockly/internal/controller"
)

type Endpoints struct {
}

func InitControllers(engine *gin.Engine) {
	// engine.POST("graphql", func(c *gin.Context) {
	// 	controller2.NewGraphQlDinamic()
	// })

	engine.GET("teste", controller.NewRestController().RestDinamic)
}

func isRouteRegistered(router *gin.Engine, method, path string) bool {
	routes := router.Routes()
	for _, route := range routes {
		if route.Method == method && route.Path == path {
			return true
		}
	}
	return false
}
