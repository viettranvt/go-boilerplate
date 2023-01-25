package date_util

import "time"

func ConvertMillisecondToTime(millisecond int64) time.Time {
	return time.Unix(0, millisecond*int64(time.Millisecond))
}

func CalculateDateDistanceByMillisecond(startDate time.Time, endDate time.Time) int64 {
	return endDate.Sub(startDate).Milliseconds()
}
