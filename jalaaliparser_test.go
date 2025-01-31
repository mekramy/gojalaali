package gojalaali_test

import (
	"testing"

	"github.com/mekramy/gojalaali"
)

func TestDateParse(t *testing.T) {
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
		datetime string
	}{
		{"2006", "1403"},
		{"06", "03"},
		{"January", "اسفند"},
		{"Jan", "اسف"},
		{"01", "07"},
		{"1", "7"},
		{"02", "08"},
		{"_2", " 9"},
		{"2", "3"},
	}

	for _, test := range tests {
		jalaali, err := gojalaali.Parse(test.layout, test.datetime)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		formatted := jalaali.Format(test.layout)
		if formatted != test.datetime {
			t.Errorf("fail %s, expected %s, got %s", test.layout, test.datetime, formatted)
		}
	}
}

func TestTimeParse(t *testing.T) {
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
	tests := []struct {
		layout   string
		datetime string
	}{
		{"15", "15"},
		{"03", "03"},
		{"3", "9"},
		{"04", "03"},
		{"4", "3"},
		{"05", "09"},
		{"5", "2"},
		{".000", ".010"},
		{".000000", ".012340"},
		{".000000000", ".012345600"},
		{".999", ".01"},
		{".999999", ".01234"},
		{".999999999", ".00123456"},
	}

	for _, test := range tests {
		jalaali, err := gojalaali.Parse(test.layout, test.datetime)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		formatted := jalaali.Format(test.layout)
		if formatted != test.datetime {
			t.Errorf("fail %s, expected %s, got %s", test.layout, test.datetime, formatted)
		}
	}
}

func TestZoneParse(t *testing.T) {
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
		datetime string
	}{
		{"MST", "UTC"},
		{"Z070000", "Z"},
		{"Z070000", "+033000"},
		{"Z0700", "Z"},
		{"Z0700", "+0330"},
		{"Z07:00:00", "Z"},
		{"Z07:00:00", "+03:30:00"},
		{"Z07:00", "Z"},
		{"Z07:00", "+03:30"},
		{"Z07", "Z"},
		{"Z07", "+03"},
		{"-070000", "+033000"},
		{"-0700", "+0330"},
		{"-07:00:00", "+03:30:00"},
		{"-07:00", "+03:30"},
		{"-07", "+03"},
	}

	for _, test := range tests {
		jalaali, err := gojalaali.Parse(test.layout, test.datetime)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		formatted := jalaali.Format(test.layout)
		if formatted != test.datetime {
			t.Errorf("fail %s, expected %s, got %s", test.layout, test.datetime, formatted)
		}
	}
}
