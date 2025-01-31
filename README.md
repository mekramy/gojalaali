# Go Jalaali

`gojalaali` is a Go package that provides implementation of Jalaali (Persian) calendar for standard go `time` interface. It supports standard Go **time** package layout for formatting and parsing jalaali dates and times, as well as converting between Jalaali and Gregorian dates.

## Installation

To install the package, use the following command:

```sh
go get github.com/mekramy/gojalaali
```

## Usage

Here is an example of how to use the `gojalaali` package:

```go
package main

import (
    "fmt"
    "time"

    "github.com/mekramy/gojalaali"
)

func main() {
    // Create a new Jalaali date from the current time
    j := gojalaali.Now()

    // Print the Jalaali date
    fmt.Println("Current Jalaali date:", j)

    // Convert Jalaali date to Gregorian date
    gregorian := j.Time()
    fmt.Println("Gregorian date:", gregorian)

    // Parse a Jalaali date from a string
    layout := "2006-01-02 15:04:05"
    datetime := "1400-07-01 12:30:45"
    parsedJalaali, err := gojalaali.Parse(layout, datetime)
    if err != nil {
        fmt.Println("Error parsing Jalaali date:", err)
        return
    }
    fmt.Println("Parsed Jalaali date:", parsedJalaali)
}
```

## Constructors

### `Parse(layout, datetime string) (Jalaali, error)`

Parses a Jalaali date from a string according to the specified layout. Returns a Jalaali instance and an error if the parsing fails.

**Example:**

```go
layout := "2006-01-02 15:04:05"
datetime := "1400-07-01 12:30:45"
parsedJalaali, err := gojalaali.Parse(layout, datetime)
if err != nil {
    fmt.Println("Error parsing Jalaali date:", err)
    return
}
fmt.Println("Parsed Jalaali date:", parsedJalaali)
```

### `New(t time.Time) Jalaali`

Creates a new Jalaali instance from a Go `time.Time` object. If the year is less than 1097, it returns a zero time instance.

**Example:**

```go
t := time.Now()
j := gojalaali.New(t)
fmt.Println("Jalaali date:", j)
```

### `Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) Jalaali`

Creates a new Jalaali instance from the specified Jalaali date and time components.

**Example:**

```go
j := gojalaali.Date(1400, gojalaali.Mehr, 1, 12, 30, 45, 0, gojalaali.TehranTz())
fmt.Println("Jalaali date:", j)
```

### `Unix(sec, nsec int64) Jalaali`

Creates a new Jalaali instance from a Unix timestamp.

**Example:**

```go
sec := time.Now().Unix()
nsec := int64(0)
j := gojalaali.Unix(sec, nsec)
fmt.Println("Jalaali date:", j)
```

### `Now() Jalaali`

Creates a new Jalaali instance from the current time.

**Example:**

```go
j := gojalaali.Now()
fmt.Println("Current Jalaali date:", j)
```

### `TehranTz() *time.Location`

Returns the time zone for Tehran.

**Example:**

```go
tehranTz := gojalaali.TehranTz()
fmt.Println("Tehran time zone:", tehranTz)
```

### `KabulTz() *time.Location`

Returns the time zone for Kabul.

**Example:**

```go
kabulTz := gojalaali.KabulTz()
fmt.Println("Kabul time zone:", kabulTz)
```

## API Documentation

### `IsZero() bool`

Returns true if the Jalaali date is a zero time instance.

### `IsLeap() bool`

Returns true if the year of the Jalaali date is a leap year.

### `Since(t2 Jalaali) time.Duration`

Returns the number of seconds between the current Jalaali date and another Jalaali date `t2`.

### `AmPm() AmPm`

Returns the 12-hour marker (AM/PM) of the Jalaali date.

### `Zone() (string, int)`

Returns the time zone name and its offset in seconds east of UTC for the Jalaali date.

### `In(loc *time.Location) Jalaali`

Sets the location of the Jalaali date and returns a new instance. If `loc` is nil, it returns the same instance.

### `Add(d time.Duration) Jalaali`

Adds a duration to the Jalaali date and returns a new instance.

### `AddTime(hour, min, sec, nsec int) Jalaali`

Adds the specified hours, minutes, seconds, and nanoseconds to the Jalaali date and returns a new instance.

### `AddDate(year, month, day int) Jalaali`

