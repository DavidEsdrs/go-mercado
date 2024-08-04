package model

type ProductCart struct {
	CartID    uint    `gorm:"primaryKey"`
	ProductID uint    `gorm:"primaryKey"`
	Quantity  int32   `gorm:"not null" json:"quantity"`
	Total     float64 `gorm:"type:decimal(10,2);not null" json:"total"`
}
