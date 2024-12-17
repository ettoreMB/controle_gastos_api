package mapper

import (
	"ettoreMB/controle_gastos/internal/entities"
	pkg_entities "ettoreMB/controle_gastos/pkg/entities"
	db "ettoreMB/controle_gastos/sqlc/db_out"
	"fmt"
)

func ExpensePgToDomainMapper(e db.Expense) (*entities.Expense, error) {
	fmt.Println(e.ID.Bytes)
	parsedId, err := parseUUID(e.ID.Bytes[:])
	if err != nil {
		return nil, err
	}

	parsedCategoryID, err := parseUUID(e.CategoryID.Bytes[:])
	if err != nil {
		return nil, err
	}

	parsedPaymentId, err := parseUUID(e.PaymentMethodID.Bytes[:])
	if err != nil {
		return nil, err
	}

	parsedCreditCardId, err := parseUUID(e.CreditCardID.Bytes[:])
	if err != nil {
		return nil, err
	}

	installment := buildInstallment(e, *parsedCreditCardId)

	expense := entities.Expense{
		ID:              *parsedId,
		Name:            e.Name,
		Value:           float32(e.Value),
		CategoryId:      entities.CategoryId(*parsedCategoryID),
		PaymentMethodId: entities.PaymentMethodId(*parsedPaymentId),
		Type:            "",
		Status:          e.Status.String,
		TransactionDate: e.TransactionDate.Time.GoString(),
		Installment:     installment,
	}
	return &expense, nil
}

func parseUUID(uuidBytes []byte) (*pkg_entities.UUID, error) {
	return pkg_entities.NewUUIDFromString(string(uuidBytes))
}

func buildInstallment(e db.Expense, creditCardId pkg_entities.UUID) *entities.InstallMent {
	countPrice := float32(e.Value) / float32(e.InstallmentCount.Int32)

	return &entities.InstallMent{
		CreditCardId:       entities.CreditcardId(creditCardId),
		InstallmentsCount:  int(e.InstallmentCount.Int32),
		TotalValue:         float32(e.Value),
		CountPrice:         countPrice,
		CurrentInstallment: uint8(e.CurrentInstallment.Int32),
	}
}
