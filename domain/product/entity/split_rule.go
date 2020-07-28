package entity

type SplitRule struct {
	PrinPercentage  int
	IntRate         int
	OvdPrinPenaRate int
	OvdIntPenaRate  int
	Entity          int
	SubSplitRule    []SplitRule
}
