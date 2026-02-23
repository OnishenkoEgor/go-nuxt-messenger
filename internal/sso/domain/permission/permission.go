package permission

type Permission struct {
	Id   string
	Name string
}

func NewPermission(Id string, Name string) (*Permission, error) {
	return &Permission{
		Id:   Id,
		Name: Name,
	}, nil
}
