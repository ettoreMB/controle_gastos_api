package entities

import (
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	"time"
)

type Expense struct {
	ID              pkg_entities.UUID
	Name            string
	Value           float64
	CategoryId      pkg_entities.UUID
	PaymentMethodId PaymentMethodId
	Type            string
	Status          string
	TransactionDate string
	Installments    int
}

func NewExpense(
	name string,
	value float64,
	categoryId pkg_entities.UUID,
	paymentMethodId PaymentMethodId,
	installments int,
) *Expense {
	return &Expense{
		ID:              pkg_entities.NewUUID(),
		Name:            name,
		Value:           value,
		CategoryId:      categoryId,
		PaymentMethodId: paymentMethodId,
		Type:            "expense",
		Status:          "pending",
		TransactionDate: time.Now().Format("2006-01-02"),
		Installments:    installments,
	}
}
