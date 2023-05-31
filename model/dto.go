package model

type ProductDetail struct {
	ID         uint     `json:"id"`
	Name       string   `json:"name"`
	Price      int      `json:"price"`
	User       User     `json:"poster"`
	Category   Category `json:"category"`
}
