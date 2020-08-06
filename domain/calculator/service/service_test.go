package service

import (
	"testing"
)

func TestCreateContract(t *testing.T) {
	payment := new(PaymentRequest)
	payment.Amount = 50000
	payment.RepayDay = 2
	payment.TermNum = 6
	payment.Rate = 300

	contract := generateContract(payment)
	t.Logf("%+v", contract)
	t.Logf("%+v", contract.SubContract)
}
