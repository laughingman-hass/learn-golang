package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	ErrNotFound = errors.New("models: resource not found")
)

func NewUserServices(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	return &UserService{
		db: db,
	}, nil
}

type UserService struct {
	db *gorm.DB
}

func (us *UserService) ByID(id int) (*User, error) {
	var user User
	err := us.db.Where("id = ?", id).First(&user).Error

	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (us *UserService) Create(user *User) error {
	return us.db.Create(user).Error
}

func (us *UserService) Close() error {
	return us.db.Close()
}

type User struct {
	gorm.Model
	Name  string
	Email string `grom:"not null;unique_index"`
}
