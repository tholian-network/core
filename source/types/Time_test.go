package types

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t.Run("AddHour()", func(t *testing.T) {

		time1 := ToTime("11:59:59")
		time2 := ToTime("23:59:59")

		time1.AddHour()
		time2.AddHour()

		if time1.String() != "12:59:59" {
			t.Errorf("Expected %s to be %s", time1.String(), "12:59:59")
		}

		if time2.String() != "24:59:59" {
			t.Errorf("Expected %s to be %s", time2.String(), "24:59:59")
		}

	})

	t.Run("AddMinute()", func(t *testing.T) {

		time1 := ToTime("11:59:59")
		time2 := ToTime("23:59:59")

		time1.AddMinute()
		time2.AddMinute()

		if time1.String() != "12:00:59" {
			t.Errorf("Expected %s to be %s", time1.String(), "12:00:59")
		}

		if time2.String() != "24:00:59" {
			t.Errorf("Expected %s to be %s", time2.String(), "24:00:59")
		}

	})

	t.Run("AddSecond()", func(t *testing.T) {

		time1 := ToTime("11:59:59")
		time2 := ToTime("23:59:59")

		time1.AddSecond()
		time2.AddSecond()

		if time1.String() != "12:00:00" {
			t.Errorf("Expected %s to be %s", time1.String(), "12:00:00")
		}

		if time2.String() != "24:00:00" {
			t.Errorf("Expected %s to be %s", time2.String(), "24:00:00")
		}

	})

	t.Run("AddTime()", func(t *testing.T) {

		time1 := ToTime("11:59:59")
		time2 := ToTime("11:59:59")
		time3 := ToTime("11:59:59")
		time4 := ToTime("23:59:59")
		time5 := ToTime("23:59:59")
		time6 := ToTime("23:59:59")

		time1.AddTime(ToTime("12:00:01"))
		time2.AddTime(ToTime("24:00:01"))
		time3.AddTime(ToTime("48:00:01"))

		time4.AddTime(ToTime("12:00:01"))
		time5.AddTime(ToTime("24:00:01"))
		time6.AddTime(ToTime("48:00:01"))

		if time1.String() != "24:00:00" {
			t.Errorf("Expected %s to be %s", time1.String(), "24:00:00")
		}

		if time2.String() != "36:00:00" {
			t.Errorf("Expected %s to be %s", time2.String(), "36:00:00")
		}

		if time3.String() != "60:00:00" {
			t.Errorf("Expected %s to be %s", time3.String(), "60:00:00")
		}

		if time4.String() != "36:00:00" {
			t.Errorf("Expected %s to be %s", time4.String(), "36:00:00")
		}

		if time5.String() != "48:00:00" {
			t.Errorf("Expected %s to be %s", time5.String(), "48:00:00")
		}

		if time6.String() != "72:00:00" {
			t.Errorf("Expected %s to be %s", time6.String(), "72:00:00")
		}

	})

	t.Run("IsAfter()", func(t *testing.T) {

		time1 := ToTime("11:59:59")
		time2 := ToTime("12:00:00")
		time3 := ToTime("12:00:01")
		time4 := ToTime("23:59:59")
		time5 := ToTime("24:00:00")
		time6 := ToTime("24:01:00")

		if time1.IsAfter(time2) != false {
			t.Errorf("Expected %s to be not after %s", time1.String(), time2.String())
		}

		if time1.IsAfter(time3) != false {
			t.Errorf("Expected %s to be not after %s", time1.String(), time3.String())
		}

		if time1.IsAfter(time4) != false {
			t.Errorf("Expected %s to be not after %s", time1.String(), time4.String())
		}

		if time1.IsAfter(time5) != false {
			t.Errorf("Expected %s to be not after %s", time1.String(), time5.String())
		}

		if time1.IsAfter(time6) != false {
			t.Errorf("Expected %s to be not after %s", time1.String(), time6.String())
		}

		if time2.IsAfter(time1) != true {
			t.Errorf("Expected %s to be after %s", time2.String(), time1.String())
		}

		if time2.IsAfter(time3) != false {
			t.Errorf("Expected %s to be not after %s", time2.String(), time3.String())
		}

		if time2.IsAfter(time4) != false {
			t.Errorf("Expected %s to be not after %s", time2.String(), time4.String())
		}

		if time2.IsAfter(time5) != false {
			t.Errorf("Expected %s to be not after %s", time2.String(), time5.String())
		}

		if time2.IsAfter(time6) != false {
			t.Errorf("Expected %s to be not after %s", time2.String(), time6.String())
		}

		if time3.IsAfter(time1) != true {
			t.Errorf("Expected %s to be after %s", time3.String(), time1.String())
		}

		if time3.IsAfter(time2) != true {
			t.Errorf("Expected %s to be after %s", time3.String(), time2.String())
		}

		if time3.IsAfter(time4) != false {
			t.Errorf("Expected %s to be not after %s", time3.String(), time4.String())
		}

		if time3.IsAfter(time5) != false {
			t.Errorf("Expected %s to be not after %s", time3.String(), time5.String())
		}

		if time3.IsAfter(time6) != false {
			t.Errorf("Expected %s to be not after %s", time3.String(), time6.String())
		}

		if time4.IsAfter(time1) != true {
			t.Errorf("Expected %s to be after %s", time4.String(), time1.String())
		}

		if time4.IsAfter(time2) != true {
			t.Errorf("Expected %s to be after %s", time4.String(), time2.String())
		}

		if time4.IsAfter(time3) != true {
			t.Errorf("Expected %s to be after %s", time4.String(), time3.String())
		}

		if time4.IsAfter(time5) != false {
			t.Errorf("Expected %s to be not after %s", time4.String(), time5.String())
		}

		if time4.IsAfter(time6) != false {
			t.Errorf("Expected %s to be not after %s", time4.String(), time6.String())
		}

		if time5.IsAfter(time1) != true {
			t.Errorf("Expected %s to be after %s", time5.String(), time1.String())
		}

		if time5.IsAfter(time2) != true {
			t.Errorf("Expected %s to be after %s", time5.String(), time2.String())
		}

		if time5.IsAfter(time3) != true {
			t.Errorf("Expected %s to be after %s", time5.String(), time3.String())
		}

		if time5.IsAfter(time4) != true {
			t.Errorf("Expected %s to be after %s", time5.String(), time4.String())
		}

		if time5.IsAfter(time6) != false {
			t.Errorf("Expected %s to be not after %s", time5.String(), time6.String())
		}

		if time6.IsAfter(time1) != true {
			t.Errorf("Expected %s to be after %s", time6.String(), time1.String())
		}

		if time6.IsAfter(time2) != true {
			t.Errorf("Expected %s to be after %s", time6.String(), time2.String())
		}

		if time6.IsAfter(time3) != true {
			t.Errorf("Expected %s to be after %s", time6.String(), time3.String())
		}

		if time6.IsAfter(time4) != true {
			t.Errorf("Expected %s to be after %s", time6.String(), time4.String())
		}

		if time6.IsAfter(time5) != true {
			t.Errorf("Expected %s to be after %s", time6.String(), time5.String())
		}

	})

	t.Run("IsBefore()", func(t *testing.T) {

		time1 := ToTime("11:59:59")
		time2 := ToTime("12:00:00")
		time3 := ToTime("12:00:01")
		time4 := ToTime("23:59:59")
		time5 := ToTime("24:00:00")
		time6 := ToTime("24:01:00")

		if time1.IsBefore(time2) != true {
			t.Errorf("Expected %s to be before %s", time1.String(), time2.String())
		}

		if time1.IsBefore(time3) != true {
			t.Errorf("Expected %s to be before %s", time1.String(), time3.String())
		}

		if time1.IsBefore(time4) != true {
			t.Errorf("Expected %s to be before %s", time1.String(), time4.String())
		}

		if time1.IsBefore(time5) != true {
			t.Errorf("Expected %s to be before %s", time1.String(), time5.String())
		}

		if time1.IsBefore(time6) != true {
			t.Errorf("Expected %s to be before %s", time1.String(), time6.String())
		}

		if time2.IsBefore(time1) != false {
			t.Errorf("Expected %s to be not before %s", time2.String(), time1.String())
		}

		if time2.IsBefore(time3) != true {
			t.Errorf("Expected %s to be before %s", time2.String(), time3.String())
		}

		if time2.IsBefore(time4) != true {
			t.Errorf("Expected %s to be before %s", time2.String(), time4.String())
		}

		if time2.IsBefore(time5) != true {
			t.Errorf("Expected %s to be before %s", time2.String(), time5.String())
		}

		if time2.IsBefore(time6) != true {
			t.Errorf("Expected %s to be before %s", time2.String(), time6.String())
		}

		if time3.IsBefore(time1) != false {
			t.Errorf("Expected %s to be not before %s", time3.String(), time1.String())
		}

		if time3.IsBefore(time2) != false {
			t.Errorf("Expected %s to be not before %s", time3.String(), time2.String())
		}

		if time3.IsBefore(time4) != true {
			t.Errorf("Expected %s to be before %s", time3.String(), time4.String())
		}

		if time3.IsBefore(time5) != true {
			t.Errorf("Expected %s to be before %s", time3.String(), time5.String())
		}

		if time3.IsBefore(time6) != true {
			t.Errorf("Expected %s to be before %s", time3.String(), time6.String())
		}

		if time4.IsBefore(time1) != false {
			t.Errorf("Expected %s to be not before %s", time4.String(), time1.String())
		}

		if time4.IsBefore(time2) != false {
			t.Errorf("Expected %s to be not before %s", time4.String(), time2.String())
		}

		if time4.IsBefore(time3) != false {
			t.Errorf("Expected %s to be not before %s", time4.String(), time3.String())
		}

		if time4.IsBefore(time5) != true {
			t.Errorf("Expected %s to be before %s", time4.String(), time5.String())
		}

		if time4.IsBefore(time6) != true {
			t.Errorf("Expected %s to be before %s", time4.String(), time6.String())
		}

		if time5.IsBefore(time1) != false {
			t.Errorf("Expected %s to be not before %s", time5.String(), time1.String())
		}

		if time5.IsBefore(time2) != false {
			t.Errorf("Expected %s to be not before %s", time5.String(), time2.String())
		}

		if time5.IsBefore(time3) != false {
			t.Errorf("Expected %s to be not before %s", time5.String(), time3.String())
		}

		if time5.IsBefore(time4) != false {
			t.Errorf("Expected %s to be not before %s", time5.String(), time4.String())
		}

		if time5.IsBefore(time6) != true {
			t.Errorf("Expected %s to be before %s", time5.String(), time6.String())
		}

		if time6.IsBefore(time1) != false {
			t.Errorf("Expected %s to be not before %s", time6.String(), time1.String())
		}

		if time6.IsBefore(time2) != false {
			t.Errorf("Expected %s to be not before %s", time6.String(), time2.String())
		}

		if time6.IsBefore(time3) != false {
			t.Errorf("Expected %s to be not before %s", time6.String(), time3.String())
		}

		if time6.IsBefore(time4) != false {
			t.Errorf("Expected %s to be not before %s", time6.String(), time4.String())
		}

		if time6.IsBefore(time5) != false {
			t.Errorf("Expected %s to be not before %s", time6.String(), time5.String())
		}

	})

	t.Run("IsFuture()", func(t *testing.T) {

		past := time.Now().Local().Add(-1 * time.Second)
		future := time.Now().Local().Add(1 * time.Second)

		time1 := ToTime(past.Format(time.TimeOnly))
		time2 := ToTime(future.Format(time.TimeOnly))

		if time1.IsFuture() != false {
			t.Errorf("Expected %s to be in the past", time1.String())
		}

		if time2.IsFuture() != true {
			t.Errorf("Expected %s to be in the future", time2.String())
		}

	})

	t.Run("IsPast()", func(t *testing.T) {

		past := time.Now().Local().Add(-1 * time.Second)
		future := time.Now().Local().Add(1 * time.Second)

		time1 := ToTime(past.Format(time.TimeOnly))
		time2 := ToTime(future.Format(time.TimeOnly))

		if time1.IsPast() != true {
			t.Errorf("Expected %s to be in the past", time1.String())
		}

		if time2.IsPast() != false {
			t.Errorf("Expected %s to be in the future", time2.String())
		}

	})

	t.Run("Offset()", func(t *testing.T) {

		time1 := ToTime("23:59:59")
		time2 := ToTime("23:59:59")
		time3 := ToTime("00:00:01")
		time4 := ToTime("00:00:01")

		time1.Offset("-12:34")
		time2.Offset("+12:34")
		time3.Offset("-12:34")
		time4.Offset("+12:34")

		want1 := "12:33:59"
		want2 := "11:25:59"
		want3 := "12:34:01"
		want4 := "11:26:01"

		if time1.String() != want1 {
			t.Errorf("Expected %s to be %s", time1.String(), want1)
		}

		if time2.String() != want2 {
			t.Errorf("Expected %s to be %s", time2.String(), want2)
		}

		if time3.String() != want3 {
			t.Errorf("Expected %s to be %s", time3.String(), want3)
		}

		if time4.String() != want4 {
			t.Errorf("Expected %s to be %s", time4.String(), want4)
		}

	})

	t.Run("Parse(HH:ii:ss format)", func(t *testing.T) {

		time1 := ToTime("12:34:56")
		time2 := ToTime("23:59:59")

		want1 := "12:34:56"
		want2 := "23:59:59"

		if time1.String() != want1 {
			t.Errorf("Expected %s to be %s", time1.String(), want1)
		}

		if time2.String() != want2 {
			t.Errorf("Expected %s to be %s", time2.String(), want2)
		}

	})

	t.Run("Parse(HH:ii format)", func(t *testing.T) {

		time1 := ToTime("01:02")
		time2 := ToTime("23:59")

		want1 := "01:02:00"
		want2 := "23:59:00"

		if time1.String() != want1 {
			t.Errorf("Expected %s to be %s", time1.String(), want1)
		}

		if time2.String() != want2 {
			t.Errorf("Expected %s to be %s", time2.String(), want2)
		}

	})

}

