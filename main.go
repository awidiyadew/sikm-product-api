package main

import (
	"log"
	"product-api/api"
	"product-api/db"
	"product-api/model"
	"product-api/repository"
	"product-api/service"

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
		Password:     "",
		DatabaseName: "",
		Port:         5432,
	})
	if err != nil {
		log.Fatal("failed to connect db", err)
	}
	db.AutoMigrate(&model.User{}, &model.Category{}, &model.Product{})

	// TODO: create gin router
}
