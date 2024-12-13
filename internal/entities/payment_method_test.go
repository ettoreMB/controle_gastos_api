package entities_test

import (
	"ettoreMB/controle_gastos/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePaymentMethod_With_NoName_Error(t *testing.T) {
	_, err := entities.NewPaymentMethod("")
	if err == nil {
		t.Error("expected error")
	}

	assert.Equal(t, entities.ErrPaymentMethodNameIsRequired, err)
}
