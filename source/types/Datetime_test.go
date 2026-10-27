package types

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestDatetime(t *testing.T) {

	t.Run("IsAfter()", func(t *testing.T) {

		datetime1 := ToDatetime("2022-12-31T23:59:59Z")
		datetime2 := ToDatetime("2023-01-01T00:00:01Z")

		if datetime1.IsAfter(datetime2) != false {
			t.Errorf("Expected %s to be after %s", datetime2.String(), datetime1.String())
		}

		if datetime2.IsAfter(datetime1) != true {
			t.Errorf("Expected %s to be after %s", datetime2.String(), datetime1.String())
		}

	})

	t.Run("IsBefore()", func(t *testing.T) {

		datetime1 := ToDatetime("2022-12-31T23:59:59Z")
		datetime2 := ToDatetime("2023-01-01T00:00:01Z")

		if datetime1.IsBefore(datetime2) != true {
			t.Errorf("Expected %s to be before %s", datetime1.String(), datetime2.String())
		}

		if datetime2.IsBefore(datetime1) != false {
			t.Errorf("Expected %s to be before %s", datetime1.String(), datetime2.String())
		}

	})

	t.Run("IsFuture()", func(t *testing.T) {

		past := time.Now().Local().Add(-1 * time.Second)
		future := time.Now().Local().Add(1 * time.Second)

		datetime1 := ToDatetime(past.Format(time.RFC3339))
		datetime2 := ToDatetime(future.Format(time.RFC3339))

		if datetime1.IsFuture() != false {
			t.Errorf("Expected %s to be in the past", datetime1.String())
		}

		if datetime2.IsFuture() != true {
			t.Errorf("Expected %s to be in the future", datetime2.String())
		}

	})

	t.Run("IsPast()", func(t *testing.T) {

		past := time.Now().Local().Add(-1 * time.Second)
		future := time.Now().Local().Add(1 * time.Second)

		datetime1 := ToDatetime(past.Format(time.RFC3339))
		datetime2 := ToDatetime(future.Format(time.RFC3339))

		if datetime1.IsPast() != true {
			t.Errorf("Expected %s to be in the past", datetime1.String())
		}

		if datetime2.IsPast() != false {
			t.Errorf("Expected %s to be in the future", datetime2.String())
		}

	})

	t.Run("Offset()", func(t *testing.T) {

		datetime1 := ToDatetime("2022-12-31T23:59:59Z")
		datetime2 := ToDatetime("2022-12-31T23:59:59Z")
		datetime3 := ToDatetime("2023-01-01T00:00:01Z")
		datetime4 := ToDatetime("2023-01-01T00:00:01Z")

		datetime1.Offset("-12:34")
		datetime2.Offset("+12:34")
		datetime3.Offset("-12:34")
		datetime4.Offset("+12:34")

		want1 := "2023-01-01 12:33:59"
		want2 := "2022-12-31 11:25:59"
		want3 := "2023-01-01 12:34:01"
		want4 := "2022-12-31 11:26:01"

		if datetime1.String() != want1 {
			t.Errorf("Expected %s to be %s", datetime1.String(), want1)
		}

		if datetime2.String() != want2 {
			t.Errorf("Expected %s to be %s", datetime2.String(), want2)
		}

		if datetime3.String() != want3 {
			t.Errorf("Expected %s to be %s", datetime3.String(), want3)
		}

		if datetime4.String() != want4 {
			t.Errorf("Expected %s to be %s", datetime4.String(), want4)
		}

	})

	t.Run("Parse(en_US local format)", func(t *testing.T) {

		current_year, _, _ := time.Now().Date()

		datetime1 := ToDatetime("Sun Jan 01 12:34:56 2023")
		datetime2 := ToDatetime("Thu Feb 29 23:59:59 2024")
		datetime3 := ToDatetime("Jan 01 12:34:56")
		datetime4 := ToDatetime("Feb 29 23:59:59")

		want1 := ToDatetime("2023-01-01 12:34:56")
		want1.ToZulu()

		want2 := ToDatetime("2024-02-29 23:59:59")
		want2.ToZulu()

		want3 := ToDatetime(strconv.FormatUint(uint64(current_year), 10) + "-01-01 12:34:56")
		want3.ToZulu()

		want4 := ToDatetime(strconv.FormatUint(uint64(current_year), 10) + "-02-29 23:59:59")
		want4.ToZulu()

		if datetime1.String() != want1.String() {
			t.Errorf("Expected %s to be %s", datetime1.String(), want1.String())
		}

		if datetime2.String() != want2.String() {
			t.Errorf("Expected %s to be %s", datetime2.String(), want2.String())
		}

		if datetime3.String() != want3.String() {
			t.Errorf("Expected %s to be %s", datetime3.String(), want3.String())
		}

		if datetime4.String() != want4.String() {
			t.Errorf("Expected %s to be %s", datetime4.String(), want4.String())
		}

	})

	t.Run("Parse(en_US meridiem format)", func(t *testing.T) {

		datetime1 := ToDatetime("Sat Dec 31 01:02:03 AM 2022")
		datetime2 := ToDatetime("Sat Dec 31 01:02:03 PM 2022")

		want1 := ToDatetime("2022-12-31 01:02:03")
		want1.ToZulu()

		want2 := ToDatetime("2022-12-31 13:02:03")
		want2.ToZulu()

		if datetime1.String() != want1.String() {
			t.Errorf("Expected %s to be %s", datetime1.String(), want1.String())
		}

		if datetime2.String() != want2.String() {
			t.Errorf("Expected %s to be %s", datetime2.String(), want2.String())
		}

	})

	t.Run("Parse(en_US timezone format)", func(t *testing.T) {

		datetime1 := ToDatetime("Sat Dec 31 01:02:03 AM CET 2022")
		datetime2 := ToDatetime("Sat Dec 31 01:02:03 PM CET 2022")
		datetime3 := ToDatetime("Sat Dec 31 01:02:03 PM CET 2022")
		datetime4 := ToDatetime("Sat Dec 31 12:13:14 CET 2022")

		want1 := ToDatetime("2022-12-31 01:02:03")
		want1.ToZulu()

		want2 := ToDatetime("2022-12-31 13:02:03")
		want2.ToZulu()

		want3 := ToDatetime("2022-12-31 13:02:03")
		want3.ToZulu()

		want4 := ToDatetime("2022-12-31 12:13:14")
		want4.ToZulu()

		if datetime1.String() != want1.String() {
			t.Errorf("Expected %s to be %s", datetime1.String(), want1.String())
		}

		if datetime2.String() != want2.String() {
			t.Errorf("Expected %s to be %s", datetime2.String(), want2.String())
		}

		if datetime3.String() != want3.String() {
			t.Errorf("Expected %s to be %s", datetime3.String(), want3.String())
		}

		if datetime4.String() != want4.String() {
			t.Errorf("Expected %s to be %s", datetime4.String(), want4.String())
		}

	})

	t.Run("Parse(ISO8601 zulu format)", func(t *testing.T) {

		datetime1 := ToDatetime("2022-12-31 01:02:03")
		datetime2 := ToDatetime("2022-12-31T01:02:03Z")
		datetime3 := ToDatetime("2022-12-31T13:14:15.123Z")

		want1 := "2022-12-31 01:02:03"
		want2 := "2022-12-31 01:02:03"
		want3 := "2022-12-31 13:14:15"

		if datetime1.String() != want1 {
			t.Errorf("Expected %s to be %s", datetime1.String(), want1)
		}

		if datetime2.String() != want2 {
			t.Errorf("Expected %s to be %s", datetime2.String(), want2)
		}

		if datetime3.String() != want3 {
			t.Errorf("Expected %s to be %s", datetime3.String(), want3)
		}

	})

	t.Run("Parse(ISO8601 juliet format)", func(t *testing.T) {

		datetime1 := ToDatetime("2022-12-31T01:02:03+12:30")
		datetime2 := ToDatetime("2022-12-31T13:14:15.123+12:30")
		datetime3 := ToDatetime("2022-12-31T01:02:03-03:45")
		datetime4 := ToDatetime("2022-12-31T13:14:15.123-03:45")

		want1 := "2022-12-31 01:02:03"
		want2 := "2022-12-31 13:14:15"
		want3 := "2022-12-31 01:02:03"
		want4 := "2022-12-31 13:14:15"

		if datetime1.String() != want1 {
			t.Errorf("Expected %s to be %s", datetime1.String(), want1)
		}

		if datetime2.String() != want2 {
			t.Errorf("Expected %s to be %s", datetime2.String(), want2)
		}

		if datetime3.String() != want3 {
			t.Errorf("Expected %s to be %s", datetime3.String(), want3)
		}

		if datetime4.String() != want4 {
			t.Errorf("Expected %s to be %s", datetime4.String(), want4)
		}

	})

	t.Run("ToDays()", func(t *testing.T) {

		datetime1 := ToDatetime("2021-01-01T23:59:59Z")
		datetime2 := ToDatetime("2022-02-14T12:34:56Z")
		datetime3 := ToDatetime("2023-02-01T00:00:01Z")
		datetime4 := ToDatetime("2024-02-14T12:34:56Z")

		if datetime1.ToDays() != 31 {
			t.Errorf("Expected %d to be %d", datetime1.ToDays(), 31)
		}

		if datetime2.ToDays() != 28 {
			t.Errorf("Expected %d to be %d", datetime2.ToDays(), 28)
		}

		if datetime3.ToDays() != 28 {
			t.Errorf("Expected %d to be %d", datetime3.ToDays(), 28)
		}

		if datetime4.ToDays() != 29 {
			t.Errorf("Expected %d to be %d", datetime4.ToDays(), 29)
		}

	})

	t.Run("ToDatetimeDifference()", func(t *testing.T) {

		datetime11 := ToDatetime("2021-01-10 12:01:01")
		datetime12 := ToDatetime("2024-03-11 11:00:00")

		datetime21 := ToDatetime("2024-01-01 12:01:01")
		datetime22 := ToDatetime("2024-03-10 12:01:01")

		datetime31 := ToDatetime("2024-01-10 12:01:01")
		datetime32 := ToDatetime("2024-01-13 11:00:00")

		datetime41 := ToDatetime("2023-03-13 12:01:01")
		datetime42 := ToDatetime("2024-01-11 11:00:00")

		diff1 := datetime11.ToDatetimeDifference(datetime12)
		want1 := "0003-02-00 22:59:59"

		diff2 := datetime21.ToDatetimeDifference(datetime22)
		want2 := "0000-02-09 00:00:00"

		diff3 := datetime31.ToDatetimeDifference(datetime32)
		want3 := "0000-00-02 22:59:59"

		diff4 := datetime41.ToDatetimeDifference(datetime42)
		want4 := "0000-09-28 22:59:59"

		if diff1.String() != want1 {
			t.Errorf("Expected %s to be %s", diff1.String(), want1)
		}

		if diff2.String() != want2 {
			t.Errorf("Expected %s to be %s", diff2.String(), want2)
		}

		if diff3.String() != want3 {
			t.Errorf("Expected %s to be %s", diff3.String(), want3)
		}

		if diff4.String() != want4 {
			t.Errorf("Expected %s to be %s", diff4.String(), want4)
		}

	})

	t.Run("ToTimeDifference()", func(t *testing.T) {

		datetime11 := ToDatetime("2021-01-10 12:01:01")
		datetime12 := ToDatetime("2024-03-11 11:00:00")

		datetime21 := ToDatetime("2024-01-01 12:01:01")
		datetime22 := ToDatetime("2024-03-10 12:01:01")

		datetime31 := ToDatetime("2024-01-10 12:01:01")
		datetime32 := ToDatetime("2024-01-13 11:00:00")

		datetime41 := ToDatetime("2024-01-10 12:01:01")
		datetime42 := ToDatetime("2024-01-11 11:00:00")

		diff1 := datetime11.ToTimeDifference(datetime12)
		want1 := "27742:59:59"

		diff2 := datetime21.ToTimeDifference(datetime22)
		want2 := "1656:00:00"

		diff3 := datetime31.ToTimeDifference(datetime32)
		want3 := "70:59:59"

		diff4 := datetime41.ToTimeDifference(datetime42)
		want4 := "22:59:59"

		if diff1.String() != want1 {
			t.Errorf("Expected %s to be %s", diff1.String(), want1)
		}

		if diff2.String() != want2 {
			t.Errorf("Expected %s to be %s", diff2.String(), want2)
		}

		if diff3.String() != want3 {
			t.Errorf("Expected %s to be %s", diff3.String(), want3)
		}

		if diff4.String() != want4 {
			t.Errorf("Expected %s to be %s", diff4.String(), want4)
		}

	})

	t.Run("ToWeekday()", func(t *testing.T) {

		datetime1 := ToDatetime("2021-01-01T23:59:59Z")
		datetime2 := ToDatetime("2022-02-14T12:34:56Z")
		datetime3 := ToDatetime("2023-02-01T00:00:01Z")
		datetime4 := ToDatetime("2024-02-14T12:34:56Z")

		if datetime1.ToWeekday() != "Friday" {
			t.Errorf("Expected %s to be %s", datetime1.String(), "Friday")
		}

		if datetime2.ToWeekday() != "Monday" {
			t.Errorf("Expected %s to be %s", datetime2.String(), "Monday")
		}

		if datetime3.ToWeekday() != "Wednesday" {
			t.Errorf("Expected %s to be %s", datetime3.String(), "Wednesday")
		}

		if datetime4.ToWeekday() != "Wednesday" {
			t.Errorf("Expected %s to be %s", datetime4.String(), "Wednesday")
		}

	})

}

