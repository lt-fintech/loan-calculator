package entity

import (
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
		log.Info.Printf("rate=%d\n", contract.Rate)
		log.Info.Printf("amountByTerm=%d\n", amountByTerm)
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
			log.Info.Printf("termDay=%d,prin=%d,interest=%d,repayDate=%d\n", termDay, prin, interest, repayDate)
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
			term.CalculateBalance()
			sub.Terms = append(sub.Terms, term)

		}
	}
	sub.UserId = contract.UserId
	sub.CreateTime = accountTime
	sub.Prin = contract.Prin
	sub.Rate = contract.Rate
	sub.AccrualTime = infra.GetTimePlusDay(accountTime, -1)
	sub.caculateBalance()
}

func (sub *SubContract) caculateBalance() {
	sub.UnpaidPrin = sub.Prin - sub.PaidPrin
	sub.UnpaidInterest = sub.Interest - sub.PaidInterest
	sub.UnpaidOvdPrinPena = sub.OvdPrinPena - sub.PaidOvdPrinPena
	sub.UnpaidOvdIntPena = sub.OvdPrinPena - sub.PaidOvdIntPena
}

func (sub *SubContract) accrual(accountTime int64) bool {
	betweenDay := infra.GetBetweenDays(sub.AccrualTime, accountTime)
	log.Trace.Printf("accrual between day=%d\n", betweenDay)
	if betweenDay == 0 {
		log.Info.Println("today have accrualed")
		return true
	}
	if betweenDay > 1 {
		log.Error.Println("can't accrual, last accrual day is ", sub.AccrualTime)
		return false
	}
	//calculate interest
	log.Info.Printf("rate=%d,unpaidPrin=%d", sub.Rate, sub.UnpaidPrin)

	// calculate normal prin
	var normalUnpaidPrin int64 = 0
	var activeTerm *Term
	for _, term := range sub.Terms {
		normalUnpaidPrin = normalUnpaidPrin + term.GetUnpaidNormalPrin()
		if term.IsActive(accountTime) {
			activeTerm = term
		}
	}
	normalInterest := infra.AccrualInterest(sub.Rate, normalUnpaidPrin)
	log.Info.Printf("normalInterest=%d", normalInterest)
	if normalInterest > 0 && activeTerm != nil {
		panic("accrual interest but no active term")
	}
	if activeTerm != nil {
		activeTerm.AccrualInterest(normalInterest)
	}
	// calculate ovd prin by term
	for _, term := range sub.Terms {
		ovdUnpaidPrin := term.GetUnpaidOvdPrin()
		ovdPrinPena := infra.AccrualInterest(sub.OvdPrinRate, ovdUnpaidPrin)
		term.AccrualPrinPena(ovdPrinPena)
		ovdUnpaidInterest := term.GetUnpaidOvdInterest()
		ovdIntPena := infra.AccrualInterest(sub.OvdIntRate, ovdUnpaidInterest)
		term.AccrualIntPena(ovdIntPena)
	}
	sub.AccrualTime = accountTime
	return true
}
