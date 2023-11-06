package user

type Repository interface {
	GetById(id uint) (User, error)
	GetAll() ([]User, error)
	Create(User User) (User, error)
	Update(pid uint, User User) (User, error)
	Delete(id uint) error
}