func TestTimeConstructorParse(t *testing.T) {

	value := NewTime()

	if value.Parse("11:59:59") == false {
		t.Fatal("Parse() returned false")
	}

	if value.String() != "11:59:59" {
		t.Errorf("Expected %s to be %s", value.String(), "11:59:59")
	}

}

func TestTimeFactories(t *testing.T) {

	zero := NewTime()

	if zero.Hour != 0 || zero.Minute != 0 || zero.Second != 0 {
		t.Fatalf("NewTime() = %+v, expected midnight", zero)
	}

	if zero.String() != "00:00:00" {
		t.Errorf("NewTime().String() = %q", zero.String())
	}

	valid := []string{"00:00:00", "11:59:59", "24:00:00", "12:34", "00:00"}
	invalid := []string{"", "not a time", "12", "12:", ":"}

	for _, value := range valid {
		if IsTime(value) == false {
			t.Errorf("IsTime(%q) = false, expected true", value)
		}
	}

	for _, value := range invalid {
		if IsTime(value) == true {
			t.Errorf("IsTime(%q) = true, expected false", value)
		}
	}

	for _, value := range []string{"", "not a time", "12"} {
		if parsed := ParseTime(value); parsed != nil {
			t.Errorf("ParseTime(%q) = %v, expected nil", value, parsed.String())
		}
	}

	parsed := ParseTime("11:59:59")

	if parsed == nil {
		t.Fatal("ParseTime() returned nil for a valid value")
	}

	if parsed.String() != "11:59:59" {
		t.Errorf("ParseTime().String() = %q", parsed.String())
	}

	empty := ToTime("")

	if empty.String() != "00:00:00" {
		t.Fatal("ToTime(\"\") should return midnight")
	}

	invalidValue := ToTime("not a time")

	if invalidValue.String() != "00:00:00" {
		t.Fatal("ToTime() with invalid input should return midnight")
	}

	partial := ToTime("12:34")

	if partial.String() != "12:34:00" {
		t.Fatal("ToTime() should accept the HH:ii format")
	}

}

