package validation

import (
	"errors"
	"strings"
)

type CategoryValidation struct {
	Name string `json:"name"`
}

func ValidateCategory(category CategoryValidation) error {
	category.Name = strings.TrimSpace(category.Name)
	if category.Name == "" {
		return errors.New("name is required")
	}
	return nil
}