package handler

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"mirmiapi/platform/models"
	"mirmiapi/platform/repository"
	"mirmiapi/platform/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ClinicHandler struct {
	Repo repository.ClinicRepository
	DB   *sql.DB
}

func NewClinicHandler(repo repository.ClinicRepository, db *sql.DB) *ClinicHandler {
	return &ClinicHandler{
		Repo: repo,
		DB:   db,
	}
}

func (h *ClinicHandler) CreateClinic(c *gin.Context) {
	var req models.CreateClinicRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	slug := utils.Slugify(req.Name)
	clinicID := uuid.New()
	safeSlug := strings.ReplaceAll(slug, "-", "_")
	dbName := fmt.Sprintf("rx_%s_%s", safeSlug, clinicID.String()[:8])

	if err := utils.CreateClinicDatabaseName(h.DB, dbName, slug); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	clinic := repository.Clinic{
		ID:        clinicID,
		Name:      req.Name,
		Slug:      slug,
		DBName:    dbName,
		CreatedBy: uuid.Nil, // To be replaced with actual user ID later
		CreatedAt: time.Now(),
	}

	if err := h.Repo.CreateClinic(c.Request.Context(), clinic); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store clinic info"})
		return
	}

	c.JSON(http.StatusCreated, models.CreateClinicResponse{
		ID:     clinic.ID.String(),
		Name:   clinic.Name,
		Slug:   clinic.Slug,
		DBName: clinic.DBName,
	})
}