func TestTimeEqualHashIsValid(t *testing.T) {

	time1 := ToTime("11:59:59")
	time2 := ToTime("11:59:59")
	time3 := ToTime("12:00:00")

	if time1.Equal(time2) == false || time1.Equal(time3) == true {
		t.Fatal("Time.Equal() failed")
	}

	if time1.Hash() == "" || time1.Hash() != time2.Hash() {
		t.Fatal("Hash() should be deterministic for equal times")
	}

	if time1.Hash() == time3.Hash() {
		t.Fatal("different times should not share a hash")
	}

	cases := []struct {
		value Time
		valid bool
	}{
		{NewTime(), true},
		{Time{Hour: 24, Minute: 60, Second: 60}, true},
		{Time{Hour: 25}, false},
		{Time{Hour: 12, Minute: 61}, false},
		{Time{Hour: 12, Minute: 0, Second: 61}, false},
	}

	for _, c := range cases {
		if result := c.value.IsValid(); result != c.valid {
			t.Errorf("Time%+v.IsValid() = %v, expected %v", c.value, result, c.valid)
		}
	}

}

func TestTimeFuturePastInvariants(t *testing.T) {

	now := time.Now()

	if now.Hour() == 23 || now.Hour() == 0 {
		t.Skip("skipping near midnight to avoid time-of-day wrapping")
	}

	past := ToTime(now.Add(-1 * time.Hour).Format(time.TimeOnly))
	future := ToTime(now.Add(1 * time.Hour).Format(time.TimeOnly))

	if past.IsPast() == false || past.IsFuture() == true {
		t.Errorf("past time %s should only be in the past", past.String())
	}

	if future.IsFuture() == false || future.IsPast() == true {
		t.Errorf("future time %s should only be in the future", future.String())
	}

}

