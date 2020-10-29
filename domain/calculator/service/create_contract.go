package service

import (
	"loan-calculator/domain/calculator/entity"
	product "loan-calculator/domain/product/entity"
)

type PaymentRequest struct {
	UserId       int64
	Amount       int64
	Rate         int
	OvdPrinRate  int
	OvdIntRate   int
	RequestTime  int64
	RepayDay     int
	TermNum      int
	InterestType product.InterestType
}

func GenerateContract(paymentRequest *PaymentRequest) *entity.Contract {
	var contract *entity.Contract
	contract = new(entity.Contract)
	contract.Prin = paymentRequest.Amount
	contract.Rate = paymentRequest.Rate
	contract.OvdPrinRate = paymentRequest.OvdPrinRate
	contract.OvdIntRate = paymentRequest.OvdIntRate
	contract.RepayDay = paymentRequest.RepayDay
	contract.TermNum = paymentRequest.TermNum
	contract.UserId = paymentRequest.UserId
	contract.GenerateSubContract()
	return contract
}
