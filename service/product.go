package service

import (
	"product-api/apperror"
	"product-api/apperror"
	"product-api/model"
	"product-api/repository"
	"strings"
)

var (
	blackListedWords = []string{
		"termurah",
		"terbaik",
		"diskon",
		"promo",
	}
	"strings"
)

var (
	blackListedWords = []string{
		"termurah",
		"terbaik",
		"diskon",
		"promo",
	}
)

type ProductService interface {
	GetList() ([]model.Product, error)
	GetByID(id int) (*model.ProductDetail, error)
	Store(*model.ProductRequest) error
	Delete(id int) error
	Update(id int, product *model.Product) error
}

type productServiceImpl struct {
	productRepo repository.ProductRepository
}

func NewProductService(prdRepo repository.ProductRepository) ProductService {
	return &productServiceImpl{
		productRepo: prdRepo,
	}
}
func (s *productServiceImpl) isValidName(name string) bool {
	for _, word := range blackListedWords {
		if strings.Contains(strings.ToLower(name), word) {
			return false
		}
	}

	return true
}
func (s *productServiceImpl) GetList() ([]model.Product, error) {
	// TODO: add code
	products, err := s.productRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *productServiceImpl) GetByID(id int) (*model.ProductDetail, error) {
	// TODO: add code
	product, err := s.productRepo.FindByID(id)
	if err != nil {
		return &model.ProductDetail{}, err
	}
	// TODO: use DTO to transform data from model
	// Mapping from DTO to Model
	// DTO : model.ProductDetail{}
	// Model : model.Product{}
	var response model.ProductDetail = model.ProductDetail{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		User:     product.User,
		Category: product.Category,
	}

	return &response, nil

}

func (s *productServiceImpl) Store(payload *model.ProductRequest) error {
	// TODO: add code
	// Validate forbidden word
	if !s.isValidName(payload.Name) {
		return apperror.ErrInvalidProductName
	}

	// Mapping from DTO to Model
	// DTO : model.ProductRequest{}
	// Model : model.Product{}
	var newProduct model.Product = model.Product{
		Name:       payload.Name,
		Price:      payload.Price,
		CategoryID: payload.CategoryID,
		PostedBy:   uint(payload.PostedBy),
	}

	err := s.productRepo.Insert(&newProduct)
	if err != nil {
		return err
	}
	return nil
}

func (s *productServiceImpl) Update(id int, product *model.Product) error {
	// TODO: add code
	if err := s.productRepo.Update(id, product); err != nil {
		return err
	}

	return nil
}

func (s *productServiceImpl) Delete(id int) error {
	// TODO: add code

	err := s.productRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
