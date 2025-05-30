package routes

import (
	"mirmiapi/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingRoute(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		if err := config.DB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "offline",
				"error":  err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "online",
			"db":     "connected",
		})
	})
}
