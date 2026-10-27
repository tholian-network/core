package matchers

import "tholian-core/types"
import utils_fmt "tholian-core/utils/fmt"
import "strings"

type Product struct {
	Vendor  string `json:"vendor"`
	Product string `json:"product"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

func NewProduct() Product {

	var product Product

	product.Vendor = Any
	product.Product = Any
	product.Version = Any
	product.Type = Any

	return product

}

func ToProduct(value string) Product {

	var product Product

	if strings.HasPrefix(value, "cpe:") {

		product.Vendor = Any
		product.Product = Any
		product.Version = Any
		product.Type = Any

		value = strings.ReplaceAll(value[4:], "\\:", "%COLON%")
		tmp := strings.Split(strings.TrimSpace(value), ":")

		for t := 0; t < len(tmp); t++ {
			tmp[t] = strings.ReplaceAll(tmp[t], "%COLON%", ":")
		}

		if len(tmp) >= 4 && tmp[0] == "2.3" {

			cpe_part := parseCPEValue(tmp[1])

			if cpe_part == "/a" || cpe_part == "a" {
				product.SetType("software")
			} else if cpe_part == "/h" || cpe_part == "h" {
				product.SetType("hardware")
			} else if cpe_part == "/o" || cpe_part == "o" {
				product.SetType("system")
			}

			cpe_vendor := parseCPEValue(tmp[2])

			if cpe_vendor != "" {
				product.SetVendor(cpe_vendor)
			}

			cpe_product := parseCPEValue(tmp[3])

			if cpe_product != "" {
				product.SetProduct(cpe_product)
			}

			if len(tmp) >= 6 {

				cpe_version := parseCPEValue(tmp[4])
				cpe_update := parseCPEValue(tmp[5])

				cpe_version = stripCPEArchitecture(cpe_version)
				cpe_update = stripCPEArchitecture(cpe_update)

				if cpe_version != "" && cpe_update != "" {
					product.SetVersion(cpe_version + "-" + cpe_update)
				} else if cpe_version != "" {
					product.SetVersion(cpe_version)
				} else if cpe_update != "" {
					product.SetVersion(cpe_update)
				}

			} else if len(tmp) == 5 {

				cpe_version := parseCPEValue(tmp[4])
				cpe_version = stripCPEArchitecture(cpe_version)

				if cpe_version != "" {
					product.SetVersion(cpe_version)
				}

			}

		} else if len(tmp) >= 3 {

			cpe_part := parseCPEValue(tmp[0])

			if cpe_part == "/a" {
				product.SetType("software")
			} else if cpe_part == "/h" {
				product.SetType("hardware")
			} else if cpe_part == "/o" {
				product.SetType("system")
			}

			cpe_vendor := parseCPEValue(tmp[1])

			if cpe_vendor != "" {
				product.SetVendor(cpe_vendor)
			}

			cpe_product := parseCPEValue(tmp[2])

			if cpe_product != "" {
				product.SetProduct(cpe_product)
			}

			if len(tmp) >= 5 {

				cpe_version := parseCPEValue(tmp[3])
				cpe_update := parseCPEValue(tmp[4])

				cpe_version = stripCPEArchitecture(cpe_version)
				cpe_update = stripCPEArchitecture(cpe_update)

				if cpe_version != "" && cpe_update != "" {
					product.SetVersion(cpe_version + "-" + cpe_update)
				} else if cpe_version != "" {
					product.SetVersion(cpe_version)
				} else if cpe_update != "" {
					product.SetVersion(cpe_update)
				}

			} else if len(tmp) == 4 {

				cpe_version := parseCPEValue(tmp[3])
				cpe_version = stripCPEArchitecture(cpe_version)

				if cpe_version != "" {
					product.SetVersion(cpe_version)
				}

			}

		}

	}

	return product

}

func (product *Product) Equal(value Product) bool {

	if product.Vendor == value.Vendor &&
		product.Product == value.Product &&
		product.Version == value.Version &&
		product.Type == value.Type {
		return true
	}

	return false

}

func (product *Product) IsValid() bool {

	if product.Vendor != Any ||
		product.Product != Any ||
		product.Version != Any ||
		product.Type != Any {
		return true
	}

	return false

}

func (product *Product) Matches(vendor string, product_ string, version string, typ string) bool {
	return product.MatchesVendor(vendor) && product.MatchesProduct(product_) && product.MatchesVersion(version) && product.MatchesType(typ)
}

func (product *Product) MatchesProduct(value string) bool {

	if product.Product == value {
		return true
	} else if product.Product == Any {
		return true
	}

	return false

}

func (product *Product) MatchesType(value string) bool {

	if product.Type == value {
		return true
	} else if product.Type == Any {
		return true
	}

	return false

}

func (product *Product) MatchesVendor(value string) bool {

	if product.Vendor == value {
		return true
	} else if product.Vendor == Any {
		return true
	}

	return false

}

func (product *Product) MatchesVersion(value string) bool {

	// Compatibility with "<operator> <version>" syntax
	if strings.Contains(value, " ") {
		value = strings.TrimSpace(value[strings.Index(value, " ")+1:])
	}


	if product.Version == Any {

		return true

	} else if strings.HasPrefix(product.Version, "<= ") {

		product_version := types.ToVersion(product.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(product_version) {
			return true
		} else if other_version.IsBefore(product_version) {
			return true
		}

	} else if strings.HasPrefix(product.Version, "< ") {

		product_version := types.ToVersion(product.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsBefore(product_version) {
			return true
		}

	} else if strings.HasPrefix(product.Version, ">= ") {

		product_version := types.ToVersion(product.Version[3:])
		other_version := types.ToVersion(value)

		if other_version.Equal(product_version) {
			return true
		} else if other_version.IsAfter(product_version) {
			return true
		}

	} else if strings.HasPrefix(product.Version, "> ") {

		product_version := types.ToVersion(product.Version[2:])
		other_version := types.ToVersion(value)

		if other_version.IsAfter(product_version) {
			return true
		}

	} else {

		product_version := types.ToVersion(product.Version)
		other_version := types.ToVersion(value)

		if other_version.Equal(product_version) {
			return true
		}

	}

	return false

}

func (product *Product) SetProduct(value string) {

	if value == Any || value == "*" || value == "-" {
		product.Product = Any
	} else if value != "" {
		product.Product = utils_fmt.FormatASCIIName(value)
	}

}

func (product *Product) SetType(value string) {

	if value == Any || value == "*" || value == "-" {
		product.Type = Any
	} else if value == "hardware" {
		product.Type = "hardware"
	} else if value == "software" {
		product.Type = "software"
	} else if value == "system" {
		product.Type = "system"
	}

}

func (product *Product) SetVendor(value string) {

	if value == Any || value == "*" || value == "-" {
		product.Vendor = Any
	} else if value != "" {
		product.Vendor = utils_fmt.FormatASCIIName(value)
	}

}

func (product *Product) SetVersion(value string) {

	if value == "all" || value == "any" || value == "*" || value == "-" {
		product.Version = Any
	} else if value != "" {
		product.Version = value
	}

}

func (product *Product) Hash() string {

	hash := ""

	if product.Vendor != "" && product.Product != "" {

		hash = toHash(
			product.Vendor,
			product.Product,
			product.Version,
			product.Type,
		)

	}

	return hash

}
