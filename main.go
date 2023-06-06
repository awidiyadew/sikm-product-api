package main

import (
	"log"
	"net/http"
	"product-api/api"
	"product-api/db"
	"product-api/model"
	"product-api/repository"
	"product-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewHandler(db *gorm.DB) *api.APIHandler {
	productRepo := repository.NewProductRepo(db)
	productService := service.NewProductService(productRepo)
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)

	return api.NewHandler(productService, userService)
}

func main() {
	db, err := db.Connect(db.DBCredential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "P@ssw0rd",
		DatabaseName: "product_api",
		Port:         5432,
	})
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})

	router := gin.Default()
	handler := NewHandler(db)

	// handling method not allowed and route not found
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "page not found or method not allowed",
		})

	})

	userRouter := router.Group("/user")
	{
		userRouter.POST("/login", handler.Login)
	}

	productRouter := router.Group("/product")
	{
		// TODO: validate auth user must login
		productRouter.GET("/list", handler.GetListProduct)
		productRouter.GET("/:id", handler.GetProductDetail)

		// TODO: validate auth user must login with role ADMIN
		productRouter.POST("/add", handler.StoreProduct)
		productRouter.PUT("/:id", handler.UpdateProduct)
		productRouter.DELETE("/:id", handler.DeleteProduct)
	}

	router.Run(":3000")
}
