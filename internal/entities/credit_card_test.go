package entities_test

import (
	"ettoreMB/controle_gastos/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateCreditCard_With_NamedRequiredError(t *testing.T) {

	_, err := entities.NewCreditCard("", 24)
	assert.Error(t, err)

}

func Test_CreateCreditCard_With_InvalidDueDateError(t *testing.T) {

	_, err := entities.NewCreditCard("card visa", 0)
	assert.Error(t, err)

	_, err = entities.NewCreditCard("card visa", 31)
	assert.Error(t, err)
}

func Test_CreateCredictCard_Success(t *testing.T) {
	card, err := entities.NewCreditCard("card visa", 11)
	expected := &entities.CreditCard{
		Id:         card.Id,
		Name:       "card visa",
		DueDateDay: 11,
	}

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expected, card)

}
