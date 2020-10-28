package service

import (
	"loan-calculator/domain/calculator/entity"
)

type RepaymentRequest struct {
	UserId    int64
	Amount    int64
	RepayTime int64
}

func (contract *entity.Contract) Repayment(repayment *RepaymentRequest) bool {
	return contract.Repayment(repayment)
}
