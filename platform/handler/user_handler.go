package handler

import (
	"mirmiapi/platform/models"
	"mirmiapi/platform/repository"
	"mirmiapi/platform/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type PlatformHandler struct {
	Repo repository.PlatformUserRepository
}

// NewPlatformHandler creates a new instance of PlatformHandler

func NewPlatformHandler(repo repository.PlatformUserRepository) *PlatformHandler {
	return &PlatformHandler{Repo: repo}
}

func (h *PlatformHandler) RegisterPlatformUser(c *gin.Context) {
	var request models.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	user := models.PlatformUser{
		ID:           uuid.New().String(),
		FullName:     request.FullName,
		Email:        request.Email,
		Phone:        request.Phone,
		PasswordHash: string(hashedPassword),
		Role:         "operator", // Default role
		IsActive:     true,
		CreatedAt:    time.Now(),
	}

	// Map models.PlatformUser to repository.PlatformUser
	repoUser := repository.PlatformUser{
		ID:           uuid.MustParse(user.ID),
		FullName:     user.FullName,
		Email:        user.Email,
		Phone:        user.Phone,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
		IsActive:     user.IsActive,
		CreatedAt:    user.CreatedAt,
	}

	// Save the user using the repository
	if err := h.Repo.Create(c.Request.Context(), repoUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, models.RegisterResponse{
		Id:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	})

}

func (h *PlatformHandler) LoginPlatformUser(c *gin.Context) {
	var request models.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	user, err := h.Repo.FindByEmail(c.Request.Context(), request.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	if user == nil || !user.IsActive {
		// 401 Unauthorized: authentication failed
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(request.Password)); err != nil {
		// 401 Unauthorized: authentication failed
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Role)
	if err != nil {
		// 500 Internal Server Error: something went wrong
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, models.LoginReponse{
		Token:    token,
		Fullname: user.FullName,
		Email:    user.Email,
		Role:     user.Role,
	})
}
