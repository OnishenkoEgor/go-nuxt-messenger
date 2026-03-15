package user

func NewUser(login string, password string) (*User, error) {
	return &User{
		login:    login,
		password: password,
	}, nil
}

type User struct {
	id       int
	login    string
	password string
	//Mail     Mail
	//Picture  Picture
}

func (user *User) GetId() int {
	return user.id
}

func (user *User) SetId(id int) {
	user.id = id
}

func (user *User) GetLogin() string { return user.login }

func (user *User) SetLogin(login string) {
	user.login = login
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) SetPassword(password string) {
	user.password = password
}