func TestDatetimeFactories(t *testing.T) {

	before := time.Now()
	current := NewDatetime()
	after := time.Now()

	if current.IsValid() == false {
		t.Fatalf("NewDatetime() = %+v, expected a valid datetime", current)
	}

	if current.Year != uint(before.Year()) && current.Year != uint(after.Year()) {
		t.Fatalf("NewDatetime().Year = %d, expected the current year", current.Year)
	}

	valid := []string{
		"2024-02-29T12:34:56Z",
		"2024-02-29",
		"20240229",
		"2024-02",
		"202402",
		"2024-02-29 12:34:56",
	}

	for _, value := range valid {
		if IsDatetime(value) == false {
			t.Errorf("IsDatetime(%q) = false, expected true", value)
		}
	}

	invalid := []string{
		"",
		"not a datetime",
		"2024/02/29",
		"1750-01-01",
		"2024",
	}

	for _, value := range invalid {
		if IsDatetime(value) == true {
			t.Errorf("IsDatetime(%q) = true, expected false", value)
		}
	}

	for _, value := range invalid {
		if parsed := ParseDatetime(value); parsed != nil {
			t.Errorf("ParseDatetime(%q) = %v, expected nil", value, parsed.String())
		}
	}

	parsed := ParseDatetime("2024-02-29T12:34:56Z")

	if parsed == nil {
		t.Fatal("ParseDatetime() returned nil for a valid value")
	}

	if parsed.String() != "2024-02-29 12:34:56" {
		t.Errorf("ParseDatetime().String() = %q", parsed.String())
	}

	empty := ToDatetime("")

	if empty.String() != "0000-00-00 00:00:00" {
		t.Fatal("ToDatetime(\"\") should return the zero value")
	}

	invalidValue := ToDatetime("not a datetime")

	if invalidValue.String() != "0000-00-00 00:00:00" {
		t.Fatal("ToDatetime() with invalid input should return the zero value")
	}

}

