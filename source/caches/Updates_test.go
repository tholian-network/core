package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupUpdates() *Updates {

	updates := NewUpdates()
	updates.Add(structs.Update{
		Name:         "glibc",
		Version:      types.ToVersion("0:1.2.3-patch4"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("apt"),
		Vendor:       "debian-bullseye",
		URL:          "http://mirror.debian.org/packages/g/glibc/glibc-0:1.2.3-patch4_amd64.deb",
	})
	updates.Add(structs.Update{
		Name:         "glibc",
		Version:      types.ToVersion("0:2.41-r47"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("apt"),
		Vendor:       "debian-bullseye",
		URL:          "http://mirror.debian.org/packages/g/glibc/glibc-0:2.41-r47_amd64.deb",
	})
	updates.Add(structs.Update{
		Name:         "left-pad",
		Version:      types.ToVersion("0:1.3.0"),
		Architecture: types.ToArchitecture("any"),
		Manager:      types.ToManager("npm"),
		Vendor:       "npm",
		URL:          "https://www.npmjs.com/package/left-pad",
	})

	return updates

}

func TestUpdates(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		updates := mockupUpdates()

		update1 := updates.Get("npm:npm:right-pad:0:0.0.1:any")

		if update1 != nil {
			t.Errorf("Expected %s to be nil", update1.Name)
		}

		updates.Add(structs.Update{
			Name:         "right-pad",
			Version:      types.ToVersion("0:0.0.1"),
			Architecture: types.ToArchitecture("any"),
			Manager:      types.ToManager("npm"),
			Vendor:       "npm",
			URL:          "https://www.npmjs.com/package/right-pad",
		})

		update2 := updates.Get("npm:npm:right-pad:0:0.0.1:any")

		if update2 == nil {
			t.Errorf("Expected nil to be %s", "right-pad")
		} else if update2.Name != "right-pad" {
			t.Errorf("Expected %s to be %s", update2.Name, "right-pad")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		updates := mockupUpdates()

		update1 := updates.Get("apt:debian-bullseye:glibc:0:1.2.3~patch4:x86_64")
		update2 := updates.Get("npm:npm:left-pad:0:1.3.0:any")

		if update1 == nil {
			t.Errorf("Expected nil to be %s", "glibc")
		} else if update1.Name != "glibc" {
			t.Errorf("Expected %s to be %s", update1.Name, "glibc")
		}

		if update2 == nil {
			t.Errorf("Expected nil to be %s", "left-pad")
		} else if update2.Name != "left-pad" {
			t.Errorf("Expected %s to be %s", update2.Name, "left-pad")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		updates := mockupUpdates()

		found1 := updates.Query(matchers.Update{
			Name: "glibc",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := updates.Query(matchers.Update{
			Name: "any",
			Version: "0:1.2.3~patch4",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found3 := updates.Query(matchers.Update{
			Name: "any",
			Version: "any",
			Architecture: "x86_64",
			Manager: "any",
			Vendor: "any",
		})

		found4 := updates.Query(matchers.Update{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "npm",
			Vendor: "any",
		})

		found5 := updates.Query(matchers.Update{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "debian-bullseye",
		})

		if len(found1) == 2 {

			if found1[0].Name != "glibc" || found1[0].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[0].Name, found1[0].Version.String(), "glibc", "0:2.41.0r47")
			}

			if found1[1].Name != "glibc" || found1[1].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[1].Name, found1[1].Version.String(), "glibc", "0:1.2.3~patch4")
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

			if found3[0].Name != "glibc" || found3[0].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[0].Name, found3[0].Version.String(), "glibc", "0:2.41.0r47")
			}

			if found3[1].Name != "glibc" || found3[1].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[1].Name, found3[1].Version.String(), "glibc", "0:1.2.3~patch4")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 2, "Architecture=x86_64")
		}

		if len(found4) == 1 {

			if found4[0].Name != "left-pad" || found4[0].Version.String() != "0:1.3.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[0].Name, found4[0].Version.String(), "left-pad", "0:1.3.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 1, "Manager=npm")
		}

		if len(found5) == 2 {

			if found5[0].Name != "glibc" || found5[0].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found5[0].Name, found5[0].Version.String(), "glibc", "0:2.41.0r47")
			}

			if found5[1].Name != "glibc" || found5[1].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found5[1].Name, found5[1].Version.String(), "glibc", "0:1.2.3~patch4")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found5), 2, "Architecture=x86_64")
		}

	})

	t.Run("Query() Versions", func(t *testing.T) {

		updates := mockupUpdates()

		found1 := updates.Query(matchers.Update{
			Name: "glibc",
			Version: "0:1.2.3~patch4",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := updates.Query(matchers.Update{
			Name: "glibc",
			Version: "< 0:2.41~r47",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found3 := updates.Query(matchers.Update{
			Name: "glibc",
			Version: "> 0:1.2.3~patch4",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found4 := updates.Query(matchers.Update{
			Name: "glibc",
			Version: "<= 0:2.41~r47",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
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

			if found4[0].Name != "glibc" || found4[0].Version.String() != "0:2.41.0r47" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[0].Name, found4[0].Version.String(), "glibc", "0:2.41.0r47")
			}

			if found4[1].Name != "glibc" || found4[1].Version.String() != "0:1.2.3~patch4" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[1].Name, found4[1].Version.String(), "glibc", "0:1.2.3~patch4")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 2, "Version=<= 0:2.41~r47")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		updates := mockupUpdates()

		update1 := updates.Get("apt:debian-bullseye:glibc:0:1.2.3~patch4:x86_64")

		if update1 == nil {
			t.Errorf("Expected nil to be %s version %s", "glibc", "0:1.3.0")
		} else if update1.Name != "glibc" || update1.Version.String() != "0:1.2.3~patch4" {
			t.Errorf("Expected %s version %s to be %s version %s", update1.Name, update1.Version.String(), "glibc", "0:1.3.0")
		}

		updates.Remove("debian-bullseye:glibc:0:1.2.3~patch4:x86_64")

		update2 := updates.Get("debian-bullseye:glibc:0:1.2.3~patch4")

		if update2 != nil {
			t.Errorf("Expected %s version %s to be nil", update2.Name, update2.Version.String())
		}

	})

}
