package types

import (
	"encoding/json"
	"math"
	"testing"
)

func TestGeolocation(t *testing.T) {

	t.Run("IsGeolocation()", func(t *testing.T) {

		is1 := IsGeolocation("geo:52.5200,13.4050")  // Berlin
		is2 := IsGeolocation("geo:38.8951,-77.0364") // Washington D.C.
		is3 := IsGeolocation("geo:55.7558,37.6176")  // Moscow
		is4 := IsGeolocation("geo:39.9042,116.4074") // Beijing
		is5 := IsGeolocation("geo:35.6892,51.3890")  // Tehran

		if is1 != true {
			t.Errorf("Expected %t to be true", is1)
		}

		if is2 != true {
			t.Errorf("Expected %t to be true", is2)
		}

		if is3 != true {
			t.Errorf("Expected %t to be true", is3)
		}

		if is4 != true {
			t.Errorf("Expected %t to be true", is4)
		}

		if is5 != true {
			t.Errorf("Expected %t to be true", is5)
		}

	})

	t.Run("ParseGeolocation()", func(t *testing.T) {

		geolocation1 := ParseGeolocation("geo:52.5200,13.4050") // Berlin
		geolocation2 := ParseGeolocation("38.8951,-77.0364")    // Washington D.C.
		geolocation3 := ParseGeolocation("35.6892:51.3890")     // invalid

		if geolocation1 == nil {
			t.Errorf("Expected %s to be valid", "geo:52.520000,13.405000")
		} else if geolocation1.String() != "geo:52.520000,13.405000" {
			t.Errorf("Expected %s to be %s", geolocation1.String(), "geo:52.520000,13.405000")
		}

		if geolocation2 == nil {
			t.Errorf("Expected %s to be valid", "38.895100,-77.036400")
		} else if geolocation2.String() != "geo:38.895100,-77.036400" {
			t.Errorf("Expected %s to be %s", geolocation2.String(), "geo:38.895100,-77.036400")
		}

		if geolocation3 != nil {
			t.Errorf("Expected %s to be invalid", "35.6892:51.3890")
		}

	})

	t.Run("ToGeolocation()", func(t *testing.T) {

		geolocation1 := ToGeolocation("geo:52.5200,13.4050") // Berlin
		geolocation2 := ToGeolocation("38.8951,-77.0364")    // Washington D.C.
		geolocation3 := ToGeolocation("35.6892:51.3890")     // invalid

		if geolocation1.IsValid() != true {
			t.Errorf("Expected %s to be valid", geolocation1.String())
		}

		if geolocation2.IsValid() != true {
			t.Errorf("Expected %s to be valid", geolocation2.String())
		}

		if geolocation3.IsValid() != false {
			t.Errorf("Expected %s to be invalid", geolocation3.String())
		}

	})

	t.Run("DistanceTo()", func(t *testing.T) {

		geolocation1 := ToGeolocation("geo:52.5200,13.4050") // Berlin
		geolocation2 := ToGeolocation("38.8951,-77.0364")    // Washington D.C.
		geolocation3 := ToGeolocation("geo:55.7558,37.6176") // Moscow

		distance1 := geolocation1.DistanceTo(geolocation2)
		distance2 := geolocation2.DistanceTo(geolocation1)
		distance3 := geolocation3.DistanceTo(geolocation1)
		distance4 := geolocation3.DistanceTo(geolocation2)

		if distance1 != 6711.1391083171275 {
			t.Errorf("Expected %.9f to be %.9f km", distance1, 6711.1391083171275)
		}

		if distance2 != 6711.1391083171275 {
			t.Errorf("Expected %.9f to be %.9f km", distance2, 6711.1391083171275)
		}

		if distance3 != 1608.8500845802332 {
			t.Errorf("Expected %.9f to be %.9f km", distance3, 1608.8500845802332)
		}

		if distance4 != 7821.9965986416655 {
			t.Errorf("Expected %.9f to be %.9f km", distance4, 7821.9965986416655)
		}

	})

	t.Run("SetLatitude()", func(t *testing.T) {

		geolocation1 := ToGeolocation("geo:0.0,0.0")
		geolocation2 := ToGeolocation("geo:0.0,0.0")
		geolocation3 := ToGeolocation("geo:0.0,0.0")

		result1 := geolocation1.SetLatitude(-89.999999999)
		result2 := geolocation2.SetLatitude(89.999999999)
		result3 := geolocation3.SetLatitude(-137.1337)

		if result1 != true {
			t.Errorf("Expected %t to be %t", result1, true)
		} else if geolocation1.Latitude != -89.999999999 {
			t.Errorf("Expected %.9f to be %.9f", geolocation1.Latitude, -89.999999999)
		}

		if result2 != true {
			t.Errorf("Expected %t to be %t", result2, true)
		} else if geolocation2.Latitude != 89.999999999 {
			t.Errorf("Expected %.9f to be %.9f", geolocation2.Latitude, 89.999999999)
		}

		if result3 != false {
			t.Errorf("Expected %t to be %t", result3, false)
		} else if geolocation3.Latitude != 0.0 {
			t.Errorf("Expected %.9f to be %.9f", geolocation3.Latitude, 0.0)
		}

	})

	t.Run("SetLongitude()", func(t *testing.T) {

		geolocation1 := ToGeolocation("geo:0.0,0.0")
		geolocation2 := ToGeolocation("geo:0.0,0.0")
		geolocation3 := ToGeolocation("geo:0.0,0.0")

		result1 := geolocation1.SetLongitude(-179.999999)
		result2 := geolocation2.SetLongitude(179.999999)
		result3 := geolocation3.SetLongitude(-1337.1337)

		if result1 != true {
			t.Errorf("Expected %t to be %t", result1, true)
		} else if geolocation1.Longitude != -179.999999 {
			t.Errorf("Expected %.9f to be %.9f", geolocation1.Longitude, -179.999999)
		}

		if result2 != true {
			t.Errorf("Expected %t to be %t", result2, true)
		} else if geolocation2.Longitude != 179.999999 {
			t.Errorf("Expected %.9f to be %.9f", geolocation2.Longitude, 179.999999)
		}

		if result3 != false {
			t.Errorf("Expected %t to be %t", result3, false)
		} else if geolocation3.Longitude != 0.0 {
			t.Errorf("Expected %.9f to be %.9f", geolocation3.Longitude, 0.0)
		}

	})

	t.Run("String()", func(t *testing.T) {

		geolocation1 := ParseGeolocation("geo:13.37,13.337")
		geolocation2 := ParseGeolocation("geo:-89.999999,-179.999999")
		geolocation3 := ParseGeolocation("geo:89.999999,179.999999")

		geolocation4 := ParseGeolocation("geo:-89.923432,-179.923432")
		geolocation5 := ParseGeolocation("geo:89.923432,179.923432")

		if geolocation1.String() != "geo:13.370000,13.337000" {
			t.Errorf("Expected %s to be %s", geolocation1.String(), "geo:13.370000,13.337000")
		}

		if geolocation2.String() != "geo:-89.999999,-179.999999" {
			t.Errorf("Expected %s to be %s", geolocation2.String(), "geo:-89.999999,-179.999999")
		}

		if geolocation3.String() != "geo:89.999999,179.999999" {
			t.Errorf("Expected %s to be %s", geolocation3.String(), "geo:89.999999,179.999999")
		}

		if geolocation4.String() != "geo:-89.923432,-179.923432" {
			t.Errorf("Expected %s to be %s", geolocation4.String(), "geo:-89.923432,-179.923432")
		}

		if geolocation5.String() != "geo:89.923432,179.923432" {
			t.Errorf("Expected %s to be %s", geolocation5.String(), "geo:89.923432,179.923432")
		}

	})

}

