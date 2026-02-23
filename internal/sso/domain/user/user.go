package user

func NewUser() (*User, error) {
	return &User{}, nil
}

type User struct {
	Id       string
	Username string
	Password string
	Mail     Mail
	Picture  Picture
}

func (user *User) GetId() string {
	return user.Id
}
