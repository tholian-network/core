package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupProviders() *Providers {

	providers := NewProviders()
	providers.Add(structs.Provider{
		Name: "amazon",
		Address: []string{
			"123 Corruption Street",
			"12345 Praxis City",
			"Praxis Nation Oblast of the United States Emirates",
		},
		Country: "US",
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name:    "AS16509",
				Country: "US",
				Address: "3.3.24.0",
				Scope:   "public",
				Type:    "ipv4",
				Prefix:  22,
			},
		},
	})
	providers.Add(structs.Provider{
		Name: "google",
		Address: []string{
			"1600 Amphitheatre Parkway",
			"94043 Mountain View",
			"CA",
		},
		Country: "US",
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name:    "AS15169",
				Country: "US",
				Address: "8.8.4.0",
				Scope:   "public",
				Type:    "ipv4",
				Prefix:  24,
			},
			structs.Subnet{
				Name:    "AS15169",
				Country: "US",
				Address: "8.8.8.0",
				Scope:   "public",
				Type:    "ipv4",
				Prefix:  24,
			},
		},
	})
	providers.Add(structs.Provider{
		Name: "ovh",
		Address: []string{
			"OVH SAS",
			"2 rue Kellermann",
			"59100 Roubaix",
		},
		Country: "FR",
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name:    "AS16276",
				Country: "FR",
				Address: "213.186.32.0",
				Scope:   "public",
				Type:    "ipv4",
				Prefix:  19,
			},
		},
	})

	return providers

}

func TestProviders(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		providers := mockupProviders()

		provider1 := providers.Get("hetzner")

		if provider1 != nil {
			t.Errorf("Expected %s to be nil", provider1.Name)
		}

		providers.Add(structs.Provider{
			Name: "hetzner",
			Address: []string{
				"Industriestrasse 25",
				"91710 Gunzenhausen",
			},
			Country: "DE",
			Subnets: []structs.Subnet{
				structs.Subnet{
					Name:    "AS24940",
					Country: "DE",
					Address: "116.202.0.0",
					Scope:   "public",
					Type:    "ipv4",
					Prefix:  16,
				},
			},
		})

		provider2 := providers.Get("hetzner")

		if provider2 == nil {
			t.Errorf("Expected nil to be %s", "hetzner")
		} else if provider2.Name != "hetzner" {
			t.Errorf("Expected %s to be %s", provider2.Name, "hetzner")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		providers := mockupProviders()

		provider1 := providers.Get("amazon")
		provider2 := providers.Get("google")

		if provider1 == nil {
			t.Errorf("Expected nil to be %s", "amazon")
		} else if provider1.Name != "amazon" {
			t.Errorf("Expected %s to be %s", provider1.Name, "amazon")
		}

		if provider2 == nil {
			t.Errorf("Expected nil to be %s", "google")
		} else if provider2.Name != "google" {
			t.Errorf("Expected %s to be %s", provider2.Name, "google")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		providers := mockupProviders()

		found1 := providers.Query(matchers.Provider{
			Name: "amazon",
			Country: "any",
			Subnet: "any",
		})

		found2 := providers.Query(matchers.Provider{
			Name: "any",
			Country: "US",
			Subnet: "any",
		})

		found3 := providers.Query(matchers.Provider{
			Name: "any",
			Country: "any",
			Subnet: "213.186.32.123/32",
		})

		if len(found1) == 1 {

			if found1[0].Name != "amazon" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "amazon")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=amazon")
		}

		if len(found2) == 2 {

			if found2[0].Name != "amazon" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "amazon")
			}

			if found2[1].Name != "google" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "google")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Country=US")
		}

		if len(found3) == 1 {

			if found3[0].Name != "ovh" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "ovh")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Subnet=213.186.32.123/32")
		}

	})

	t.Run("QueryByIP()", func(t *testing.T) {

		providers := mockupProviders()

		found1 := providers.QueryByIP("3.3.24.25")
		found2 := providers.QueryByIP("8.8.8.9")

		if len(found1) == 1 {

			if found1[0].Name != "amazon" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "amazon")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "3.3.24.25")
		}

		if len(found2) == 1 {

			if found2[0].Name != "google" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "google")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "8.8.8.9")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		providers := mockupProviders()

		provider1 := providers.Get("amazon")

		if provider1 == nil {
			t.Errorf("Expected nil to be %s", "amazon")
		} else if provider1.Name != "amazon" {
			t.Errorf("Expected %s to be %s", provider1.Name, "amazon")
		}

		providers.Remove("amazon")

		provider2 := providers.Get("amazon")

		if provider2 != nil {
			t.Errorf("Expected %s to be nil", provider2.Name)
		}

	})

}
