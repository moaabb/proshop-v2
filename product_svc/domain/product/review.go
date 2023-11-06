package product

import (
	"time"
)

type Review struct {
	ID        int        `json:"id"`
	Rating    float64    `json:"rating"`
	Comment   string     `json:"comment"`
	UserID    int        `json:"userId"`
	ProductID int        `json:"productId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
