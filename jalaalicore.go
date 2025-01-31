package gojalaali

import "time"

type jTime struct {
	year  int
	month Month
	day   int
	hour  int
	min   int
	sec   int
	nsec  int
	loc   *time.Location
	wday  Weekday
}

func (driver *jTime) setTime(t time.Time) {
	var year, month, day int
	driver.nsec = t.Nanosecond()
	driver.sec = t.Second()
	driver.min = t.Minute()
	driver.hour = t.Hour()
	driver.loc = t.Location()
	driver.wday = JWeekday(t.Weekday())

	var jdn int
	gy, gmm, gd := t.Date()
	gm := int(gmm)

	if isAfterGregorianReform(gy, gm, gd) {
		jdn = convertGregorianPostReformToJDN(gy, gm, gd)
	} else {
		jdn = convertGregorianPreReformToJDN(gy, gm, gd)
	}

	year, month, day = convertJDNToShamsi(jdn)

	driver.year = year
	driver.month = Month(month)
	driver.day = day
}

func (driver *jTime) set(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) {
	// helpers
	norm := func(hi, lo, base int) (int, int) {
		if lo < 0 {
			n := (-lo-1)/base + 1
			hi -= n
			lo += n * base
		}
		if lo >= base {
			n := lo / base
			hi += n
			lo -= n * base
		}
		return hi, lo
	}
	normDay := func(hi, lo, base int) (int, int) {
		if lo < 1 {
			n := (-lo-1)/base + 1
			hi -= n
			lo += n * base
		}
		if lo > base {
			n := lo / base
			hi += n
			lo -= n * base
		}
		return hi, lo
	}

	// Get current location if not passed
	if loc == nil {
		loc = time.Local
	}

	// Normalize nsec, sec, min, hour, overflowing into day.
	sec, nsec = norm(sec, nsec, 1e9)
	min, sec = norm(min, sec, 60)
	hour, min = norm(hour, min, 60)
	day, hour = norm(day, hour, 24)

	// Normalize month, overflowing into year.
	m := int(month) - 1
	year, m = norm(year, m, 12)

	if m < 0 {
		m = 0
	} else if m > 11 {
		m = 11
	}

	if isLeap(year) {
		m, day = normDay(m, day, monthMeta[m][1])
	} else {
		m, day = normDay(m, day, monthMeta[m][0])
	}
	year, m = norm(year, m, 12)
	month = Month(m) + 1
	driver.year = year
	driver.month = month
	driver.day = day
	driver.hour = hour
	driver.min = min
	driver.sec = sec
	driver.nsec = nsec
	driver.loc = loc
	driver.resetWeekday()
	driver.normalize()
}

func (driver jTime) clone() *jTime {
	return &jTime{
		year:  driver.year,
		month: driver.month,
		day:   driver.day,
		hour:  driver.hour,
		min:   driver.min,
		sec:   driver.sec,
		nsec:  driver.nsec,
		loc:   driver.loc,
		wday:  driver.wday,
	}
}

func (driver *jTime) normalize() {
	driver.normalizeNano()
	driver.normalizeSec()
	driver.normalizeMin()
	driver.normalizeHour()
	driver.normalizeMonth()
	driver.normalizeDay()
}

func (driver *jTime) normalizeNano() {
	if driver.nsec < 0 {
		driver.nsec = 0
	} else if driver.nsec > 999999999 {
		driver.nsec = 999999999
	}
}

func (driver *jTime) normalizeSec() {
	if driver.sec < 0 {
		driver.sec = 0
	} else if driver.sec > 59 {
		driver.sec = 59
	}
}

func (driver *jTime) normalizeMin() {
	if driver.min < 0 {
		driver.min = 0
	} else if driver.min > 59 {
		driver.min = 59
	}
}

func (driver *jTime) normalizeHour() {
	if driver.hour < 0 {
		driver.hour = 0
	} else if driver.hour > 23 {
		driver.hour = 23
	}
}

func (driver *jTime) normalizeMonth() {
	if driver.month < Farvardin {
		driver.month = Farvardin
	} else if driver.month > Esfand {
		driver.month = Esfand
	}
}

func (driver *jTime) normalizeDay() {
	index := 0
	if driver.IsLeap() {
		index = 1
	}

	mIndex := driver.month - 1
	if mIndex < 0 {
		mIndex = 0
	} else if mIndex > 11 {
		mIndex = 11
	}

	days := monthMeta[mIndex][index]
	if driver.day < 1 {
		driver.day = 1
	} else if driver.day > days {
		driver.day = days
	}
}

func (driver *jTime) resetWeekday() {
	driver.wday = JWeekday(driver.Time().Weekday())
}
