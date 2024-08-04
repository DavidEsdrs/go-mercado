package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"type:varchar(100);not null" json:"username"`
	Email        string `gorm:"type:varchar(100);not null;unique" json:"email"`
	HashPassword string `gorm:"type:varchar(50);not null" json:"hash_password"`
}
