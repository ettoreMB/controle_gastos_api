package entities

type InstallMent struct {
	CreditCardId       CreditcardId
	InstallmentsCount  int
	CurrentInstallment uint8
	TotalValue         float32
	CountPrice         float32
}

func NewStallment(creditcardId CreditcardId, counts int, value float32) *InstallMent {
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
