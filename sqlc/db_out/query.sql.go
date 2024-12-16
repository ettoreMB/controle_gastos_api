// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
)

const gettExpenses = `-- name: GettExpenses :many
SELECT id, name, value, category_id, payment_method_id, status, credit_card_id, installment_count, current_installment, installment_count_value, transaction_date FROM expenses
ORDER BY transaction_date DESC
`

func (q *Queries) GettExpenses(ctx context.Context) ([]Expense, error) {
	rows, err := q.db.Query(ctx, gettExpenses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Value,
			&i.CategoryID,
			&i.PaymentMethodID,
			&i.Status,
			&i.CreditCardID,
			&i.InstallmentCount,
			&i.CurrentInstallment,
			&i.InstallmentCountValue,
			&i.TransactionDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}