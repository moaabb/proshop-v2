package user

type Repository interface {
	GetUserByEmail(email string) (User, error)
}
