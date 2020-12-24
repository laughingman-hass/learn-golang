package models

import (
	"errors"
	"fmt"

	"learn-golang/hash"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound        = errors.New("models: resource not found")
	ErrInvalidID       = errors.New("models: ID provided was invalid")
	ErrInvalidPassword = errors.New("models: incorrect password provided")
)

const userPwPepper = "random-secret-pepper"
const hmacSecretKey = "secret-hmac-key"

func NewUserServices(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(User{})
	hmac := hash.NewHMAC(hmacSecretKey)
	return &UserService{
		db:   db,
		hmac: hmac,
	}, nil
}

type UserService struct {
	db   *gorm.DB
	hmac hash.HMAC
}

func (us *UserService) ByID(id int) (*User, error) {
	var user User
	db := us.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) ByEmail(email string) (*User, error) {
	var user User
	db := us.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

func (us *UserService) BySession(token string) (*User, error) {
	tokenHash := us.hmac.Hash(token)
	var user User
	db := us.db.Where("session_token_hash", tokenHash)
	err := first(db, &user)
	return &user, err

}

func (us *UserService) Authenticate(email, password string) (*User, error) {
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+userPwPepper))
	fmt.Println(err)
	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		default:
			return nil, err
		}
	}

	return foundUser, nil
}

func first(db *gorm.DB, user *User) error {
	err := db.First(user).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err

}

func (us *UserService) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	if user.SessionToken != "" {
		user.SessionTokenHash = us.hmac.Hash(user.SessionToken)
	}
	return us.db.Create(user).Error
}

func (us *UserService) Update(user *User) error {
	if user.SessionToken != "" {
		user.SessionTokenHash = us.hmac.Hash(user.SessionToken)
	}
	return us.db.Save(user).Error
}

func (us *UserService) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return us.db.Delete(&user).Error
}

func (us *UserService) Close() error {
	return us.db.Close()
}

type User struct {
	gorm.Model
	Name             string
	Email            string `gorm:"not null;unique_index"`
	Password         string `gorm:"-"`
	PasswordHash     string `gorm:"not null"`
	SessionToken     string `gorm:"-"`
	SessionTokenHash string `gorm:"not null;unique_index"`
}
