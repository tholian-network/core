package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupVulnerabilities() *Vulnerabilities {

	vulnerabilities := NewVulnerabilities()
	vulnerabilities.Add(structs.Vulnerability{
		Name:        "CVE-2021-44228",
		Description: "Apache Log4j2 \u003c=2.14.1 JNDI features used in configuration, log messages, and parameters do not protect against attacker controlled LDAP and other JNDI related endpoints. An attacker who can control log messages or log message parameters can execute arbitrary code loaded from LDAP servers when message lookup substitution is enabled. From log4j 2.15.0, this behavior has been disabled by default.\r\n\r\nIn previous releases (\u003e2.10) this behavior can be mitigated by setting system property \"log4j2.formatMsgNoLookups\" to \"true\" or it can be mitigated in prior releases (\u003c2.10) by removing the JndiLookup class from the classpath (example: zip -q -d log4j-core-*.jar org/apache/logging/log4j/core/lookup/JndiLookup.class).",
		Type:        "network",
		Severity:    "critical",
		State:       "published",
		References: []string{
			"https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-44228",
		},
		Datetime:    types.ToDatetime("2021-12-10 00:00:00"),
		Connections: []matchers.Connection{},
		Devices:     []matchers.Device{},
		Distributions: []matchers.Distribution{
			matchers.Distribution{
				Name:    "archlinux",
				Version: "any",
				Vendor:  "any",
			},
			matchers.Distribution{
				Name:    "debian-bullseye",
				Version: "11",
				Vendor:  "debian-bullseye",
			},
			matchers.Distribution{
				Name:    "debian-bookworm",
				Version: "12",
				Vendor:  "debian-bookworm",
			},
		},
		Packages: []matchers.Package{
			matchers.Package{
				Name:         "elasticsearch",
				Version:      "\u003c= 7.10.2-2",
				Architecture: "any",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
			matchers.Package{
				Name:         "ghidra",
				Version:      "\u003c 10.1-1",
				Architecture: "any",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
			matchers.Package{
				Name:         "apache-log4j2",
				Version:      "\u003c 0:2.15.0-1~deb10u1",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bookworm",
			},
			matchers.Package{
				Name:         "apache-log4j2",
				Version:      "\u003c 0:2.15.0-1~deb10u1",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bullseye",
			},
			matchers.Package{
				Name:         "apache-log4j1.2",
				Version:      "\u003c 1.2.17-10+deb11u1",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bookworm",
			},
			matchers.Package{
				Name:         "apache-log4j1.2",
				Version:      "\u003c 1.2.17-10+deb11u1",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bullseye",
			},
		},
		Products: []matchers.Product{
			matchers.Product{
				Vendor:  "apache",
				Product: "log4j",
				Version: "\u003c 2.15.0",
				Type:    "software",
			},
		},
		Programs: []matchers.Program{
			matchers.Program{
				Name:    "elasticsearch",
				Command: "/usr/lib/java/bin/GenericElasticSearchDaemonFactory",
			},
			matchers.Program{
				Name:    "tomcat",
				Command: "/usr/bin/tomcat",
			},
		},
		Weaknesses: []matchers.Weakness{
			matchers.Weakness{
				Name:   "CWE-20",
				Impact: "any",
				Scope:  "any",
			},
			matchers.Weakness{
				Name:   "CWE-264",
				Impact: "any",
				Scope:  "any",
			},
			matchers.Weakness{
				Name:   "CWE-400",
				Impact: "any",
				Scope:  "any",
			},
			matchers.Weakness{
				Name:   "CWE-502",
				Impact: "any",
				Scope:  "any",
			},
			matchers.Weakness{
				Name:   "CWE-917",
				Impact: "any",
				Scope:  "any",
			},
		},
	})
	vulnerabilities.Add(structs.Vulnerability{
		Name:        "CVE-2001-0169",
		Description: "When using the LD_PRELOAD environmental variable in SUID or SGID applications, glibc does not verify that preloaded libraries in /etc/ld.so.cache are also SUID/SGID, which could allow a local user to overwrite arbitrary files by loading a library from /lib or /usr/lib.",
		Type:        "local",
		Severity:    "high",
		State:       "published",
		References: []string{
			"https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2001-0169",
		},
		Datetime:    types.ToDatetime("2001-01-11 00:00:00"),
		Connections: []matchers.Connection{},
		Devices:     []matchers.Device{},
		Distributions: []matchers.Distribution{
			matchers.Distribution{
				Name:    "archlinux",
				Version: "any",
				Vendor:  "any",
			},
			matchers.Distribution{
				Name:    "debian-bullseye",
				Version: "11",
				Vendor:  "debian-bullseye",
			},
			matchers.Distribution{
				Name:    "debian-bookworm",
				Version: "12",
				Vendor:  "debian-bookworm",
			},
			matchers.Distribution{
				Name:    "ubuntu-xenial",
				Version: "16.04",
				Vendor:  "ubuntu-xenial",
			},
		},
		Packages: []matchers.Package{
			matchers.Package{
				Name:         "glibc",
				Version:      "any",
				Architecture: "any",
				Manager:      "pacman",
				Vendor:       "archlinux",
			},
			matchers.Package{
				Name:         "libc6",
				Version:      "any",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bullseye",
			},
			matchers.Package{
				Name:         "libc6",
				Version:      "any",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bookworm",
			},
			matchers.Package{
				Name:         "libc6-dev",
				Version:      "any",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bullseye",
			},
			matchers.Package{
				Name:         "libc6-dev",
				Version:      "any",
				Architecture: "any",
				Manager:      "apt",
				Vendor:       "debian-bookworm",
			},
		},
		Products: []matchers.Product{
			matchers.Product{
				Vendor:  "gnu",
				Product: "glibc",
				Version: "any",
				Type:    "software",
			},
		},
		Programs: []matchers.Program{
			matchers.Program{
				Name:    "tomcat",
				Command: "/usr/bin/tomcat",
			},
		},
		Weaknesses: []matchers.Weakness{
			matchers.Weakness{
				Name:   "CWE-264",
				Impact: "any",
				Scope:  "any",
			},
		},
	})

	return vulnerabilities

}

func TestVulnerabilities(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		vulnerabilities := mockupVulnerabilities()

		vulnerability1 := vulnerabilities.Get("CVE-2020-1472")

		if vulnerability1 != nil {
			t.Errorf("Expected %s to be nil", vulnerability1.Name)
		}

		vulnerabilities.Add(structs.Vulnerability{
			Name:        "CVE-2020-1472",
			Description: "An elevation of privilege vulnerability exists when an attacker establishes a vulnerable Netlogon secure channel connection to a domain controller, using the Netlogon Remote Protocol (MS-NRPC). An attacker who successfully exploited the vulnerability could run a specially crafted application on a device on the network.",
			Type:        "network",
			Severity:    "critical",
			State:       "published",
			References: []string{
				"https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-1472",
			},
			Datetime:    types.ToDatetime("2020-08-17 00:00:00"),
			Connections: []matchers.Connection{},
			Distributions: []matchers.Distribution{
				matchers.Distribution{
					Name:    "windows-10",
					Version: "10",
					Vendor:  "microsoft",
				},
				matchers.Distribution{
					Name:    "windows-11",
					Version: "11",
					Vendor:  "microsoft",
				},
			},
			Packages: []matchers.Package{},
			Products: []matchers.Product{
				matchers.Product{
					Vendor:  "microsoft",
					Product: "windows-server",
					Version: "1903",
					Type:    "software",
				},
			},
			Programs: []matchers.Program{
				matchers.Program{
					Name:    "netlogon.exe",
					Command: "C:/Windows/system32/netlogon.exe",
				},
			},
			Weaknesses: []matchers.Weakness{},
		})

		vulnerability2 := vulnerabilities.Get("CVE-2020-1472")

		if vulnerability2 == nil {
			t.Errorf("Expected nil to be %s", "CVE-2020-1472")
		} else if vulnerability2.Name != "CVE-2020-1472" {
			t.Errorf("Expected %s to be %s", vulnerability2.Name, "CVE-2020-1472")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		vulnerabilities := mockupVulnerabilities()

		vulnerability1 := vulnerabilities.Get("CVE-2001-0169")
		vulnerability2 := vulnerabilities.Get("CVE-2021-44228")

		if vulnerability1 == nil {
			t.Errorf("Expected nil to be %s", "CVE-2001-0169")
		} else if vulnerability1.Name != "CVE-2001-0169" {
			t.Errorf("Expected %s to be %s", vulnerability1.Name, "CVE-2001-0169")
		}

		if vulnerability2 == nil {
			t.Errorf("Expected nil to be %s", "CVE-2021-44228")
		} else if vulnerability2.Name != "CVE-2021-44228" {
			t.Errorf("Expected %s to be %s", vulnerability2.Name, "CVE-2021-44228")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		vulnerabilities := mockupVulnerabilities()

		found1 := vulnerabilities.Query(matchers.Vulnerability{
			Name: "CVE-2021-44228",
			Type: "any",
			Severity: "any",
			State: "any",
			Datetime: "any",
		})

		found2 := vulnerabilities.Query(matchers.Vulnerability{
			Name: "any",
			Type: "network",
			Severity: "any",
			State: "published",
			Datetime: "any",
		})

		found3 := vulnerabilities.Query(matchers.Vulnerability{
			Name: "any",
			Type: "any",
			Severity: "any",
			State: "published",
			Datetime: "any",
		})

		found4 := vulnerabilities.Query(matchers.Vulnerability{
			Name: "any",
			Type: "any",
			Severity: ">= high",
			State: "any",
			Datetime: "any",
		})

		found5 := vulnerabilities.Query(matchers.Vulnerability{
			Name: "any",
			Type: "any",
			Severity: "any",
			State: "any",
			Datetime: ">= 2021-12-10 00:00:00",
		})

		if len(found1) == 1 {

			if found1[0].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=CVE-2021-44228")
		}

		if len(found2) == 1 {

			if found2[0].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Type=network,State=published")
		}

		if len(found3) == 2 {

			if found3[0].Name != "CVE-2001-0169" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "CVE-2001-0169")
			}

			if found3[1].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found3[1].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 2, "State=published")
		}

		if len(found4) == 2 {

			if found4[0].Name != "CVE-2001-0169" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "CVE-2001-0169")
			}

			if found4[1].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found4[1].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 2, "Severity>=high")
		}

		if len(found5) == 1 {

			if found5[0].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found5[0].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found5), 1, "Datetime>=2021-12-10 00:00:00")
		}

	})

	t.Run("QueryByConnection()", func(t *testing.T) {

		// TODO

	})

	t.Run("QueryByDevice()", func(t *testing.T) {

		vulnerabilities := NewVulnerabilities()
		vulnerabilities.Add(structs.Vulnerability{
			Name:        "CVE-2025-1337",
			Description: "Only affects Siemens Hardware",
			Type:        "local",
			Severity:    "high",
			State:       "published",
			Devices: []matchers.Device{
				matchers.Device{
					Vendor: "0101",
					Device: "1337",
					Name:   "Siemens Router 1337 US802B",
					Bus:    "other",
				},
				matchers.Device{
					Vendor: "0101",
					Device: "1338",
					Name:   "Siemens Router 1337 US802B 2.0",
					Bus:    "other",
				},
			},
		})
		vulnerabilities.Add(structs.Vulnerability{
			Name:        "CVE-2025-13337",
			Description: "Only affects Cisco Hardware",
			Type:        "local",
			Severity:    "high",
			State:       "published",
			Devices: []matchers.Device{
				matchers.Device{
					Vendor: "0102",
					Device: "1337",
					Name:   "Cisco Router 1337 Wi-Fi ASA",
					Bus:    "other",
				},
				matchers.Device{
					Vendor: "0102",
					Device: "1338",
					Name:   "Cisco Router 1337 Wi-Fi ASA Enterprise",
					Bus:    "other",
				},
			},
		})

		found1 := vulnerabilities.QueryByDevice(matchers.Device{
			Name: "any",
			Vendor: "0101",
			Device: "1337",
			Bus: "any",
		})

		found2 := vulnerabilities.QueryByDevice(matchers.Device{
			Name: "Siemens Router 1337 US802B",
			Vendor: "any",
			Device: "any",
			Bus: "any",
		})

		if len(found1) == 1 {

			if found1[0].Name != "CVE-2025-1337" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "CVE-2025-1337")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Vendor=0101,Device=1337")
		}

		if len(found2) == 1 {

			if found2[0].Name != "CVE-2025-1337" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "CVE-2025-1337")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Name=Siemens Router 1337 US802B")
		}

	})

	t.Run("QueryByDistribution()", func(t *testing.T) {

		vulnerabilities := mockupVulnerabilities()

		found1 := vulnerabilities.QueryByDistribution(matchers.Distribution{
			Name: "debian-bookworm",
			Version: "12",
			Manager: "any",
			Vendor: "any",
		})

		found2 := vulnerabilities.QueryByDistribution(matchers.Distribution{
			Name: "ubuntu-xenial",
			Version: "16.04",
			Manager: "any",
			Vendor: "any",
		})

		if len(found1) == 2 {

			if found1[0].Name != "CVE-2001-0169" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "CVE-2001-0169")
			}

			if found1[1].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found1[1].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Name=debian-bookworm,Version=12")
		}

		if len(found2) == 1 {

			if found2[0].Name != "CVE-2001-0169" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "CVE-2001-0169")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Name=ubuntu-xenial,Version=16.04")
		}

	})

	t.Run("QueryByPackage()", func(t *testing.T) {

		vulnerabilities := mockupVulnerabilities()

		found1 := vulnerabilities.QueryByPackage(matchers.Package{
			Name: "any",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "archlinux",
		})

		found2 := vulnerabilities.QueryByPackage(matchers.Package{
			Name: "libc6",
			Version: "any",
			Architecture: "any",
			Manager: "any",
			Vendor: "debian-bookworm",
		})

		if len(found1) == 2 {

			if found1[0].Name != "CVE-2001-0169" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "CVE-2001-0169")
			}

			if found1[1].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found1[1].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "Vendor=archlinux")
		}

		if len(found2) == 1 {

			if found2[0].Name != "CVE-2001-0169" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "CVE-2001-0169")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Vendor=debian-bookworm,Name=libc6")
		}

	})

	t.Run("QueryByProgram()", func(t *testing.T) {

		vulnerabilities := mockupVulnerabilities()

		found1 := vulnerabilities.QueryByProgram(matchers.Program{
			Name: "elasticsearch",
			Command: "any",
			Arguments: []string{},
		})

		found2 := vulnerabilities.QueryByProgram(matchers.Program{
			Name: "tomcat",
			Command: "any",
			Arguments: []string{},
		})

		if len(found1) == 1 {

			if found1[0].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=elasticsearch")
		}

		if len(found2) == 2 {

			if found2[0].Name != "CVE-2001-0169" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "CVE-2001-0169")
			}

			if found2[1].Name != "CVE-2021-44228" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "CVE-2021-44228")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Name=tomcat")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		vulnerabilities := mockupVulnerabilities()

		vulnerability1 := vulnerabilities.Get("CVE-2001-0169")

		if vulnerability1 == nil {
			t.Errorf("Expected nil to be %s", "CVE-2001-0169")
		} else if vulnerability1.Name != "CVE-2001-0169" {
			t.Errorf("Expected %s to be %s", vulnerability1.Name, "CVE-2001-0169")
		}

		vulnerabilities.Remove("CVE-2001-0169")

		vulnerability2 := vulnerabilities.Get("CVE-2001-0169")

		if vulnerability2 != nil {
			t.Errorf("Expected %s to be nil", vulnerability2.Name)
		}

	})

}
