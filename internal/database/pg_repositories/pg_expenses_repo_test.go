package pg_repositories_test

import (
	"database/sql"
	"ettoreMB/controle_gastos/internal/database/pg_repositories"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgxtest"
	"github.com/stretchr/testify/assert"
)

func (t *testing.T) Test_GetAllExpenses_Should_Return_All_Expenses(t *testing.T) {

	conn := pgxpool.
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	defer db.Close()

	repo := pg_repositories.NewPGExpensesRepository(db)
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()
}
