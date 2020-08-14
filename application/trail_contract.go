package application

import (
	service "loan-calculator/domain/calculator/service"
	product "loan-calculator/domain/product/entity"
)

type PaymentTrailRequest struct {
	Amount       int64  `json:amount`
	Rate         int    `json:rate`
	RequestTime  int64  `json:requestTime`
	RepayDay     int    `json:repayDay`
	TermNum      int    `json:termNum`
	InterestType string `json:interestType`
}

type RepayPlan struct {
	Prin      int64 `json:prin`
	Interest  int64 `json:interest`
	RepayDate int64 `json:repayDate`
	TermNo    int   `json:termNo`
}

type PaymentTrailResponse struct {
	Prin       int64       `json:prin`
	Interest   int64       `json:interest`
	Rate       int         `json:rate`
	RepayDay   int         `json:repayDay`
	TermNum    int         `json:termNum`
	RepayPlans []RepayPlan `json:repayPlans`
}

func TrailPayment(request *PaymentTrailRequest) *PaymentTrailResponse {
	var serviceRequest *service.PaymentRequest = new(service.PaymentRequest)
	serviceRequest.Amount = request.Amount
	serviceRequest.Rate = request.Rate
	serviceRequest.RepayDay = request.RepayDay
	serviceRequest.RequestTime = request.RequestTime
	serviceRequest.TermNum = request.TermNum
	serviceRequest.InterestType = product.ConvertInterestType(request.InterestType)
	contract := service.GenerateContract(serviceRequest)
	subContract := contract.SubContract
	var response *PaymentTrailResponse = new(PaymentTrailResponse)
	response.Prin = subContract.Prin
	response.Rate = subContract.Rate
	response.RepayDay = contract.RepayDay
	response.TermNum = len(subContract.Terms)
	var sumInterest int64 = 0
	for _, item := range subContract.Terms {
		var repayPlan RepayPlan
		repayPlan.Prin = item.Prin
		repayPlan.Interest = item.Interest
		repayPlan.RepayDate = item.EndTime
		repayPlan.TermNo = item.TermNo
		sumInterest = sumInterest + item.Interest
	}
	response.Interest = sumInterest
	return response
}
