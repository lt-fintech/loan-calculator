package service

import (
	"loan-calculator/domain/calculator/event"
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
	contract.Accrual(infra.GetTimestamp())
}

func TestLastDayAccrual(t *testing.T) {
	payment := new(PaymentRequest)
	payment.Amount = 5000
	payment.RepayDay = 1
	payment.TermNum = 6
	payment.Rate = 300

	contract := GenerateContract(payment)
	curTime := infra.GetTimestamp()
	for i := curTime; i < infra.GetNextRepayDate(curTime, payment.RepayDay, 20); i = i + 86400000 {
		// t.Logf("accountTime=%d", i)
		contract.Accrual(i)
	}
	t.Logf("%+v/n", contract.SubContract.Terms[0])
}

func TestFirstOvdAccrual(t *testing.T) {
	payment := new(PaymentRequest)
	payment.Amount = 50000
	payment.RepayDay = 1
	payment.TermNum = 6
	payment.Rate = 300
	payment.OvdPrinRate = 300 * 1.5
	payment.OvdIntRate = 300 * 1.5

	contract := GenerateContract(payment)
	t.Logf("%+v/n", contract.SubContract)
	curTime := infra.GetTimestamp()
	for i := curTime; i < infra.GetNextRepayDate(curTime, payment.RepayDay, 20)+86400000*2; i = i + 86400000 {
		// t.Logf("accountTime=%d", i)
		contract.Accrual(i)
	}
	t.Logf("%+v/n", contract.SubContract.Terms[0])
}

func TestRepayment(t *testing.T) {
	payment := new(PaymentRequest)
	payment.Amount = 50000
	payment.RepayDay = 1
	payment.TermNum = 6
	payment.Rate = 300
	payment.OvdPrinRate = 300 * 1.5
	payment.OvdIntRate = 300 * 1.5

	contract := GenerateContract(payment)
	curTime := infra.GetTimestamp()

	var repaytime int64
	for i := curTime; i < infra.GetNextRepayDate(curTime, payment.RepayDay, 20); i = i + 86400000 {
		// t.Logf("accountTime=%d", i)
		contract.Accrual(i)
		repaytime = i
	}
	repayment := new(event.RepaymentRequest)
	repayment.Amount = 50000
	repayment.RepayTime = repaytime
	contract.Repayment(repayment)
}
