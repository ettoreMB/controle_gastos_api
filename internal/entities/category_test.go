package entities_test

import (
	"ettoreMB/controle_gastos/internal/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateCategory_With_Err_RequiredNameError(t *testing.T) {
	_, err := entities.NewCategory("")
	assert.Error(t, err)
	assert.Equal(t, entities.ErrCategoryNameIsRequired, err)
}

func Test_CreateCategory_Success(t *testing.T) {
	category, err := entities.NewPaymentMethod("nova categoria")
	if err != nil {
		t.Error(err)
	}
	assert.NoError(t, err)
	assert.Equal(t, category.Name, "nova categoria")

}
