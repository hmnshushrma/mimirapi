package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func CreateClinicDatabaseName(db *sql.DB, dbName, slug string) error {
	// PostgreSQL does not support "IF NOT EXISTS" in CREATE DATABASE
	createQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
	_, err := db.Exec(createQuery)
	if err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	clinicConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	clinicDB, err := sql.Open("postgres", clinicConnStr)
	if err != nil {
		return fmt.Errorf("failed to connect to clinic DB: %w", err)
	}
	defer clinicDB.Close()

	if err := clinicDB.Ping(); err != nil {
		return fmt.Errorf("failed to ping clinic DB: %w", err)
	}

	log.Println("✅ Connected to new clinic DB:", dbName)

	// Run schema (users table only for now)
	_, err = clinicDB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id UUID PRIMARY KEY,
            full_name TEXT NOT NULL,
            email TEXT NOT NULL UNIQUE,
            password_hash TEXT NOT NULL,
            role TEXT CHECK (role IN ('admin', 'doctor', 'receptionist')) NOT NULL,
            is_active BOOLEAN DEFAULT TRUE,
            created_at TIMESTAMP DEFAULT NOW()
        );
    `)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// Create default admin user
	password := "Admin#2030"
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	adminEmail := fmt.Sprintf("admin@%s.com", slug)

	_, err = clinicDB.Exec(`
        INSERT INTO users (id, full_name, email, password_hash, role, is_active, created_at)
        VALUES (gen_random_uuid(), 'Admin', $1, $2, 'admin', TRUE, NOW());
    `, adminEmail, string(hashed))

	if err != nil {
		return fmt.Errorf("failed to create default admin user: %w", err)
	}

	log.Println("✅ Default admin user created for clinic:", adminEmail)
	return nil
}
