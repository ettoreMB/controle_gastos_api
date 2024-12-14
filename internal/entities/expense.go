package entities

import (
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	"time"
)

type Expense struct {
	ID                pkg_entities.UUID
	Name              string
	Value             float64
	CategoryId        CategoryId
	PaymentMethodId   PaymentMethodId
	Type              string
	Status            string
	TransactionDate   string
	InstallmentsCount int
	IsInstallmentPlan bool
}

// CreateExpenseCommand - Name, Value CategoryId

type CreateExpenseCommand struct {
	Name              string
	Value             float64
	CategoryId        CategoryId
	PaymentMethodId   PaymentMethodId
	IsInstallmentPlan bool
	InstallmentsCount int
}

func NewExpense(
	command *CreateExpenseCommand,
) *Expense {
	return &Expense{
		ID:                pkg_entities.NewUUID(),
		Name:              command.Name,
		Value:             command.Value,
		CategoryId:        command.CategoryId,
		PaymentMethodId:   command.PaymentMethodId,
		Type:              "expense",
		Status:            "pending",
		TransactionDate:   time.Now().Format("2006-01-02"),
		IsInstallmentPlan: false,
		InstallmentsCount: command.InstallmentsCount,
	}
}
