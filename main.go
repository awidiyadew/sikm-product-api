package main

import (
	"log"
	"net/http"
	"product-api/api"
	"product-api/config"
	"product-api/db"
	"product-api/middleware"
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
	config.Init()
	db, err := db.Connect(db.DBCredential{
		Host:         config.DBHost,
		Username:     config.DBUsername,
		Password:     config.DBPassword,
		DatabaseName: config.DBName,
		Port:         config.DBPort,
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
		productRouter.Use(middleware.Auth()) // API after this line  will be protected with auth (within code scope {productRouter})
		productRouter.GET("/list", handler.GetListProduct)
		productRouter.GET("/:id", handler.GetProductDetail)
	}

	productAdmin := router.Group("/product")
	{
		productAdmin.Use(middleware.AuthAdmin()) // API after this line  will be protected with auth admin (within code scope {productAdmin})
		productAdmin.POST("/add", handler.StoreProduct)
		productAdmin.PUT("/:id", handler.UpdateProduct)
		productAdmin.DELETE("/:id", handler.DeleteProduct)
	}

	router.Run(":3000")
}
