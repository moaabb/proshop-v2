package userdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/moaabb/ecommerce/auth_svc/domain/user"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (ur *Repository) GetUserByEmail(email string) (user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var u user.User

	err := ur.db.QueryRowContext(ctx, GetUserByEmail, email).Scan(
		&u.Id,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.IsAdmin,
		&u.CreatedAt,
		&u.UpdatedAt,
	)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}
