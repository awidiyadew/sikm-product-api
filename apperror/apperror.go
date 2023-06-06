package apperror

import "errors"

var (
	ErrProductNotFound           = errors.New("product not found")
	ErrInvalidProductName        = errors.New("product name contains forbidden words")
	ErrInvalidUserIdOrCategoryId = errors.New("invalid category_id or posted_by")
	ErrUserNotFound              = errors.New("user not found")
	ErrInvalidPassword           = errors.New("invalid password")
)
