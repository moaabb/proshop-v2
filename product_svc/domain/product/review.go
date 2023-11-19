package product

import (
	"time"
)

type Review struct {
	ID        int        `json:"id"`
	Rating    float64    `json:"rating"`
	Comment   string     `json:"comment"`
	UserID    uint       `json:"userId"`
	User      User       `json:"user"`
	ProductID int        `json:"productId"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
