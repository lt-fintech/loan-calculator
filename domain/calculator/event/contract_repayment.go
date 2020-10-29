package event

type RepaymentRequest struct {
	UserId    int64
	Amount    int64
	RepayTime int64
}
