package account

import "time"

type User struct {
	ID          uint      `json:"id"`
	FullName    string    `json;"full_name"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phone_number"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
}

type passwordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}
