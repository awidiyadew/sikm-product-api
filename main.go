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
		Password:     "Postgres",
		DatabaseName: "product_api1",
		Port:         5432,
	})
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})

	router := gin.Default()
	handler := NewHandler(db)

	router.GET("/hello/:name", func(ctx *gin.Context) {
		name, _ := ctx.Params.Get("name")
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello " + name,
		})
	})

	productRouter := router.Group("/product")
	{
		productRouter.GET("/list", handler.GetListProduct)
		productRouter.GET("/:id", handler.GetProductDetail)
		productRouter.POST("/add", handler.StoreProduct)
		productRouter.DELETE("/:id", handler.DeleteProduct)
		productRouter.PUT("/:id", handler.UpdateProduct)
	}

	router.Run(":3000")
}
