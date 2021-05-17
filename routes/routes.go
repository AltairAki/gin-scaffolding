package routes

import (
	"net/http"
	"web_app/logger"

	"github.com/gin-gonic/gin"
)

func Setup() (route *gin.Engine) {
	route = gin.New()
	route.Use(logger.GinLogger())
	route.Use(logger.GinRecovery(true))
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world.")
	})
	return route
}
