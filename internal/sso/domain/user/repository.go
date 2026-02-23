package user

type Repository interface {
	Create(user *User) error
	GetById(id string) (*User, error)
	Get() ([]*User, error)
	Update(user *User) error
	Delete(id string) error
}
