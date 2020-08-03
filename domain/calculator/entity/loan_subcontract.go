package entity

import (
	"fmt"
	"math"
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

	Status     LoanStatus
	CreateTime int64

	SubContracts []*SubContract
	Terms        []*Term
}

func (sub *SubContract) generateSubContract(contract *Contract, parent *SubContract) {

	if sub.Terms == nil {
		//calculate pmt per term
		var amountByTerm int64
		p := math.Pow(1.0+float64(sub.Rate)/float64(1000000), float64(contract.TermNum))
		amountByTerm = int64((float64(sub.Prin*int64(sub.Rate)*30) * p / float64(1000000)) / (p - float64(1)))
		fmt.Println("amountByTerm=", amountByTerm)

	}
}
