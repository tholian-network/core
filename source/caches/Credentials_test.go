package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "testing"

func mockupCredentials() *Credentials {

	credentials := NewCredentials()
	credentials.Add(structs.Credential{
		Name: "root",
		Passwords: []string{
			"root",
			"ubuntu",
			"admin",
			"adminadmin",
		},
		Type: "user",
	})
	credentials.Add(structs.Credential{
		Name: "admin",
		Passwords: []string{
			"admin",
			"admin123",
			"adminadmin",
		},
		Type: "user",
	})
	credentials.Add(structs.Credential{
		Name: "postgres",
		Passwords: []string{
			"mysql",
			"postgres",
			"postgres123",
		},
		Type: "program",
	})

	return credentials

}

func TestCredentials(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		credentials := mockupCredentials()

		credential1 := credentials.Get("ubuntu")

		if credential1 != nil {
			t.Errorf("Expected %s to be nil", credential1.Name)
		}

		credentials.Add(structs.Credential{
			Name: "ubuntu",
			Passwords: []string{
				"ubuntu",
			},
			Type: "user",
		})

		credential2 := credentials.Get("ubuntu")

		if credential2 == nil {
			t.Errorf("Expected nil to be %s", "ubuntu")
		} else if credential2.Name != "ubuntu" {
			t.Errorf("Expected %s to be %s", credential2.Name, "ubuntu")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		credentials := mockupCredentials()

		credential1 := credentials.Get("root")
		credential2 := credentials.Get("postgres")

		if credential1 == nil {
			t.Errorf("Expected nil to be %s", "root")
		} else if credential1.Name != "root" {
			t.Errorf("Expected %s to be %s", credential1.Name, "root")
		}

		if credential2 == nil {
			t.Errorf("Expected nil to be %s", "postgres")
		} else if credential2.Name != "postgres" {
			t.Errorf("Expected %s to be %s", credential2.Name, "postgres")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		credentials := mockupCredentials()

		found1 := credentials.Query(matchers.Credential{
			Name: "postgres",
			Password: "any",
			Type: "any",
		})

		found2 := credentials.Query(matchers.Credential{
			Name: "any",
			Password: "admin",
			Type: "any",
		})

		found3 := credentials.Query(matchers.Credential{
			Name: "any",
			Password: "any",
			Type: "program",
		})

		if len(found1) == 1 {

			if found1[0].Name != "postgres" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "postgres")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=postgres")
		}

		if len(found2) == 2 {

			if found2[0].Name != "admin" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "admin")
			}

			if found2[1].Name != "root" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "root")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "Password=admin")
		}

		if len(found3) == 1 {

			if found3[0].Name != "postgres" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "postgres")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Type=program")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		credentials := mockupCredentials()

		found1 := credentials.Get("admin")

		if found1 == nil {
			t.Errorf("Expected nil to be %s", "admin")
		} else if found1.Name != "admin" {
			t.Errorf("Expected %s to be %s", found1.Name, "admin")
		}

		credentials.Remove("admin")

		found2 := credentials.Get("admin")

		if found2 != nil {
			t.Errorf("Expected %s to be nil", found2.Name)
		}

	})

}
