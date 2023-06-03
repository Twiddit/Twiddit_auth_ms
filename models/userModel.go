package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string `gorm:"unique"`
	Password     string
	Birthday     time.Time
	Phone        string
	ProfilePhoto string
	Description  string
	Username     string
	Token        string
}
