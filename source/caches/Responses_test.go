package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupResponses() *Responses {

	responses := NewResponses()

	responses.Add(structs.Response{
		System:   "firewall-01",
		Type:     "Permit",
		Datetime: types.ToDatetime("2025-01-02 00:01:00"),
		Connections: []matchers.Connection{
			matchers.Connection{
				Host:     "any",
				Port:     22,
				Protocol: "tcp",
				Type:     "server",
			},
		},
		Products: []matchers.Product{
			matchers.Product{
				Vendor:  "apache",
				Product: "log4j",
				Version: "2.15.0",
				Type:    "software",
			},
		},
		Weaknesses: []matchers.Weakness{
			matchers.Weakness{
				Name:   "CWE-20",
				Impact: "Read Application Data",
				Scope:  "Confidentiality",
			},
		},
	})

	responses.Add(structs.Response{
		System:   "firewall-02",
		Type:     "Forbid",
		Datetime: types.ToDatetime("2025-01-02 00:01:01"),
		Connections: []matchers.Connection{
			matchers.Connection{
				Host:     "any",
				Port:     3306,
				Protocol: "tcp",
				Type:     "server",
			},
		},
		Products: []matchers.Product{
			matchers.Product{
				Vendor:  "gnu",
				Product: "glibc",
				Version: "2.31",
				Type:    "software",
			},
		},
		Weaknesses: []matchers.Weakness{
			matchers.Weakness{
				Name:   "CWE-264",
				Impact: "Gain Privileges or Assume Identity",
				Scope:  "Access Control",
			},
		},
	})

	return responses

}

func TestResponses(t *testing.T) {

	t.Run("QueryByConnection()", func(t *testing.T) {

		responses := mockupResponses()

		found1 := responses.QueryByConnection(matchers.Connection{
			Host:     "any",
			Port:     22,
			Protocol: "tcp",
			Type:     "server",
		})

		found2 := responses.QueryByConnection(matchers.Connection{
			Host:     "any",
			Port:     3306,
			Protocol: "tcp",
			Type:     "server",
		})

		found3 := responses.QueryByConnection(matchers.Connection{
			Host:     "any",
			Port:     22,
			Protocol: "any",
			Type:     "server",
		})

		if len(found1) == 1 {

			if found1[0].System != "firewall-01" || found1[0].Type != "Permit" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[0].System, found1[0].Type, "firewall-01", "Permit")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Port=22")
		}

		if len(found2) == 1 {

			if found2[0].System != "firewall-02" || found2[0].Type != "Forbid" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found2[0].System, found2[0].Type, "firewall-02", "Forbid")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Port=3306")
		}

		if len(found3) != 1 {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Protocol=any")
		}

	})

	t.Run("QueryByProduct()", func(t *testing.T) {

		responses := mockupResponses()

		found1 := responses.QueryByProduct(matchers.Product{
			Vendor:  "apache",
			Product: "any",
			Version: "any",
			Type:    "any",
		})

		found2 := responses.QueryByProduct(matchers.Product{
			Vendor:  "any",
			Product: "glibc",
			Version: "any",
			Type:    "any",
		})

		if len(found1) == 1 {

			if found1[0].System != "firewall-01" || found1[0].Type != "Permit" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[0].System, found1[0].Type, "firewall-01", "Permit")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Vendor=apache")
		}

		if len(found2) == 1 {

			if found2[0].System != "firewall-02" || found2[0].Type != "Forbid" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found2[0].System, found2[0].Type, "firewall-02", "Forbid")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Product=glibc")
		}

	})

	t.Run("QueryByWeakness()", func(t *testing.T) {

		responses := mockupResponses()

		found1 := responses.QueryByWeakness(matchers.Weakness{
			Name:   "CWE-20",
			Impact: "any",
			Scope:  "any",
		})

		found2 := responses.QueryByWeakness(matchers.Weakness{
			Name:   "CWE-264",
			Impact: "Gain Privileges or Assume Identity",
			Scope:  "Access Control",
		})

		found3 := responses.QueryByWeakness(matchers.Weakness{
			Name:   "any",
			Impact: "any",
			Scope:  "any",
		})

		if len(found1) == 1 {

			if found1[0].System != "firewall-01" || found1[0].Type != "Permit" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[0].System, found1[0].Type, "firewall-01", "Permit")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=CWE-20")
		}

		if len(found2) == 1 {

			if found2[0].System != "firewall-02" || found2[0].Type != "Forbid" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found2[0].System, found2[0].Type, "firewall-02", "Forbid")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Name=CWE-264")
		}

		if len(found3) != 2 {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 2, "Name=any")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		responses := mockupResponses()

		responses.Remove("firewall-01:Permit:2025-01-02 00:01:00")

		found1 := responses.QueryByConnection(matchers.Connection{
			Host:     "any",
			Port:     22,
			Protocol: "tcp",
			Type:     "server",
		})

		found2 := responses.QueryByProduct(matchers.Product{
			Vendor:  "apache",
			Product: "any",
			Version: "any",
			Type:    "any",
		})

		found3 := responses.QueryByWeakness(matchers.Weakness{
			Name:   "CWE-20",
			Impact: "any",
			Scope:  "any",
		})

		if len(found1) != 0 {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 0, "Port=22")
		}

		if len(found2) != 0 {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 0, "Vendor=apache")
		}

		if len(found3) != 0 {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 0, "Name=CWE-20")
		}

	})

}