func TestDatetimeEqualHashIsValid(t *testing.T) {

	base := ToDatetime("2024-02-29T12:34:56Z")
	same := ToDatetime("2024-02-29T12:34:56Z")
	different := ToDatetime("2024-03-01T00:00:00Z")

	if base.Equal(same) == false || base.Equal(different) == true {
		t.Fatal("Datetime.Equal() failed")
	}

	if base.IsAfter(same) || base.IsBefore(same) {
		t.Fatal("equal datetimes should be neither after nor before each other")
	}

	if base.Hash() == "" || base.Hash() != same.Hash() {
		t.Fatal("Hash() should be deterministic for equal datetimes")
	}

	if base.Hash() == different.Hash() {
		t.Fatal("different datetimes should not share a hash")
	}

	cases := []struct {
		value Datetime
		valid bool
	}{
		{Datetime{}, false},
		{Datetime{Year: 1752, Month: 1, Day: 1}, false},
		{Datetime{Year: 1753, Month: 1, Day: 1}, true},
		{Datetime{Year: 2024, Month: 0, Day: 1}, false},
		{Datetime{Year: 2024, Month: 13, Day: 1}, false},
		{Datetime{Year: 2024, Month: 1, Day: 0}, false},
		{Datetime{Year: 2024, Month: 1, Day: 32}, false},
		{Datetime{Year: 2024, Month: 12, Day: 31}, true},
	}

	for _, c := range cases {
		if result := c.value.IsValid(); result != c.valid {
			t.Errorf("Datetime%+v.IsValid() = %v, expected %v", c.value, result, c.valid)
		}
	}

}

