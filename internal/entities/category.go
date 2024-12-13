package entities

import (
	"errors"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
)

var (
	ErrCategoryNameIsRequired = errors.New("category name is required")
)

type CategoryId pkg_entities.UUID

type Category struct {
	ID   CategoryId
	Name string
}

func NewCategory(name string) (*Category, error) {
	category := &Category{
		ID:   CategoryId(pkg_entities.NewUUID()),
		Name: name,
	}

	if err := category.validate(); err != nil {
		return nil, err
	}
	return category, nil
}

func (c Category) NewCategoryId() CategoryId {
	id := CategoryId(pkg_entities.NewUUID())
	return id
}

func (c *Category) validate() error {
	if c.Name == "" {
		return ErrCategoryNameIsRequired
	}
	return nil
}
