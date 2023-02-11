package model

type Product struct {
	ProductId     int     `json:"id"`
	ProductPrice  int     `json:"price"`
	ProductName   string  `json:"name"`
	ProductRating float32 `json:"rating"`
}
