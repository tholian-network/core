package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupAntiques() *Antiques {

	antiques := NewAntiques()
	antiques.Add(structs.Antique{
		Name:         "glibc",
		Version:      types.ToVersion("0:1.2.3-patch4"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("apt"),
		Vendor:       "debian-bullseye",
		Service:      "apache2",
		URL:          "http://mirror.debian.org/packages/g/glibc/glibc-0:1.2.3-patch4_amd64.deb",
	})
	antiques.Add(structs.Antique{
		Name:         "glibc",
		Version:      types.ToVersion("0:2.41-r47"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("apt"),
		Vendor:       "debian-bullseye",
		Service:      "apache2",
		URL:          "http://mirror.debian.org/packages/g/glibc/glibc-0:2.41-r47_amd64.deb",
	})
	antiques.Add(structs.Antique{
		Name:         "left-pad",
		Version:      types.ToVersion("0:1.3.0"),
		Architecture: types.ToArchitecture("any"),
		Manager:      types.ToManager("npm"),
		Vendor:       "npm",
		Service:      "backend",
		URL:          "https://www.npmjs.com/package/left-pad",
	})
	antiques.Add(structs.Antique{
		Name:         "is-even",
		Version:      types.ToVersion("0:1.0.0"),
		Architecture: types.ToArchitecture("any"),
		Manager:      types.ToManager("npm"),
		Vendor:       "npm",
		Service:      "backend",
		URL:          "https://www.npmjs.com/package/is-even",
	})

	return antiques

}

func TestAntiques(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		antiques := mockupAntiques()

		antique1 := antiques.Get("npm:npm:right-pad:0:0.0.1:any")

		if antique1 != nil {
			t.Errorf("Expected %s to be nil", antique1.Name)
		}

		antiques.Add(structs.Antique{
			Name:         "right-pad",
			Version:      types.ToVersion("0:0.0.1"),
			Architecture: types.ToArchitecture("any"),
			Manager:      types.ToManager("npm"),
			Vendor:       "npm",
			Service:      "backend",
			URL:          "https://www.npmjs.com/package/right-pad",
		})

		antique2 := antiques.Get("npm:npm:right-pad:0:0.0.1:any")

		if antique2 == nil {
			t.Errorf("Expected nil to be %s", "right-pad")
		} else if antique2.Name != "right-pad" {
			t.Errorf("Expected %s to be %s", antique2.Name, "right-pad")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		antiques := mockupAntiques()

		antique1 := antiques.Get("apt:debian-bullseye:glibc:0:1.2.3~patch4:x86_64")
		antique2 := antiques.Get("npm:npm:left-pad:0:1.3.0:any")

		if antique1 == nil {
			t.Errorf("Expected nil to be %s", "glibc")
		} else if antique1.Name != "glibc" {
			t.Errorf("Expected %s to be %s", antique1.Name, "glibc")
		}

		if antique2 == nil {
			t.Errorf("Expected nil to be %s", "left-pad")
		} else if antique2.Name != "left-pad" {
			t.Errorf("Expected %s to be %s", antique2.Name, "left-pad")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		antiques := mockupAntiques()

		found1 := antiques.Query(matchers.Antique{
			Name: "glibc",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
			Service: "any",
		})

		found2 := antiques.Query(matchers.Antique{
			Name: "any",
			Version: "0:1.2.3~patch4",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
			Service: "any",
		})

		found3 := antiques.Query(matchers.Antique{
			Name: "any",
			Version: "any",
			Architecture: "x86_64",
			Manager: "any",
			Vendor: "any",
			Service: "any",
		})

		found4 := antiques.Query(matchers.Antique{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "npm",
			Vendor: "any",
			Service: "any",
		})

		found5 := antiques.Query(matchers.Antique{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "debian-bullseye",
			Service: "any",
		})

		found6 := antiques.Query(matchers.Antique{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
			Service: "backend",
		})

		if len(found1) == 2 {

			if found1[0].Name != "glibc" || found1[0].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[0].Name, found1[0].Version.String(), "glibc", "0:1.2.3~patch4")
			}

			if found1[1].Name != "glibc" || found1[1].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[1].Name, found1[1].Version.String(), "glibc", "0:2.41.0r47")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=glibc")
		}

		if len(found2) == 1 {

			if found2[0].Name != "glibc" || found2[0].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found2[0].Name, found2[0].Version.String(), "glibc", "0:1.2.3~patch4")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Version=0:1.2.3~patch4")
		}

		if len(found3) == 2 {

			if found3[0].Name != "glibc" || found3[0].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[0].Name, found3[0].Version.String(), "glibc", "0:1.2.3~patch4")
			}

			if found3[1].Name != "glibc" || found3[1].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[1].Name, found3[1].Version.String(), "glibc", "0:2.41.0r47")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 2, "Architecture=x86_64")
		}

		if len(found4) == 2 {

			if found4[0].Name != "is-even" || found4[0].Version.String() != "0:1.0.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[0].Name, found4[0].Version.String(), "is-even", "0:1.0.0")
			}

			if found4[1].Name != "left-pad" || found4[1].Version.String() != "0:1.3.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[1].Name, found4[1].Version.String(), "left-pad", "0:1.3.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 2, "Manager=npm")
		}

		if len(found5) == 2 {

			if found5[0].Name != "glibc" || found5[0].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found5[0].Name, found5[0].Version.String(), "glibc", "0:1.2.3~patch4")
			}

			if found5[1].Name != "glibc" || found5[1].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found5[1].Name, found5[1].Version.String(), "glibc", "0:2.41.0r47")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found5), 2, "Architecture=x86_64")
		}

		if len(found6) == 2 {

			if found6[0].Name != "is-even" || found6[0].Version.String() != "0:1.0.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found6[0].Name, found6[0].Version.String(), "is-even", "0:1.0.0")
			}

			if found6[1].Name != "left-pad" || found6[1].Version.String() != "0:1.3.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found6[1].Name, found6[1].Version.String(), "left-pad", "0:1.3.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found6), 2, "Service=backend")
		}

	})

	t.Run("Query() Versions", func(t *testing.T) {

		antiques := mockupAntiques()

		found1 := antiques.Query(matchers.Antique{
			Name: "glibc",
			Version: "0:1.2.3~patch4",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
			Service: "any",
		})

		found2 := antiques.Query(matchers.Antique{
			Name: "glibc",
			Version: "< 0:2.41~r47",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
			Service: "any",
		})

		found3 := antiques.Query(matchers.Antique{
			Name: "glibc",
			Version: "> 0:1.2.3~patch4",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
			Service: "any",
		})

		found4 := antiques.Query(matchers.Antique{
			Name: "glibc",
			Version: "<= 0:2.41~r47",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
			Service: "any",
		})

		if len(found1) == 1 {

			if found1[0].Name != "glibc" || found1[0].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[0].Name, found1[0].Version.String(), "glibc", "0:1.2.3~patch4")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Version=0:1.2.3~patch4")
		}

		if len(found2) == 1 {

			if found2[0].Name != "glibc" || found2[0].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found2[0].Name, found2[0].Version.String(), "glibc", "0:1.2.3~patch4")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Version=< 0:2.41~r47")
		}

		if len(found3) == 1 {

			if found3[0].Name != "glibc" || found3[0].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[0].Name, found3[0].Version.String(), "glibc", "0:2.41.0r47")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Version=> 0:1.2.3~patch4")
		}

		if len(found4) == 2 {

			if found4[0].Name != "glibc" || found4[0].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[0].Name, found4[0].Version.String(), "glibc", "0:1.2.3~patch4")
			}

			if found4[1].Name != "glibc" || found4[1].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[1].Name, found4[1].Version.String(), "glibc", "0:2.41.0r47")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 2, "Version=<= 0:2.41~r47")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		antiques := mockupAntiques()

		antique1 := antiques.Get("apt:debian-bullseye:glibc:0:1.2.3~patch4:x86_64")

		if antique1 == nil {
			t.Errorf("Expected nil to be %s version %s", "glibc", "0:1.3.0")
		} else if antique1.Name != "glibc" || antique1.Version.String() != "0:1.2.3~patch4" {
			t.Errorf("Expected %s version %s to be %s version %s", antique1.Name, antique1.Version.String(), "glibc", "0:1.3.0")
		}

		antiques.Remove("debian-bullseye:glibc:0:1.2.3~patch4")

		antique2 := antiques.Get("debian-bullseye:glibc:0:1.2.3~patch4:x86_64")

		if antique2 != nil {
			t.Errorf("Expected %s version %s to be nil", antique2.Name, antique2.Version.String())
		}

	})

}
