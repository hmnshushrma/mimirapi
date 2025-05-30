package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "register endpoint hit"})
}

func PlatformLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "login endpoint hit"})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "logout endpoint hit"})
}

func Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "profile endpoint hit"})
}
