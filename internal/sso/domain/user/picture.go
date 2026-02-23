package user

func NewPicture() (*Picture, error) {
	return nil, nil
}

type Picture struct {
	name     string
	path     string
	size     PictureSize
	mimeType string
}

type PictureSize struct {
	width  int
	height int
}
