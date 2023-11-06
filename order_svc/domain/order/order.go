package order

import (
	"time"
)

type Order struct {
	ID                  int64     `json:"id"`
	UserID              int       `json:"userId"`
	ShippingAddress     string    `json:"shippingAddress"`
	ShippingCity        string    `json:"shippingCity"`
	ShippingPostalCode  string    `json:"shippingPostalCode"`
	ShippingCountry     string    `json:"shippingCountry"`
	PaymentMethod       string    `json:"paymentMethod"`
	PaymentID           string    `json:"paymentID"`
	PaymentStatus       string    `json:"paymentStatus"`
	PaymentUpdateTime   string    `json:"paymentUpdateTime"`
	PaymentEmailAddress string    `json:"paymentEmailAddress"`
	ItemsPrice          float64   `json:"itemsPrice"`
	TaxPrice            float64   `json:"taxPrice"`
	ShippingPrice       float64   `json:"shippingPrice"`
	TotalPrice          float64   `json:"totalPrice"`
	IsPaid              bool      `json:"isPaid"`
	PaidAt              time.Time `json:"paidAt"`
	IsDelivered         bool      `json:"isDelivered"`
	DeliveredAt         time.Time `json:"deliveredAt"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	User                User      `json:"user,omitempty"`
}
