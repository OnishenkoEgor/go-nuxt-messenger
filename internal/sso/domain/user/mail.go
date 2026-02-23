package user

func NewMail(address string, domain string) (*Mail, error) {
	return &Mail{
		address: address,
		domain:  domain,
	}, nil
}

type Mail struct {
	address string
	domain  string
}

func (m *Mail) String() string {
	return m.address + "@" + m.domain
}
