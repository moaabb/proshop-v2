package order

import (
	"time"
)

type Order struct {
	ID                  int64           `json:"id"`
	UserID              int             `json:"-"`
	ShippingAddress     ShippingAddress `json:"shippingAddress"`
	PaymentMethod       string          `json:"paymentMethod"`
	PaymentID           string          `json:"paymentID"`
	PaymentStatus       string          `json:"paymentStatus"`
	PaymentUpdateTime   string          `json:"paymentUpdateTime"`
	PaymentEmailAddress string          `json:"paymentEmailAddress"`
	ItemsPrice          string          `json:"itemsPrice"`
	TaxPrice            string          `json:"taxPrice"`
	ShippingPrice       string          `json:"shippingPrice"`
	TotalPrice          string          `json:"totalPrice"`
	IsPaid              bool            `json:"isPaid"`
	PaidAt              time.Time       `json:"paidAt"`
	IsDelivered         bool            `json:"isDelivered"`
	OrderItems          []OrderItem     `json:"orderItems"`
	DeliveredAt         time.Time       `json:"deliveredAt"`
	CreatedAt           time.Time       `json:"createdAt"`
	UpdatedAt           time.Time       `json:"updatedAt"`
	User                User            `json:"user,omitempty"`
}
