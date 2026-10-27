package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupTimeslots() *Timeslots {

	timeslots := NewTimeslots()
	timeslots.Add(structs.Timeslot{
		Offset:     "+00:00",
		Population: 300000,
		Weight:     0.30,
		Countries:  []string{
			"Gondor",
			"Rohan",
		},
		Timezones: []string{
			"Gondor/Minas-Tirith",
			"Rohan/Isengard",
		},
	})
	timeslots.Add(structs.Timeslot{
		Offset:     "+01:30",
		Population: 200000,
		Weight:     0.20,
		Countries:  []string{
			"Rohan",
		},
		Timezones: []string{
			"Rohan/Edoras",
		},
	})
	timeslots.Add(structs.Timeslot{
		Offset:     "+05:00",
		Population: 500000,
		Weight:     0.50,
		Countries:  []string{
			"Mordor",
		},
		Timezones: []string{
			"Mordor/Orodruin",
			"Mordor/Barad-Dur",
			"Mordor/Nurn",
		},
	})

	return timeslots

}

func TestTimeslots(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		timeslots := mockupTimeslots()

		timeslot1 := timeslots.Get("+03:00")

		if timeslot1 != nil {
			t.Errorf("Expected %s to be nil", timeslot1.Offset)
		}

		timeslots.Add(structs.Timeslot{
			Offset:     "+03:00",
			Population: 1000,
			Weight:     0.001,
			Countries:  []string{
				"South Gondor",
			},
			Timezones: []string{
				"South Gondor/Dol Amroth",
			},
		})

		timeslot2 := timeslots.Get("+03:00")

		if timeslot2 == nil {
			t.Errorf("Expected nil to be %s", "+03:00")
		} else if timeslot2.Offset != "+03:00" {
			t.Errorf("Expected %s to be %s", timeslot2.Offset, "+03:00")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		timeslots := mockupTimeslots()

		timeslot1 := timeslots.Get("+00:00")
		timeslot2 := timeslots.Get("+05:00")

		if timeslot1 == nil {
			t.Errorf("Expected nil to be %s", "+00:00")
		} else if timeslot1.Offset != "+00:00" {
			t.Errorf("Expected %s to be %s", timeslot1.Offset, "+00:00")
		}

		if timeslot2 == nil {
			t.Errorf("Expected nil to be %s", "+05:00")
		} else if timeslot2.Offset != "+05:00" {
			t.Errorf("Expected %s to be %s", timeslot2.Offset, "+05:00")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		timeslots := mockupTimeslots()

		found1 := timeslots.Query(matchers.Timeslot{
			Offset: "+00:00",
			Country: "any",
			Timezone: "any",
		})

		found2 := timeslots.Query(matchers.Timeslot{
			Offset: "any",
			Country: "Rohan",
			Timezone: "any",
		})

		found3 := timeslots.Query(matchers.Timeslot{
			Offset: "any",
			Country: "any",
			Timezone: "Mordor/Orodruin",
		})

		if len(found1) == 1 {

			if found1[0].Offset != "+00:00" {
				t.Errorf("Expected %s to be %s", found1[0].Offset, "+00:00")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Offset=+00:00")
		}

		if len(found2) == 2 {

			if found2[0].Offset != "+00:00" {
				t.Errorf("Expected %s to be %s", found2[0].Offset, "+00:00")
			}

			if found2[1].Offset != "+01:30" {
				t.Errorf("Expected %s to be %s", found2[1].Offset, "+01:30")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Country=Rohan")
		}

		if len(found3) == 1 {

			if found3[0].Offset != "+05:00" {
				t.Errorf("Expected %s to be %s", found3[0].Offset, "+05:00")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Timezone=Mordor/Orodruin")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		timeslots := mockupTimeslots()

		timeslot1 := timeslots.Get("+00:00")

		if timeslot1 == nil {
			t.Errorf("Expected nil to be %s", "+00:00")
		} else if timeslot1.Offset != "+00:00" {
			t.Errorf("Expected nil to be %s", "+00:00")
		}

		timeslots.Remove("+00:00")

		timeslot2 := timeslots.Get("+00:00")

		if timeslot2 != nil {
			t.Errorf("Expected %s to be nil", timeslot2.Offset)
		}

	})

}

