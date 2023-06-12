package model

import "gorm.io/gorm"

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name" gorm:"type:varchar(255);not null;"`
	Role       string `json:"role" gorm:"type:varchar(50);not null;"`
	Email      string `json:"email" gorm:"type:varchar(255);unique;not null;"`
	Password   string `json:"-" gorm:"type:varchar(255);not null;"`
}

type Product struct {
	gorm.Model `json:"-"`
	ID         uint     `json:"id" gorm:"primarykey;"`
	Name       string   `json:"name" gorm:"type:varchar(255);not null;"`
	Price      int      `json:"price" gorm:"type:numeric(10);not null;"`
	PostedBy   uint     `json:"posted_by" gorm:"type:varchar(255);not null;"`
	User       User     `json:"-" gorm:"foreignKey:PostedBy"` // product belong to user
	CategoryID int      `json:"category_id"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"` // product belong to category
}

type Category struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primarykey"`
	Name       string `json:"name" gorm:"type:varchar(255);not null;"`
}
