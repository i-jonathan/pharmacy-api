package model

import "time"

type Permission struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Slug        string    `json:"slug"`
}

type Role struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	CreatedAt   time.Time    `json:"created_at"`
	Permissions []Permission `json:"permissions,omitempty"`
	Slug        string       `json:"slug"`
}

type Account struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	Role        Role      `json:"role"`
	RoleID      int       `json:"role_id"`
	Slug        string    `json:"slug"`
}
