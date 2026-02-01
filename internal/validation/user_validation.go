package validation

import (
	"errors"
	"net/mail"
	"strings"
)

type UserValidation struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func ValidateUser(user UserValidation) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if _, err := mail.ParseAddress(user.Email); err != nil {
		return errors.New("email is invalid")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if len(user.Password) < 6 {
		return errors.New("password must be at least 6 characters long")
	}

	return nil
}