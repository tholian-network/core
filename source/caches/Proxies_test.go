package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "testing"

func mockupProxies() *Proxies {

	proxies := NewProxies()
	proxies.Add(structs.Proxy{
		Name:      "dynamic-ip-ads1.viettel.vn",
		Addresses: []string{"171.249.209.84"},
		Port:      10001,
		Protocol:  types.ProtocolHTTPS,
	})
	proxies.Add(structs.Proxy{
		Name:      "31.40.248.2",
		Addresses: []string{"31.40.248.2"},
		Port:      8080,
		Protocol:  types.ProtocolHTTP,
	})

	return proxies

}

func TestProxies(t *testing.T) {

	t.Run("Add()", func(t *testing.T) {

		proxies := mockupProxies()

		proxy1 := proxies.Get("postmaster.bitcommand.com")

		if proxy1 != nil {
			t.Errorf("Expected %s to be nil", proxy1.Name)
		}

		proxies.Add(structs.Proxy{
			Name:      "postmaster.bitcommand.com",
			Addresses: []string{"154.91.170.145"},
			Port:      8080,
			Protocol:  types.ProtocolHTTP,
		})

		proxy2 := proxies.Get("postmaster.bitcommand.com")

		if proxy2 == nil {
			t.Errorf("Expected nil to be %s", "postmaster.bitcommand.com")
		} else if proxy2.Name != "postmaster.bitcommand.com" {
			t.Errorf("Expected %s to be %s", proxy2.Name, "postmaster.bitcommand.com")
		}

	})

	t.Run("Get()", func(t *testing.T) {

		proxies := mockupProxies()

		proxy1 := proxies.Get("dynamic-ip-ads1.viettel.vn")
		proxy2 := proxies.Get("31.40.248.2")

		if proxy1 == nil {
			t.Errorf("Expected nil to be %s", "dynamic-ip-ads1.viettel.vn")
		} else if proxy1.Name != "dynamic-ip-ads1.viettel.vn" {
			t.Errorf("Expected %s to be %s", proxy1.Name, "dynamic-ip-ads1.viettel.vn")
		}

		if proxy2 == nil {
			t.Errorf("Expected nil to be %s", "31.40.248.2")
		} else if proxy2.Name != "31.40.248.2" {
			t.Errorf("Expected %s to be %s", proxy2.Name, "31.40.248.2")
		}

	})

	t.Run("Query()", func(t *testing.T) {

		proxies := mockupProxies()

		found1 := proxies.Query(matchers.Proxy{
			Name: "dynamic-ip-ads1.viettel.vn",
			Host: "any",
			Port: 0,
			Protocol: "any",
		})

		found2 := proxies.Query(matchers.Proxy{
			Name: "any",
			Host: "31.40.248.2/32",
			Port: 0,
			Protocol: "any",
		})

		found3 := proxies.Query(matchers.Proxy{
			Name: "any",
			Host: "any",
			Port: 8080,
			Protocol: "any",
		})

		found4 := proxies.Query(matchers.Proxy{
			Name: "any",
			Host: "any",
			Port: 0,
			Protocol: "https",
		})

		if len(found1) == 1 {

			if found1[0].Name != "dynamic-ip-ads1.viettel.vn" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "dynamic-ip-ads1.viettel.vn")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "Name=dynamic-ip-ads1.viettel.vn")
		}

		if len(found2) == 1 {

			if found2[0].Name != "31.40.248.2" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "31.40.248.2")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "Subnet=31.40.248.2/32")
		}

		if len(found3) == 1 {

			if found3[0].Name != "31.40.248.2" {
				t.Errorf("Expected %s to be %s", found3[0].Name, "31.40.248.2")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found3), 1, "Protocol=https")
		}

		if len(found4) == 1 {

			if found4[0].Name != "dynamic-ip-ads1.viettel.vn" {
				t.Errorf("Expected %s to be %s", found4[0].Name, "dynamic-ip-ads1.viettel.vn")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found4), 1, "Protocol=https")
		}

	})

	t.Run("QueryByIP()", func(t *testing.T) {

		proxies := mockupProxies()

		found1 := proxies.QueryByIP("171.249.209.84")
		found2 := proxies.QueryByIP("31.40.248.2")

		if len(found1) == 1 {

			if found1[0].Name != "dynamic-ip-ads1.viettel.vn" {
				t.Errorf("Expected %s to be %s", found1[0].Name, "dynamic-ip-ads1.viettel.vn")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found1), 1, "171.249.209.84")
		}

		if len(found2) == 1 {

			if found2[0].Name != "31.40.248.2" {
				t.Errorf("Expected %s to be %s", found2[0].Name, "31.40.248.2")
			}

		} else {
			t.Errorf("Expected %d results to be %d for query %s", len(found2), 1, "31.40.248.2")
		}

	})

	t.Run("Remove()", func(t *testing.T) {

		proxies := mockupProxies()

		proxy1 := proxies.Get("dynamic-ip-ads1.viettel.vn")

		if proxy1 == nil {
			t.Errorf("Expected nil to be %s", "dynamic-ip-ads1.viettel.vn")
		} else if proxy1.Name != "dynamic-ip-ads1.viettel.vn" {
			t.Errorf("Expected %s to be %s", proxy1.Name, "dynamic-ip-ads1.viettel.vn")
		}

		proxies.Remove("dynamic-ip-ads1.viettel.vn")

		proxy2 := proxies.Get("dynamic-ip-ads1.viettel.vn")

		if proxy2 != nil {
			t.Errorf("Expected %s to be nil", proxy2.Name)
		}

	})

}
