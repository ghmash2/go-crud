package models

import (
	"time"
)

type Category struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
