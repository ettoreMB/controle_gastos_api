package entities

import (
	"errors"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
)

var (
	ErrPaymentMethodNameIsRequired = errors.New("payment method name is required")
)

type PaymentMethodId pkg_entities.UUID

type PaymentMethod struct {
	ID   PaymentMethodId
	Name string
}

func NewPaymentMethod(name string) (*PaymentMethod, error) {
	method := &PaymentMethod{
		ID:   PaymentMethodId(pkg_entities.NewUUID()),
		Name: name,
	}

	if err := method.validate(); err != nil {
		return nil, err
	}

	return method, nil
}

func (p *PaymentMethod) validate() error {
	if p.Name == "" {
		return ErrPaymentMethodNameIsRequired
	}
	return nil
}
