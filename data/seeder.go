package main

import (
	"encoding/json"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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
	NumReviews   int32     `json:"numReviews"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	IsAdmin   bool      `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func main() {
	var products []Product
	var users []User

	json.Unmarshal(productsData, &products)
	json.Unmarshal(userData, &users)

	db, err := gorm.Open(postgres.Open("postgres://moab:supersecure@localhost:5432/ecommerce"), &gorm.Config{})
	if err != nil {
		panic("deu ruim")
	}

	for _, product := range products {
		db.Exec("INSERT INTO products (name, description, brand, category, image, num_reviews, rating, price, count_in_stock, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
			product.Name,
			product.Description,
			product.Brand,
			product.Category,
			product.Image,
			product.NumReviews,
			product.Rating,
			product.Price,
			product.CountInStock,
			time.Now(),
			time.Now(),
		)
	}

	for _, user := range users {
		db.Exec("INSERT INTO users (name, email, password, is_admin, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
			user.Name,
			user.Email,
			user.Password,
			user.IsAdmin,
			time.Now(),
			time.Now(),
		)
	}

	log.Println("Data seed completed!")
}
