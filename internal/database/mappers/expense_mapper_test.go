package mapper_test

import (
	mapper "ettoreMB/controle_gastos/internal/database/mappers"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	db "ettoreMB/controle_gastos/sqlc/db_out"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func TestExpensePgToDomainMapper(t *testing.T) {
	uuid := pkg_entities.NewUUID()

	dbExpense := db.Expense{
		ID:                    uuid,
		Status:                pgtype.Text{Valid: true, String: "Paid"},
		Name:                  "Test Expense",
		Value:                 100.0,
		CategoryID:            uuid,
		PaymentMethodID:       uuid,
		CreditCardID:          uuid,
		InstallmentCount:      pgtype.Int4{Int32: 1, Valid: true},
		CurrentInstallment:    pgtype.Int4{Int32: 1, Valid: true},
		InstallmentCountValue: pgtype.Float8{Float64: 100.0, Valid: true},
		TransactionDate:       time.Now(),
	}

	expense, err := mapper.ExpensePgToDomainMapper(dbExpense)
	if err != nil {
		t.Errorf("Error mapping Expense: %v", err)
	}

	assert.Equal(t, "Test Expense", expense.Name)
}
