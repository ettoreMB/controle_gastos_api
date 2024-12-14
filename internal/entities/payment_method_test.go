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

func Test_CreatePaymentMethod_Success(t *testing.T) {
	method, err := entities.NewPaymentMethod("novo metodo")
	assert.NoError(t, err)

	assert.Equal(t, method.Name, "novo metodo")
	assert.NotNil(t, method.ID)
}
