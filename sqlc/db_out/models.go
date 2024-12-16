// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Expense struct {
	ID                    pgtype.UUID
	Name                  string
	Value                 float64
	CategoryID            pgtype.UUID
	PaymentMethodID       pgtype.UUID
	Status                pgtype.Text
	CreditCardID          pgtype.UUID
	InstallmentCount      pgtype.Int4
	CurrentInstallment    pgtype.Int4
	InstallmentCountValue pgtype.Float8
	TransactionDate       pgtype.Timestamptz
}

type Status struct {
	Name string
}