func TestDatetimeYesterdayTomorrow(t *testing.T) {

	yesterdayCases := []struct {
		value    string
		expected string
	}{
		{"2024-03-15", "2024-03-14 00:00:00"},
		{"2024-03-01", "2024-02-29 00:00:00"},
		{"2023-03-01", "2023-02-28 00:00:00"},
		{"2024-01-01", "2023-12-31 00:00:00"},
		{"2024-01-15", "2024-01-14 00:00:00"},
	}

	for _, c := range yesterdayCases {
		value := ToDatetime(c.value)

		yesterday := value.Yesterday()
		result := yesterday.String()

		if result != c.expected {
			t.Errorf("ToDatetime(%q).Yesterday() = %s, expected %s", c.value, result, c.expected)
		}
	}

	tomorrowCases := []struct {
		value    string
		expected string
	}{
		{"2024-03-15", "2024-03-16 00:00:00"},
		{"2024-02-29", "2024-03-01 00:00:00"},
		{"2023-02-28", "2023-03-01 00:00:00"},
		{"2024-12-31", "2025-01-01 00:00:00"},
		{"2024-01-31", "2024-02-01 00:00:00"},
	}

	for _, c := range tomorrowCases {
		value := ToDatetime(c.value)

		tomorrow := value.Tomorrow()
		result := tomorrow.String()

		if result != c.expected {
			t.Errorf("ToDatetime(%q).Tomorrow() = %s, expected %s", c.value, result, c.expected)
		}
	}

}

