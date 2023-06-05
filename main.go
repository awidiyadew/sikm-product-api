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

	return api.NewHandler(productService)
}

func main() {
	db, err := db.Connect(db.DBCredential{
		Host:         "localhost",
		Username:     "postgres",
		Password:     "Agustin22",
		DatabaseName: "tutor_db",
		Port:         5432,
	})
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})

	// TODO: create gin router
	router := gin.Default()
	handler := NewHandler(db)

	// handling method not allowed and route not found
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "page not found or method not allowed",
		})

	})

	productRouter := router.Group("/product")
	{
		productRouter.GET("/list", handler.GetListProduct)
		productRouter.GET("/:id", handler.GetProductDetail)
		productRouter.POST("/add", handler.StoreProduct)
		productRouter.PUT("/:id", handler.UpdateProduct)
	}
	router.Run(":3000")
}
