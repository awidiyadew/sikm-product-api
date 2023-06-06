package repository

import (
	"errors"
	"product-api/apperror"
	"product-api/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmail(email string) (*model.User, error)
}

type userRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepoImpl {
	return &userRepoImpl{
		db: db,
	}
}

func (r *userRepoImpl) FindByEmail(email string) (*model.User, error) {
	var user model.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperror.ErrUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
