package user

import "time"

type User struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	IsAdmin   bool      `json:"isAdmin"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
