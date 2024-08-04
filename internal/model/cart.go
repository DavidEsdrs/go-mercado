package model

import "gorm.io/gorm"

type CartStatus string

const (
	StatusOpen      CartStatus = "open"
	StatusClosed    CartStatus = "closed"
	StatusCancelled CartStatus = "cancelled"
)

type Cart struct {
	gorm.Model
	Status     CartStatus `gorm:"type:enum('open', 'closed', 'cancelled')" json:"status"`
	TotalValue float64    `gorm:"type:decimal(10,2);not null" json:"total_value"`
	Products   []Product  `gorm:"many2many:cart_products"`
}
