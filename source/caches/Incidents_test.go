package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupIncidents() *Incidents {

	networks := NewNetworks()
	incidents := NewIncidents(networks)

	incidents.Add(structs.Incident{
		System: "database-01",
		Type: "Intel",
		Datetime: types.ToDatetime("2025-01-02 00:01:00"),
		Connections: []matchers.Connection{
			matchers.Connection{
				Host: "any",
				Port: 22,
				Protocol: "tcp",
				Type: "server",
			},
		},
		Distributions: []matchers.Distribution{
			matchers.Distribution{
				Name: "archlinux",
				Version: "any",
				Manager: "pacman",
				Vendor: "archlinux",
			},
		},
		Networks: []matchers.Network{
			matchers.Network{
				Name: "AS210644",
				Subnet: "185.217.197.0/24",
			},
		},
		Packages: []matchers.Package{
			matchers.Package{
				Name: "openssh",
				Version: "10.0p1-3",
				Architecture: "x86_64",
				Manager: "pacman",
				Vendor: "archlinux",
			},
		},
		Programs: []matchers.Program{
			matchers.Program{
				Name: "sshd",
				Command: "/usr/bin/sshd",
				Arguments: []string{},
			},
		},
		Users: []matchers.User{
			matchers.User{
				Name: "postgres",
				Password: "postgres",
				Type: "program",
			},
		},
		Vulnerabilities: []matchers.Vulnerability{},
	})
	incidents.Add(structs.Incident{
		System: "database-01",
		Type: "Recon",
		Datetime: types.ToDatetime("2025-01-02 00:01:30"),
		Connections: []matchers.Connection{},
		Distributions: []matchers.Distribution{
			matchers.Distribution{
				Name: "archlinux",
				Version: "any",
				Manager: "pacman",
				Vendor: "archlinux",
			},
		},
		Networks: []matchers.Network{
			matchers.Network{
				Name: "AS210644",
				Subnet: "185.217.197.0/24",
			},
			matchers.Network{
				Name: "AS210644",
				Subnet: "185.229.65.0/24",
			},
		},
		Packages: []matchers.Package{
			matchers.Package{
				Name: "postgres",
				Version: "17.5-2",
				Architecture: "x86_64",
				Manager: "pacman",
				Vendor: "archlinux",
			},
			matchers.Package{
				Name: "postgresql-libs",
				Version: "17.5-2",
				Architecture: "x86_64",
				Manager: "pacman",
				Vendor: "archlinux",
			},
		},
		Programs: []matchers.Program{
			matchers.Program{
				Name: "postgres",
				Command: "/usr/bin/postgres",
				Arguments: []string{"-D", "any"},
			},
		},
		Users: []matchers.User{
			matchers.User{
				Name: "admin",
				Password: "admin123",
				Type: "user",
			},
			matchers.User{
				Name: "postgres",
				Password: "postgres",
				Type: "program",
			},
		},
		Vulnerabilities: []matchers.Vulnerability{},
	})
	incidents.Add(structs.Incident{
		System: "database-02",
		Type: "Intel",
		Datetime: types.ToDatetime("2025-01-02 00:01:01"),
		Connections: []matchers.Connection{
			matchers.Connection{
				Host: "any",
				Port: 22,
				Protocol: "tcp",
				Type: "server",
			},
		},
		Distributions: []matchers.Distribution{
			matchers.Distribution{
				Name: "debian-buster",
				Version: "10",
				Manager: "apt",
				Vendor: "debian-buster",
			},
		},
		Networks: []matchers.Network{
			matchers.Network{
				Name: "AS210644",
				Subnet: "185.217.197.0/24",
			},
		},
		Packages: []matchers.Package{
			matchers.Package{
				Name: "openssh",
				Version: "7.6p1-3",
				Architecture: "x86_64",
				Manager: "apt",
				Vendor: "debian-buster",
			},
		},
		Programs: []matchers.Program{
			matchers.Program{
				Name: "sshd",
				Command: "/usr/bin/sshd",
				Arguments: []string{},
			},
		},
		Users: []matchers.User{
			matchers.User{
				Name: "admin",
				Password: "admin123",
				Type: "user",
			},
		},
		Vulnerabilities: []matchers.Vulnerability{
			matchers.Vulnerability{
				Name: "CVE-2018-15473",
			},
		},
	})
	incidents.Add(structs.Incident{
		System: "database-03",
		Type: "Conquer",
		Datetime: types.ToDatetime("2025-01-02 00:03:45"),
		Connections: []matchers.Connection{
			matchers.Connection{
				Host: "any",
				Port: 3306,
				Protocol: "tcp",
				Type: "server",
			},
		},
		Distributions: []matchers.Distribution{
			matchers.Distribution{
				Name: "debian-buster",
				Version: "10",
				Manager: "apt",
				Vendor: "debian-buster",
			},
		},
		Networks: []matchers.Network{
			matchers.Network{
				Name: "AS12722",
				Subnet: "46.22.48.0/20",
			},
		},
		Packages: []matchers.Package{
			matchers.Package{
				Name: "mysql",
				Version: "8.0.23",
				Architecture: "x86_64",
				Manager: "apt",
				Vendor: "debian-buster",
			},
		},
		Programs: []matchers.Program{
			matchers.Program{
				Name: "mysqld",
				Command: "/usr/bin/mysqld",
				Arguments: []string{},
			},
		},
		Users: []matchers.User{
			matchers.User{
				Name: "mysql",
				Password: "mysql",
				Type: "program",
			},
		},
		Vulnerabilities: []matchers.Vulnerability{
			matchers.Vulnerability{
				Name: "CVE-2023-21971",
			},
		},
	})

	return incidents

}

