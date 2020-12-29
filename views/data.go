package views

import (
	"learn-golang/models"
)

const (
	AlertLevelError   = "danger"
	AlertLevelWarning = "warning"
	AlertLevelInfo    = "info"
	AlertLevelSuccess = "success"

	AlertMessageGeneric = "Something went wrong. Please try again, and contact us if the problem persists."
)

type Alert struct {
	Level   string
	Message string
}
type Data struct {
	Alert *Alert
	User  *models.User
	Yield interface{}
}

func (d *Data) SetAlert(err error) {
	if publicErr, ok := err.(PublicError); ok {
		d.Alert = &Alert{
			Level:   AlertLevelError,
			Message: publicErr.Public(),
		}
	} else {
		d.AlertError(AlertMessageGeneric)
	}
}

func (d *Data) AlertError(msg string) {
	d.Alert = &Alert{
		Level:   AlertLevelError,
		Message: msg,
	}
}

type PublicError interface {
	error
	Public() string
}
