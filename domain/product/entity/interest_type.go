package entity

type InterestType int

const (
	EQUAL_LOAN      InterestType = 1
	EQUAL_PRIN      InterestType = 2
	INTEREST_BEFORE InterestType = 3
)

func ConvertInterestType(str string) InterestType {
	switch str {
	case "EQUAL_LOAN":
		return EQUAL_LOAN
	case "EQUAL_PRIN":
		return EQUAL_PRIN
	case "INTEREST_BEFORE":
		return INTEREST_BEFORE
	default:
		panic("unrecognized interest type")
	}
}
