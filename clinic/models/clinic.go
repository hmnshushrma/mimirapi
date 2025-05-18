package models

import "time"

type Clinic struct {
	ID            string    `db:"id"`
	Name          string    `db:"name"`
	Slug          string    `db:"slug"`
	DBName        string    `db:"db_name"`
	ConnectionURI string    `db:"connection_uri"` // optional
	CreatedBy     string    `db:"created_by"`     // FK to platform_users
	CreatedAt     time.Time `db:"created_at"`
}
