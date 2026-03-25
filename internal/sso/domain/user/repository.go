package user

type Repository interface {
	Create(user *User) error
	GetById(id int) (*User, error)
	Get() ([]*User, error)
	Update(user *User) error
	Delete(id int) error
}
