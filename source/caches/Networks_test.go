package caches

import "tholian-core/matchers"
import "tholian-core/structs"
// import "tholian-core/types"
import "testing"

func mockupNetworks() *Networks {

	networks := NewNetworks()
	networks.Add(structs.Network{
		Name: "AS1337",
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name: "AS1337",
				Country: "US",
				Address: "13.37.0.0",
				Scope: "public",
				Type: "ipv4",
				Prefix: 16,
			},
		},
	})
	networks.Add(structs.Network{
		Name: "AS1338",
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name: "AS1338",
				Country: "DE",
				Address: "13.38.0.0",
				Scope: "public",
				Type: "ipv4",
				Prefix: 16,
			},
		},
	})
	networks.Add(structs.Network{
		Name: "AS13",
		Subnets: []structs.Subnet{
			structs.Subnet{
				Name: "AS13",
				Country: "US",
				Address: "13.0.0.0",
				Scope: "public",
				Type: "ipv4",
				Prefix: 8,
			},
		},
	})

	return networks

}

func TestNetworks(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		networks := mockupNetworks()

		network1 := networks.Get("AS13337")

		if network1 != nil {
			t.Errorf("Expected %s to be nil", network1.Name)
		}

		networks.Add(structs.Network{
			Name: "AS133337",
			Subnets: []structs.Subnet{
				structs.Subnet{
					Name: "AS133337",
					Country: "US",
					Address: "13.33.37.0",
					Scope: "public",
					Type: "ipv4",
					Prefix: 24,
				},
			},
		})

		network2 := networks.Get("AS133337")

		if network2 == nil {
			t.Errorf("Expected nil to be %s", "AS133337")
		} else if network2.Name != "AS133337" {
			t.Errorf("Expected %s to be %s", network2.Name, "AS133337")
		}

	})

	t.Run("AddSubnet()", func(t *testing.T) {

		networks := mockupNetworks()

		network1 := networks.GetByIP("13.33.33.34")

		if network1 == nil {
			t.Errorf("Expected nil to be %s", network1.Name)
		} else if network1.Name != "AS13" {
			t.Errorf("Expected %s to be %s", network1.Name, "AS13")
		}

		networks.Add(structs.Network{
			Name: "AS133333",
			Subnets: []structs.Subnet{},
		})

		network2 := networks.GetByIP("13.33.33.34")

		if network2 == nil {
			t.Errorf("Expected nil to be %s", network2.Name)
		} else if network2.Name != "AS13" {
			t.Errorf("Expected %s to be %s", network2.Name, "AS13")
		}

		networks.AddSubnet(structs.Subnet{
			Name: "AS133333",
			Country: "US",
			Address: "13.33.33.0",
			Scope: "public",
			Type: "ipv4",
			Prefix: 24,
		})

		network3 := networks.GetByIP("13.33.33.34")

		if network3 == nil {
			t.Errorf("Expected nil to be %s", network3.Name)
		} else if network3.Name != "AS133333" {
			t.Errorf("Expected %s to be %s", network3.Name, "AS133333")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		networks := mockupNetworks()

		network1 := networks.Get("AS1337")
		network2 := networks.Get("AS1338")

		if network1 == nil {
			t.Errorf("Expected nil to be %s", "AS1337")
		} else if network1.Name != "AS1337" {
			t.Errorf("Expected %s to be %s", network1.Name, "AS1337")
		}

		if network2 == nil {
			t.Errorf("Expected nil to be %s", "AS1338")
		} else if network2.Name != "AS1338" {
			t.Errorf("Expected %s to be %s", network1.Name, "AS1338")
		}

	})

	t.Run("GetByIP()", func(t *testing.T) {

		networks := mockupNetworks()

		network1 := networks.GetByIP("13.30.30.34")
		network2 := networks.GetByIP("13.37.12.34")
		network3 := networks.GetByIP("13.38.12.34")

		if network1 == nil {
			t.Errorf("Expected nil to be %s", "AS13")
		} else if network1.Name != "AS13" {
			t.Errorf("Expected %s to be %s", network1.Name, "AS13")
		}

		if network2 == nil {
			t.Errorf("Expected nil to be %s", "AS1337")
		} else if network2.Name != "AS1337" {
			t.Errorf("Expected %s to be %s", network2.Name, "AS1337")
		}

		if network3 == nil {
			t.Errorf("Expected nil to be %s", "AS1338")
		} else if network3.Name != "AS1338" {
			t.Errorf("Expected %s to be %s", network3.Name, "AS1338")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		networks := mockupNetworks()

		found1 := networks.Query(matchers.Network{
			Name: "AS1337",
			Subnet: "any",
		})

		found2 := networks.Query(matchers.Network{
			Name: "AS1337",
			Subnet: "13.37.12.34/32",
		})

		found3 := networks.Query(matchers.Network{
			Name: "any",
			Subnet: "13.38.12.34/32",
		})

		if len(found1) == 1 {

			if found1[0].Name != "AS1337" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "AS1337")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=AS1337")
		}

		if len(found2) == 1 {

			if found2[0].Name != "AS1337" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "AS1337")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Name=AS1337,Subnet=13.37.12.34/32")
		}

		if len(found3) == 2 {

			if found3[0].Name != "AS13" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "AS13")
			}

			if found3[1].Name != "AS1338" {
				t.Errorf("Expected %s to be %s", found3[1].Name, "AS1338")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Name=AS1338,Subnet=13.38.12.34/32")
		}

	})

	t.Run("QueryByIP()", func(t *testing.T) {

		networks := mockupNetworks()

		found1 := networks.QueryByIP("13.37.12.34")
		found2 := networks.QueryByIP("13.38.12.34")

		if len(found1) == 2 {

			if found1[0].Name != "AS1337" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "AS1337")
			}

			if found1[1].Name != "AS13" {
				t.Errorf("Expected %s to be %s", found1[1].Name, "AS13")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 2, "13.37.12.34")
		}

		if len(found2) == 2 {

			if found2[0].Name != "AS1338" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "AS1338")
			}

			if found2[1].Name != "AS13" {
				t.Errorf("Expected %s to be %s", found2[1].Name, "AS13")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 2, "13.38.12.34")
		}
	})

	t.Run("Remove()", func(t *testing.T) {

		networks := mockupNetworks()

		network1 := networks.Get("AS1337")

		if network1 == nil {
			t.Errorf("Expected nil to be %s", "AS1337")
		} else if network1.Name != "AS1337" {
			t.Errorf("Expected %s to be %s", network1.Name, "AS1337")
		}

		networks.Remove("AS1337")

		network2 := networks.Get("AS1337")

		if network2 != nil {
			t.Errorf("Expected %s to be nil", network2.Name)
		}

	})

	t.Run("RemoveSubnet()", func(t *testing.T) {

		networks := mockupNetworks()

		network1 := networks.GetByIP("13.37.12.34")

		if network1 == nil {
			t.Errorf("Expected nil to be %s", "AS1337")
		} else if network1.Name != "AS1337" {
			t.Errorf("Expected %s to be %s", network1.Name, "AS1337")
		}

		subnets := networks.QueryByIP("13.37.12.34")

		if len(subnets) != 2 {
			t.Errorf("Expected %d results to be %d for query %s", len(subnets), 2, "13.37.12.34")
		}

		networks.RemoveSubnet(*subnets[0])

		network2 := networks.GetByIP("13.37.12.34")

		if network2 == nil {
			t.Errorf("Expected nil to be %s", "AS13")
		} else if network2.Name != "AS13" {
			t.Errorf("Expected %s to be %s", network2.Name, "AS13")
		}

	})

}
