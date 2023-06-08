package service

import (
	"product-api/apperror"
	"product-api/model"
	"product-api/repository"
	"time"

	"github.com/golang-jwt/jwt"
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

func (s *userServiceImpl) generateToken(user *model.User) (string, error) {
	now := time.Now()
	claims := model.Claims{
		UserID: int(user.ID),
		Name:   user.Name,
		Email:  user.Email,
		Scope:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(15 * time.Minute).Unix(),
			Issuer:    "product-api",
			IssuedAt:  now.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(model.JwtKey)
}

func (s *userServiceImpl) Login(payload *model.LoginRequest) (token string, err error) {
	user, err := s.userRepo.FindByEmail(payload.Email)
	if err != nil {
		return "", err
	}
	if user.Password != payload.Password {
		return "", apperror.ErrInvalidPassword
	}
	token, err = s.generateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil
}
