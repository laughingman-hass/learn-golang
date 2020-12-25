package models

import "strings"

var (
	ErrNotFound             modelError = "models: resource not found"
	ErrIDInvalid            modelError = "models: ID provided was invalid"
	ErrEmailRequired        modelError = "models: email address is required"
	ErrEmailInvalid         modelError = "models: email address is not valid"
	ErrEmailNotUnique       modelError = "models: email address already in use"
	ErrPasswordIncorrect    modelError = "models: incorrect password provided"
	ErrPasswordTooShort     modelError = "models: password must be 8 characters long"
	ErrPasswordRequired     modelError = "models: password is required"
	ErrSessionTokenInvalid  modelError = "models: session token must be at least 32 bytes"
	ErrSessionTokenRequired modelError = "models: session token is required"
)

type modelError string

func (e modelError) Error() string {
	return string(e)
}

func (e modelError) Public() string {
	s := strings.Replace(string(e), "models: ", "", 1)
	return strings.Title(s)
}
