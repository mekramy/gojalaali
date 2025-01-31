package gojalaali_test

import (
	"testing"
	"time"

	"github.com/mekramy/gojalaali"
)

func TestJalaali(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.TehranTz())
		expected := "1403-01-15T20:14:00+03:30"
		result := date.String()
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("FormatDate", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.TehranTz())
		expected := "1403-01-15 20:14:00"
		result := date.Format(time.DateTime)
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("AddDate", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 1, gojalaali.TehranTz())
		newDate := date.AddDate(1, 1, 1)
		expected := "1404-02-16T20:14:00+03:30"
		result := newDate.String()
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("AddTime", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.TehranTz()).
			AddTime(1, 30, 0, 0)
		expected := "1403-01-15T21:44:00+03:30"
		result := date.String()
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("SetTime", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.TehranTz())
		date.SetTime(-1, 30, 10, 10)
		expected := "1403-01-15T20:30:10.000000010+03:30"
		result := date.Format(time.RFC3339Nano)
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("BeginningOfMonth", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.TehranTz()).
			BeginningOfMonth()
		expected := "1403-01-01T00:00:00+03:30"
		result := date.String()
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("EndOfMonth", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.KabulTz())
		newDate := date.EndOfMonth()
		expected := "1403-01-31T23:59:59.999999999+04:30"
		result := newDate.Format(time.RFC3339Nano)
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("IsLeap", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.TehranTz())
		expected := true
		result := date.IsLeap()
		if expected != result {
			t.Errorf("Expect %v but get %v", expected, result)
		}
	})

	t.Run("Unix", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 0, gojalaali.TehranTz())
		expected := date.Unix()
		result := gojalaali.Unix(date.Unix(), 0).Unix()
		if expected != result {
			t.Errorf("Expect %d but get %d", expected, result)
		}
	})

	t.Run("Format", func(t *testing.T) {
		date := gojalaali.Date(1403, 01, 15, 20, 14, 0, 9841223, gojalaali.KabulTz())
		expected := ".009841 Asia/Kabul"
		result := date.Format(".999999 MST")
		if expected != result {
			t.Errorf("Expect %s but get %s", expected, result)
		}
	})

	t.Run("Now", func(t *testing.T) {
		now := gojalaali.Now()
		expected := time.Now().Year()
		result := now.Time().Year()
		if expected != result {
			t.Errorf("Expect %d but get %d", expected, result)
		}
	})
}
