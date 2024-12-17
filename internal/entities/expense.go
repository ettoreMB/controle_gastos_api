package entities

import (
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	"time"
)

type Expense struct {
	ID              pkg_entities.UUID
	Name            string
	Value           float32
	CategoryId      pkg_entities.UUID
	PaymentMethodId pkg_entities.UUID
	Type            string
	Status          string
	TransactionDate string
	Installment     *InstallMent
}

// CreateExpenseCommand - Name, Value CategoryId

type CreateExpenseCommand struct {
	Name            string
	Value           float32
	CategoryId      pkg_entities.UUID
	PaymentMethodId pkg_entities.UUID
	Installment     *InstallMent
}

func NewExpense(
	command *CreateExpenseCommand,
) *Expense {
	expense := &Expense{
		ID:              pkg_entities.NewUUID(),
		Name:            command.Name,
		Value:           command.Value,
		CategoryId:      command.CategoryId,
		PaymentMethodId: command.PaymentMethodId,
		Type:            "expense",
		Status:          "pending",
		TransactionDate: time.Now().Format("2006-01-02"),
		Installment:     command.Installment,
	}

	return expense
}
