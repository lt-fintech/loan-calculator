package entity

type LoanStatus int

const (
	NORMAL LoanStatus = 1
	CLEAR  LoanStatus = 2
	OVD    LoanStatus = 3
)
