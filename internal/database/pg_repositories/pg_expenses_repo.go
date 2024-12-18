package pg_repositories

import (
	"context"
	mapper "ettoreMB/controle_gastos/internal/database/mappers"
	"ettoreMB/controle_gastos/internal/entities"
	db "ettoreMB/controle_gastos/sqlc/db_out"

	"github.com/jackc/pgx/v5"
)

type PGEpensesRepository struct {
	ctx  context.Context
	conn *pgx.Conn
}

func NewPGExpensesRepository(ctx context.Context, conn *pgx.Conn) *PGEpensesRepository {
	repo := &PGEpensesRepository{
		ctx:  ctx,
		conn: conn,
	}
	return repo

}

func (r *PGEpensesRepository) Create(expense entities.Expense) error {

	// _, err := r.db.Exec("INSERT INTO expenses () VALUES (?, ?)", expense.Name, expense.Value, expense.CategoryId, expense.PaymentMethodId, expense.TransactionDate, expense) // Adjust the SQL
	return nil
}

func (r *PGEpensesRepository) GetByID(id string) (*entities.Expense, error) {
	return nil, nil
}

func (r *PGEpensesRepository) GetAll() ([]entities.Expense, error) {

	expensesdb, err := db.Queries.(r.ctx)
	if err != nil {
		return nil, err
	}
	expenses := make([]entities.Expense, len(expensesdb))

	for i, e := range expensesdb {
		ex, err := mapper.ExpensePgToDomainMapper(e)
		expenses[i] = *ex
		if err != nil {
			return nil, err
		}
	}

	return expenses, nil
}

func (r *PGEpensesRepository) Update(item entities.Expense) error {
	return nil
}

func (r *PGEpensesRepository) Delete(id string) error {
	return nil
}
