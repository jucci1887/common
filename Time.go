package common

import "time"

type times struct {
	standard  string
	timestamp int64
	dateTime  time.Time
}

const DateFormat = "2006-01-02"
const TimeFormat = "2006-01-02 15:04:05"

var Time = new(times)

func (t *times) getNowDateTime() {
	t.dateTime = time.Now()
}

func (t *times) getNowStandardTime() {
	t.standard = t.dateTime.Format(TimeFormat)
}

func (t *times) getNowTimeUnix() {
	toUnixTime, _ := time.Parse(TimeFormat, t.standard)
	t.timestamp = toUnixTime.Unix()
}

func (t *times) Now() *times {
	t.getNowDateTime()
	t.getNowStandardTime()
	t.getNowTimeUnix()

	return t
}

func (t *times) ToTimestamp() int64 {
	return t.timestamp
}

func (t *times) ToStandard() string {
	return t.standard
}

func (t *times) ToDateTime() time.Time {
	return t.dateTime
}

// 获取当天日期字符串
func (t *times) GetTodayDate() string {
	return t.dateTime.Format(DateFormat)
}
