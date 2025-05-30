package platform

import (
	"mirmiapi/platform/handler"

	"github.com/gin-gonic/gin"
)

type PlatformRouter struct {
	Handler       *handler.PlatformHandler
	CLinicHandler *handler.ClinicHandler
}

func NewPlatformHandler(h *handler.PlatformHandler) *PlatformRouter {
	return &PlatformRouter{Handler: h}
}

func (pr *PlatformRouter) RegisterPlatformRoutes(router *gin.Engine) {
	// Register platform-specific routes here

	group := router.Group("/platform")

	group.POST("/create", pr.Handler.RegisterPlatformUser)
	group.POST("/login", pr.Handler.LoginPlatformUser)
	group.GET("/logout", Logout)
	group.GET("/profile", Profile)

	group.GET("/users", ListUsers)
	group.POST("/users", CreateUser)
	group.PATCH("/users/:id", UpdateUser)
	group.DELETE("/users/:id", DeleteUser)

	group.GET("/clinics", ListClinics)
	group.POST("/clinics", pr.CLinicHandler.CreateClinic)
	group.PUT("/clinics", CreateClinic)

}
