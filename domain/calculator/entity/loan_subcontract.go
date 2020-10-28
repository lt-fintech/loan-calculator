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
			term.Interest = 0
			term.Prin = prin
			term.Status = NORMAL
			term.CreateTime = accountTime
			term.SubContractId = sub.Id
			term.TrailInterest = interest
			term.StartTime = termStartDate
			term.AccrualStartTime = termStartDate
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
	sub.OvdPrinRate = contract.OvdPrinRate
	sub.OvdIntRate = contract.OvdIntRate
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
	if betweenDay <= 0 {
		log.Info.Println("today have accrualed")
		return true
	}
	if betweenDay > 1 {
		log.Error.Println("can't accrual, last accrual day is ", sub.AccrualTime)
		return false
	}
	//calculate if ovd
	for _, term := range sub.Terms {
		if term.IsOvd(accountTime) {
			term.Status = OVD
			sub.Status = OVD
		}
	}

	//calculate interest

	// calculate normal prin
	var normalUnpaidPrin int64 = 0
	var activeTerm *Term
	for _, term := range sub.Terms {
		normalUnpaidPrin = normalUnpaidPrin + term.GetUnpaidNormalPrin()
		if term.IsActive(accountTime) {
			activeTerm = term
		}
	}
	if activeTerm != nil {
		accrualedDay := infra.GetBetweenDays(activeTerm.AccrualStartTime, accountTime) + 1
		normalInterest := infra.AccrualInterest(sub.Rate, normalUnpaidPrin, accrualedDay, activeTerm.Interest)
		log.Info.Printf("termNo=%d,accrualedDay=%d,accountTime=%d,normalInterest=%d", activeTerm.TermNo, accrualedDay, accountTime, normalInterest)
		activeTerm.AccrualInterest(normalInterest)
	}
	// calculate ovd prin by term
	for _, term := range sub.Terms {
		accrualedDay := infra.GetBetweenDays(term.EndTime, accountTime) + 1
		ovdUnpaidPrin := term.GetUnpaidOvdPrin()
		log.Info.Printf("rate=%d,unpaidprin=%d", sub.OvdPrinRate, ovdUnpaidPrin)
		ovdPrinPena := infra.AccrualOvdInterest(sub.OvdPrinRate, ovdUnpaidPrin)
		term.AccrualPrinPena(ovdPrinPena)
		ovdUnpaidInterest := term.GetUnpaidOvdInterest()
		ovdIntPena := infra.AccrualOvdInterest(sub.OvdIntRate, ovdUnpaidInterest)
		term.AccrualIntPena(ovdIntPena)
		if ovdPrinPena > 0 || ovdIntPena > 0 {
			log.Info.Printf("accrualedDay=%d,accountTime=%d,ovdPrinPena=%d,ovdIntPena=%d", accrualedDay, accountTime, ovdPrinPena, ovdIntPena)
		}
	}
	sub.AccrualTime = accountTime
	sub.caculateBalance()
	return true
}

func (sub *SubContract) Repayment(repayment *RepaymentRequest) bool {
	return false

}
