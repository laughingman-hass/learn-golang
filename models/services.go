package models

func NewServices(connectionInfo string) (*Services, error) {
	db, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}

    return &Services {

    }, nil

	return nil, nil
}

type Services struct {
	Gallery GalleryService
	User    UserService
}
