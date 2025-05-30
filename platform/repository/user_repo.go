package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type PlatformUser struct {
	ID           uuid.UUID
	FullName     string
	Email        string
	Phone        string
	PasswordHash string
	Role         string
	IsActive     bool
	CreatedAt    time.Time
}

type PlatformUserRepository interface {
	Create(ctx context.Context, user PlatformUser) error
	FindByEmail(ctx context.Context, email string) (*PlatformUser, error)
}

type platformUserRepo struct {
	db *sql.DB
}

func NewPlatformUserRepo(db *sql.DB) PlatformUserRepository {
	return &platformUserRepo{db: db}
}

func (r *platformUserRepo) Create(ctx context.Context, user PlatformUser) error {
	_, err := r.db.ExecContext(ctx, `
        INSERT INTO platform_users (
            id, full_name, email, phone, password_hash, role, is_active, created_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `, user.ID, user.FullName, user.Email, user.Phone, user.PasswordHash, user.Role, user.IsActive, user.CreatedAt)

	return err
}

func (r *platformUserRepo) FindByEmail(ctx context.Context, email string) (*PlatformUser, error) {
	row := r.db.QueryRowContext(ctx, `
        SELECT id, full_name, email, phone, password_hash, role, is_active, created_at
        FROM platform_users
        WHERE email = $1
    `, email)

	var user PlatformUser
	err := row.Scan(&user.ID, &user.FullName, &user.Email, &user.Phone, &user.PasswordHash, &user.Role, &user.IsActive, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
