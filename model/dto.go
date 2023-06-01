package model

type ProductDetail struct {
	ID       uint     `json:"id"`
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	User     User     `json:"poster"`
	Category Category `json:"category"`
}

type ProductRequest struct {
	Name       string `json:"name" binding:"required"`
	Price      int    `json:"price" binding:"required,gt=0"`
	CategoryID int    `json:"category_id" binding:"required"`
	PostedBy   int    `json:"posted_by" binding:"required"`
}
