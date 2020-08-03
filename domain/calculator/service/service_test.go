package service

import (
	"testing"
)

func TestCreateContract(t *testing.T) {
	payment := new(PaymentRequest)
	payment.Amount = 50000
	payment.RepayDay = 2
	payment.TermNum = 6

	contract := generateContract(payment)
	t.Log(contract)
}
