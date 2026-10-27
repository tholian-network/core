package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupDistributions() *Distributions {

	distributions := NewDistributions()
	distributions.Add(structs.Distribution{
		Name:    "archlinux",
		Version: "any",
		Kernel:  "Linux",
		Manager: "pacman",
		Vendor:  "archlinux",
		Keywords: &map[string]string{
			"ID":   "arch",
			"NAME": "Arch Linux",
		},
	})
	distributions.Add(structs.Distribution{
		Name:    "debian",
		Version: "any",
		Kernel:  "any",
		Manager: "apt",
		Vendor:  "debian",
		Keywords: &map[string]string{
			"ID":   "debian",
			"NAME": "Debian GNU/Linux",
		},
	})
	distributions.Add(structs.Distribution{
		Name:    "debian-bullseye",
		Version: "11",
		Kernel:  "any",
		Manager: "apt",
		Vendor:  "debian-bullseye",
		Keywords: &map[string]string{
			"ID":               "debian",
			"NAME":             "Debian GNU/Linux",
			"VERSION_CODENAME": "bullseye",
		},
	})
	distributions.Add(structs.Distribution{
		Name:    "debian-bookworm",
		Version: "12",
		Kernel:  "any",
		Manager: "apt",
		Vendor:  "debian-bookworm",
		Keywords: &map[string]string{
			"ID":               "debian",
			"NAME":             "Debian GNU/Linux",
			"VERSION_CODENAME": "bookworm",
		},
	})

	return distributions

}

func TestDistributions(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		distributions := mockupDistributions()

		distribution1 := distributions.Get("ubuntu")

		if distribution1 != nil {
			t.Errorf("Expected %s to be nil", distribution1.Name)
		}

		distributions.Add(structs.Distribution{
			Name:    "ubuntu",
			Version: "any",
			Kernel:  "any",
			Manager: "apt",
			Vendor:  "ubuntu",
			Keywords: &map[string]string{
				"ID":   "ubuntu",
				"NAME": "Ubuntu",
			},
		})

		distribution2 := distributions.Get("ubuntu")

		if distribution2 == nil {
			t.Errorf("Expected nil to be %s", "ubuntu")
		} else if distribution2.Name != "ubuntu" {
			t.Errorf("Expected %s to be %s", distribution2.Name, "ubuntu")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		distributions := mockupDistributions()

		distribution1 := distributions.Get("archlinux")
		distribution2 := distributions.Get("debian")
		distribution3 := distributions.Get("debian-bookworm")

		if distribution1 == nil {
			t.Errorf("Expected nil to be %s", "archlinux")
		} else if distribution1.Name != "archlinux" {
			t.Errorf("Expected %s to be %s", distribution1.Name, "archlinux")
		}

		if distribution2 == nil {
			t.Errorf("Expected nil to be %s", "debian")
		} else if distribution2.Name != "debian" {
			t.Errorf("Expected %s to be %s", distribution2.Name, "debian")
		}

		if distribution3 == nil {
			t.Errorf("Expected nil to be %s", "debian-bookworm")
		} else if distribution3.Name != "debian-bookworm" {
			t.Errorf("Expected %s to be %s", distribution3.Name, "debian-bookworm")
		}

	})

	t.Run("QueryByDistribution()", func(t *testing.T) {

		distributions := mockupDistributions()

		found1 := distributions.Query(matchers.Distribution{
			Name: "debian-bullseye",
			Version: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := distributions.Query(matchers.Distribution{
			Name: "any",
			Version: "12",
			Manager: "any",
			Vendor: "any",
		})

		found3 := distributions.Query(matchers.Distribution{
			Name: "any",
			Version: "any",
			Manager: "apt",
			Vendor: "any",
		})

		found4 := distributions.Query(matchers.Distribution{
			Name: "any",
			Version: "any",
			Manager: "any",
			Vendor: "archlinux",
		})

		if len(found1) == 1 {

			if found1[0].Name != "debian-bullseye" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "debian-bullseye")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=debian-bullseye")
		}

		if len(found2) == 1 {

			if found2[0].Name != "debian-bookworm" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "debian-bookworm")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "debian-bookworm")
		}

		if len(found3) == 3 {

			if found3[0].Name != "debian" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "debian")
			}

			if found3[1].Name != "debian-bookworm" {
				t.Errorf("Expected %s to be %s", found3[1].Name, "debian-bookworm")
			}

			if found3[2].Name != "debian-bullseye" {
				t.Errorf("Expected %s to be %s", found3[2].Name, "debian-bullseye")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 3, "Manager=apt")
		}

		if len(found4) == 1 {

			if found4[0].Name != "archlinux" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "archlinux")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 1, "Vendor=archlinux")
		}

	})

	t.Run("QueryByKernel()", func(t *testing.T) {

		distributions := mockupDistributions()

		found1 := distributions.QueryByKernel("any", "any")
		found2 := distributions.QueryByKernel("Linux", "any")

		if len(found1) == 4 {

			if found1[0].Name != "archlinux" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "archlinux")
			}

			if found1[1].Name != "debian" {
				t.Errorf("Expected %s to be %s", found1[1].Name, "debian")
			}

			if found1[2].Name != "debian-bookworm" {
				t.Errorf("Expected %s to be %s", found1[2].Name, "debian-bookworm")
			}

			if found1[3].Name != "debian-bullseye" {
				t.Errorf("Expected %s to be %s", found1[3].Name, "debian-bullseye")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 4, "any")
		}

		if len(found2) == 1 {

			if found2[0].Name != "archlinux" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "archlinux")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Linux")
		}

	})

	t.Run("QueryByKeyword()", func(t *testing.T) {

		distributions := mockupDistributions()

		found1 := distributions.QueryByKeyword("ID", "arch")
		found2 := distributions.QueryByKeyword("ID", "debian")

		if len(found1) == 1 {

			if found1[0].Name != "archlinux" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "archlinux")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "ID=arch")
		}

		if len(found2) == 3 {

			if found2[0].Name != "debian" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "debian")
			}

			if found2[1].Name != "debian-bookworm" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "debian-bookworm")
			}

			if found2[2].Name != "debian-bullseye" {
				t.Errorf("Expected %s to be %s", found2[2].Name, "debian-bullseye")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 3, "ID=debian")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		distributions := mockupDistributions()

		found1 := distributions.Get("debian-bookworm")

		distributions.Remove("debian-bookworm")

		found2 := distributions.Get("debian-bookworm")

		if found1 == nil {
			t.Errorf("Expected nil to be %s", "debian-bookworm")
		} else if found1.Name != "debian-bookworm" {
			t.Errorf("Expected %s to be %s", found1.Name, "debian-bookworm")
		}

		if found2 != nil {
			t.Errorf("Expected %s to be nil", "debian-bookworm")
		}

	})

}
