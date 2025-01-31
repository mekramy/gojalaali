package gojalaali

import (
	"math"
	"time"
)

func (driver jTime) IsZero() bool {
	return driver == jTime{}
}

func (driver jTime) IsLeap() bool {
	return isLeap(driver.year)
}

func (driver jTime) Since(t2 Jalaali) time.Duration {
	return time.Duration(math.Abs(float64(t2.Unix()-driver.Unix()))) * time.Second
}

func (driver jTime) AmPm() AmPm {
	if driver.hour > 12 || (driver.hour == 12 && (driver.min > 0 || driver.sec > 0)) {
		return Pm
	}
	return Am
}

func (driver jTime) Zone() (string, int) {
	return driver.Time().Zone()
}

func (driver jTime) In(loc *time.Location) Jalaali {
	res := driver.clone()
	if loc != nil {
		res.loc = loc
	}
	res.resetWeekday()
	return res
}

func (driver jTime) Add(d time.Duration) Jalaali {
	return New(driver.Time().Add(d))
}

func (driver jTime) AddTime(hour, min, sec, nsec int) Jalaali {
	hours := time.Duration(hour) * time.Hour
	mins := time.Duration(min) * time.Minute
	secs := time.Duration(sec) * time.Second
	nanos := time.Duration(nsec) * time.Nanosecond
	return driver.Add(hours + mins + secs + nanos)
}

func (driver jTime) AddDate(year, month, day int) Jalaali {
	return Date(
		driver.year+year, driver.month+Month(month), driver.day+day,
		driver.hour, driver.min, driver.sec, driver.nsec, driver.loc,
	)
}

func (driver jTime) AddDatetime(year, month, day, hour, min, sec, nsec int) Jalaali {
	return Date(
		driver.year+year, driver.month+Month(month), driver.day+day,
		driver.hour+hour, driver.min+min, driver.sec+sec,
		driver.nsec+nsec, driver.loc,
	)
}

func (driver jTime) Yesterday() Jalaali {
	return driver.AddDate(0, 0, -1)
}

func (driver jTime) Tomorrow() Jalaali {
	return driver.AddDate(0, 0, 1)
}