Adds the specified years, months, and days to the Jalaali date and returns a new instance.

### `AddDatetime(year, month, day, hour, min, sec, nsec int) Jalaali`

Adds the specified years, months, days, hours, minutes, seconds, and nanoseconds to the Jalaali date and returns a new instance.

### `Yesterday() Jalaali`

Returns a new instance of the Jalaali date representing the day before the current instance.

### `Tomorrow() Jalaali`

Returns a new instance of the Jalaali date representing the day after the current instance.

### `BeginningOfDay() Jalaali`

Returns a new instance of the Jalaali date representing the 00:00:00.000000000 time of today.

### `EndOfDay() Jalaali`

Returns a new instance of the Jalaali date representing the 23:59:59.999999999 time of today.

### `FirstWeekDay() Jalaali`

Returns a new instance of the Jalaali date representing the first day of the week of the current instance.

### `LastWeekDay() Jalaali`

Returns a new instance of the Jalaali date representing the last day of the week of the current instance.

### `BeginningOfWeek() Jalaali`

Returns a new instance of the Jalaali date representing the first day of the week of the current instance, with the time set to 00:00:00.000000000.

### `EndOfWeek() Jalaali`

Returns a new instance of the Jalaali date representing the last day of the week of the current instance, with the time set to 23:59:59.999999999.

### `FirstMonthDay() Jalaali`

Returns a new instance of the Jalaali date representing the first day of the month of the current instance.

### `LastMonthDay() Jalaali`

Returns a new instance of the Jalaali date representing the last day of the month of the current instance.

### `BeginningOfMonth() Jalaali`

Returns a new instance of the Jalaali date representing the first day of the month of the current instance, with the time set to 00:00:00.000000000.

### `EndOfMonth() Jalaali`

Returns a new instance of the Jalaali date representing the last day of the month of the current instance, with the time set to 23:59:59.999999999.

### `FirstYearDay() Jalaali`

Returns a new instance of the Jalaali date representing the first day of the year of the current instance.

### `LastYearDay() Jalaali`

Returns a new instance of the Jalaali date representing the last day of the year of the current instance.

### `BeginningOfYear() Jalaali`

Returns a new instance of the Jalaali date representing the first day of the year of the current instance, with the time set to 00:00:00.000000000.

### `EndOfYear() Jalaali`

Returns a new instance of the Jalaali date representing the last day of the year of the current instance, with the time set to 23:59:59.999999999.

### `SetYear(year int)`

Sets the year of the Jalaali date.

### `SetMonth(month Month)`

Sets the month of the Jalaali date.

### `SetDay(day int)`

Sets the day of the Jalaali date.

### `SetHour(hour int)`

Sets the hour of the Jalaali time.

### `SetMinute(min int)`

Sets the minute of the Jalaali time.

### `SetSecond(sec int)`

Sets the second of the Jalaali time.

### `SetNanosecond(nsec int)`

Sets the nanosecond of the Jalaali time.

### `SetTime(hour, min, sec, nsec int)`

Sets the hour, minute, second, and nanosecond of the Jalaali time. Pass -1 to ignore a parameter.

### `SetDate(year, month, day int)`

Sets the year, month, and day of the Jalaali date. Pass -1 to ignore a parameter.

### `SetDateTime(year, month, day, hour, min, sec, nsec int)`

Sets the year, month, day, hour, minute, second, and nanosecond of the Jalaali date and time. Pass -1 to ignore a parameter.

### `Year() int`

Returns the year of the Jalaali date.

### `YearDay() int`

Returns the day of the year of the Jalaali date.

### `YearRemainDays() int`

Returns the number of remaining days in the year of the Jalaali date.

### `Month() Month`

Returns the month of the Jalaali date in the range [1, 12].

### `Weekday() Weekday`

Returns the weekday of the Jalaali date.

### `MonthWeek() int`

Returns the week of the month of the Jalaali date.

### `YearWeek() int`

Returns the week of the year of the Jalaali date.

### `YearRemainWeeks() int`

Returns the number of remaining weeks in the year of the Jalaali date.

### `Day() int`

Returns the day of the month of the Jalaali date.

### `MonthRemainDays() int`

Returns the number of remaining days in the month of the Jalaali date.

### `Hour() int`

