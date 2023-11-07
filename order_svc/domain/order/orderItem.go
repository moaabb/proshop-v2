package order

import (
	"time"
)

type OrderItem struct {
	ID          int64     `json:"-"`
	OrderID     int       `json:"-"`
	ProductID   int       `json:"productId"`
	Name        string    `json:"name"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
