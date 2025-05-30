package platform

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListClinics(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "list clinics"})
}

func CreateClinic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "create clinic"})
}

func UpdateClinic(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update clinic"})
}