func TestIncidents(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		incidents := mockupIncidents()

		incident1 := incidents.Get("database-01:Persist:2025-01-02 00:13:37")

		if incident1 != nil {
			t.Errorf("Expected system %s type %s to be nil", incident1.System, incident1.Type)
		}

		incidents.Add(structs.Incident{
			System: "database-01",
			Type: "Persist",
			Datetime: types.ToDatetime("2025-01-02 00:13:37"),
			Connections:   []matchers.Connection{},
			Distributions: []matchers.Distribution{
				matchers.Distribution{
					Name: "archlinux",
					Version: "any",
					Manager: "pacman",
					Vendor: "archlinux",
				},
			},
			Networks: []matchers.Network{},
			Packages: []matchers.Package{},
			Programs: []matchers.Program{
				matchers.Program{
					Name: "7z",
					Command: "/usr/bin/7z",
					Arguments: []string{
						"a",
						"-ttar",
						"-an",
						"-so",
						"/etc/passwd",
					},
				},
			},
			Users: []matchers.User{
				matchers.User{
					Name: "postgres",
					Password: "any",
					Type: "program",
				},
			},
			Vulnerabilities: []matchers.Vulnerability{},
		})

		incident2 := incidents.Get("database-01:Persist:2025-01-02 00:13:37")

		if incident2 == nil {
			t.Errorf("Expected nil to be %s", "database-01")
		} else if incident2.System != "database-01" || incident2.Type != "Persist" {
			t.Errorf("Expected system %s type %s to be system %s type %s", incident2.System, incident2.Type, "database-01", "Persistence")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		incidents := mockupIncidents()

		incident1 := incidents.Get("database-01:Intel:2025-01-02 00:01:00")
		incident2 := incidents.Get("database-01:Recon:2025-01-02 00:01:30")

		if incident1 == nil {
			t.Errorf("Expected nil to be system %s type %s", "database-01", "Intel")
		} else if incident1.System != "database-01" || incident1.Type != "Intel" {
			t.Errorf("Expected system %s type %s to be system %s type %s", incident1.System, incident1.Type, "database-01", "Intel")
		}

		if incident2 == nil {
			t.Errorf("Expected nil to be system %s type %s", "database-01", "Intel")
		} else if incident2.System != "database-01" || incident2.Type != "Recon" {
			t.Errorf("Expected system %s type %s to be system %s type %s", incident2.System, incident2.Type, "database-01", "Recon")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		incidents := mockupIncidents()

		found1 := incidents.Query(matchers.Incident{
			System: "database-01",
			Type: "any",
			Datetime: "any",
		})

		found2 := incidents.Query(matchers.Incident{
			System: "any",
			Type: "Recon",
			Datetime: "any",
		})

		found3 := incidents.Query(matchers.Incident{
			System: "any",
			Type: "any",
			Datetime: "> 2025-01-02 00:01:00",
		})

		if len(found1) == 2 {

			if found1[0].System != "database-01" || found1[0].Type != "Recon" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[0].System, found1[0].Type, "database-01", "Recon")
			}

			if found1[1].System != "database-01" || found1[1].Type != "Intel" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[1].System, found1[1].Type, "database-01", "Intel")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "System=database-01")
		}

		if len(found2) == 1 {

			if found2[0].System != "database-01" || found2[0].Type != "Recon" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found2[0].System, found2[0].Type, "database-01", "Recon")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Type=Recon")
		}

		if len(found3) == 3 {

			if found3[0].System != "database-03" || found3[0].Type != "Conquer" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found3[0].System, found3[0].Type, "database-03", "Conquer")
			}

			if found3[1].System != "database-01" || found3[1].Type != "Recon" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found3[1].System, found3[1].Type, "database-01", "Recon")
			}

			if found3[2].System != "database-02" || found3[2].Type != "Intel" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found3[2].System, found3[2].Type, "database-02", "Intel")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 3, "Datetime>2025-01-02 00:01:00")
		}

	})

	t.Run("QueryByDistribution()", func(t *testing.T) {

		incidents := mockupIncidents()

		found1 := incidents.QueryByDistribution(matchers.Distribution{
			Name: "archlinux",
			Version: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := incidents.QueryByDistribution(matchers.Distribution{
			Name: "debian-buster",
			Version: "any",
			Manager: "any",
			Vendor: "any",
		})

		if len(found1) == 2 {

			if found1[0].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found1[0].System, "database-01")
			}

			if found1[1].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found1[1].System, "database-01")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=archlinux")
		}
		
		if len(found2) == 2 {

			if found2[0].System != "database-03" {
				t.Errorf("Expected system %s to be system %s", found2[0].System, "database-03")
			}

			if found2[1].System != "database-02" {
				t.Errorf("Expected system %s to be system %s", found2[1].System, "database-02")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Name=debian-buster")
		}

	})

	t.Run("QueryByIP()", func(t *testing.T) {

		incidents := mockupIncidents()

		found1 := incidents.QueryByIP("185.217.197.123")
		found2 := incidents.QueryByIP("185.229.65.100")
		found3 := incidents.QueryByIP("46.22.48.1")

		if len(found1) == 3 {

			if found1[0].System != "database-01" || found1[0].Type != "Recon" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[0].System, found1[0].Type, "database-01", "Recon")
			}

			if found1[1].System != "database-02" || found1[1].Type != "Intel" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[1].System, found1[1].Type, "database-02", "Intel")
			}

			if found1[2].System != "database-01" || found1[2].Type != "Intel" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found1[2].System, found1[2].Type, "database-01", "Intel")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 3, "185.217.197.123")
		}

		if len(found2) == 3 {

			if found2[0].System != "database-01" || found2[0].Type != "Recon" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found2[0].System, found2[0].Type, "database-01", "Recon")
			}

			if found2[1].System != "database-02" || found2[1].Type != "Intel" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found2[1].System, found2[1].Type, "database-02", "Intel")
			}

			if found2[2].System != "database-01" || found2[2].Type != "Intel" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found2[2].System, found2[2].Type, "database-01", "Intel")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 3, "185.229.65.100")
		}

		if len(found3) == 1 {

			if found3[0].System != "database-03" || found3[0].Type != "Conquer" {
				t.Errorf("Expected system %s type %s to be system %s type %s", found3[0].System, found3[0].Type, "database-03", "Conquer")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "46.22.48.1")
		}

	})

	t.Run("QueryByPackage()", func(t *testing.T) {

		incidents := mockupIncidents()

		found1 := incidents.QueryByPackage(matchers.Package{
			Name: "openssh",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		found2 := incidents.QueryByPackage(matchers.Package{
			Name: "postgres",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "any",
		})

		if len(found1) == 2 {

			if found1[0].System != "database-02" {
				t.Errorf("Expected system %s to be system %s", found1[0].System, "database-02")
			}

			if found1[1].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found1[1].System, "database-01")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=openssh")
		}

		if len(found2) == 1 {

			if found2[0].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found2[0].System, "database-01")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Name=postgres")
		}

	})

	t.Run("QueryByProgram()", func(t *testing.T) {

		incidents := mockupIncidents()

		found1 := incidents.QueryByProgram(matchers.Program{
			Name: "sshd",
			Command: "any",
			Arguments: []string{},
		})

		found2 := incidents.QueryByProgram(matchers.Program{
			Name: "any",
			Command: "/usr/bin/postgres",
			Arguments: []string{},
		})

		if len(found1) == 2 {

			if found1[0].System != "database-02" {
				t.Errorf("Expected system %s to be system %s", found1[0].System, "database-02")
			}

			if found1[1].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found1[1].System, "database-01")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=openssh")
		}

		if len(found2) == 1 {

			if found2[0].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found2[0].System, "database-01")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Name=postgres")
		}

	})

	t.Run("QueryByUser()", func(t *testing.T) {

		incidents := mockupIncidents()

		found1 := incidents.QueryByUser(matchers.User{
			Name: "postgres",
			Password: "any",
			Type: "any",
		})

		found2 := incidents.QueryByUser(matchers.User{
			Name: "any",
			Password: "any",
			Type: "user",
		})

		if len(found1) == 2 {

			if found1[0].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found1[0].System, "database-01")
			}

			if found1[1].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found1[1].System, "database-01")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=postgres")
		}

		if len(found2) == 2 {

			if found2[0].System != "database-01" {
				t.Errorf("Expected system %s to be system %s", found2[0].System, "database-01")
			}

			if found2[1].System != "database-02" {
				t.Errorf("Expected system %s to be system %s", found2[1].System, "database-02")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Type=user")
		}

	})

	t.Run("QueryByVulnerability()", func(t *testing.T) {

		incidents := mockupIncidents()

		found1 := incidents.QueryByVulnerability(matchers.Vulnerability{
			Name: "CVE-2018-15473",
		})

		if len(found1) == 1 {

			if found1[0].System != "database-02" || found1[0].Datetime.String() != "2025-01-02 00:01:01" {
				t.Errorf("Expected system %s at %s to be system %s at %s", found1[0].System, found1[0].Datetime.String(), "database-02", "2025-01-02 00:01:01")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=CVE-2018-15473")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		incidents := mockupIncidents()

		incident1 := incidents.Get("database-01:Recon:2025-01-02 00:01:30")

		if incident1 == nil {
			t.Errorf("Expected nil to be system %s", "database-01")
		} else if incident1.System != "database-01" || incident1.Type != "Recon" {
			t.Errorf("Expected system %s type %s to be system %s type %s", incident1.System, incident1.Type, "database-01", "Recon")
		}

		incidents.Remove("database-01:Recon:2025-01-02 00:01:30")

		incident2 := incidents.Get("database-01:Recon:2025-01-02 00:01:30")

		if incident2 != nil {
			t.Errorf("Expected system %s type %s to be nil", incident2.System, incident2.Type)
		}

	})

}
