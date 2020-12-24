package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"learn-golang/hash"
	"learn-golang/rand"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNotFound        = errors.New("models: resource not found")
	ErrInvalidID       = errors.New("models: ID provided was invalid")
	ErrInvalidPassword = errors.New("models: incorrect password provided")
	ErrEmailRequired   = errors.New("models: email address is required")
	ErrInvalidEmail    = errors.New("models: email address is not valid")
	ErrEmailNotUnique  = errors.New("models: email address already in use")
)

const userPwPepper = "random-secret-pepper"
const hmacSecretKey = "secret-hmac-key"

type User struct {
	gorm.Model
	Name             string
	Email            string `gorm:"not null;unique_index"`
	Password         string `gorm:"-"`
	PasswordHash     string `gorm:"not null"`
	SessionToken     string `gorm:"-"`
	SessionTokenHash string `gorm:"not null;unique_index"`
}

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

	hmac := hash.NewHMAC(hmacSecretKey)
	uv := newUserValidator(ug, hmac)

	return &userService{
		UserDB: uv,
	}, nil
}

var _ UserService = &userService{}

type userService struct {
	UserDB
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

type userValFunc func(*User) error

func runUserValFuncs(user *User, fns ...userValFunc) error {
	for _, fn := range fns {
		if err := fn(user); err != nil {
			return err
		}
	}
	return nil
}

var _ UserDB = &userValidator{}

func newUserValidator(udb UserDB, hmac hash.HMAC) *userValidator {
	return &userValidator{
		UserDB:     udb,
		hmac:       hmac,
		emailRegex: regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"),
	}
}

type userValidator struct {
	UserDB
	hmac       hash.HMAC
	emailRegex *regexp.Regexp
}

// BySession will hash the session token and then call
// BySession on the subsequent UserDB layer.
func (uv *userValidator) BySession(token string) (*User, error) {
	user := User{
		SessionToken: token,
	}
	if err := runUserValFuncs(&user, uv.hmacSessionToken); err != nil {
		return nil, err
	}
	return uv.UserDB.BySession(user.SessionTokenHash)
}

func (uv *userValidator) ByEmail(email string) (*User, error) {
	user := User{
		Email: email,
	}
	if err := runUserValFuncs(&user, uv.normalizeEmail); err != nil {
		return nil, err
	}
	return uv.UserDB.ByEmail(user.Email)
}

func (uv *userValidator) Create(user *User) error {

	err := runUserValFuncs(
		user,
		uv.bcryptPassword,
		uv.defaultSessionToken,
		uv.hmacSessionToken,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emailUnique,
	)
	if err != nil {
		return err
	}
	return uv.UserDB.Create(user)
}

func (uv *userValidator) Update(user *User) error {
	err := runUserValFuncs(
		user,
		uv.bcryptPassword,
		uv.hmacSessionToken,
		uv.normalizeEmail,
		uv.requireEmail,
		uv.emailFormat,
		uv.emailUnique,
	)
	if err != nil {
		return err
	}
	return uv.UserDB.Update(user)
}

func (uv *userValidator) Delete(id uint) error {
	var user User
	user.ID = id
	err := runUserValFuncs(&user, uv.idGreaterThan(0))
	if err != nil {
		return err
	}
	return uv.UserDB.Delete(id)
}

var _ UserDB = &userGorm{}

type userGorm struct {
	db *gorm.DB
}

// bcryptPassword will hash a user's password with a predefined
// pepper (userPwPepper) and bcrypt if the Password field is
// not an empty string
func (uv *userValidator) bcryptPassword(user *User) error {
	if user.Password == "" {
		return nil
	}

	pwBytes := []byte(user.Password + userPwPepper)
	hashedBytes, err := bcrypt.GenerateFromPassword(pwBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedBytes)
	user.Password = ""
	return nil
}

func (uv *userValidator) hmacSessionToken(user *User) error {
	if user.SessionToken == "" {
		return nil
	}
	user.SessionTokenHash = uv.hmac.Hash(user.SessionToken)
	return nil
}

func (uv *userValidator) defaultSessionToken(user *User) error {
	if user.SessionToken != "" {
		return nil
	}
	token, err := rand.NewSessionToken()
	if err != nil {
		return err
	}
	user.SessionToken = token
	return nil
}

func (uv *userValidator) idGreaterThan(n uint) userValFunc {
	return userValFunc(func(user *User) error {
		if user.ID <= n {
			return ErrInvalidID
		}
		return nil
	})
}

func (uv *userValidator) normalizeEmail(user *User) error {
	user.Email = strings.ToLower(user.Email)
	user.Email = strings.TrimSpace(user.Email)
	return nil
}

func (uv *userValidator) requireEmail(user *User) error {
	if user.Email == "" {
		return ErrEmailRequired
	}
	return nil
}

func (uv *userValidator) emailFormat(user *User) error {
	if user.Email == "" {
		return nil
	}
	if !uv.emailRegex.MatchString(user.Email) {
		return ErrInvalidEmail
	}
	return nil
}

func (uv *userValidator) emailUnique(user *User) error {
	existing, err := uv.ByEmail(user.Email)
	if err == ErrNotFound {
		return nil
	}
	if err != nil {
		return err
	}

	if user.ID != existing.ID {
		return ErrEmailNotUnique
	}
	return nil
}

func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(User{})
	return &userGorm{
		db: db,
	}, nil
}

// Query the database for a User by an ID
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id)
	err := first(db, &user)
	return &user, err
}

// Query the database for a User by an email
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email)
	err := first(db, &user)
	return &user, err
}

// Query the database for a User by a session token
func (ug *userGorm) BySession(tokenHash string) (*User, error) {
	var user User
	db := ug.db.Where("session_token_hash = ?", tokenHash)
	err := first(db, &user)
	return &user, err
}

func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(user).Error
}

func (ug *userGorm) Update(user *User) error {
	return ug.db.Save(user).Error
}

func (ug *userGorm) Delete(id uint) error {
	user := User{Model: gorm.Model{ID: id}}
	return ug.db.Delete(&user).Error
}

func (ug *userGorm) Close() error {
	return ug.db.Close()
}

func first(db *gorm.DB, user *User) error {
	err := db.First(user).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err

}
