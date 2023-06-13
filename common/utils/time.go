package util

import (
	"database/sql/driver"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

const TimeLayout = "2006-01-02 15:04:05"
const MinuteLayout = "2006-01-02 15:04"
const DayLayout = "2006-01-02"
const MonthLayout = "2006-01"
const UtcTimeLayout = "2006-01-02T15:04:05+08:00"
const CNMonthTimeLayout = "2006年01月"
const CNDayTimeLayout = "2006年01月02日"
const CNSimpleDayTimeLayout = "01月02日"
const YMDLayout = "20060102"
const YearLayout = "2006年"
const DayDotLayout = "2006.01.02"
const AllNumLayout = "20060102150405"
const ZeroTimeout = "0001-01-01 00:00:00"

type JsonTime struct {
	time.Time
}

func (t JsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", t.Time.Format(TimeLayout))
	return []byte(stamp), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) error {
	var err error
	str, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	if len(str) == 0 {
		return nil
	}
	i, err := time.ParseInLocation(TimeLayout, str, time.Local)
	t.Time = i
	return err
}

func (t JsonTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

func (t JsonTime) Empty() bool {
	return t == JsonTime{} || t.Time.Unix() == 0
}

func (t *JsonTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JsonTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

type JsonTimeSlice []JsonTime

func (s JsonTimeSlice) Len() int { return len(s) }

func (s JsonTimeSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s JsonTimeSlice) Less(i, j int) bool { return s[i].Before(s[j].Time) }

func CheckContinuity(l JsonTimeSlice) bool {
	sort.Sort(l)
	for k := range l {
		if k > 0 {
			if l[k].Day()-l[k-1].Day() > 1 {
				return false
			}
		}
	}
	return true
}

func GetCNMothDay(ti time.Time) string {
	return ti.Format("06年01月")
}

// GetCurrentWeekMondayAndSunday 获取当前周的周一和周日
func GetCurrentWeekMondayAndSunday() (monday time.Time, sunday time.Time) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekMonday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekSunday := weekMonday.AddDate(0, 0, 6)
	return weekMonday, weekSunday
}

// GetWeekMondayAndSundayByTime 根据给定时间获取本周的周一和周日
func GetWeekMondayAndSundayByTime(baseTime time.Time) (monday time.Time, sunday time.Time) {
	now := baseTime
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekMonday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekSunday := weekMonday.AddDate(0, 0, 6).Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)

	return weekMonday, weekSunday
}

func GetDiffDays(startTime, endTime time.Time) int {
	if startTime.After(endTime) {
		return -1
	}
	return int(math.Ceil(float64(endTime.Unix()-startTime.Unix()) / (3600 * 24)))
}

func GetBuTimeStr(from, to time.Time) string {
	seconds := int64(to.Sub(from).Seconds())
	if seconds < 60 {
		return "刚刚"
	}
	if seconds < 3600 {
		return strconv.FormatInt(seconds/60, 10) + "分钟前"
	}
	if seconds < 86400 {
		return strconv.FormatInt(seconds/3600, 10) + "小时前"
	}
	if seconds < 86400*7 {
		return strconv.FormatInt(seconds/86400, 10) + "天前"
	}
	if seconds < 86400*30 {
		return strconv.FormatInt(seconds/(86400*7), 10) + "周前"
	}
	if seconds < 86400*360 {
		return strconv.FormatInt(seconds/(86400*30), 10) + "个月前"
	}
	return strconv.FormatInt(seconds/(86400*360), 10) + "年前"
}

func GetManyHourTimeStr(now time.Time, hour int) string {
	h := fmt.Sprintf("%dh", hour)
	dh, _ := time.ParseDuration(h)
	return now.Add(dh).Format(TimeLayout)
}
