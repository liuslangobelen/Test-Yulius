package router

import (
	"net/http"

	"tog.test/no6/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	config.ExposeHeaders = []string{"Content-Length", "Content-Type"}

	r.Use(cors.New(config))
	v1 := r.Group("/api")
	{
		v1.GET("ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, "ping")
		})
		v1.GET("show-cart", controller.ShowCart)
		v1.POST("add-cart", controller.AddCart)
		v1.DELETE("delete-cart/:code", controller.DeleteCart)
	}
	return r

}
