package entity

import (
	"fmt"
	infra "loan-calculator/infrastructure"
	"loan-calculator/infrastructure/log"
)

type SubContract struct {
	Id          int64
	UserId      int64
	ProductId   int
	ProjectId   int
	SplitRuleId int
	OwnerId     int

	Prin       int64
	PaidPrin   int64
	UnpaidPrin int64

	TrailInterest  int64
	Interest       int64
	PaidInterest   int64
	UnpaidInterest int64

	OvdPrinPena       int64
	PaidOvdPrinPena   int64
	UnpaidOvdPrinPena int64

	OvdIntPena       int64
	PaidOvdIntPena   int64
	UnpaidOvdIntPena int64

	Rate        int
	OvdPrinRate int
	OvdIntRate  int

	Status      LoanStatus
	AccrualTime int64
	CreateTime  int64

	SubContracts []*SubContract
	Terms        []*Term
}

func (sub *SubContract) generateSubContract(contract *Contract, parent *SubContract) {

	accountTime := infra.GetTimestamp()
	if sub.Terms == nil {
		//calculate pmt per term
		var amountByTerm int64
		amountByTerm = infra.PMTTermRepayAmount(sub.Rate, contract.TermNum, sub.Prin)
		fmt.Printf("rate=%d\n", contract.Rate)
		fmt.Printf("amountByTerm=%d\n", amountByTerm)
		remainPrin := sub.Prin

		//determind first term day
		termStartDate := accountTime

		for i := 0; i < contract.TermNum; i++ {
			var interest int64
			repayDate := infra.GetNextRepayDate(termStartDate, contract.RepayDay, 20)
			// every term acrual day
			termDay := infra.GetBetweenDays(termStartDate, repayDate)
			// every term 30 day

			// var termDay int
			// if i == 0 {
			// 	termDay = infra.GetBetweenDays(termStartDate, repayDate)
			// } else {
			// 	termDay = 30
			// }
			interest = infra.PMTTermInterst(sub.Rate, termDay, remainPrin)

			var prin int64
			if i != contract.TermNum-1 {
				prin = amountByTerm - interest
			} else {
				prin = remainPrin
			}
			remainPrin = remainPrin - prin
			fmt.Printf("termDay=%d,prin=%d,interest=%d,repayDate=%d\n", termDay, prin, interest, repayDate)
			var term *Term
			term = new(Term)
			term.UserId = contract.UserId
			term.Interest = interest
			term.Prin = prin
			term.Status = NORMAL
			term.CreateTime = accountTime
			term.SubContractId = sub.Id
			term.TrailInterest = interest
			term.StartTime = termStartDate
			term.EndTime = repayDate
			// next term start date
			termStartDate = repayDate
			term.TermNo = i + 1
			sub.Terms = append(sub.Terms, term)

		}
	}
	sub.UserId = contract.UserId
	sub.CreateTime = accountTime
	sub.Prin = contract.Prin
	sub.Rate = contract.Rate
	sub.AccrualTime = infra.GetTimePlusDay(accountTime, -1)
}

func (sub *SubContract) accrual(accountTime int64) bool {
	betweenDay := infra.GetBetweenDays(sub.AccrualTime, accountTime)
	if betweenDay > 1 {
		log.Error.Println("can't accrual, last accrual day is ", sub.AccrualTime)
		return false
	}
	//calculat interest
	interest := infra.AccrualInterest(sub.Rate, sub.UnpaidPrin)
	// for term := range sub.Terms {

	// }
	log.Info.Println(interest)
	return true
}
