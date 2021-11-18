package user

import (
	"time"
)

type FormatUser struct {
	ID int `json:"id"`
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Address string `json:"address"`
	Role string `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func UserFormat(user User) FormatUser {
	return FormatUser{
		ID:          user.ID,
		Name:        user.Name,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Role:        user.Role,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}
}

func UsersFormat(users []User) []FormatUser {
	var formats []FormatUser

	for _, user := range users {
		format := UserFormat(user)
		formats = append(formats, format)
	}
	return formats
}

