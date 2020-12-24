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

type UserDB interface {
	// query for single user
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	BySession(email string) (*User, error)

	// for altering user
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	// close a DB connection
	Close() error
}

// UserService is a set of methods used to manipulate and
// work with the user model
type UserService interface {
	// Authenticate will verify the provided email address and
	// password are correct. If they are correct, the user
	// corresponding to that email will be returned. Otherwise
	// it will return either: ErrNotFound, ErrInvalidPassword,
	// or another error if something goes wrong.
	Authenticate(email, password string) (*User, error)
	UserDB
}

func NewUserServices(connectionInfo string) (UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	return &userService{
		UserDB: &userValidator{
			UserDB: ug,
		},
	}, nil
}

type userService struct {
	UserDB
}

type userValidator struct {
	UserDB
}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(User{})
	hmac := hash.NewHMAC(hmacSecretKey)
	return &userGorm{
		db:   db,
		hmac: hmac,
	}, nil
}

type userGorm struct {
	db   *gorm.DB
	hmac hash.HMAC
}

func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

func (ug *userGorm) BySession(token string) (*User, error) {
	tokenHash := ug.hmac.Hash(token)
	var user User
	db := ug.db.Where("session_token_hash = ?", tokenHash)
	err := first(db, &user)
	return &user, err
}

func (us *userService) Authenticate(email, password string) (*User, error) {
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

func (ug *userGorm) Create(user *User) error {
	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	if user.SessionToken != "" {
		user.SessionTokenHash = ug.hmac.Hash(user.SessionToken)
	}
	return ug.db.Create(user).Error
}

func (ug *userGorm) Update(user *User) error {
	if user.SessionToken != "" {
		user.SessionTokenHash = ug.hmac.Hash(user.SessionToken)
	}
	return ug.db.Save(user).Error
}

func (ug *userGorm) Delete(id uint) error {
	if id == 0 {
		return ErrInvalidID
	}
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

func (ug *userGorm) Close() error {
	return ug.db.Close()
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
