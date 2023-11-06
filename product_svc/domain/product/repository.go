package product

type Repository interface {
	GetById(id uint) (Product, error)
	GetAll() ([]Product, error)
	GetTopProducts() ([]Product, error)
	Create(product Product) (Product, error)
	Update(pid uint, product Product) (Product, error)
	Delete(id uint) error
}
