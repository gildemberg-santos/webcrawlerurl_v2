package googleshopping_test

import (
	"testing"

	googleshopping "github.com/gildemberg-santos/webcrawlerurl_v2/util/google_shopping"
	"github.com/stretchr/testify/assert"
)

func TestInstallment_Normalize(t *testing.T) {
	installment := googleshopping.NewInstallment("\n 1 \n", "\n 1 \n", "\n 1 \n", "\n 1 \n")

	assert.Equal(t, "1", installment.Amount.Value)
	assert.Equal(t, "1", installment.Downpayment.Value)
	assert.Equal(t, "1", installment.CreditType.Value)
	assert.Equal(t, "1", installment.Months.Value)
}
