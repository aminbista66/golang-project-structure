package user

type Repository interface {
	FindByID(id int64) (*User, error)
	Create(u *User) error
}
