package pg_repositories

import (
	"context"
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
	result := make([]entities.Expense, len(expensesdb))

	for _, e := range expensesdb {
		ex := entities.NewExpense(&entities.CreateExpenseCommand{
			Name:                  e.Name,
			Value:                 float32(e.Value),
			CategoryId:            entities.New,
			PaymentMethodId:       e.PaymentMethodID.String(),
			Status:                e.Status.String(),
			CreditCardId:          e.CreditCardID.String(),
			InstallmentCount:      e.InstallmentCount.Int32,
			CurrentInstallment:    e.CurrentInstallment.Int32,
			InstallmentCountValue: e.InstallmentCountValue.Float64,
			TransactionDate:       e.TransactionDate.Time,
		})
		expenses = append(expenses, entities.Expense{
			ID:                    expensesdb.ID.String(),
			Name:                  expensesdb.Name,
			Value:                 expensesdb.Value,
			CategoryId:            expensesdb.CategoryID.String(),
			PaymentMethodId:       expensesdb.PaymentMethodID.String(),
			Status:                expensesdb.Status.String(),
			CreditCardId:          expensesdb.CreditCardID.String(),
			InstallmentCount:      expensesdb.InstallmentCount.Int32,
			CurrentInstallment:    expensesdb.CurrentInstallment.Int32,
			InstallmentCountValue: expensesdb.InstallmentCountValue.Float64,
			TransactionDate:       expensesdb.TransactionDate.Time,
		})
	}

	return expenses, nil
}

func (r *PGpensesRepository) Update(item entities.Expense) error {
	return nil
}

func (r *PGpensesRepository) Delete(id string) error {
	return nil
}
