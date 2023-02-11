package model

type GradeProductParams struct {
	Token     string
	ProductId int     `json:"product-id"`
	Rating    float32 `json:"rating"`
}
