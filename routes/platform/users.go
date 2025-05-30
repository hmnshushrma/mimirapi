package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list users"})
}

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "create user"})
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update user"})
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "delete user"})
}
