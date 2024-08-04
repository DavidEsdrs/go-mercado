package model

import (
	"gorm.io/gorm"
)

// represents whether the specified product is precified by unit or kg/mg
type PriceType string

const (
	PriceTypeWeight PriceType = "kg"
	PriceTypeUnit   PriceType = "unit"
)

type Product struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description *string   `gorm:"type:varchar(255)" json:"description"`
	Value       float64   `gorm:"type:decimal(10,2);not null" json:"value"`
	Stock       int32     `gorm:"not null" json:"stock"`
	Type        PriceType `gorm:"type:enum('kg', 'unit');not null" json:"type"`
	VendorID    int       `gorm:"not null" json:"vendor_id"`
}

func (p *Product) IsValid() bool {
	return p.ID != 0 && len(p.Name) < 50 && p.Value > 0
}
