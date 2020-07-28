package entity

type Contract struct {
	Id     int64
	UserId int64

	Prin       int64
	PaidPrin   int64
	UnpaidPrin int64

	Interest       int64
	PaidInterest   int64
	UnpaidInterest int64

	OvdPrinPena       int64
	PaidOvdPrinPena   int64
	UnpaidOvdPrinPena int64

	OvdIntPena       int64
	PaidOvdIntPena   int64
	UnpaidOvdIntPena int64

	CreateTime int64

	SubContracts []Contract
}
