package api

import "product-api/service"

type APIHandler struct {
	productService service.ProductService
	userService    service.UserService
}

func NewHandler(p service.ProductService, u service.UserService) *APIHandler {
	return &APIHandler{
		productService: p,
		userService:    u,
	}
}
