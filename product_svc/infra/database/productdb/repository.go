package productdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/moaabb/ecommerce/product_svc/domain/product"
	"go.uber.org/zap"
)

type Repository struct {
	db *sql.DB
	l  *zap.Logger
}

func NewRepository(db *sql.DB, l *zap.Logger) *Repository {
	return &Repository{
		db: db,
		l:  l,
	}
}

func (pr *Repository) getAll(query string, args ...interface{}) ([]product.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := pr.db.QueryContext(ctx, query, args...)
	if err != nil {
		pr.l.Error("error getting data", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var products []product.Product
	for rows.Next() {
		var p product.Product
		err := scanProduct(&p, rows)
		if err != nil {
			pr.l.Error("error getting product", zap.Error(err))
			return nil, err
		}
		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		pr.l.Error("error getting data", zap.Error(err))
		return nil, err
	}

	return products, nil
}

func (pr *Repository) GetTopProducts() ([]product.Product, error) {
	return pr.getAll(GetTopProducts)
}

func (pr *Repository) GetAll() ([]product.Product, error) {
	return pr.getAll(GetProducts)
}

func (pr *Repository) GetById(id uint) (product.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var p product.Product

	err := scanProduct(&p, pr.db.QueryRowContext(ctx, GetProductById, id))
	if err != nil {
		pr.l.Error("error getting product", zap.Error(err))
		return product.Product{}, err
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	rows, err := pr.db.QueryContext(ctx, GetReviewsByProduct, p.Id)
	if err != nil {
		pr.l.Error("error getting review data", zap.Error(err))
		return product.Product{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var r product.Review
		err = rows.Scan(
			&r.ID,
			&r.Rating,
			&r.Comment,
			&r.UpdatedAt,
			&r.ProductID,
			&r.CreatedAt,
			&r.UpdatedAt,
		)
		if err != nil {
			pr.l.Error("error getting data", zap.Error(err))
			return product.Product{}, err
		}

		p.Reviews = append(p.Reviews, r)
	}
	if err := rows.Err(); err != nil {
		pr.l.Error("error getting data", zap.Error(err))
		return product.Product{}, err
	}

	if len(p.Reviews) == 0 {
		p.Reviews = []product.Review{}
	}

	return p, nil
}

func (pr *Repository) Create(p product.Product) (product.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var newProduct product.Product
	err := scanProduct(&newProduct, pr.db.QueryRowContext(ctx, CreateProduct))
	if err != nil {
		pr.l.Error("error creating product", zap.Error(err))
		return product.Product{}, err
	}

	return newProduct, nil
}

func (pr *Repository) Update(pid uint, p product.Product) (product.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var updatedProduct product.Product
	err := scanProduct(
		&updatedProduct,
		pr.db.QueryRowContext(
			ctx, UpdateProduct,
			p.Name,
			p.Description,
			p.Brand,
			p.Category,
			p.Image,
			p.NumReviews,
			p.Rating,
			p.Price,
			p.CountInStock,
			time.Now(),
			pid,
		),
	)

	if err != nil {
		pr.l.Error("error updating product", zap.Error(err))
		return product.Product{}, err
	}

	return updatedProduct, nil
}

func (pr *Repository) Delete(id uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := pr.db.ExecContext(ctx, DeleteProduct, id)
	if err != nil {
		pr.l.Error("error deleting product", zap.Error(err))
	}

	return err
}

func scanProduct(p *product.Product, row interface{}) error {
	var err error
	switch row := row.(type) {
	case *sql.Row:
		err = row.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Brand,
			&p.Category,
			&p.Image,
			&p.NumReviews,
			&p.Rating,
			&p.Price,
			&p.CountInStock,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
	case *sql.Rows:
		err = row.Scan(
			&p.Id,
			&p.Name,
			&p.Description,
			&p.Brand,
			&p.Category,
			&p.Image,
			&p.NumReviews,
			&p.Rating,
			&p.Price,
			&p.CountInStock,
			&p.CreatedAt,
			&p.UpdatedAt,
		)

	}

	return err

}
