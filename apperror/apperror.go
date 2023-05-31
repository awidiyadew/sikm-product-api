package apperror

import "errors"

// TODO: use this sentinel errors in repo, service, or handler
var (
	ErrProductNotFound           = errors.New("product not found")
	ErrInvalidProductName        = errors.New("product name contains forbidden words")
	ErrInvalidUserIdOrCategoryId = errors.New("invalid category_id or posted_by")

	// TODO: add other sentinel if needed
)
