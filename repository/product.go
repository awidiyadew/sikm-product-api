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
	var products []model.Product
	if err := r.db.Preload("Category").Find(&products).Error; err != nil {
		return []model.Product{}, err
	}
	return products, nil
}

func (r *productRepoImpl) FindByID(id int) (*model.Product, error) {
	var product model.Product
	if err := r.db.Preload("User").Preload("Category").First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrProductNotFound
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepoImpl) Insert(product *model.Product) error {
	if err := r.db.Create(product).Error; err != nil {
		var pgErr *pgconn.PgError
		// casting apakah error dari postgres
		if errors.As(err, &pgErr) {
			// cek apakah error codenya SQL state 23503
			if pgErr.Code == pgerrcode.ForeignKeyViolation {
				return apperror.ErrInvalidUserIdOrCategoryId
			}
		}
		return err
	}
	return nil
}

func (r *productRepoImpl) Update(id int, product *model.Product) error {
	result := r.db.Model(&model.Product{}).Where("id = ?", id).Updates(map[string]interface{}{
		"Name":        product.Name,
		"Price":       product.Price,
		"CategoryID":  product.CategoryID,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperror.ErrProductNotFound
	}
	return nil
}

func (r *productRepoImpl) Delete(id int) error {
	result := r.db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return apperror.ErrProductNotFound
	}
	return nil
}
