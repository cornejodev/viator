package domain

import (
	"errors"
)

var ErrDemoHasNoName = errors.New("the demo has no name")
var ErrProductNotFound = errors.New("product not found")

type Demo struct {
	ID   int64
	Name string

	// CreatedAt time.Time
	// UpdatedAt time.Time
}

func (d Demo) HasName() bool {
	return d.Name != ""
}

type AddDemoForm struct {
	Name string `json:"name"`
}
