package infrastructure

import "time"

func GetTimestamp() int64 {
	return time.Now().UnixNano()
}

func GetNextRepayDate(repayDay int, minInterval int) int64 {

	now := time.Now()
	nowday := now.Day()
	if nowday < repayDay && repayDay-nowday > minInterval {
		return time.Date(now.Year(), now.Month(), repayDay, 0, 0, 0, 0, time.Local).UnixNano()
	} else {
		nextMonth := now.AddDate(0, 1, 0)
		return time.Date(nextMonth.Year(), nextMonth.Month(), repayDay, 0, 0, 0, 0, time.Local).UnixNano()
	}

}

func GetBetweenDays(from int64, to int64) int {
	t1 := time.UnixNano(from)
	t2 := time.UnixNano(to)
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)

	return int(t1.Sub(t2).Hours() / 24)
}
