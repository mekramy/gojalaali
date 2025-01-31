package gojalaali_test

import (
	"testing"

	"github.com/mekramy/gojalaali"
)

func TestDateFormat(t *testing.T) {
	// Year
	// 2006				Four-digit year								"1403"
	// 06				Two-digit year 								"03"
	//
	// Month
	// January			Full month name								"اسفند"
	// Jan				Three-letter abbreviation of the month		"اسف"
	// 01				Two-digit month with a leading 0			"07"
	// 1				One-digit month								"7"
	//
	// Day
	// 02				Two-digit month day with a leading 0		"08"
	// _2				Two-digit month day with a leading space	" 9"
	// 2				One-digit month day							"3"
	//
	// Weekday
	// Monday			Full weekday name							"شنبه"
	// Mon				abbreviation of the weekday					"ش"
	tests := []struct {
		layout   string
		expected string
	}{
		// Year
		{"2006", "1403"},
		{"06", "03"},
		// Month
		{"January", "شهریور"},
		{"Jan", "شهر"},
		{"01", "06"},
		{"1", "6"},
		// Day
		{"02", "03"},
		{"_2", " 3"},
		{"2", "3"},
		// Weekday
		{"Monday", "شنبه"},
		{"Mon", "ش"},
	}

	date := gojalaali.Date(1403, gojalaali.Shahrivar, 3, 0, 0, 0, 0, nil)
	for _, test := range tests {
		formatted := date.Format(test.layout)
		if formatted != test.expected {
			t.Errorf("fail %s, expected %s, got %s", test.layout, test.expected, formatted)
		}
	}
}

func TestTimeFormat(t *testing.T) {
	//
	// Hour
	// 15				Two-digit 24 hour format					"15"
	// 03				Two-digit 12 hour format					"03"
	// 3				One-digit 12 hour format					"9"
	//
	// Minute
	// 04				Two-digit minute with leading 0				"03"
	// 4				One-digit minute							"3"
	//
	// Second
	// 05				Two-digit second with leading 0				"09"
	// 5				One-digit second							"2"
	//
	// Milliseconds
	// .000				Millisecond									".120"
	// .000000			Microsecond									".123400"
	// .000000000		Nanosecond									".123456000"
	// .999				Trailing zeros removed millisecond			".12"
	// .999999			Trailing zeros removed microsecond			".1234"
	// .999999999		Trailing zeros removed nanosecond			".123456"
	//
	// Daytime
	// Morning			day time									"صبح"
	// PM				Full 12-Hour marker							"قبل از ظهر"
	// pm				Short 12-Hour marker						"ق.ظ"
	tests := []struct {
		layout   string
		expected string
	}{
		// Hour
		{"15", "14"},
		{"03", "02"},
		{"3", "2"},
		// Minute
		{"04", "05"},
		{"4", "5"},
		// Second
		{"05", "06"},
		{"5", "6"},
		// Milliseconds
		{".999999999", ".01034567"},
		{".999999", ".010345"},
		{".999", ".01"},
		{".000000000", ".010345670"},
		{".000000", ".010345"},
		{".000", ".010"},
		// Daytime
		{"Morning", "ظهر"},
		{"PM", "بعد از ظهر"},
		{"pm", "ب.ظ"},
	}

	date := gojalaali.Date(1400, gojalaali.Farvardin, 1, 14, 5, 6, 10345670, gojalaali.TehranTz())
	for _, test := range tests {
		formatted := date.Format(test.layout)
		if formatted != test.expected {
			t.Errorf("fail %s, expected %s, got %s", test.layout, test.expected, formatted)
		}
	}
}

func TestZoneFormat(t *testing.T) {
	// Timezone
	// MST				Abbreviation of the time zone				"UTC"
	// Z070000			zone offset	Hour, Minute and second			"Z" or "+033000"
	// Z0700			zone offset Hour and Minute					"Z" or "+0330"
	// Z07:00:00		zone offset	Hour, Minute and second			"Z" or "+03:30:00"
	// Z07:00			zone offset Hour and Minute					"Z" or "+03:30"
	// Z07				zone offset Hour							"Z" or "+03"
	// -070000			zone offset	Hour, Minute and second			"+033000"
	// -0700			zone offset Hour and Minute					"+0330"
	// -07:00:00		zone offset Hour, Minute and second			"+03:30:00"
	// -07:00			zone offset Hour and Minute					"+03:30"
	// -07				zone offset Hour							"+03"
	tests := []struct {
		layout   string
		expected string
	}{
		// Timezone
		{"MST", "Asia/Tehran"},
		{"Z070000", "+033000"},
		{"Z0700", "+0330"},
		{"Z07:00:00", "+03:30:00"},
		{"Z07:00", "+03:30"},
		{"Z07", "+03"},
		{"-070000", "+033000"},
		{"-0700", "+0330"},
		{"-07:00:00", "+03:30:00"},
		{"-07:00", "+03:30"},
		{"-07", "+03"},
	}

	date := gojalaali.Date(1400, gojalaali.Farvardin, 1, 0, 0, 0, 0, gojalaali.TehranTz())
	for _, test := range tests {
		formatted := date.Format(test.layout)
		if formatted != test.expected {
			t.Errorf("fail %s, expected %s, got %s", test.layout, test.expected, formatted)
		}
	}
}
