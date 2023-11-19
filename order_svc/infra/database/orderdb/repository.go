package orderdb

import (
	"context"
	"database/sql"
	"time"

	"github.com/moaabb/ecommerce/order_svc/domain/order"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (or *Repository) GetAll() ([]order.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := or.db.QueryContext(ctx, GetOrders)
	if err != nil {
		return nil, err
	}

	var orders []order.Order
	for rows.Next() {
		o, err := scanOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (or *Repository) GetByUserId(id uint) ([]order.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := or.db.QueryContext(ctx, GetByUserId, id)
	if err != nil {
		return nil, err
	}

	var orders []order.Order
	for rows.Next() {
		var o order.Order
		o, err := scanOrder(rows)
		if err != nil {
			return nil, err
		}

		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}
func (or *Repository) GetById(id uint) (order.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	o, err := scanOrder(or.db.QueryRowContext(ctx, GetOrderById, id))
	if err != nil {
		return order.Order{}, err
	}

	var orderItems []order.OrderItem
	rows, err := or.db.QueryContext(ctx, GetOrderItems, o.ID)
	if err != nil {
		return order.Order{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var oi order.OrderItem
		err = rows.Scan(
			&oi.ProductID,
			&oi.Name,
			&oi.Image,
			&oi.Description,
			&oi.Price,
			&oi.Quantity,
		)
		if err != nil {
			return order.Order{}, err
		}

		orderItems = append(orderItems, oi)
	}
	if rows.Err() != nil {
		return order.Order{}, err
	}

	o.OrderItems = orderItems

	return o, nil
}
func (or *Repository) Create(o order.Order) (order.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*8)
	defer cancel()

	var newOrder order.Order
	err := or.db.QueryRowContext(ctx, CreateOrder,
		o.UserID,
		o.ShippingAddress.Address,
		o.ShippingAddress.City,
		o.ShippingAddress.PostalCode,
		o.ShippingAddress.Country,
		o.PaymentMethod,
		o.PaymentID,
		o.PaymentStatus,
		o.PaymentUpdateTime,
		o.PaymentEmailAddress,
		o.ItemsPrice,
		o.TaxPrice,
		o.ShippingPrice,
		o.TotalPrice,
		o.IsPaid,
		o.PaidAt,
		o.IsDelivered,
		o.DeliveredAt,
		time.Now(),
		time.Now(),
	).Scan(&newOrder.ID)
	if err != nil {
		return order.Order{}, err
	}

	for _, v := range o.OrderItems {
		_, err = or.db.ExecContext(ctx, CreateOrderItem,
			newOrder.ID,
			v.Image,
			v.ProductID,
			v.Quantity,
			v.Price,
			time.Now(),
			time.Now(),
		)
		if err != nil {
			return order.Order{}, err
		}
	}

	newOrder, _ = or.GetById(uint(newOrder.ID))

	return newOrder, nil
}
func (or *Repository) Update(orderID uint, updatedOrder order.Order) (order.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Execute the update query
	updatedOrder, err := scanOrder(or.db.QueryRowContext(ctx, UpdateOrder,
		updatedOrder.UserID,
		updatedOrder.ShippingAddress.Address,
		updatedOrder.ShippingAddress.City,
		updatedOrder.ShippingAddress.PostalCode,
		updatedOrder.ShippingAddress.Country,
		updatedOrder.PaymentMethod,
		updatedOrder.PaymentID,
		updatedOrder.PaymentStatus,
		updatedOrder.PaymentUpdateTime,
		updatedOrder.PaymentEmailAddress,
		updatedOrder.ItemsPrice,
		updatedOrder.TaxPrice,
		updatedOrder.ShippingPrice,
		updatedOrder.TotalPrice,
		updatedOrder.IsPaid,
		updatedOrder.PaidAt,
		updatedOrder.IsDelivered,
		updatedOrder.DeliveredAt,
		time.Now(),
		orderID,
	))

	if err != nil {
		return order.Order{}, err
	}

	return updatedOrder, nil
}

func (or *Repository) UpdateToPaid(orderID uint, updatedOrder order.Order) (order.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	// Execute the update query
	_, err := or.db.ExecContext(ctx, UpdateOrderToPaid,
		updatedOrder.PaymentID,
		updatedOrder.PaymentStatus,
		updatedOrder.PaymentUpdateTime,
		updatedOrder.PaymentEmailAddress,
		updatedOrder.IsPaid,
		updatedOrder.PaidAt,
		time.Now(),
		orderID,
	)

	if err != nil {
		return order.Order{}, err
	}

	updatedOrder, _ = or.GetById(orderID)

	return updatedOrder, nil
}

func (or *Repository) Delete(id uint) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := or.db.ExecContext(ctx, DeleteOrder, id)
	if err != nil {
		return err
	}

	return nil
}

func scanOrder(row interface{}) (order.Order, error) {
	var o order.Order
	var err error

	switch row := row.(type) {
	case *sql.Row:
		err = row.Scan(
			&o.ID,
			&o.UserID,
			&o.ShippingAddress.Address,
			&o.ShippingAddress.City,
			&o.ShippingAddress.PostalCode,
			&o.ShippingAddress.Country,
			&o.PaymentMethod,
			&o.PaymentID,
			&o.PaymentStatus,
			&o.PaymentUpdateTime,
			&o.PaymentEmailAddress,
			&o.ItemsPrice,
			&o.TaxPrice,
			&o.ShippingPrice,
			&o.TotalPrice,
			&o.IsPaid,
			&o.PaidAt,
			&o.IsDelivered,
			&o.DeliveredAt,
			&o.CreatedAt,
			&o.UpdatedAt,
			&o.User.Id,
			&o.User.Name,
			&o.User.Email,
			&o.User.CreatedAt,
			&o.User.UpdatedAt,
		)
	case *sql.Rows:
		err = row.Scan(
			&o.ID,
			&o.UserID,
			&o.ShippingAddress.Address,
			&o.ShippingAddress.City,
			&o.ShippingAddress.PostalCode,
			&o.ShippingAddress.Country,
			&o.PaymentMethod,
			&o.PaymentID,
			&o.PaymentStatus,
			&o.PaymentUpdateTime,
			&o.PaymentEmailAddress,
			&o.ItemsPrice,
			&o.TaxPrice,
			&o.ShippingPrice,
			&o.TotalPrice,
			&o.IsPaid,
			&o.PaidAt,
			&o.IsDelivered,
			&o.DeliveredAt,
			&o.CreatedAt,
			&o.UpdatedAt,
			&o.User.Id,
			&o.User.Name,
			&o.User.Email,
			&o.User.CreatedAt,
			&o.User.UpdatedAt,
		)

	}

	return o, err
}
