package model

import "gorm.io/gorm"

type Vendor struct {
	gorm.Model
	Name string `gorm:"type:varchar(100);not null" json:"name"`
}
