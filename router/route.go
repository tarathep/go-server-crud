package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tarathep/go-server-crud/apis"
)

type Router struct {
	Api apis.HelloHandler
}

// Route is setup router
func (router Router) Route() *gin.Engine {
	r := gin.Default()

	r.GET("/hello", router.Api.GetHello)
	r.POST("/hello", router.Api.PostHello)

	return r
}
