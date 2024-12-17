package pg_repositories

import (
	"context"
	mapper "ettoreMB/controle_gastos/internal/database/mappers"
	"ettoreMB/controle_gastos/internal/entities"
	db "ettoreMB/controle_gastos/sqlc/db_out"

	"github.com/jackc/pgx/v5"
)

type PGpensesRepository struct {
	db  *db.Queries
	ctx context.Context
}

func NewPGExpensesRepository(conn *pgx.Conn) *PGpensesRepository {
	repo := &PGpensesRepository{db: nil, ctx: context.Background()}
	repo.db = db.New(conn)
	return repo

}

func (r *PGpensesRepository) Create(expense entities.Expense) error {

	// _, err := r.db.Exec("INSERT INTO expenses () VALUES (?, ?)", expense.Name, expense.Value, expense.CategoryId, expense.PaymentMethodId, expense.TransactionDate, expense) // Adjust the SQL
	return nil
}

func (r *PGpensesRepository) GetByID(id string) (*entities.Expense, error) {
	return nil, nil
}

func (r *PGpensesRepository) GetAll() ([]entities.Expense, error) {

	expensesdb, err := r.db.GettExpenses(context.Background())
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

func (r *PGpensesRepository) Update(item entities.Expense) error {
	return nil
}

func (r *PGpensesRepository) Delete(id string) error {
	return nil
}
