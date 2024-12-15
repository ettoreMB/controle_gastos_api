package entities

import (
	"errors"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
)

var (
	errNamerequiredError = errors.New("credit card name is required")
	errInvalidDueDateDay = errors.New("duedateday is invalid")
)

type CreditcardId pkg_entities.UUID

type CreditCard struct {
	Id         CreditcardId
	Name       string
	DueDateDay int
}

func NewCreditCard(name string, duedate int) (*CreditCard, error) {
	creditcard := &CreditCard{
		Id:         CreditcardId(pkg_entities.NewUUID()),
		Name:       name,
		DueDateDay: duedate,
	}
	if err := creditcard.validate(); err != nil {
		return nil, err
	}

	return creditcard, nil

}

func (c *CreditCard) validate() error {
	if c.Name == "" {
		return errNamerequiredError
	}
	if c.DueDateDay == 0 || c.DueDateDay > 30 {
		return errInvalidDueDateDay
	}

	return nil
}
