package product

type Repository interface {
	GetById(id uint) (Product, error)
	GetAll(page int) (GetProductsDto, error)
	GetTopProducts() ([]Product, error)
	Create(product Product) (Product, error)
	CreateReview(review Review) (Review, error)
	Update(pid uint, product Product) (Product, error)
	Delete(id uint) error
}
