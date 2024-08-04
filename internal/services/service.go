package service

import (
	"fmt"

	"github.com/DavidEsdrs/go-mercado/internal/model"
)

type Repository[T any] interface {
	Insert(*T) error
	Read(id uint) (T, error)
	Update(*T) error
	Delete(id uint) error
	FindAll() ([]T, error)
}

type ProductService struct {
	productRepository Repository[model.Product]
}

func NewProductService(repo Repository[model.Product]) *ProductService {
	return &ProductService{
		productRepository: repo,
	}
}

func (pc *ProductService) InsertProduct(p *model.Product) error {
	if !p.IsValid() {
		return fmt.Errorf("invalid product given")
	}
	pc.productRepository.Insert(p)
	return nil
}

func (pc *ProductService) ReadProduct(id uint) (product model.Product, err error) {
	if product, err = pc.productRepository.Read(id); err != nil {
		return product, fmt.Errorf("no entity found for id %v - error: %w", id, err)
	}
	return product, nil
}

func (pc *ProductService) ReadProducts() (products []model.Product, err error) {
	return pc.productRepository.FindAll()
}
