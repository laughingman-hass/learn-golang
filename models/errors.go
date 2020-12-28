package models

import "strings"

var (
	ErrNotFound          modelError = "models: resource not found"
	ErrEmailRequired     modelError = "models: email address is required"
	ErrEmailInvalid      modelError = "models: email address is not valid"
	ErrEmailNotUnique    modelError = "models: email address already in use"
	ErrPasswordIncorrect modelError = "models: incorrect password provided"
	ErrPasswordTooShort  modelError = "models: password must be 8 characters long"
	ErrPasswordRequired  modelError = "models: password is required"
	ErrTitleRequired     modelError = "models: title is required"

	ErrIDInvalid            privateError = "models: ID provided was invalid"
	ErrSessionTokenInvalid  privateError = "models: session token must be at least 32 bytes"
	ErrSessionTokenRequired privateError = "models: session token is required"
	ErrUserIDRequired       privateError = "models: user ID is required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	return strings.Title(s)
}

type privateError string

func (e privateError) Error() string {
	return string(e)
}
