package entities

import pkg_entities "ettoreMB/controle_gastos/pkg/entities"

type InstallMent struct {
	CreditCardId       pkg_entities.UUID
	InstallmentsCount  int
	CurrentInstallment uint8
	TotalValue         float32
	CountPrice         float32
}

func NewStallment(creditcardId pkg_entities.UUID, counts int, value float32) *InstallMent {
	st := &InstallMent{
		CreditCardId:       creditcardId,
		InstallmentsCount:  counts,
		CurrentInstallment: 1,
		TotalValue:         value,
		CountPrice:         0,
	}

	if counts > 1 {
		st.CountPrice = value / float32(counts)
	} else {
		st.CountPrice = value
	}

	return st
}
