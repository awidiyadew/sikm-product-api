package service

import (
	"product-api/apperror"
	"product-api/model"
	"product-api/repository"
)

type UserService interface {
	Login(payload *model.LoginRequest) (token string, err error)
}

type userServiceImpl struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (s *userServiceImpl) Login(payload *model.LoginRequest) (token string, err error) {
	user, err := s.userRepo.FindByEmail(payload.Email)
	if err != nil {
		return "", err
	}
	if user.Password != payload.Password {
		return "", apperror.ErrInvalidPassword
	}
	// TODO: generate token
	token = "dummy-token"
	return token, nil
}
