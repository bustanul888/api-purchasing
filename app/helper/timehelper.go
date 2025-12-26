package helper

import (
	"time"
)

func UtcTime() time.Time {
	return time.Now().UTC()
}

func StringToDateTime(dateStr string, zone float64) (time.Time, error) {
	layout := "2006-01-02 15:04:05"
	localTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	utcTime := localTime.UTC()
	offset := time.Duration(zone * float64(time.Hour))
	return utcTime.Add(offset * -1), nil
	//
}

func StringToDate(dateStr string) (time.Time, error) {
	layout := "2006-01-02"
	localTime, err := time.Parse(layout, dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return localTime, nil
	//
}

func FormatDateToString(date time.Time) string {
	return date.Format("2006-01-02")
}

func TodayDateString() string {
	return time.Now().Format("2006-01-02")
}

func FormatDateTimeToStr(date time.Time) string {
	return date.Format("2006-01-02 15:04:00")
}

func GetTimeStartEnd() (time.Time, time.Time) {
	loc, _ := time.LoadLocation("Asia/Makassar")
	t := time.Now().In(loc)
	y, m, d := t.In(loc).Date()
	return time.Date(y, m, d, 0, 0, 0, 0, loc), time.Date(y, m, d, 23, 59, 59, 999999999, loc)
}

func ConvertUtcStringToUTcTime(time_ string) (time.Time, error) {
	layout := "2006-01-02T15:04:05Z07:00"
	res, err := time.Parse(layout, time_)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}

func NowUtcTime(zone float64) time.Time {
	offset := time.Duration(zone * float64(time.Hour))
	return time.Now().UTC().Add(offset).Truncate(time.Second)
}
