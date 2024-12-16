-- name: GettExpenses :many
SELECT * FROM expenses
ORDER BY transaction_date DESC
;