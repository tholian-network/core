package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupPackages() *Packages {

	packages := NewPackages()
	packages.Add(structs.Package{
		Name:         "binutils",
		Version:      types.ToVersion("0:2.44+r94"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("pacman"),
		Vendor:       "archlinux",
		URL:          "https://www.gnu.org/software/binutils/",
		Datetime:     types.ToDatetime("2025-04-27 16:00:09"),
		Maintainers:  []types.Maintainer{},
		Filesystem:   []string{},
		Conflicts:    []matchers.Package{},
		Dependencies: []matchers.Package{
			matchers.Package{
				Name:         "glibc",
				Version:      ">= 2.41+r48",
				Architecture: "x86_64",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
			matchers.Package{
				Name:         "zlib",
				Version:      "any",
				Architecture: "x86_64",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
		},
		Provides:     []matchers.Package{
			matchers.Package{
				Name:         "libctf.so",
				Version:      "any",
				Architecture: "x86_64",
				Manager:      "any",
				Vendor:       "any",
			},
			matchers.Package{
				Name:         "libgprofng.so",
				Version:      "any",
				Architecture: "x86_64",
				Manager:      "any",
				Vendor:       "any",
			},
			matchers.Package{
				Name:         "libsframe.so",
				Version:      "any",
				Architecture: "x86_64",
				Manager:      "any",
				Vendor:       "any",
			},
		},
		Replaces:     []matchers.Package{},
		Unresolved:   []matchers.Unresolved{},
	})
	packages.Add(structs.Package{
		Name:         "glibc",
		Version:      types.ToVersion("0:2.41+r48"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("pacman"),
		Vendor:       "archlinux",
		URL:          "https://www.gnu.org/software/libc/",
		Datetime:     types.ToDatetime("2025-04-27 15:33:36"),
		Maintainers:  []types.Maintainer{},
		Filesystem:   []string{},
		Conflicts:    []matchers.Package{},
		Dependencies: []matchers.Package{
			matchers.Package{
				Name:         "tzdata",
				Version:      "any",
				Architecture: "x86_64",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
		},
		Provides:     []matchers.Package{},
		Replaces:     []matchers.Package{},
		Unresolved:   []matchers.Unresolved{},
	})
	packages.Add(structs.Package{
		Name:         "glibc",
		Version:      types.ToVersion("0:1.23"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("pacman"),
		Vendor:       "archlinux",
		URL:          "https://www.gnu.org/software/libc/",
		Datetime:     types.ToDatetime("2020-01-02 03:04:05"),
		Maintainers:  []types.Maintainer{},
		Filesystem:   []string{},
		Conflicts:    []matchers.Package{},
		Dependencies: []matchers.Package{
			matchers.Package{
				Name:         "tzdata",
				Version:      "any",
				Architecture: "x86_64",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
		},
		Provides:     []matchers.Package{},
		Replaces:     []matchers.Package{},
		Unresolved:   []matchers.Unresolved{},
	})
	packages.Add(structs.Package{
		Name:         "tzdata",
		Version:      types.ToVersion("2025b-1"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("pacman"),
		Vendor:       "archlinux",
		URL:          "https://www.iana.org/time-zones",
		Datetime:     types.ToDatetime("2025-03-23 10:28:56"),
		Maintainers:  []types.Maintainer{},
		Filesystem:   []string{},
		Conflicts:    []matchers.Package{},
		Dependencies: []matchers.Package{},
		Provides:     []matchers.Package{},
		Replaces:     []matchers.Package{},
		Unresolved:   []matchers.Unresolved{},
	})
	packages.Add(structs.Package{
		Name:         "zlib",
		Version:      types.ToVersion("1:1.3.1-2"),
		Architecture: types.ToArchitecture("amd64"),
		Manager:      types.ToManager("pacman"),
		Vendor:       "archlinux",
		URL:          "https://www.gnu.org/software/libc/",
		Datetime:     types.ToDatetime("2025-05-02 09:47:11"),
		Maintainers:  []types.Maintainer{},
		Filesystem:   []string{},
		Conflicts:    []matchers.Package{},
		Dependencies: []matchers.Package{
			matchers.Package{
				Name:         "glibc",
				Version:      "2.41+r48",
				Architecture: "x86_64",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
		},
		Provides:     []matchers.Package{
			matchers.Package{
				Name:         "libz.so",
				Version:      "1",
				Architecture: "x86_64",
				Manager:      "any",
				Vendor:       "any",
			},
		},
		Replaces:     []matchers.Package{},
		Unresolved:   []matchers.Unresolved{},
	})
	packages.Add(structs.Package{
		Name:         "zlib",
		Version:      types.ToVersion("1:1.3.1-2"),
		Architecture: types.ToArchitecture("i686"),
		Manager:      types.ToManager("pacman"),
		Vendor:       "archlinux",
		URL:          "https://www.gnu.org/software/libc/",
		Datetime:     types.ToDatetime("2025-05-02 09:47:11"),
		Maintainers:  []types.Maintainer{},
		Filesystem:   []string{},
		Conflicts:    []matchers.Package{},
		Dependencies: []matchers.Package{
			matchers.Package{
				Name:         "glibc",
				Version:      ">= 2.41+r48",
				Architecture: "x86_64",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
		},
		Provides:     []matchers.Package{
			matchers.Package{
				Name:         "libz.so",
				Version:      "1",
				Architecture: "x86_64",
				Manager:      "any",
				Vendor:       "any",
			},
		},
		Replaces:     []matchers.Package{},
		Unresolved:   []matchers.Unresolved{},
	})

	return packages

}

func TestPackages(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		packages := mockupPackages()

		package1 := packages.Get("pacman:archlinux:libelf:0:0.193.2:x86_64")

		if package1 != nil {
			t.Errorf("Expected %s to be nil", package1.Name)
		}

		packages.Add(structs.Package{
			Name:         "libelf",
			Version:      types.ToVersion("0.193-2"),
			Architecture: types.ToArchitecture("amd64"),
			Manager:      types.ToManager("pacman"),
			Vendor:       "archlinux",
			URL:          "https://sourceware.org/elfutils/",
			Datetime:     types.ToDatetime("2025-04-29 09:53:38"),
			Maintainers:  []types.Maintainer{},
			Filesystem:   []string{},
			Conflicts:    []matchers.Package{},
			Dependencies: []matchers.Package{
				matchers.Package{
					Name:         "glibc",
					Version:      ">= 2.41+r48",
					Architecture: "x86_64",
					Manager:      "pacman",
					Vendor:       "archlinux",
				},
				matchers.Package{
					Name:         "zlib",
					Version:      "any",
					Architecture: "x86_64",
					Manager:      "pacman",
					Vendor:       "archlinux",
				},
			},
			Provides:     []matchers.Package{},
			Replaces:     []matchers.Package{},
			Unresolved:   []matchers.Unresolved{},
		})

		package2 := packages.Get("pacman:archlinux:libelf:0:0.193.2:x86_64")

		if package2 == nil {
			t.Errorf("Expected nil to be %s", "libelf")
		} else if package2.Name != "libelf" {
			t.Errorf("Expected %s to be %s", package2.Name, "libelf")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		packages := mockupPackages()

		package1 := packages.Get("pacman:archlinux:binutils:0:2.44.0r94:x86_64")

		if package1 == nil {
			t.Errorf("Expected nil to be %s version %s", "binutils", "0:2.44.0r94")
		} else if package1.Name != "binutils" || package1.Version.String() != "0:2.44.0r94" {
			t.Errorf("Expected %s version %s to be %s version %s", package1.Name, package1.Version.String(), "binutils", "0:2.44.0r94")
		}

		package2 := packages.Get("pacman:archlinux:glibc:0:2.41.0r48:x86_64")

		if package2 == nil {
			t.Errorf("Expected nil to be %s version %s", "glibc", "0:2.41.0r48")
		} else if package2.Name != "glibc" || package2.Version.String() != "0:2.41.0r48" {
			t.Errorf("Expected %s version %s to be %s version %s", package2.Name, package2.Version.String(), "glibc", "0:2.41.0r48")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		packages := mockupPackages()

		found1 := packages.Query(matchers.Package{
			Name: "glibc",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := packages.Query(matchers.Package{
			Name: "any",
			Version: "0:2.44.0r94",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found3 := packages.Query(matchers.Package{
			Name: "any",
			Version: "any",
			Architecture: "x86_64",
			Manager: "any",
			Vendor: "any",
		})

		found4 := packages.Query(matchers.Package{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "pacman",
			Vendor: "any",
		})

		found5 := packages.Query(matchers.Package{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "archlinux",
		})

		if len(found1) == 2 {

			if found1[0].Name != "glibc" || found1[0].Version.String() != "0:1.23.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[0].Name, found1[0].Version.String(), "glibc", "0:1.23.0")
			}

			if found1[1].Name != "glibc" || found1[1].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[1].Name, found1[1].Version.String(), "glibc", "0:2.41.0r48")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=glibc")
		}

		if len(found2) == 1 {

			if found2[0].Name != "binutils" || found2[0].Version.String() != "0:2.44.0r94" {
				t.Errorf("Expected %s version %s to be %s version %s", found2[0].Name, found2[0].Version.String(), "binutils", "0:2.44.0r94")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Version=0:2.44.0r94")
		}

		if len(found3) == 5 {

			if found3[0].Name != "binutils" || found3[0].Version.String() != "0:2.44.0r94" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[0].Name, found3[0].Version.String(), "binutils", "0:2.44.0r94")
			}

			if found3[1].Name != "glibc" || found3[1].Version.String() != "0:1.23.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[1].Name, found3[1].Version.String(), "glibc", "0:1.23.0")
			}

			if found3[2].Name != "glibc" || found3[2].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[2].Name, found3[2].Version.String(), "glibc", "0:2.41.0r48")
			}

			if found3[3].Name != "tzdata" || found3[3].Version.String() != "0:2025.0.0b~1" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[3].Name, found3[3].Version.String(), "tzdata", "0:2025.0.0b~1")
			}

			if found3[4].Name != "zlib" || found3[4].Version.String() != "1:1.3.1~2" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[4].Name, found3[4].Version.String(), "zlib", "1:1.3.1~2")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 5, "Architecture=x86_64")
		}

		if len(found4) == 6 {

			if found4[0].Name != "binutils" || found4[0].Version.String() != "0:2.44.0r94" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[0].Name, found4[0].Version.String(), "binutils", "0:2.44.0r94")
			}

			if found4[1].Name != "glibc" || found4[1].Version.String() != "0:1.23.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[1].Name, found4[1].Version.String(), "glibc", "0:1.23.0")
			}

			if found4[2].Name != "glibc" || found4[2].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[2].Name, found4[2].Version.String(), "glibc", "0:2.41.0r48")
			}

			if found4[3].Name != "tzdata" || found4[3].Version.String() != "0:2025.0.0b~1" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[3].Name, found4[3].Version.String(), "tzdata", "0:2025.0.0b~1")
			}

			if found4[4].Name != "zlib" || found4[4].Version.String() != "1:1.3.1~2" || found4[4].Architecture.String() != "x86" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[4].Name, found4[4].Version.String(), "zlib", "1:1.3.1~2")
			}

			if found4[5].Name != "zlib" || found4[5].Version.String() != "1:1.3.1~2" || found4[5].Architecture.String() != "x86_64" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[5].Name, found4[5].Version.String(), "zlib", "1:1.3.1~2")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 6, "Manager=pacman")
		}

		if len(found5) != 6 {
			t.Errorf("Expected %d results to be %d for query %s", len(found5), 6, "Vendor=archlinux")
		}

	})

	t.Run("Query() Versions", func(t *testing.T) {

		packages := mockupPackages()

		found1 := packages.Query(matchers.Package{
			Name: "glibc",
			Version: "0:2.41.0r48",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := packages.Query(matchers.Package{
			Name: "glibc",
			Version: "> 0:1.23.0",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found3 := packages.Query(matchers.Package{
			Name: "glibc",
			Version: ">= 0:1.23.0",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found4 := packages.Query(matchers.Package{
			Name: "glibc",
			Version: "< 0:2.41.0r48",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found5 := packages.Query(matchers.Package{
			Name: "glibc",
			Version: "<= 0:2.41.0r48",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		if len(found1) == 1 {

			if found1[0].Name != "glibc" || found1[0].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[0].Name, found1[0].Version.String(), "glibc", "0:2.41.0r48")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Version=0:2.41.0r48")
		}

		if len(found2) == 1 {

			if found2[0].Name != "glibc" || found2[0].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found2[0].Name, found2[0].Version.String(), "glibc", "0:2.41.0r48")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Version>0:1.23.0")
		}

		if len(found3) == 2 {

			if found3[0].Name != "glibc" || found3[0].Version.String() != "0:1.23.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[0].Name, found3[0].Version.String(), "glibc", "0:1.23.0")
			}

			if found3[1].Name != "glibc" || found3[1].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found3[1].Name, found3[1].Version.String(), "glibc", "0:2.41.0r48")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 2, "Version>=0:1.23.0")
		}

		if len(found4) == 1 {

			if found4[0].Name != "glibc" || found4[0].Version.String() != "0:1.23.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found4[0].Name, found4[0].Version.String(), "glibc", "0:1.23.0")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 1, "Version<0:2.41.0r48")
		}

		if len(found5) == 2 {

			if found5[0].Name != "glibc" || found5[0].Version.String() != "0:1.23.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found5[0].Name, found5[0].Version.String(), "glibc", "0:1.23.0")
			}

			if found5[1].Name != "glibc" || found5[1].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found5[1].Name, found5[1].Version.String(), "glibc", "0:2.41.0r48")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found5), 2, "Version<0:2.41.0r48")
		}

	})

	t.Run("QueryByDependency()", func(t *testing.T) {

		packages := mockupPackages()

		found1 := packages.QueryByDependency(matchers.Package{
			Name: "glibc",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := packages.QueryByDependency(matchers.Package{
			Name: "tzdata",
			Version: "<= 0:2.41.0r48",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		if len(found1) == 3 {

			if found1[0].Name != "binutils" || found1[0].Version.String() != "0:2.44.0r94" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[0].Name, found1[0].Version.String(), "binutils", "0:2.44.0r94")
			}

			if found1[1].Name != "zlib" || found1[1].Version.String() != "1:1.3.1~2" || found1[1].Architecture.String() != "x86" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[1].Name, found1[1].Version.String(), "zlib", "1:1.3.1~2")
			}

			if found1[2].Name != "zlib" || found1[2].Version.String() != "1:1.3.1~2" || found1[2].Architecture.String() != "x86_64" {
				t.Errorf("Expected %s version %s to be %s version %s", found1[2].Name, found1[2].Version.String(), "zlib", "1:1.3.1~2")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 3, "Name=glibc")
		}

		if len(found2) == 2 {

			if found2[0].Name != "glibc" || found2[0].Version.String() != "0:1.23.0" {
				t.Errorf("Expected %s version %s to be %s version %s", found2[0].Name, found2[0].Version.String(), "glibc", "0:1.23.0")
			}

			if found2[1].Name != "glibc" || found2[1].Version.String() != "0:2.41.0r48" {
				t.Errorf("Expected %s version %s to be %s version %s", found2[1].Name, found2[1].Version.String(), "glibc", "0:2.41.0r48")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Version<=0:2.41.0r48")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		packages := mockupPackages()

		package1 := packages.Get("pacman:archlinux:binutils:0:2.44.0r94:x86_64")

		if package1 == nil {
			t.Errorf("Expected nil to be %s version %s", "binutils", "0:2.44.0r94")
		} else if package1.Name != "binutils" || package1.Version.String() != "0:2.44.0r94" {
			t.Errorf("Expected %s version %s to be %s version %s", package1.Name, package1.Version.String(), "binutils", "0:2.44.0r94")
		}

		packages.Remove("pacman:archlinux:binutils:0:2.44.0r94:x86_64")

		package2 := packages.Get("pacman:archlinux:binutils:0:2.44.0r94:x86_64")

		if package2 != nil {
			t.Errorf("Expected %s version %s to be nil", package2.Name, package2.Version.String())
		}

	})

}