func TestTimeOffsetExtra(t *testing.T) {

	for _, offset := range []string{"", "1234", "12:34", "+1:2", "+aa:bb"} {

		value := ToTime("12:34:56")
		value.Offset(offset)

		if value.String() != "12:34:56" {
			t.Errorf("Offset(%q) changed the time to %s", offset, value.String())
		}

	}

	value := ToTime("00:00:00")
	value.Offset("+00:01")

	if value.String() != "23:59:00" {
		t.Errorf("Offset(+00:01) = %s, expected 23:59:00", value.String())
	}

	value = ToTime("24:00:00")
	value.Offset("-01:00")

	if value.String() != "01:00:00" {
		t.Errorf("Offset(-01:00) = %s, expected 01:00:00", value.String())
	}

}

func TestTimeZulu(t *testing.T) {

	offset := strings.Split(time.Now().String(), " ")[2]
	value := ToTime("12:34:56")
	expected := ToTime("12:34:56")

	if len(offset) == 5 {
		expected.Offset(offset)
	}

	value.ToZulu()

	if value.Equal(expected) == false {
		t.Errorf("ToZulu() = %s, expected %s", value.String(), expected.String())
	}

}

func TestTimeJSON(t *testing.T) {

	original := ToTime("11:59:59")

	data, err := json.Marshal(original)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	if string(data) != `"11:59:59"` {
		t.Errorf("Marshal() = %s, expected \"11:59:59\"", string(data))
	}

	var decoded Time

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if decoded.Equal(original) == false {
		t.Errorf("Unmarshal() = %+v, expected %+v", decoded, original)
	}

	var malformed Time

	if err := json.Unmarshal([]byte("123"), &malformed); err == nil {
		t.Fatal("Unmarshal() should fail for malformed JSON")
	}

	unchanged := original

	if err := json.Unmarshal([]byte(`"25:00:00"`), &unchanged); err != nil {
		t.Fatalf("Unmarshal() of an invalid string returned error: %v", err)
	}

	if unchanged.Equal(original) == false {
		t.Fatal("Unmarshal() of an invalid string should leave the time unchanged")
	}

}
