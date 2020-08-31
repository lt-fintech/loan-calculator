package service

import (
	infra "loan-calculator/infrastructure"
	"testing"
)

func TestCreateContract(t *testing.T) {
	payment := new(PaymentRequest)
	payment.Amount = 50000
	payment.RepayDay = 1
	payment.TermNum = 6
	payment.Rate = 300

	contract := GenerateContract(payment)
	t.Logf("%+v", contract)
	t.Logf("%+v", contract.SubContract)
}

func TestAccrual(t *testing.T) {
	payment := new(PaymentRequest)
	payment.Amount = 50000
	payment.RepayDay = 1
	payment.TermNum = 6
	payment.Rate = 300

	contract := GenerateContract(payment)
	contract.Accrual(infra.GetTimestamp())
}
