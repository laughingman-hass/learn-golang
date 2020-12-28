package models

import "github.com/jinzhu/gorm"

func NewServices(connectionInfo string) (*Services, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)

	db.AutoMigrate(User{})

	return &Services{
		User: NewUserServices(db),
	}, nil
}

type Services struct {
	Gallery GalleryService
	User    UserService
}
