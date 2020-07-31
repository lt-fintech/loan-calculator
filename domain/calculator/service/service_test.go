package service

import (
	"testing"
)

func TestCreateContract(t *testing.T) {
	payment := new(PaymentRequest)

	contract := generateContract(payment)
	t.Log(contract)
}
