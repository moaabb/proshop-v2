package order

type Repository interface {
	GetById(id uint) (Order, error)
	GetByUserId(id uint) ([]Order, error)
	GetAll() ([]Order, error)
	Create(order Order) (Order, error)
	Update(orderId uint, order Order) (Order, error)
	UpdateToPaid(orderId uint, order Order) (Order, error)
	UpdateToDelivered(orderId uint, order Order) (Order, error)
	Delete(id uint) error
}
