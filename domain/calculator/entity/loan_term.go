package entity

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

	TermNo     int
	StartTime  int64
	EndTime    int64
	Status     LoanStatus
	CreateTime int64
}

func (term *Term) CalculateBalance() {
	term.UnpaidPrin = term.Prin - term.PaidPrin
	term.UnpaidInterest = term.Interest - term.PaidInterest
	term.UnpaidOvdPrinPena = term.OvdPrinPena - term.PaidOvdPrinPena
	term.UnpaidOvdIntPena = term.OvdPrinPena - term.PaidOvdIntPena
}
