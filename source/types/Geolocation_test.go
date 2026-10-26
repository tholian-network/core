package types

import "testing"

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
