package types

import "strconv"
import "strings"

type Domain string

func isASCIILabel(value string) bool {

	var result bool = true

	for v := 0; v < len(value); v++ {

		character := string(value[v])

		if character >= "0" && character <= "9" {
			continue
		} else if character >= "A" && character <= "Z" {
			continue
		} else if character >= "a" && character <= "z" {
			continue
		} else if character == "-" || character == "_" {
			continue
		} else {
			result = false
			break
		}

	}

	return result

}

func IsDomain(value string) bool {

	value = strings.TrimSpace(value)
	value = strings.TrimSuffix(value, ".")

	if value == "" || len(value) > 253 || strings.Contains(value, ".") == false {
		return false
	}

	labels := strings.Split(value, ".")

	if len(labels) < 2 {
		return false
	}

	for l := 0; l < len(labels); l++ {

		if len(labels[l]) < 1 || len(labels[l]) > 63 {
			return false
		}

		if isASCIILabel(labels[l]) == false {
			return false
		}

	}

	tld := labels[len(labels)-1]

	return len(tld) >= 2

}

func IsDomainAndPort(value string) bool {

	if strings.Contains(value, ":") {

		tmp := strings.Split(value, ":")

		if len(tmp) == 2 {

			_, err := strconv.ParseUint(tmp[1], 10, 16)

			if IsDomain(tmp[0]) && err == nil {
				return true
			}

		}

	}

	return false

}

func NewDomain() Domain {
	return Domain("")
}

func ParseDomain(value string) *Domain {

	var result *Domain = nil
	var domain Domain

	if domain.Parse(value) == true {
		result = &domain
	}

	return result

}

func (domain *Domain) Parse(value string) bool {

	value = strings.TrimSpace(value)
	value = strings.TrimSuffix(value, ".")

	if IsDomain(value) {
		*domain = Domain(value)
		return true
	}

	return false

}

func ParseDomainAndPort(value string) (*Domain, uint16) {

	var result_domain *Domain = nil
	var result_port uint16 = 0

	if strings.Contains(value, ":") {

		tmp1 := strings.Split(value, ":")

		if len(tmp1) == 2 {

			tmp2 := ParseDomain(tmp1[0])
			num, err := strconv.ParseUint(tmp1[1], 10, 16)

			if tmp2 != nil && err == nil {
				result_domain = tmp2
				result_port = uint16(num)
			}

		}

	}

	return result_domain, result_port

}

func ToDomain(value string) Domain {

	var domain Domain

	domain.Parse(value)

	return domain

}

func (domain Domain) Equal(other Domain) bool {
	return domain == other
}

func (domain Domain) IsValid() bool {
	return IsDomain(string(domain))
}

func (domain Domain) Bytes() []byte {
	return []byte(domain)
}

func (domain Domain) Scope() string {

	tmp := string(domain)

	if strings.HasSuffix(tmp, ".in-addr.arpa") {

		suffixes := []string{
			// RFC 6761
			"10.in-addr.arpa",
			"16.172.in-addr.arpa",
			"17.172.in-addr.arpa",
			"18.172.in-addr.arpa",
			"19.172.in-addr.arpa",
			"20.172.in-addr.arpa",
			"21.172.in-addr.arpa",
			"22.172.in-addr.arpa",
			"23.172.in-addr.arpa",
			"24.172.in-addr.arpa",
			"25.172.in-addr.arpa",
			"26.172.in-addr.arpa",
			"27.172.in-addr.arpa",
			"28.172.in-addr.arpa",
			"29.172.in-addr.arpa",
			"30.172.in-addr.arpa",
			"31.172.in-addr.arpa",
			"168.192.in-addr.arpa",

			// RFC 6762
			"254.169.in-addr.arpa",

			// RFC 8880
			"170.0.0.192.in-addr.arpa",
			"171.0.0.192.in-addr.arpa",
		}

		scope := "public"

		for s := 0; s < len(suffixes); s++ {

			if strings.HasSuffix(tmp, suffixes[s]) {
				scope = "private"
				break
			}

		}

		return scope

	} else if strings.HasSuffix(tmp, ".ip6.arpa") {

		suffixes := []string{
			// RFC 6762
			"8.e.f.ip6.arpa",
			"9.e.f.ip6.arpa",
			"a.e.f.ip6.arpa",
			"b.e.f.ip6.arpa",
		}

		scope := "public"

		for s := 0; s < len(suffixes); s++ {

			if strings.HasSuffix(tmp, suffixes[s]) {
				scope = "private"
				break
			}

		}

		return scope

	} else if strings.Contains(tmp, ".") {

		suffixes := []string{
			// RFC 6761
			".example",
			".invalid",
			".localhost",
			".test",

			// RFC 6762
			".local",

			// RFC 8375
			".home.arpa",

			// RFC 9476
			".alt",
		}

		scope := "public"

		for s := 0; s < len(suffixes); s++ {

			if strings.HasSuffix(tmp, suffixes[s]) {
				scope = "private"
				break
			}

		}

		return scope

	} else {

		return "private"

	}

}

func (domain Domain) String() string {
	return string(domain)
}

func (domain Domain) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(domain))), nil
}

func (domain *Domain) UnmarshalJSON(data []byte) error {

	unquoted, err := strconv.Unquote(string(data))

	if err != nil {
		return err
	}

	tmp := ParseDomain(unquoted)

	if tmp != nil {
		*domain = *tmp
	}

	return nil

}
