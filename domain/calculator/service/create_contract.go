package service

import (
	"loan-calculator/domain/calculator/entity"
	product "loan-calculator/domain/product/entity"
)

type PaymentRequest struct {
	Amount       int64
	RequestTime  int64
	RepayDay     int
	TermNum      int
	InterestType product.InterestType
}

func generateContract(paymentRequest *PaymentRequest) *entity.Contract {
	var contract *entity.Contract
	contract = new(entity.Contract)
	contract.Prin = paymentRequest.Amount
	contract.RepayDay = paymentRequest.RepayDay
	contract.TermNum = paymentRequest.TermNum
	contract.GenerateSubContract()
	return contract
}
