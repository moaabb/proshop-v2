package userdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/moaabb/ecommerce/user_svc/domain/user"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (ur *Repository) GetAll() ([]user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := ur.db.QueryContext(ctx, GetUsers)
	if err != nil {
		return nil, err
	}

	var users []user.User
	for rows.Next() {
		var u user.User
		err = rows.Scan(
			&u.Id,
			&u.Name,
			&u.Email,
			&u.IsAdmin,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
func (ur *Repository) GetById(id uint) (user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var u user.User

	err := ur.db.QueryRowContext(ctx, GetUserById, id).Scan(
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
func (ur *Repository) Create(u user.User) (user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var newUser user.User
	err := ur.db.QueryRowContext(ctx, CreateUser,
		u.Name,
		u.Email,
		u.Password,
		u.IsAdmin,
		time.Now(),
		time.Now(),
	).Scan(
		&newUser.Id,
		&newUser.Name,
		&newUser.Email,
		&newUser.Password,
		&newUser.IsAdmin,
		&newUser.CreatedAt,
		&newUser.UpdatedAt,
	)
	if err != nil {
		return user.User{}, err
	}

	return newUser, nil
}
func (ur *Repository) Update(uid uint, u user.User) (user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var updatedUser user.User
	err := ur.db.QueryRowContext(ctx, UpdateUser,
		u.Name,
		u.Email,
		u.Password,
		u.IsAdmin,
		time.Now(),
		uid,
	).Scan(
		&updatedUser.Id,
		&updatedUser.Name,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.IsAdmin,
		&updatedUser.CreatedAt,
		&updatedUser.UpdatedAt,
	)
	if err != nil {
		return user.User{}, err
	}

	return updatedUser, nil
}
func (ur *Repository) Delete(id uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := ur.db.ExecContext(ctx, DeleteUser, id)
	if err != nil {
		return err
	}

	return nil
}
