package repository

import (
	"errors"
	"product-api/apperror"
	"product-api/model"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
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
	// Preload() -> to join table when querying data which is have association with another table
	// key -> field which is referencing to another table
	var listProduct []model.Product
	result := r.db.Preload("Category").Find(&listProduct)
	if result.Error != nil {
		return nil, result.Error
	}

	return listProduct, nil
}

func (r *productRepoImpl) FindByID(id int) (*model.Product, error) {
	// Preload() -> to join table when querying data which is have association with another table
	// key -> field which is referencing to another table
	var product model.Product
	result := r.db.Preload("Category").Preload("User").Where("id = ?", id).First(&product)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, apperror.ErrProductNotFound

	}
	return &product, nil
}

func (r *productRepoImpl) Insert(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		// Custom error with PgErr
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return apperror.ErrInvalidUserIdOrCategoryId
			}
			return err
		}
	}
	return nil
}

func (r *productRepoImpl) Update(id int, product *model.Product) error {
	// Dont need initializer method
	// Reason :  if we have passed a valid database model on the finisher method,
	// then we don't need to call method .Model(&model.Product{}) or .Table("")
	if err := r.db.Where("id = ?", id).Updates(model.Product{
		Name:       product.Name,
		Price:      product.Price,
		CategoryID: product.CategoryID,
	}).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return apperror.ErrInvalidUserIdOrCategoryId
			}
			return err
		}
	}
	return nil
}

func (r *productRepoImpl) Delete(id int) error {
	result := r.db.Where("id", id).Delete(&model.Product{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
