package api

import "product-api/service"

type APIHandler struct {
	productService service.ProductService
}

func NewHandler(p service.ProductService) *APIHandler {
	return &APIHandler{
		productService: p,
	}
}
