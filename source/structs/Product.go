package structs

import utils_fmt "tholian-core/utils/fmt"
import "strings"

type Product struct {
	Name    string `json:"name"`
	Vendor  string `json:"vendor"`
	Product string `json:"product"`
	Version string `json:"version"`
	Type    string `json:"type"`
	State   string `json:"state"`
}

func NewProduct(typ string) Product {

	var product Product

	product.Name = ""
	product.Vendor = "any"
	product.Product = "any"
	product.Version = "any"
	product.SetType(typ)
	product.State = "invalid"

	return product

}

func ToProduct(value string) Product {

	var product Product

	if strings.HasPrefix(value, "cpe:") {

		product.Version = "any"
		product.Vendor = "any"
		product.Type = "any"

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

	if product.Name == value.Name &&
		product.Version == value.Version &&
		product.Vendor == value.Vendor &&
		product.Type == value.Type {
		return true
	}

	return false

}

func (product *Product) IsValid() bool {

	if product.Name != "" && product.State != "invalid" {

		if product.Vendor != "" && product.Type != "" {
			return true
		}

	}

	return false

}

func (product *Product) SetProduct(value string) {

	if value == "*" || value == "-" {
		product.Product = ""
		product.Name = ""
	} else if value != "" {
		product.Product = utils_fmt.FormatASCIIName(value)
		product.Name = product.Vendor + ":" + product.Product + ":" + product.Version
	}

}

func (product *Product) SetState(value string) {

	if value == "edited" {
		product.State = "edited"
	} else if value == "published" {
		product.State = "published"
	} else if value == "invalid" {
		product.State = "invalid"
	}

}

func (product *Product) SetType(value string) {

	if value == "hardware" {
		product.Type = "hardware"
	} else if value == "software" {
		product.Type = "software"
	} else if value == "system" {
		product.Type = "system"
	}

}

func (product *Product) SetVendor(value string) {

	if value == "*" || value == "-" {
		product.Vendor = ""
		product.Name = ""
	} else if value != "" {
		product.Vendor = utils_fmt.FormatASCIIName(value)
		product.Name = product.Vendor + ":" + product.Product + ":" + product.Version
	}

}

func (product *Product) SetVersion(value string) {

	if value == "any" || value == "*" || value == "-" {
		product.Version = "any"
		product.Name = ""
	} else if value != "" {
		product.Version = utils_fmt.FormatASCIIName(value)
		product.Name = product.Vendor + ":" + product.Product + ":" + product.Version
	}

}
