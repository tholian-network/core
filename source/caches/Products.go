package caches

import "tholian-core/matchers"
import "tholian-core/structs"
import "tholian-core/types"
import "encoding/json"
import "sort"
import "strings"
import "sync"

type Products struct {
	Products map[string]*structs.Product             `json:"products"`
	mapcpe   map[string]map[string]map[string]string `json:"-"`
	mutex    *sync.RWMutex                           `json:"-"`
}

func NewProducts() *Products {

	var cache Products

	cache.Products = make(map[string]*structs.Product)
	cache.mapcpe = make(map[string]map[string]map[string]string)
	cache.mutex = &sync.RWMutex{}

	return &cache

}

func (cache *Products) Add(value structs.Product) bool {

	cache.mutex.Lock()

	var result bool

	if value.Type != "" {

		cache.Products[value.Name] = &value
		result = true

		_, ok1 := cache.mapcpe[value.Vendor]

		if ok1 == false {
			cache.mapcpe[value.Vendor] = make(map[string]map[string]string)
		}

		_, ok2 := cache.mapcpe[value.Vendor][value.Product]

		if ok2 == false {
			cache.mapcpe[value.Vendor][value.Product] = make(map[string]string)
		}

		_, ok3 := cache.mapcpe[value.Vendor][value.Product][value.Version]

		if ok3 == false {
			cache.mapcpe[value.Vendor][value.Product][value.Version] = value.Name
		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Products) Get(name string) *structs.Product {

	cache.mutex.RLock()

	var result *structs.Product = nil

	if strings.Contains(name, ":") {

		tmp := strings.Split(name, ":")

		if len(tmp) == 3 {

			product, ok := cache.Products[name]

			if ok == true {
				result = product
			}

		}

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Products) Length() int {
	return len(cache.Products)
}

func (cache *Products) Query(query matchers.Product) []*structs.Product {

	cache.mutex.RLock()

	result := make([]*structs.Product, 0)
	found := make(map[string]*structs.Product)

	if query.Vendor != "any" && query.Product != "any" && query.Version != "any" {

		// Cache Optimizations for specified query.Vendor, query.Product, query.Version
		_, ok1 := cache.mapcpe[query.Vendor]

		if ok1 == true {

			_, ok2 := cache.mapcpe[query.Vendor][query.Product]

			if ok2 == true {

				name, ok3 := cache.mapcpe[query.Vendor][query.Product][query.Version]

				if ok3 == true {

					product, ok := cache.Products[name]

					if ok == true {

						matches_type := query.MatchesType(product.Type)

						if matches_type {
							found[product.Name] = product
						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Product != "any" && query.Version == "any" {

		// Cache Optimizations for specified query.Vendor, query.Product

		_, ok1 := cache.mapcpe[query.Vendor]

		if ok1 == true {

			_, ok2 := cache.mapcpe[query.Vendor][query.Product]

			if ok2 == true {

				for _, name := range cache.mapcpe[query.Vendor][query.Product] {

					product, ok := cache.Products[name]

					if ok == true {

						matches_type := query.MatchesType(product.Type)

						if matches_type {
							found[product.Name] = product
						}

					}

				}

			}

		}

	} else if query.Vendor != "any" && query.Product == "any" && query.Version == "any" {

		// Cache Optimizations for specified query.Vendor

		_, ok1 := cache.mapcpe[query.Vendor]

		if ok1 == true {

			for _, versions := range cache.mapcpe[query.Vendor] {

				for _, name := range versions {

					product, ok := cache.Products[name]

					if ok == true {

						matches_type := query.MatchesType(product.Type)

						if matches_type {
							found[product.Name] = product
						}

					}

				}

			}

		}

	} else {

		for _, product := range cache.Products {

			matches_vendor := query.MatchesVendor(product.Vendor)
			matches_product := query.MatchesProduct(product.Product)
			matches_version := query.MatchesVersion(product.Version)
			matches_type := query.MatchesType(product.Type)

			if matches_vendor && matches_product && matches_version && matches_type {
				found[product.Name] = product
			}

		}

	}

	if len(found) > 0 {

		for _, product := range found {
			result = append(result, product)
		}

	}

	if len(result) > 1 {

		sort.Slice(result, func(a int, b int) bool {

			if result[a].Vendor == result[b].Vendor {

				if result[a].Product == result[b].Product {

					if result[a].Version == result[b].Version {

						return false

					} else {

						version_a := types.ToVersion(result[a].Version)
						version_b := types.ToVersion(result[b].Version)

						return version_a.IsBefore(version_b)

					}

				} else {
					return result[a].Product < result[b].Product
				}

			} else {
				return result[a].Vendor < result[b].Vendor
			}

		})

	}

	cache.mutex.RUnlock()

	return result

}

func (cache *Products) Remove(name string) bool {

	cache.mutex.Lock()

	var result bool

	if name != "" {

		product, ok := cache.Products[name]

		if ok == true {

			delete(cache.Products, name)
			result = true

			_, ok1 := cache.mapcpe[product.Vendor]

			if ok1 == true {

				_, ok2 := cache.mapcpe[product.Vendor][product.Product]

				if ok2 == true {

					_, ok3 := cache.mapcpe[product.Vendor][product.Product]

					if ok3 == true {

						delete(cache.mapcpe[product.Vendor][product.Product], product.Version)

						if len(cache.mapcpe[product.Vendor][product.Product]) == 0 {

							delete(cache.mapcpe[product.Vendor], product.Product)

							if len(cache.mapcpe[product.Vendor]) == 0 {
								delete(cache.mapcpe, product.Vendor)
							}

						}

					}

				}

			}

		}

	}

	cache.mutex.Unlock()

	return result

}

func (cache *Products) MarshalJSON() ([]byte, error) {

	tmp := make(map[string]*structs.Product)
	names := make([]string, 0)

	for name, _ := range cache.Products {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		tmp[name] = cache.Products[name]
	}

	return json.MarshalIndent(&struct {
		Products map[string]*structs.Product `json:"products"`
	}{
		Products: tmp,
	}, "", "\t")

}

func (cache *Products) UnmarshalJSON(data []byte) error {

	var tmp struct {
		Products map[string]*structs.Product `json:"products"`
	}

	err := json.Unmarshal(data, &tmp)

	if err != nil {
		return err
	}

	for _, product := range tmp.Products {
		cache.Add(*product)
	}

	return nil

}
