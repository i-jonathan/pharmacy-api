package model

type Permission struct {
	baseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Role struct {
	baseModel
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions,omitempty"`
}

type Account struct {
	baseModel
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Role        Role   `json:"role"`
	RoleID      int    `json:"role_id"`
}
