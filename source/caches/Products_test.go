package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupProducts() *Products {

	products := NewProducts()
	products.Add(structs.Product{
		Name:    "Microsoft:Internet Explorer:7",
		Vendor:  "Microsoft",
		Product: "Internet Explorer",
		Version: "7",
		Type:    "software",
		State:   "published",
	})
	products.Add(structs.Product{
		Name:    "Microsoft:Internet Explorer:8",
		Vendor:  "Microsoft",
		Product: "Internet Explorer",
		Version: "8",
		Type:    "software",
		State:   "published",
	})
	products.Add(structs.Product{
		Name:    "Microsoft:Windows:7",
		Vendor:  "Microsoft",
		Product: "Windows",
		Version: "7",
		Type:    "software",
		State:   "published",
	})
	products.Add(structs.Product{
		Name:    "Microsoft:Windows:8",
		Vendor:  "Microsoft",
		Product: "Windows",
		Version: "8",
		Type:    "software",
		State:   "published",
	})
	products.Add(structs.Product{
		Name:    "Microsoft:Windows:10",
		Vendor:  "Microsoft",
		Product: "Windows",
		Version: "10",
		Type:    "software",
		State:   "published",
	})
	products.Add(structs.Product{
		Name:    "Canonical:Ubuntu:6",
		Vendor:  "Canonical",
		Product: "Ubuntu",
		Version: "6",
		Type:    "software",
		State:   "published",
	})

	return products

}

func TestProducts(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		products := mockupProducts()

		product1 := products.Get("Red Hat:RHEL:8")

		if product1 != nil {
			t.Errorf("Expected %s to be nil", product1.Name)
		}

		products.Add(structs.Product{
			Name:    "Red Hat:RHEL:8",
			Vendor:  "Red Hat",
			Product: "RHEL",
			Version: "8",
			Type:    "software",
			State:   "published",
		})

		product2 := products.Get("Red Hat:RHEL:8")

		if product2 == nil {
			t.Errorf("Expected nil to be %s", "Red Hat:RHEL:8")
		} else if product2.Name != "Red Hat:RHEL:8" {
			t.Errorf("Expected %s to be %s", product2.Name, "Red Hat:RHEL:8")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		products := mockupProducts()

		product1 := products.Get("Microsoft:Internet Explorer:7")
		product2 := products.Get("Canonical:Ubuntu:6")

		if product1 == nil {
			t.Errorf("Expected nil to be %s", "Microsoft:Internet Explorer:7")
		} else if product1.Name != "Microsoft:Internet Explorer:7" {
			t.Errorf("Expected %s to be %s", product1.Name, "Microsoft:Internet Explorer:7")
		}

		if product2 == nil {
			t.Errorf("Expected nil to be %s", "Canonical:Ubuntu:6")
		} else if product2.Name != "Canonical:Ubuntu:6" {
			t.Errorf("Expected %s to be %s", product2.Name, "Canonical:Ubuntu:6")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		products := mockupProducts()

		found1 := products.Query(matchers.Product{
			Product: "Internet Explorer",
			Vendor: "any",
			Version: "any",
			Type: "software",
		})

		found2 := products.Query(matchers.Product{
			Product: "any",
			Vendor: "Microsoft",
			Version: "any",
			Type: "software",
		})

		found3 := products.Query(matchers.Product{
			Product: "any",
			Vendor: "any",
			Version: "10",
			Type: "software",
		})

		found4 := products.Query(matchers.Product{
			Product: "any",
			Vendor: "any",
			Version: "any",
			Type: "software",
		})

		if len(found1) == 2 {

			if found1[0].Name != "Microsoft:Internet Explorer:7" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "Microsoft:Internet Explorer:7")
			}

			if found1[1].Name != "Microsoft:Internet Explorer:8" {
				t.Errorf("Expected %s to be %s", found1[1].Name, "Microsoft:Internet Explorer:8")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Product=Internet Explorer")
		}

		if len(found2) == 5 {

			if found2[0].Name != "Microsoft:Internet Explorer:7" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "Microsoft:Internet Explorer:7")
			}

			if found2[1].Name != "Microsoft:Internet Explorer:8" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "Microsoft:Internet Explorer:8")
			}

			if found2[2].Name != "Microsoft:Windows:7" {
				t.Errorf("Expected %s to be %s", found2[2].Name, "Microsoft:Windows:7")
			}

			if found2[3].Name != "Microsoft:Windows:8" {
				t.Errorf("Expected %s to be %s", found2[3].Name, "Microsoft:Windows:8")
			}

			if found2[4].Name != "Microsoft:Windows:10" {
				t.Errorf("Expected %s to be %s", found2[4].Name, "Microsoft:Windows:10")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 5, "Vendor=Microsoft")
		}

		if len(found3) == 1 {

			if found3[0].Name != "Microsoft:Windows:10" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "Microsoft:Windows:10")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Version=10")
		}

		if len(found4) == 6 {

			if found4[0].Name != "Canonical:Ubuntu:6" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "Canonical:Ubuntu:6")
			}

			if found4[1].Name != "Microsoft:Internet Explorer:7" {
				t.Errorf("Expected %s to be %s", found4[1].Name, "Microsoft:Internet Explorer:7")
			}

			if found4[2].Name != "Microsoft:Internet Explorer:8" {
				t.Errorf("Expected %s to be %s", found4[2].Name, "Microsoft:Internet Explorer:8")
			}

			if found4[3].Name != "Microsoft:Windows:7" {
				t.Errorf("Expected %s to be %s", found4[3].Name, "Microsoft:Windows:7")
			}

			if found4[4].Name != "Microsoft:Windows:8" {
				t.Errorf("Expected %s to be %s", found4[4].Name, "Microsoft:Windows:8")
			}

			if found4[5].Name != "Microsoft:Windows:10" {
				t.Errorf("Expected %s to be %s", found4[5].Name, "Microsoft:Windows:10")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 6, "Type=software")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		products := mockupProducts()

		product1 := products.Get("Microsoft:Internet Explorer:7")

		if product1 == nil {
			t.Errorf("Expected nil to be %s", "Microsoft:Internet Explorer:7")
		} else if product1.Name != "Microsoft:Internet Explorer:7" {
			t.Errorf("Expected %s to be %s", product1.Name, "Microsoft:Internet Explorer:7")
		}

		products.Remove("Microsoft:Internet Explorer:7")

		product2 := products.Get("Microsoft:Internet Explorer:7")

		if product2 != nil {
			t.Errorf("Expected %s to be nil", product2.Name)
		}

	})

}
