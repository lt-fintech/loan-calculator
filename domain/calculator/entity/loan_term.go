package entity

import (
	infra "loan-calculator/infrastructure"
)

type Term struct {
	Id            int64
	UserId        int64
	SubContractId int64

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

	TermNo           int
	StartTime        int64
	AccrualStartTime int64
	EndTime          int64
	Status           LoanStatus
	CreateTime       int64
}

func (term *Term) IsActive(accountTime int64) bool {
	// log.Info.Printf("start=%d end=%d accounttime=%d %t", term.StartTime, term.EndTime, accountTime, accountTime >= term.StartTime && accountTime < term.EndTime)
	if accountTime >= term.StartTime && accountTime < term.EndTime {
		return true
	}
	return false
}
func (term *Term) IsLastDay(accountTime int64) bool {
	if infra.GetBetweenDays(accountTime, term.EndTime) == 1 {
		return true
	}
	return false
}
func (term *Term) IsRepayDay(accountTime int64) bool {
	if infra.GetBetweenDays(term.EndTime, accountTime) == 0 {
		return true
	}
	return false
}
func (term *Term) IsOvd(accountTime int64) bool {
	if term.Status == OVD || (term.Status == NORMAL && infra.GetBetweenDays(term.EndTime, accountTime) > 0) {
		return true
	}
	return false
}
func (term *Term) CalculateBalance() {
	term.UnpaidPrin = term.Prin - term.PaidPrin
	term.UnpaidInterest = term.Interest - term.PaidInterest
	term.UnpaidOvdPrinPena = term.OvdPrinPena - term.PaidOvdPrinPena
	term.UnpaidOvdIntPena = term.OvdIntPena - term.PaidOvdIntPena
}

func (term *Term) GetUnpaidNormalPrin() int64 {
	if term.Status == NORMAL {
		return term.UnpaidPrin
	}
	return 0
}
func (term *Term) GetUnpaidOvdPrin() int64 {
	if term.Status == OVD {
		return term.UnpaidPrin
	}
	return 0
}
func (term *Term) GetUnpaidOvdInterest() int64 {
	if term.Status == OVD {
		return term.UnpaidInterest
	}
	return 0
}

func (term *Term) AccrualInterest(interest int64) {
	term.Interest = term.Interest + interest
	term.CalculateBalance()
}
func (term *Term) AccrualPrinPena(ovdPrinPena int64) {
	term.OvdPrinPena = term.OvdPrinPena + ovdPrinPena
	term.CalculateBalance()
}
func (term *Term) AccrualIntPena(ovdIntPena int64) {
	term.OvdIntPena = term.OvdIntPena + ovdIntPena
	term.CalculateBalance()
}
