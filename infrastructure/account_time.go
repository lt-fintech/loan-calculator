package infrastructure

import "time"

func GetTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}

func GetNextRepayDate(start int64, repayDay int, minInterval int) int64 {

	now := time.Unix(0, start*1e6)
	nowday := now.Day()
	if nowday < repayDay && repayDay-nowday > minInterval {
		return time.Date(now.Year(), now.Month(), repayDay, 0, 0, 0, 0, time.Local).UnixNano() / 1e6
	} else {
		nextMonth := now.AddDate(0, 1, 0)
		nextMonthRepayDay := time.Date(nextMonth.Year(), nextMonth.Month(), repayDay, 0, 0, 0, 0, time.Local).UnixNano() / 1e6
		if GetBetweenDays(start, nextMonthRepayDay) < minInterval {
			next2Month := now.AddDate(0, 2, 0)
			return time.Date(next2Month.Year(), next2Month.Month(), repayDay, 0, 0, 0, 0, time.Local).UnixNano() / 1e6
		} else {
			return nextMonthRepayDay
		}
	}

}

func GetBetweenDays(from int64, to int64) int {
	fromTime := time.Unix(0, from*1e6)
	t1 := time.Date(fromTime.Year(), fromTime.Month(), fromTime.Day(), 0, 0, 0, 0, time.Local)
	toTime := time.Unix(0, to*1e6)
	t2 := time.Date(toTime.Year(), toTime.Month(), toTime.Day(), 0, 0, 0, 0, time.Local)

	return int(t2.Sub(t1).Hours() / 24)
}

func GetZeroTimeFromTimestamp(tm int64) int64 {
	t := time.Unix(0, tm*1e6)
	newT := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return newT.UnixNano() / 1e6

}
