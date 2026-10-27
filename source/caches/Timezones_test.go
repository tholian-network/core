package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupTimezones() *Timezones {

	timezones := NewTimezones()
	timezones.Add(structs.Timezone{
		Name: "Gondor/Minas-Tirith",
		Offset: "+00:00",
	})
	timezones.Add(structs.Timezone{
		Name: "Rohan/Isengard",
		Offset: "+00:00",
	})
	timezones.Add(structs.Timezone{
		Name: "Rohan/Edoras",
		Offset: "+01:30",
	})
	timezones.Add(structs.Timezone{
		Name: "Mordor/Orodruin",
		Offset: "+05:00",
	})

	return timezones

}

func TestTimezones(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		timezones := mockupTimezones()

		timezone1 := timezones.Get("South Gondor/Dol Amroth")

		if timezone1 != nil {
			t.Errorf("Expected %s to be nil", timezone1.Name)
		}

		timezones.Add(structs.Timezone{
			Name: "South Gondor/Dol Amroth",
			Offset: "+03:00",
		})

		timezone2 := timezones.Get("South Gondor/Dol Amroth")

		if timezone2 == nil {
			t.Errorf("Expected nil to be %s", "South Gondor/Dol Amroth")
		} else if timezone2.Name != "South Gondor/Dol Amroth" {
			t.Errorf("Expected %s to be %s", timezone2.Name, "South Gondor/Dol Amroth")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		timezones := mockupTimezones()

		timezone1 := timezones.Get("Rohan/Edoras")
		timezone2 := timezones.Get("Rohan/Isengard")

		if timezone1 == nil {
			t.Errorf("Expected nil to be %s", "Rohan/Edoras")
		} else if timezone1.Name != "Rohan/Edoras" {
			t.Errorf("Expected %s to be %s", timezone1.Name, "Rohan/Edoras")
		}

		if timezone2 == nil {
			t.Errorf("Expected nil to be %s", "Rohan/Isengard")
		} else if timezone2.Name != "Rohan/Isengard" {
			t.Errorf("Expected %s to be %s", timezone2.Name, "Rohan/Isengard")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		timezones := mockupTimezones()

		found1 := timezones.Query(matchers.Timezone{
			Name: "Mordor/Orodruin",
			Offset: "any",
		})

		found2 := timezones.Query(matchers.Timezone{
			Name: "Rohan/*",
			Offset: "any",
		})

		found3 := timezones.Query(matchers.Timezone{
			Name: "any",
			Offset: "+01:30",
		})

		if len(found1) == 1 {

			if found1[0].Name != "Mordor/Orodruin" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "Mordor/Orodruin")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=Mordor/Orodruin")
		}

		if len(found2) == 2 {

			if found2[0].Name != "Rohan/Edoras" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "Rohan/Edoras")
			}

			if found2[1].Name != "Rohan/Isengard" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "Rohan/Isengard")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Name=Rohan/*")
		}

		if len(found3) == 1 {

			if found3[0].Name != "Rohan/Edoras" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "Rohan/Edoras")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Offset=+01:30")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		timezones := mockupTimezones()

		timezone1 := timezones.Get("Rohan/Edoras")

		if timezone1 == nil {
			t.Errorf("Expected nil to be %s", "Rohan/Edoras")
		} else if timezone1.Name != "Rohan/Edoras" {
			t.Errorf("Expected %s to be %s", timezone1.Name, "Rohan/Edoras")
		}

		timezones.Remove("Rohan/Edoras")

		timezone2 := timezones.Get("Rohan/Edoras")

		if timezone2 != nil {
			t.Errorf("Expected %s to be nil", "Rohan/Edoras")
		}

	})

}
