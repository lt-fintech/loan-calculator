package infrastructure

import "math"

func Round(x float64) int64 {
	return int64(math.Floor(x + 0.5))
}

func PMTTermRepayAmount(rate int, termNum int, prin int64) int64 {
	var p float64 = math.Pow(1.0+float64(rate*30)/float64(1000000), float64(termNum))

	amountByTerm := Round((float64(prin*int64(rate)*30) * p / float64(1000000)) / (p - float64(1)))
	return amountByTerm
}

func PMTTermInterst(rate int, day int, remainPrin int64) int64 {
	return Round(float64(remainPrin*int64(rate)*int64(day)) / float64(1000000))
}
