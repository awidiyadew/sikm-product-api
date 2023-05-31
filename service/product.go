package service

import (
	"product-api/model"
	"product-api/repository"
)

type ProductService interface {
	GetList() ([]model.Product, error)
	GetByID(id int) (*model.ProductDetail, error)
	Store(*model.Product) error
	Delete(*model.Product) error
	Update(*model.Product) error
}

type productServiceImpl struct {
	productRepo repository.ProductRepository
}

func NewProductService(prdRepo repository.ProductRepository) ProductService {
	return &productServiceImpl{
		productRepo: prdRepo,
	}
}

func (s *productServiceImpl) GetList() ([]model.Product, error) {
	// TODO: add code
	return []model.Product{}, nil
}

func (s *productServiceImpl) GetByID(id int) (*model.ProductDetail, error) {
	// TODO: add code
	// TODO: use DTO to transform data from model
	return nil, nil
}

func (s *productServiceImpl) Store(product *model.Product) error {
	// TODO: add code
	return nil
}

func (s *productServiceImpl) Update(*model.Product) error {
	// TODO: add code
	return nil
}

func (s *productServiceImpl) Delete(*model.Product) error {
	// TODO: add code
	return nil
}
