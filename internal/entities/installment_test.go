package entities_test

import (
	"ettoreMB/controle_gastos/internal/entities"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateStallment_Without_Counts(t *testing.T) {
	cardId := entities.CreditcardId(pkg_entities.NewUUID())
	st := entities.NewStallment(cardId, 0, 100.00)

	expected := &entities.InstallMent{
		CreditCardId:       cardId,
		InstallmentsCount:  0,
		CurrentInstallment: 1,
		TotalValue:         100.00,
		CountPrice:         100.00,
	}

	assert.Equal(t, expected, st)
}

func Test_CreateStallment_with_counts(t *testing.T) {
	cardId := entities.CreditcardId(pkg_entities.NewUUID())
	st := entities.NewStallment(cardId, 2, 100.00)
	expected := &entities.InstallMent{
		CreditCardId:       cardId,
		InstallmentsCount:  2,
		CurrentInstallment: 1,
		TotalValue:         100.00,
		CountPrice:         50.00,
	}

	assert.Equal(t, st, expected)
}
