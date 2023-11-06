package product

import "time"

type Product struct {
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	Image        string    `json:"image"`
	Description  string    `json:"description"`
	Brand        string    `json:"brand"`
	Category     string    `json:"category"`
	Price        float64   `json:"price"`
	CountInStock int32     `json:"countInStock"`
	Rating       float32   `json:"rating"`
	Reviews      []Review  `json:"reviews"`
	NumReviews   int32     `json:"numReviews"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
