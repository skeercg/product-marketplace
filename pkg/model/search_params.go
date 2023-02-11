package model

type SearchProductParams struct {
	Token             string
	ProductPriceFrom  int     `json:"product-price-from"`
	ProductPriceTo    int     `json:"product-price-to"`
	ProductName       string  `json:"product-name"`
	ProductRatingFrom float32 `json:"product-rating-from"`
}
