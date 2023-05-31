package repository

import (
	"product-api/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(id int) (*model.Product, error)
	Insert(product *model.Product) error
	Update(id int, product *model.Product) error
	Delete(id int) error
}

type productRepoImpl struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) *productRepoImpl {
	return &productRepoImpl{
		db: db,
	}
}

func (r *productRepoImpl) FindAll() ([]model.Product, error) {
	// TODO: select all product with category
	return []model.Product{}, nil
}

func (r *productRepoImpl) FindByID(id int) (*model.Product, error) {
	// TODO: select product with category and user
	return nil, nil
}

func (r *productRepoImpl) Insert(product *model.Product) error {
	// TODO: create product and check pgConn.PgError with code 23503 FK Violation
	return nil
}

func (r *productRepoImpl) Update(id int, product *model.Product) error {
	// TODO: Update only specific fields name, price, and category_id
	return nil
}

func (r *productRepoImpl) Delete(id int) error {
	// TODO: delete by id
	return nil
}
