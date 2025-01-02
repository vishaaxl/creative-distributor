package data

import (
	"database/sql"
	"time"
)

type User struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Activated   bool      `json:"activated"`
	Version     int       `json:"version"`
}

type UserModel struct {
	DB *sql.DB
}