Returns the hour of the Jalaali time in the range [0, 23].

### `Hour12() int`

Returns the hour of the Jalaali time in the range [0, 11].

### `Minute() int`

Returns the minute of the Jalaali time in the range [0, 59].

### `Second() int`

Returns the second of the Jalaali time in the range [0, 59].

### `Nanosecond() int`

Returns the nanosecond of the Jalaali time in the range [0, 999999999].

### `DayTime() DayTime`

Returns the part of the day for the Jalaali time.

### `Location() *time.Location`

Returns a pointer to the time location of the Jalaali date.

### `Date() (int, Month, int)`

Returns the year, month, and day of the Jalaali date.

### `Clock() (int, int, int)`

Returns the hour, minute, and second of the Jalaali time.

### `Unix() int64`

Returns the number of seconds since January 1, 1970 UTC.

### `UnixNano() int64`

Returns the number of nanoseconds since January 1, 1970 UTC.

### `Time() time.Time`

Converts the Jalaali date to a Gregorian date and returns it as a Go `time.Time` object.

### `String() string`

Returns the Jalaali date in RFC3339 format.

### `Format(layout string) string`

Formats the Jalaali date according to the specified layout.

| Layout           | Description                              | Example            |
| ---------------- | ---------------------------------------- | ------------------ |
| **Year**         |                                          |                    |
| 2006             | Four-digit year                          | "1403"             |
| 06               | Two-digit year                           | "03"               |
| **Month**        |                                          |                    |
| January          | Full month name                          | "اسفند"            |
| Jan              | Three-letter abbreviation of the month   | "اسف"              |
| 01               | Two-digit month with a leading 0         | "07"               |
| 1                | One-digit month                          | "7"                |
| **Day**          |                                          |                    |
| 02               | Two-digit month day with a leading 0     | "08"               |
| \_2              | Two-digit month day with a leading space | " 9"               |
| 2                | One-digit month day                      | "3"                |
| **Weekday**      |                                          |                    |
| Monday           | Full weekday name                        | "شنبه"             |
| Mon              | Abbreviation of the weekday              | "ش"                |
| **Hour**         |                                          |                    |
| 15               | Two-digit 24 hour format                 | "15"               |
| 03               | Two-digit 12 hour format                 | "03"               |
| 3                | One-digit 12 hour format                 | "9"                |
| **Minute**       |                                          |                    |
| 04               | Two-digit minute with leading 0          | "03"               |
| 4                | One-digit minute                         | "3"                |
| **Second**       |                                          |                    |
| 05               | Two-digit second with leading 0          | "09"               |
| 5                | One-digit second                         | "2"                |
| **Milliseconds** |                                          |                    |
| .000             | Millisecond                              | ".120"             |
| .000000          | Microsecond                              | ".123400"          |
| .000000000       | Nanosecond                               | ".123456000"       |
| .999             | Trailing zeros removed millisecond       | ".12"              |
| .999999          | Trailing zeros removed microsecond       | ".1234"            |
| .999999999       | Trailing zeros removed nanosecond        | ".123456"          |
| **Daytime**      |                                          |                    |
| Morning          | Day time                                 | "صبح"              |
| PM               | Full 12-Hour marker                      | "قبل از ظهر"       |
| pm               | Short 12-Hour marker                     | "ق.ظ"              |
| **Timezone**     |                                          |                    |
| MST              | Abbreviation of the time zone            | "UTC"              |
| Z070000          | Zone offset Hour, Minute and second      | "Z" or "+033000"   |
| Z0700            | Zone offset Hour and Minute              | "Z" or "+0330"     |
| Z07:00:00        | Zone offset Hour, Minute and second      | "Z" or "+03:30:00" |
| Z07:00           | Zone offset Hour and Minute              | "Z" or "+03:30"    |
| Z07              | Zone offset Hour                         | "Z" or "+03"       |
| -070000          | Zone offset Hour, Minute and second      | "+033000"          |
| -0700            | Zone offset Hour and Minute              | "+0330"            |
| -07:00:00        | Zone offset Hour, Minute and second      | "+03:30:00"        |
| -07:00           | Zone offset Hour and Minute              | "+03:30"           |
| -07              | Zone offset Hour                         | "+03"              |

## License

This package jalaali conversion inspired from `github.com/yaa110/go-persian-calendar` library.