func TestGeolocationNew(t *testing.T) {

	geolocation := NewGeolocation()

	if geolocation.Latitude != -90.0 || geolocation.Longitude != 0.0 {
		t.Fatalf("NewGeolocation() = %+v, expected Antarctica", geolocation)
	}

	if geolocation.String() != "geo:-90.000000,0.000000" {
		t.Errorf("NewGeolocation().String() = %q", geolocation.String())
	}

	if geolocation.IsValid() == true {
		t.Fatal("the default -90 latitude is outside the strict validation bounds")
	}

}

func TestGeolocationValidation(t *testing.T) {

	cases := []struct {
		value string
		valid bool
	}{
		{"geo:52.5200,13.4050", true},
		{"38.8951,-77.0364", true},
		{"0,0", true},
		{"89.999,-179.999", true},
		{"-89.999,179.999", true},
		{"geo:90,0", false},
		{"geo:-90,0", false},
		{"geo:0,180", false},
		{"geo:0,-180", false},
		{"geo:91,0", false},
		{"geo:0,181", false},
		{"geo:abc,0", false},
		{"geo:1,2,3", false},
		{"geo:1", false},
		{"", false},
	}

	for _, c := range cases {
		if result := IsGeolocation(c.value); result != c.valid {
			t.Errorf("IsGeolocation(%q) = %v, expected %v", c.value, result, c.valid)
		}
	}

	valid := ToGeolocation("geo:52.520000,13.405000")

	if valid.IsValid() == false {
		t.Fatal("a valid geolocation should be valid")
	}

	outOfBounds := ToGeolocation("geo:200,300")

	if outOfBounds.IsValid() == true {
		t.Fatal("an out-of-range geolocation should be invalid")
	}

	latitudeBoundary := ToGeolocation("geo:90,13")

	if latitudeBoundary.IsValid() == true {
		t.Fatal("a latitude of exactly 90 should be invalid")
	}

	longitudeBoundary := ToGeolocation("geo:52,180")

	if longitudeBoundary.IsValid() == true {
		t.Fatal("a longitude of exactly 180 should be invalid")
	}

}

