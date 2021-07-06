package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID int64 `json:"id" gorm:"primaryKey"`
	Username string `json:"username"`
	Password string	`json:"-"`
}

