package service

import (
	"fmt"

	models "github.com/DavidEsdrs/go-mercado/internal/model"
)

type Repository[T any] interface {
	Insert(*T) error
	Read(id uint) (T, error)
	Update(*T) error
	Delete(id uint) error
}

type ProductService struct {
	productRepository Repository[models.Product]
}

func NewProductService(repo Repository[models.Product]) *ProductService {
	return &ProductService{
		productRepository: repo,
	}
}

func (pc *ProductService) InsertProduct(p *models.Product) error {
	if !p.IsValid() {
		return fmt.Errorf("invalid product given")
	}
	pc.productRepository.Insert(p)
	return nil
}
