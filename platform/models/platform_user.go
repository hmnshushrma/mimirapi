package models

import "time"

type PlatformUser struct {
	ID           string    `db:"id"`
	FullName     string    `db:"full_name"`
	Email        string    `db:"email"`
	Phone        string    `db:"phone"`
	PasswordHash string    `db:"password_hash"`
	Role         string    `db:"role"` // admin or operator
	IsActive     bool      `db:"is_active"`
	CreatedAt    time.Time `db:"created_at"`
}
