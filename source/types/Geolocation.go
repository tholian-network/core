package types

import "math"
import "strconv"
import "strings"

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func IsGeolocation(value string) bool {

	if strings.HasPrefix(value, "geo:") {

		if strings.Contains(value, ",") {

			tmp := strings.Split(value[4:], ",")

			if len(tmp) == 2 {

				num1, err1 := strconv.ParseFloat(tmp[0], 64)
				num2, err2 := strconv.ParseFloat(tmp[1], 64)

				if err1 == nil && num1 > -90.0 && num1 < 90.0 {

					if err2 == nil && num2 > -180.0 && num2 < 180.0 {
						return true
					}

				}

			}

		}

	} else {

		if strings.Contains(value, ",") {

			tmp := strings.Split(value, ",")

			if len(tmp) == 2 {

				num1, err1 := strconv.ParseFloat(tmp[0], 64)
				num2, err2 := strconv.ParseFloat(tmp[1], 64)

				if err1 == nil && num1 > -90.0 && num1 < 90.0 {

					if err2 == nil && num2 > -180.0 && num2 < 180.0 {
						return true
					}

				}

			}

		}

	}

	return false

}

func NewGeolocation() Geolocation {

	var geolocation Geolocation

	// Antarctica
	geolocation.Latitude = -90.0
	geolocation.Longitude = 0.0

	return geolocation

}

func ParseGeolocation(value string) *Geolocation {

	var result *Geolocation = nil

	if strings.HasPrefix(value, "geo:") {

		if strings.Contains(value, ",") {

			tmp1 := strings.Split(value[4:], ",")

			if len(tmp1) == 2 {

				num1, err1 := strconv.ParseFloat(tmp1[0], 64)
				num2, err2 := strconv.ParseFloat(tmp1[1], 64)

				if err1 == nil && err2 == nil {

					geolocation := Geolocation{
						Latitude:  num1,
						Longitude: num2,
					}

					result = &geolocation

				}

			}

		}

	} else {

		if strings.Contains(value, ",") {

			tmp1 := strings.Split(value, ",")

			if len(tmp1) == 2 {

				num1, err1 := strconv.ParseFloat(tmp1[0], 64)
				num2, err2 := strconv.ParseFloat(tmp1[1], 64)

				if err1 == nil && err2 == nil {

					geolocation := Geolocation{
						Latitude:  num1,
						Longitude: num2,
					}

					result = &geolocation

				}

			}

		}

	}

	return result

}

func ToGeolocation(value string) Geolocation {

	var geolocation Geolocation

	// Antarctica
	geolocation.Latitude = -90.0
	geolocation.Longitude = 0.0

	if value != "" {

		tmp := ParseGeolocation(value)

		if tmp != nil {
			geolocation = *tmp
		}

	}

	return geolocation

}

func (geolocation *Geolocation) DistanceTo(other Geolocation) float64 {

	lat1_radian := geolocation.Latitude * (math.Pi / 180.0)
	lon1_radian := geolocation.Longitude * (math.Pi / 180.0)
	lat2_radian := other.Latitude * (math.Pi / 180.0)
	lon2_radian := other.Longitude * (math.Pi / 180.0)

	lat_delta := lat2_radian - lat1_radian
	lon_delta := lon2_radian - lon1_radian

	a := math.Sin(lat_delta/2.0)*math.Sin(lat_delta/2.0) + math.Cos(lat1_radian)*math.Cos(lat2_radian)*math.Sin(lon_delta/2.0)*math.Sin(lon_delta/2.0)
	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1.0-a))

	return 6371.00 * c

}

func (geolocation *Geolocation) IsValid() bool {

	if IsGeolocation(geolocation.String()) {
		return true
	}

	return false

}

func (geolocation *Geolocation) SetLatitude(latitude float64) bool {

	if latitude > -90.0 && latitude < 90.0 {

		geolocation.Latitude = latitude

		return true

	}

	return false

}

func (geolocation *Geolocation) SetLongitude(longitude float64) bool {

	if longitude > -180.0 && longitude < 180.0 {

		geolocation.Longitude = longitude

		return true

	}

	return false

}

func (geolocation *Geolocation) String() string {

	lat_encoded := strconv.FormatFloat(geolocation.Latitude, 'f', 6, 64)
	long_encoded := strconv.FormatFloat(geolocation.Longitude, 'f', 6, 64)

	return "geo:" + lat_encoded + "," + long_encoded

}

func (geolocation Geolocation) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(geolocation.String())), nil
}

func (geolocation *Geolocation) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := ParseGeolocation(unquoted)

	if tmp != nil {
		*geolocation = *tmp
	}

	return nil

}

func (geolocation Geolocation) Hash() string {
	return toHash(geolocation.String())
}
