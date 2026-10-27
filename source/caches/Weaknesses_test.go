package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupWeaknesses() *Weaknesses {

	weaknesses := NewWeaknesses()
	weaknesses.Add(structs.Weakness{
		Name: "CWE-13",
		Description: "Storing a plaintext password in a configuration file allows anyone who can read the file access to the password-protected resource making them an easy target for attackers.",
		Impacts: []string{
			"Gain Privileges or Assume Identity",
		},
		References: []string{
			"https://samate.nist.gov/SSATTM_Content/papers/Seven%20Pernicious%20Kingdoms%20-%20Taxonomy%20of%20Sw%20Security%20Errors%20-%20Tsipenyuk%20-%20Chess%20-%20McGraw.pdf",
			"https://learn.microsoft.com/en-us/previous-versions/msp-n-p/ff647398(v=pandp.10)?redirectedfrom=MSDN",
			"https://learn.microsoft.com/en-us/previous-versions/msp-n-p/ff650304(v=pandp.10)?redirectedfrom=MSDN",
			"http://msdn.microsoft.com/en-us/library/89211k9b(VS.80).aspx",
		},
		Scopes: []string{
			"Access Control",
			"Integrity",
		},
	})
	weaknesses.Add(structs.Weakness{
		Name: "CWE-37",
		Description: "The product accepts input in the form of a slash absolute path ('/absolute/pathname/here') without appropriate validation, which can allow an attacker to traverse the file system to unintended locations or access arbitrary files.",
		Impacts: []string{
			"Read Files or Directories",
			"Modify Files or Directories",
			"Modify Application Data",
		},
		References: []string{
			"https://www.cve.org/CVERecord?id=CVE-2002-1345",
			"https://www.cve.org/CVERecord?id=CVE-2001-1269",
			"https://www.cve.org/CVERecord?id=CVE-2002-1818",
			"https://www.cve.org/CVERecord?id=CVE-2002-1913",
			"https://www.cve.org/CVERecord?id=CVE-2005-2147",
			"https://www.cve.org/CVERecord?id=CVE-2000-0614",
		},
		Scopes: []string{
			"Confidentiality",
			"Integrity",
		},
	})
	weaknesses.Add(structs.Weakness{
		Name: "CWE-370",
		Description: "The product does not check the revocation status of a certificate after its initial revocation check, which can cause the product to perform privileged actions even after the certificate is revoked at a later time.",
		Impacts: []string{
			"Gain Privileges or Assume Identity",
			"Modify Application Data",
			"Read Application Data",
		},
		References: []string{
			"https://cwe.mitre.org/documents/sources/TheCLASPApplicationSecurityProcess.pdf",
		},
		Scopes: []string{
			"Access Control",
			"Integrity",
			"Confidentiality",
		},
	})

	return weaknesses

}

func TestWeaknesses(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		weaknesses := mockupWeaknesses()

		weakness1 := weaknesses.Get("CWE-1337")

		if weakness1 != nil {
			t.Errorf("Expected %s to be nil", weakness1.Name)
		}

		weaknesses.Add(structs.Weakness{
			Name: "CWE-1337",
			Description: "This is a special 1337 h4xx0r weakness",
			Impacts: []string{
				"Gain Privileges or Assume Identity",
				"Modify Application Data",
				"Read Application Data",
			},
			References: []string{
			},
			Scopes: []string{
				"Access Control",
				"Integrity",
				"Confidentiality",
			},
		})

		weakness2 := weaknesses.Get("CWE-1337")

		if weakness2 == nil {
			t.Errorf("Expected nil to be %s", "CWE-1337")
		} else if weakness2.Name != "CWE-1337" {
			t.Errorf("Expected %s to be %s", weakness2.Name, "CWE-1337")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		weaknesses := mockupWeaknesses()

		weakness1 := weaknesses.Get("CWE-13")
		weakness2 := weaknesses.Get("CWE-37")

		if weakness1 == nil {
			t.Errorf("Expected nil to be %s", "CWE-13")
		} else if weakness1.Name != "CWE-13" {
			t.Errorf("Expected %s to be %s", weakness1.Name, "CWE-13")
		}

		if weakness2 == nil {
			t.Errorf("Expected nil to be %s", "CWE-37")
		} else if weakness2.Name != "CWE-37" {
			t.Errorf("Expected %s to be %s", weakness2.Name, "CWE-37")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		weaknesses := mockupWeaknesses()

		found1 := weaknesses.Query(matchers.Weakness{
			Name: "CWE-13",
			Impact: "any",
			Scope: "any",
		})

		found2 := weaknesses.Query(matchers.Weakness{
			Name: "any",
			Impact: "Modify Application Data",
			Scope: "any",
		})

		found3 := weaknesses.Query(matchers.Weakness{
			Name: "any",
			Impact: "any",
			Scope: "Access Control",
		})

		if len(found1) == 1 {

			if found1[0].Name != "CWE-13" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "CWE-13")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=CWE-13")
		}

		if len(found2) == 2 {

			if found2[0].Name != "CWE-37" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "CWE-37")
			}

			if found2[1].Name != "CWE-370" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "CWE-370")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Impact=Modify Application Data")
		}

		if len(found3) == 2 {

			if found3[0].Name != "CWE-13" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "CWE-13")
			}

			if found3[1].Name != "CWE-370" {
				t.Errorf("Expected %s to be %s", found3[1].Name, "CWE-370")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 2, "Scope=Access Control")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		weaknesses := mockupWeaknesses()

		weakness1 := weaknesses.Get("CWE-13")

		if weakness1 == nil {
			t.Errorf("Expected nil to be %s", "CWE-13")
		} else if weakness1.Name != "CWE-13" {
			t.Errorf("Expected %s to be %s", weakness1.Name, "CWE-13")
		}

		weaknesses.Remove("CWE-13")

		weakness2 := weaknesses.Get("CWE-13")

		if weakness2 != nil {
			t.Errorf("Expected %s to be nil", weakness2.Name)
		}

	})

}
