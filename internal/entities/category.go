package entities

import (
	"errors"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
)

var (
	ErrCategoryNameIsRequired = errors.New("category name is required")
)

type Category struct {
	ID   pkg_entities.UUID
	Name string
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		ID:   pkg_entities.NewUUID(),
		Name: name,
	}

	if err := category.validate(); err != nil {
		return nil, err
	}
	return category, nil
}

func (c *Category) validate() error {
	if c.Name == "" {
		return ErrCategoryNameIsRequired
	}
	return nil
}