func (driver jTime) BeginningOfDay() Jalaali {
	res := driver.clone()
	res.SetTime(0, 0, 0, 0)
	return res
}
func (driver jTime) EndOfDay() Jalaali {
	res := driver.clone()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (driver jTime) FirstWeekDay() Jalaali {
	if driver.wday == Shanbeh {
		return driver.clone()
	}
	return driver.AddDate(0, 0, int(Shanbeh-driver.wday))
}

func (driver jTime) LastWeekDay() Jalaali {
	if driver.wday == Jomeh {
		return driver.clone()
	}
	return driver.AddDate(0, 0, int(Jomeh-driver.wday))
}

func (driver jTime) BeginningOfWeek() Jalaali {
	res := driver.FirstWeekDay()
	res.SetTime(0, 0, 0, 0)
	return res
}

func (driver jTime) EndOfWeek() Jalaali {
	res := driver.LastWeekDay()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (driver jTime) FirstMonthDay() Jalaali {
	if driver.day == 1 {
		return driver.clone()
	}
	return Date(
		driver.year, driver.month, 1,
		driver.hour, driver.min, driver.sec,
		driver.nsec, driver.loc,
	)
}

func (driver jTime) LastMonthDay() Jalaali {
	dIndex := 0
	if driver.IsLeap() {
		dIndex = 1
	}

	mIndex := driver.month - 1
	if mIndex < 0 {
		mIndex = 0
	} else if mIndex > 11 {
		mIndex = 11
	}

	lastDay := monthMeta[mIndex][dIndex]
	if lastDay == driver.day {
		return driver.clone()
	}
	return Date(
		driver.year, driver.month, lastDay,
		driver.hour, driver.min, driver.sec,
		driver.nsec, driver.loc,
	)
}

func (driver jTime) BeginningOfMonth() Jalaali {
	res := driver.FirstMonthDay()
	res.SetTime(0, 0, 0, 0)
	return res
}

func (driver jTime) EndOfMonth() Jalaali {
	res := driver.LastMonthDay()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (driver jTime) FirstYearDay() Jalaali {
	if driver.month == Farvardin && driver.day == 1 {
		return driver.clone()
	}
	return Date(
		driver.year, Farvardin, 1,
		driver.hour, driver.min, driver.sec,
		driver.nsec, driver.loc,
	)
}

func (driver jTime) LastYearDay() Jalaali {
	dIndex := 0
	if driver.IsLeap() {
		dIndex = 1
	}
	lastDay := monthMeta[Esfand-1][dIndex]
	if driver.month == Esfand && driver.day == lastDay {
		return driver.clone()
	}
	return Date(
		driver.year, Esfand, lastDay,
		driver.hour, driver.min, driver.sec,
		driver.nsec, driver.loc,
	)
}

func (driver jTime) BeginningOfYear() Jalaali {
	res := driver.FirstYearDay()
	res.SetTime(0, 0, 0, 0)
	return res
}

func (driver jTime) EndOfYear() Jalaali {
	res := driver.LastYearDay()
	res.SetTime(23, 59, 59, 999999999)
	return res
}

func (driver *jTime) SetYear(year int) {
	driver.year = year
	driver.normalizeDay()
	driver.resetWeekday()
}

func (driver *jTime) SetMonth(month Month) {
	driver.month = month
	driver.normalizeMonth()
	driver.normalizeDay()
	driver.resetWeekday()
}

func (driver *jTime) SetDay(day int) {
	driver.day = day
	driver.normalizeDay()
	driver.resetWeekday()
}

func (driver *jTime) SetHour(hour int) {
	driver.hour = hour
	driver.normalizeHour()
}

func (driver *jTime) SetMinute(min int) {
	driver.min = min
	driver.normalizeMin()
}

func (driver *jTime) SetSecond(sec int) {
	driver.sec = sec
	driver.normalizeSec()
}

func (driver *jTime) SetNanosecond(nsec int) {
	driver.nsec = nsec
	driver.normalizeNano()
}

func (driver *jTime) SetTime(hour, min, sec, nsec int) {
	if hour >= 0 {
		driver.SetHour(hour)
	}
	if min >= 0 {
		driver.SetMinute(min)
	}
	if sec >= 0 {
		driver.SetSecond(sec)
	}
	if nsec >= 0 {
		driver.SetNanosecond(nsec)
	}
}

func (driver *jTime) SetDate(year, month, day int) {
	if year > 0 {
		driver.SetYear(year)
	}
	if month > 0 {
		driver.SetMonth(Month(month))
	}
	if day > 0 {
		driver.SetDay(day)
	}
}

func (driver *jTime) SetDateTime(year, month, day, hour, min, sec, nsec int) {
	driver.SetDate(year, month, day)
	driver.SetTime(hour, min, sec, nsec)
}

func (driver jTime) Year() int {
	return driver.year
}

func (driver jTime) YearDay() int {
	month := driver.month - 1
	if month < 0 {
		month = 0
	} else if month > 11 {
		month = 11
	}
	return monthMeta[month][2] + driver.day
}

func (driver jTime) YearRemainDays() int {
	days := 365
	if driver.IsLeap() {
		days = 366
	}
	return days - driver.YearDay()
}

func (driver jTime) Month() Month {
	return driver.month
}

func (driver jTime) Weekday() Weekday {
	return driver.wday
}

func (driver jTime) MonthWeek() int {
	return int(math.Ceil(float64(driver.day+int(driver.FirstMonthDay().Weekday())) / 7.0))
}

func (driver jTime) YearWeek() int {
	return int(math.Ceil(float64(driver.YearDay()+int(driver.FirstYearDay().Weekday())) / 7.0))
}

func (driver jTime) YearRemainWeeks() int {
	return 52 - driver.YearWeek()
}

func (driver jTime) Day() int {
	return driver.day
}

func (driver jTime) MonthRemainDays() int {
	dIndex := 0
	if driver.IsLeap() {
		dIndex = 1
	}

	mIndex := driver.month - 1
	if mIndex < 0 {
		mIndex = 0
	} else if mIndex > 11 {
		mIndex = 11
	}

	return monthMeta[mIndex][dIndex] - driver.day
}

func (driver jTime) Hour() int {
	return driver.hour
}

func (driver jTime) Hour12() int {
	if driver.hour > 12 {
		return driver.hour - 12
	} else {
		return driver.hour
	}
}

func (driver jTime) Minute() int {
	return driver.min
}

func (driver jTime) Second() int {
	return driver.sec
}

func (driver jTime) Nanosecond() int {
	return driver.nsec
}

func (driver jTime) DayTime() DayTime {
	return DayTime(driver.hour / 3)
}

func (driver jTime) Location() *time.Location {
	return driver.loc
}

func (driver jTime) Date() (int, Month, int) {
	return driver.year, driver.month, driver.day
}

func (driver jTime) Clock() (int, int, int) {
	return driver.hour, driver.min, driver.sec
}

func (driver jTime) Unix() int64 {
	return driver.Time().Unix()
}

func (driver jTime) UnixNano() int64 {
	return driver.Time().UnixNano()
}

func (driver jTime) Time() time.Time {
	// Handle empty date
	if driver.IsZero() {
		return time.Time{}
	}

	var year, month, day int

	// Convert the Shamsi to the corresponding Julian Day Number (JDN)
	jdn := convertShamsiToJDN(driver.year, int(driver.month), driver.day)

	// Convert the JDN to a Gregorian testDate
	if jdn > gregorianReformJulianDay {
		year, month, day = convertJDNToGregorianPostReform(jdn)
	} else {
		year, month, day = convertJDNToGregorianPreReform(jdn)
	}

	// Use the location stored in the Time struct, or default to the local time zone
	loc := driver.loc
	if loc == nil {
		loc = time.Local
	}

	// Return the corresponding time.Time object
	return time.Date(
		year, time.Month(month), day,
		driver.hour, driver.min, driver.sec,
		driver.nsec, loc,
	)
}

func (driver jTime) String() string {
	return driver.Format(time.RFC3339)
}
