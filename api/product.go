package api

import (
	"errors"
	"net/http"
	"product-api/apperror"
	"product-api/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *APIHandler) GetListProduct(ctx *gin.Context) {
	products, err := h.productService.GetList()
	if err != nil {
		ctx.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, products)
}

func (h *APIHandler) GetProductDetail(ctx *gin.Context) {
	s := ctx.Param("id")
	id, err1 := strconv.Atoi(s)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("id must contain number"))
		return
	}
	productDetail, err2 := h.productService.GetByID(id)
	if err2 != nil {
		ctx.JSON(http.StatusNotFound, model.NewErrorResponse("data not found"))
		return
	}

	ctx.JSON(http.StatusOK, productDetail)
}

func (h *APIHandler) StoreProduct(ctx *gin.Context) {
	var payload model.ProductRequest
	err1 := ctx.ShouldBindJSON(&payload)
	if err1 != nil {
		ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid request body"))
		return
	}

	err2 := h.productService.Store(&payload)
	if err2 != nil {
		if errors.Is(err2, apperror.ErrInvalidUserIdOrCategoryId) {
			ctx.JSON(http.StatusBadRequest, model.NewErrorResponse(err2.Error()))
			return
		}

		if errors.Is(err2, apperror.ErrInvalidProductName) {
			ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("product name contains forbidden word"))
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse(err2.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, model.NewSuccessResponse("product added"))

}

func (h *APIHandler) DeleteProduct(ctx *gin.Context) {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("id must contains number"))
		return
	}

	err2 := h.productService.Delete(id)
	if err2 != nil {
		if errors.Is(err2, apperror.ErrProductNotFound) {
			ctx.JSON(http.StatusNotFound, model.NewErrorResponse("data not found"))
			return
		}

		if errors.Is(err2, apperror.ErrInvalidUserIdOrCategoryId) {
			ctx.JSON(http.StatusNotFound, model.NewErrorResponse(err2.Error()))
		}

	}

	ctx.JSON(http.StatusOK, model.NewSuccessResponse("product deleted"))

}

func (h *APIHandler) UpdateProduct(ctx *gin.Context) {
	s := ctx.Param("id")
	id, err := strconv.Atoi(s)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("id must contains number"))
		return
	}
	var UpdateProductRequest model.UpdateProductRequest
	err2 := ctx.ShouldBindJSON(&UpdateProductRequest)
	if err2 != nil {
		ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid request body"))
		return
	}
	// Mapping from DTO to Model
	// DTO : model.UpdateProductRequest{}
	// Model : model.Product{}
	updatedProduct := model.Product{
		Name:       UpdateProductRequest.Name,
		Price:      UpdateProductRequest.Price,
		CategoryID: UpdateProductRequest.CategoryID,
	}

	err3 := h.productService.Update(id, &updatedProduct)
	if err3 != nil {
		ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse(apperror.ErrInvalidUserIdOrCategoryId.Error()))
		return
	}
	ctx.JSON(http.StatusCreated, model.NewSuccessResponse("product updated"))
}