func TestDatetimeParseFormatsExtra(t *testing.T) {

	cases := []struct {
		value    string
		expected string
	}{
		{"2024-02-29T12:34:56Z", "2024-02-29 12:34:56"},
		{"2024-02-29T12:34:56.789Z", "2024-02-29 12:34:56"},
		{"2024-02-29T12:34:56+02:00", "2024-02-29 12:34:56"},
		{"2024-02-29T12:34:56-03:45", "2024-02-29 12:34:56"},
		{"2024-02-29 12:34:56", "2024-02-29 12:34:56"},
		{"2024-02-29", "2024-02-29 00:00:00"},
		{"20240229", "2024-02-29 00:00:00"},
		{"2024-02", "2024-02-01 00:00:00"},
		{"202402", "2024-02-01 00:00:00"},
	}

	for _, c := range cases {

		value := ToDatetime(c.value)

		if result := value.String(); result != c.expected {
			t.Errorf("ToDatetime(%q).String() = %q, expected %q", c.value, result, c.expected)
		}

	}

	for _, value := range []string{"", "not a datetime", "2024/02/29", "1750-01-01"} {

		var datetime Datetime

		if datetime.Parse(value) == true {
			t.Errorf("Datetime.Parse(%q) = true, expected false", value)
		}

	}

}

