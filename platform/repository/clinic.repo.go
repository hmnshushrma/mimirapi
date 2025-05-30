package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Clinic struct {
	ID            uuid.UUID
	Name          string
	Slug          string
	DBName        string
	ConnectionURI string
	CreatedBy     uuid.UUID
	CreatedAt     time.Time
}

type ClinicRepository interface {
	CreateClinic(ctx context.Context, clinic Clinic) error
}

type clinicRepo struct {
	db *sql.DB
}

func NewClinicRepo(db *sql.DB) ClinicRepository {
	return &clinicRepo{db: db}
}

func (r *clinicRepo) CreateClinic(ctx context.Context, clinic Clinic) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO clinics (id, name, slug, db_name, connection_uri, created_by, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, clinic.ID, clinic.Name, clinic.Slug, clinic.DBName, clinic.ConnectionURI, clinic.CreatedBy, clinic.CreatedAt)

	return err
}
