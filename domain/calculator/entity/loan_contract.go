package entity

type Contract struct {
	Id          int64
	UserId      int64
	ProductId   int
	ProjectId   int
	SplitRuleId int

	Prin        int64
	Rate        int
	OvdPrinRate int
	OvdIntRate  int
	AccountTime int64
	RepayDay    int
	TermNum     int
	CreateTime  int64

	SubContract *SubContract
}

func (contract *Contract) GenerateSubContract() {
	if contract.SubContract == nil {
		sub := new(SubContract)
		sub.UserId = contract.UserId
		sub.Prin = contract.Prin
		sub.Rate = contract.Rate
		sub.OvdPrinRate = contract.OvdPrinRate
		sub.OvdIntRate = contract.OvdIntRate
		sub.generateSubContract(contract, nil)
		contract.SubContract = sub

	}
}
