package permission

import "errors"

type Role struct {
	Id          string
	Name        string
	Permissions []Permission
}

func NewRole(Id string, Name string, Permissions []Permission) (*Role, error) {
	if len(Permissions) == 0 {
		return nil, errors.New("permissions cannot be empty")
	}

	return &Role{
		Id:          Id,
		Name:        Name,
		Permissions: Permissions,
	}, nil
}
