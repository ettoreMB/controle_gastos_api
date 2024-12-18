package mapper

import (
	"ettoreMB/controle_gastos/internal/entities"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	db "ettoreMB/controle_gastos/sqlc/db_out"

	"github.com/google/uuid"
)

func ExpensePgToDomainMapper(e db.Expense) (*entities.Expense, error) {

	installment := buildInstallment(e, e.CreditCardID)

	expense := entities.Expense{
		ID:              pkg_entities.UUID(e.ID),
		Name:            e.Name,
		Value:           float32(e.Value),
		CategoryId:      e.CategoryID,
		PaymentMethodId: e.PaymentMethodID,
		Type:            "",
		Status:          e.Status.String,
		TransactionDate: e.TransactionDate.String(),
		Installment:     installment,
	}
	return &expense, nil
}

func parseUUID(id *uuid.UUID) pkg_entities.UUID {
	return pkg_entities.UUID(*id)
}
func buildInstallment(e db.Expense, creditCardId pkg_entities.UUID) *entities.InstallMent {
	countPrice := float32(e.Value) / float32(e.InstallmentCount.Int32)

	return &entities.InstallMent{
		CreditCardId:       creditCardId,
		InstallmentsCount:  int(e.InstallmentCount.Int32),
		TotalValue:         float32(e.Value),
		CountPrice:         countPrice,
		CurrentInstallment: uint8(e.CurrentInstallment.Int32),
	}
}
