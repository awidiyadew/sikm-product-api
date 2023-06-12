package api

import (
	"errors"
	"net/http"
	"product-api/apperror"
	"product-api/model"

	"github.com/gin-gonic/gin"
)

func (h *APIHandler) Login(c *gin.Context) {
	var payload model.LoginRequest
	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.NewErrorResponse(err.Error()))
		return
	}

	token, err := h.userService.Login(&payload)
	if err != nil {
		if errors.Is(err, apperror.ErrUserNotFound) {
			c.JSON(http.StatusBadRequest, model.NewErrorResponse("unregistered email error"))
			return
		}
		if errors.Is(err, apperror.ErrInvalidPassword) {
			c.JSON(http.StatusBadRequest, model.NewErrorResponse("invalid password"))
			return
		}
		c.JSON(http.StatusInternalServerError, model.NewErrorResponse("login error"))
		return
	}

	expiry15s := 15 * 60
	c.SetCookie("session_token", token, expiry15s, "", "", false, true)

	c.JSON(http.StatusOK, model.NewSuccessResponse("login success"))
}
