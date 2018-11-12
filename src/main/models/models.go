package models

import "github.com/satori/go.uuid"

type User struct {
	ID           uuid.UUID `json:"id"`
	UserName     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password-hash"`
	Active       bool      `json:"active"`
	Deleted      bool      `json:"deleted"`
}

