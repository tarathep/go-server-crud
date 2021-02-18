package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tarathep/go-server-crud/apis"
)

type Router struct {
	apis.HelloHandler
	apis.TutorialHandler
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Route is setup router
func (router Router) Route() *gin.Engine {

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/hello", router.GetHello)
	r.POST("/hello", router.PostHello)

	r.GET("/api/tutorials", router.ReadTutorials)
	r.GET("/api/tutorials/:id", router.ReadTutorial)
	r.POST("/api/tutorials", router.CreateTutorial)
	r.PUT("/api/tutorials", router.UpdateTutorial)
	r.PUT("/api/tutorials/:id", router.UpdateTutorial)

	r.DELETE("/api/tutorials/:id", router.DeleteTutorial)
	r.DELETE("/api/tutorials", router.DeleteTutorials)

	return r
}