func TestDatetimeString(t *testing.T) {

	cases := []struct {
		value    Datetime
		expected string
	}{
		{Datetime{}, "0000-00-00 00:00:00"},
		{ToDatetime("2024-02-29T12:34:56Z"), "2024-02-29 12:34:56"},
		{Datetime{Year: 2024, Month: 2, Day: 9, Hour: 1, Minute: 2, Second: 3}, "2024-02-09 01:02:03"},
	}

	for _, c := range cases {
		if result := c.value.String(); result != c.expected {
			t.Errorf("Datetime.String() = %q, expected %q", result, c.expected)
		}
	}

}

func TestDatetimeDifferenceZero(t *testing.T) {

	earlier := ToDatetime("2024-02-29T12:34:56Z")
	later := ToDatetime("2024-03-01T12:34:56Z")

	if diff := earlier.ToDatetimeDifference(earlier); diff != (Datetime{}) {
		t.Errorf("ToDatetimeDifference(equal) = %s, expected zero", diff.String())
	}

	if diff := later.ToDatetimeDifference(earlier); diff != (Datetime{}) {
		t.Errorf("ToDatetimeDifference(later, earlier) = %s, expected zero", diff.String())
	}

	if diff := earlier.ToTimeDifference(earlier); diff != (Time{}) {
		t.Errorf("ToTimeDifference(equal) = %s, expected zero", diff.String())
	}

	if diff := later.ToTimeDifference(earlier); diff != (Time{}) {
		t.Errorf("ToTimeDifference(later, earlier) = %s, expected zero", diff.String())
	}

}

func TestDatetimeFuturePastInvariants(t *testing.T) {

	now := time.Now()

	if now.Hour() == 23 || now.Hour() == 0 {
		t.Skip("skipping near midnight to avoid time-of-day wrapping")
	}

	past := ToDatetime(now.Add(-1 * time.Hour).Format(time.RFC3339))
	future := ToDatetime(now.Add(1 * time.Hour).Format(time.RFC3339))

	if past.IsPast() == false || past.IsFuture() == true {
		t.Errorf("past datetime %s should only be in the past", past.String())
	}

	if future.IsFuture() == false || future.IsPast() == true {
		t.Errorf("future datetime %s should only be in the future", future.String())
	}

}

func TestDatetimeZulu(t *testing.T) {

	offset := strings.Split(time.Now().String(), " ")[2]
	value := ToDatetime("2024-02-29T12:34:56Z")
	expected := ToDatetime("2024-02-29 12:34:56")

	if len(offset) == 5 {
		expected.Offset(offset)
	}

	value.ToZulu()

	if value.String() != expected.String() {
		t.Errorf("ToZulu() = %s, expected %s", value.String(), expected.String())
	}

	var zero Datetime

	zero.ToZulu()

	if zero != (Datetime{}) {
		t.Fatal("ToZulu() on the zero value should be a no-op")
	}

}

func TestDatetimeJSON(t *testing.T) {

	original := ToDatetime("2026-10-19T11:59:59Z")

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	if string(data) != `"2026-10-19 11:59:59"` {
		t.Errorf("Marshal() = %s, expected \"2026-10-19 11:59:59\"", string(data))
	}

	var decoded Datetime

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if decoded.Equal(original) == false {
		t.Errorf("Unmarshal() = %+v, expected %+v", decoded, original)
	}

	var malformed Datetime

	if err := json.Unmarshal([]byte("123"), &malformed); err == nil {
		t.Fatal("Unmarshal() should fail for malformed JSON")
	}

	unchanged := original

	if err := json.Unmarshal([]byte(`"not a datetime"`), &unchanged); err != nil {
		t.Fatalf("Unmarshal() of an invalid string returned error: %v", err)
	}

	if unchanged.Equal(original) == false {
		t.Fatal("Unmarshal() of an invalid string should leave the datetime unchanged")
	}

}
