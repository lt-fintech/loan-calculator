package entity

type Contract struct {
	Id          int64
	UserId      int64
	ProductId   int
	ProjectId   int
	SplitRuleId int

	Prin        int64
	AccountTime int64
	RepayDay    int
	CreateTime  int64

	SubContract SubContract
}
