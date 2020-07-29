package entity

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

	Status     LoanStatus
	CreateTime int64

	SubContracts []SubContract
}
