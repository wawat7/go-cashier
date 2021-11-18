package user

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID int `gorm:"not null;uniqueIndex;primary_key"`
	Username string `gorm:"size:100;not null"`
	Password string `gorm:"size:255;not null"`
	Name string `gorm:"size:255;not null"`
	PhoneNumber string `gorm:"size:15;not null"`
	Address string `gorm:"size:255;not null"`
	Role string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}


