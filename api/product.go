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
		ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (h *APIHandler) GetProductDetail(ctx *gin.Context) {
	idParam, _ := ctx.Params.Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid id param"))
		return
	}

	pDetail, err := h.productService.GetByID(id)
	if err != nil {
		if errors.Is(err, apperror.ErrProductNotFound) {
			ctx.JSON(http.StatusNotFound, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, pDetail)
}

func (h *APIHandler) StoreProduct(ctx *gin.Context) {
	var payload model.ProductRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		// error validasi payload dengan `binding` pd struct model.ProductRequest
		ctx.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}

	err = h.productService.Store(&payload)
	if err != nil {
		if errors.Is(err, apperror.ErrInvalidUserIdOrCategoryId) {
			ctx.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
			return
		}

		ctx.JSON(http.StatusInternalServerError, model.NewErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, model.NewSuccessResponse("product added"))
}

func (h *APIHandler) DeleteProduct(ctx *gin.Context) {
	// TODO: handle delete product
}

func (h *APIHandler) UpdateProduct(ctx *gin.Context) {
	// TODO: handle update product
}