func TestGeolocationParseEdges(t *testing.T) {

	for _, value := range []string{"geo:52.52,13.405", "52.52,13.405"} {
		geolocation := ParseGeolocation(value)

		if geolocation == nil {
			t.Fatalf("ParseGeolocation(%q) returned nil", value)
		}

		if geolocation.Latitude != 52.52 || geolocation.Longitude != 13.405 {
			t.Errorf("ParseGeolocation(%q) = %+v", value, geolocation)
		}
	}

	outOfRange := ParseGeolocation("geo:200,300")

	if outOfRange == nil {
		t.Fatal("ParseGeolocation() should parse out-of-range coordinates")
	}

	if outOfRange.IsValid() == true {
		t.Fatal("out-of-range coordinates should not be valid")
	}

	for _, value := range []string{"", "geo:", "geo:52.52", "geo:52.52,", "geo:a,b", "52.52:13.405"} {
		if geolocation := ParseGeolocation(value); geolocation != nil {
			t.Errorf("ParseGeolocation(%q) = %+v, expected nil", value, geolocation)
		}
	}

	empty := ToGeolocation("")

	if empty != NewGeolocation() {
		t.Fatalf("ToGeolocation(\"\") = %+v, expected the default", empty)
	}

	invalid := ToGeolocation("not a geolocation")

	if invalid != NewGeolocation() {
		t.Fatalf("ToGeolocation() with invalid input = %+v, expected the default", invalid)
	}

	parsed := ToGeolocation("geo:52.52,13.405")

	if parsed.String() != "geo:52.520000,13.405000" {
		t.Fatal("ToGeolocation() did not parse a valid value")
	}

}

func TestGeolocationDistanceToEdges(t *testing.T) {

	berlin := ToGeolocation("geo:52.5200,13.4050")
	washington := ToGeolocation("geo:38.8951,-77.0364")

	if distance := berlin.DistanceTo(berlin); distance != 0.0 {
		t.Errorf("DistanceTo(self) = %f, expected 0", distance)
	}

	forward := berlin.DistanceTo(washington)
	backward := washington.DistanceTo(berlin)

	if math.Abs(forward-backward) > 1e-9 {
		t.Errorf("DistanceTo() is not symmetric: %f vs %f", forward, backward)
	}

	if forward <= 0.0 {
		t.Fatalf("DistanceTo() = %f, expected a positive distance", forward)
	}

}

func TestGeolocationHashJSON(t *testing.T) {

	berlin := ToGeolocation("geo:52.5200,13.4050")
	berlin2 := ToGeolocation("geo:52.520000,13.405000")
	washington := ToGeolocation("geo:38.8951,-77.0364")

	if berlin.Hash() == "" || berlin.Hash() != berlin2.Hash() {
		t.Fatal("Hash() should be deterministic for equal coordinates")
	}

	if berlin.Hash() == washington.Hash() {
		t.Fatal("different coordinates should not share a hash")
	}

	data, err := json.Marshal(berlin)

	if err != nil {
		t.Fatalf("Marshal() returned error: %v", err)
	}

	if string(data) != `"geo:52.520000,13.405000"` {
		t.Errorf("Marshal() = %s", string(data))
	}

	var decoded Geolocation

	if err := json.Unmarshal(data, &decoded); err != nil {
		t.Fatalf("Unmarshal() returned error: %v", err)
	}

	if decoded.Latitude != berlin.Latitude || decoded.Longitude != berlin.Longitude {
		t.Errorf("Unmarshal() = %+v, expected %+v", decoded, berlin)
	}

	var malformed Geolocation

	if err := json.Unmarshal([]byte("123"), &malformed); err == nil {
		t.Fatal("Unmarshal() should fail for malformed JSON")
	}

	unchanged := berlin

	if err := json.Unmarshal([]byte(`"not a geolocation"`), &unchanged); err != nil {
		t.Fatalf("Unmarshal() of an invalid string returned error: %v", err)
	}

	if unchanged != berlin {
		t.Fatal("Unmarshal() of an invalid string should leave the geolocation unchanged")
	}

}
