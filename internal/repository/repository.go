package repository

import (
	"github.com/DavidEsdrs/go-mercado/internal/model"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(conn *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: conn,
	}
}

func (pr *ProductRepository) Insert(product *model.Product) error {
	return pr.db.Create(product).Error
}

func (pr *ProductRepository) Read(id uint) (model.Product, error) {
	var product model.Product
	if err := pr.db.First(&product, id).Error; err != nil { // First founds first match
		return product, err
	}
	return product, nil
}

func (pr *ProductRepository) Update(product *model.Product) error {
	return pr.db.Save(product).Error // upsert
}

func (pr *ProductRepository) Delete(id uint) error {
	return pr.db.Delete(&model.Product{}, id).Error // deletes by id
}

func (pr *ProductRepository) FindAll() (products []model.Product, err error) {
	tx := pr.db.Find(&products)
	return products, tx.Error
}
