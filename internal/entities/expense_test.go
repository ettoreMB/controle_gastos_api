package entities_test

import (
	"ettoreMB/controle_gastos/internal/entities"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_CreateNewExpense(t *testing.T) {
	paymentID := entities.PaymentMethodId(pkg_entities.NewUUID())
	categoryID := entities.CategoryId(pkg_entities.NewUUID())

	createCommand := &entities.CreateExpenseCommand{
		Name:              "food",
		PaymentMethodId:   paymentID,
		Value:             10.50,
		CategoryId:        categoryID,
		IsInstallmentPlan: false,
		InstallmentsCount: 0,
	}

	expense := entities.NewExpense(createCommand)
	expected := entities.Expense{
		ID:                expense.ID,
		Name:              "food",
		Value:             10.50,
		CategoryId:        categoryID,
		PaymentMethodId:   paymentID,
		Type:              "expense",
		Status:            "pending",
		TransactionDate:   time.Now().Format("2006-01-02"),
		IsInstallmentPlan: false,
		InstallmentsCount: 0,
	}

	assert.Equal(t, expense, &expected)
}
