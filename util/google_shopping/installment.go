package googleshopping

type Installment struct {
	Months      Chardata `xml:"months"`
	Amount      Chardata `xml:"amount"`
	Downpayment Chardata `xml:"downpayment"`
	CreditType  Chardata `xml:"credit_type"`
}

func NewInstallment(months, amount, downpayment, creditType string) *Installment {
	i := Installment{
		Months:      *NewChardata(months),
		Amount:      *NewChardata(amount),
		Downpayment: *NewChardata(downpayment),
		CreditType:  *NewChardata(creditType),
	}
	// i.Normalize()
	return &i
}

func (i *Installment) Normalize() *Installment {
	return NewInstallment(i.Months.Value, i.Amount.Value, i.Downpayment.Value, i.CreditType.Value)
}
