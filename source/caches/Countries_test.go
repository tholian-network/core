package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupCountries() *Countries {

	countries := NewCountries()
	countries.Add(structs.Country{
		ISO:         "DE",
		Name:        "Germany",
		Continent:   "Europe",
		Geolocation: types.ToGeolocation("geo:51.0,9.0"),
		Population:  79127551,
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name:    "AS136264",
				Country: "DE",
				Address: "45.112.84.0",
				Scope:   "public",
				Type:    "ipv4",
				Prefix:  22,
			},
		},
		Allegiances: []string{
			"NATO",
		},
		Timezones: []structs.Timezone{
			structs.NewTimezone("Europe/Berlin", "+02:00"),
			structs.NewTimezone("Europe/Busingen", "+02:00"),
		},
		Registry: nil,
	})
	countries.Add(structs.Country{
		ISO:         "RU",
		Name:        "Russia",
		Continent:   "Europe",
		Geolocation: types.ToGeolocation("geo:60.0,100.0"),
		Population:  124630000,
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name:    "AS210063",
				Country: "RU",
				Address: "2.58.212.0",
				Scope:   "public",
				Type:    "ipv4",
				Prefix:  24,
			},
		},
		Allegiances: []string{
			"SCO",
		},
		Timezones: []structs.Timezone{
			structs.NewTimezone("Europe/Moscow", "+03:00"),
			structs.NewTimezone("Europe/Kaliningrad", "+02:00"),
		},
		Registry: nil,
	})

	return countries

}

func TestCountries(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		countries := mockupCountries()

		country1 := countries.Get("PL")

		if country1 != nil {
			t.Errorf("Expected %s to be nil", country1.Name)
		}

		countries.Add(structs.Country{
			ISO:         "PL",
			Name:        "Poland",
			Continent:   "Europe",
			Geolocation: types.ToGeolocation("geo:52.0,20.0"),
			Population:  34697848,
			Subnets: []structs.Subnet{
				structs.ToSubnet("5.206.240.0/20"),
			},
			Allegiances: []string{
				"NATO",
			},
			Registry: nil,
		})

		country2 := countries.Get("PL")

		if country2 == nil {
			t.Errorf("Expected nil to be %s", "Poland")
		} else if country2.Name != "Poland" {
			t.Errorf("Expected %s to be %s", country2.Name, "Poland")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		countries := mockupCountries()

		country1 := countries.Get("DE")
		country2 := countries.Get("RU")

		if country1 == nil {
			t.Errorf("Expected nil to be %s", "Germany")
		} else if country1.Name != "Germany" {
			t.Errorf("Expected %s to be %s", country1.Name, "Germany")
		}

		if country2 == nil {
			t.Errorf("Expected nil to be %s", "Russia")
		} else if country2.Name != "Russia" {
			t.Errorf("Expected %s to be %s", country2.Name, "Russia")
		}

	})

	t.Run("GetByName()", func(t *testing.T) {

		countries := mockupCountries()

		country1 := countries.GetByName("Germany")
		country2 := countries.GetByName("Russia")

		if country1 == nil {
			t.Errorf("Expected nil to be %s", "Germany")
		} else if country1.Name != "Germany" {
			t.Errorf("Expected %s to be %s", country1.Name, "Germany")
		}

		if country2 == nil {
			t.Errorf("Expected nil to be %s", "Russia")
		} else if country2.Name != "Russia" {
			t.Errorf("Expected %s to be %s", country2.ISO, "Russia")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		countries := mockupCountries()

		found1 := countries.Query(matchers.Country{
			Name: "Russia",
			Continent: "any",
			Allegiance: "any",
			Subnet: "any",
			Timezone: "any",
		})

		found2 := countries.Query(matchers.Country{
			Name: "any",
			Continent: "Europe",
			Allegiance: "any",
			Subnet: "any",
			Timezone: "any",
		})

		found3 := countries.Query(matchers.Country{
			Name: "any",
			Continent: "any",
			Allegiance: "NATO",
			Subnet: "any",
			Timezone: "any",
		})

		found4 := countries.Query(matchers.Country{
			Name: "any",
			Continent: "any",
			Allegiance: "any",
			Subnet: "2.58.212.123/32",
			Timezone: "any",
		})

		found5 := countries.Query(matchers.Country{
			Name: "any",
			Continent: "any",
			Allegiance: "any",
			Subnet: "any",
			Timezone: "Europe/Berlin",
		})

		if len(found1) == 1 {

			if found1[0].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "Russia")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=Russia")
		}

		if len(found2) == 2 {

			if found2[0].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "Germany")
			}

			if found2[1].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "Russia")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Continent=Europe")
		}

		if len(found3) == 1 {

			if found3[0].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "Germany")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Allegiance=NATO")
		}

		if len(found4) == 1 {

			if found4[0].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "Russia")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 1, "Subnet=2.58.212.123/32")
		}

		if len(found5) == 1 {

			if found5[0].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found5[0].Name, "Germany")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found5), 1, "Timezone=Europe/Berlin")
		}

	})

	t.Run("QueryByGeolocation()", func(t *testing.T) {

		countries := mockupCountries()

		found1 := countries.QueryByGeolocation("geo:55.6,54.5")
		found2 := countries.QueryByGeolocation("geo:55.5,54.6")
		found3 := countries.QueryByGeolocation("geo:55.5,50.0")
		found4 := countries.QueryByGeolocation("geo:45.0,54.5")

		if len(found1) >= 2 {

			if found1[0].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "Russia")
			}

			if found1[1].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found1[1].Name, "Germany")
			}

		} else {
			t.Errorf("Expected more than %d", len(found1))
		}

		if len(found2) >= 2 {

			if found2[0].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "Russia")
			}

			if found2[1].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "Germany")
			}

		} else {
			t.Errorf("Expected more than %d", len(found2))
		}

		if len(found3) >= 2 {

			if found3[0].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "Germany")
			}

			if found3[1].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found3[1].Name, "Russia")
			}

		} else {
			t.Errorf("Expected more than %d", len(found2))
		}

		if len(found4) >= 2 {

			if found4[0].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "Germany")
			}

			if found4[1].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found4[1].Name, "Russia")
			}

		} else {
			t.Errorf("Expected more than %d", len(found2))
		}

	})

	t.Run("QueryByIP()", func(t *testing.T) {

		countries := mockupCountries()

		found1 := countries.QueryByIP("45.112.84.123")
		found2 := countries.QueryByIP("2.58.212.123")

		if len(found1) == 1 {

			if found1[0].Name != "Germany" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "Germany")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "45.112.84.123")
		}

		if len(found2) == 1 {

			if found2[0].Name != "Russia" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "Russia")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "2.58.212.123")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		countries := mockupCountries()

		country1 := countries.Get("RU")

		if country1 == nil {
			t.Errorf("Expected nil to be %s", "Russia")
		} else if country1.Name != "Russia" {
			t.Errorf("Expected %s to be %s", country1.Name, "Russia")
		}

		countries.Remove("RU")

		country2 := countries.Get("RU")

		if country2 != nil {
			t.Errorf("Expected %s to be nil", country2.Name)
		}

	})

}
