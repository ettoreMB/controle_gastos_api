package mapper_test

import (
	"crypto/rand"
	mapper "ettoreMB/controle_gastos/internal/database/mappers"
	db "ettoreMB/controle_gastos/sqlc/db_out"
	"log"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/assert"
)

func generateUUID() [16]byte {
	var uuid [16]byte
	_, err := rand.Read(uuid[:]) // Fill the array with random bytes
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}

	// Set version (4) in the most significant 4 bits of byte 6
	uuid[6] = (uuid[6] & 0x0F) | 0x40

	// Set variant (RFC 4122) in the most significant 2 bits of byte 8
	uuid[8] = (uuid[8] & 0x3F) | 0x80

	return uuid
}

func TestExpensePgToDomainMapper(t *testing.T) {
	uuid := generateUUID()

	dbExpense := db.Expense{
		ID: pgtype.UUID{
			Bytes: uuid,
			Valid: true,
		},
		Name:  "Test Expense",
		Value: 100.0,
		CategoryID: pgtype.UUID{
			Bytes: uuid,
			Valid: true,
		},
		PaymentMethodID: pgtype.UUID{
			Bytes: uuid,
			Valid: true,
		},
		Status:          pgtype.Text{Valid: true, String: "Paid"},
		TransactionDate: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	}

	expense, err := mapper.ExpensePgToDomainMapper(dbExpense)
	if err != nil {
		t.Errorf("Error mapping Expense: %v", err)
	}

	assert.Equal(t, "Test Expense", expense.Name)
}
