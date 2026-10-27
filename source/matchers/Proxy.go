package matchers

import "tholian-core/types"
import "encoding/binary"
import "encoding/hex"
import "hash/crc32"
import "strconv"
import "strings"

type Proxy struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Protocol string `json:"protocol"`
}

func NewProxy() Proxy {

	var proxy Proxy

	proxy.Name = Any
	proxy.Host = Any
	proxy.Port = 0
	proxy.Protocol = Any

	return proxy

}

func ToProxy(value string) Proxy {

	var proxy Proxy

	proxy.Name = Any
	proxy.Host = Any
	proxy.Port = 0
	proxy.Protocol = Any

	proxy.SetName(value)

	return proxy

}

func (proxy *Proxy) IsIdentical(value Proxy) bool {

	if proxy.Name == value.Name &&
		proxy.Host == value.Host &&
		proxy.Port == value.Port &&
		proxy.Protocol == value.Protocol {
		return true
	}

	return false

}

func (proxy *Proxy) IsValid() bool {

	if proxy.Name != Any ||
		proxy.Host != Any ||
		(uint(proxy.Port) > 0 && uint(proxy.Port) < 65535) ||
		proxy.Protocol != Any {
		return true
	}

	return false

}

func (proxy *Proxy) Matches(name string, host string, port uint16, protocol string) bool {
	return proxy.MatchesName(name) && proxy.MatchesHost(host) && proxy.MatchesPort(port) && proxy.MatchesProtocol(protocol)
}

func (proxy *Proxy) MatchesHost(value string) bool {

	if proxy.Host == value {
		return true
	} else if proxy.Host != Any && value != "any" {
		return containsSubnet(value, proxy.Host)
	} else if proxy.Host == Any {
		return true
	}

	return false

}

func (proxy *Proxy) MatchesName(value string) bool {

	if proxy.Name == value {
		return true
	} else if proxy.Name == Any {
		return true
	}

	return false

}

func (proxy *Proxy) MatchesPort(value uint16) bool {

	if proxy.Port == value {
		return true
	} else if proxy.Port == 0 {
		return true
	}

	return false

}

func (proxy *Proxy) MatchesProtocol(value string) bool {

	if proxy.Protocol == value {
		return true
	} else if proxy.Protocol == Any {
		return true
	}

	return false

}

func (proxy *Proxy) SetHost(value string) {

	if value == "all" || value == "any" || value == "*" {

		proxy.Host = Any

	} else if types.IsIPv6(value) {

		ipv6 := types.ParseIPv6(value)

		if ipv6 != nil {
			proxy.Host = ipv6.String()
		}

	} else if types.IsIPv4(value) {

		ipv4 := types.ParseIPv4(value)

		if ipv4 != nil {
			proxy.Host = ipv4.String()
		}

	} else if strings.Contains(value, "/") {

		address, prefix := toSubnet(value)

		if address != "" && prefix != 0 {
			proxy.Host = value
		}

	}

}

func (proxy *Proxy) SetName(value string) {
	proxy.Name = strings.TrimSpace(value)
}

func (proxy *Proxy) SetProtocol(value string) {

	if value == "all" || value == "any" || value == "*" {
		proxy.Protocol = Any
	} else if value != "" {
		proxy.Protocol = value
	}

}

func (proxy *Proxy) Hash() string {

	hash := ""

	if proxy.Name != "" && proxy.Port != 0 {

		checksum := crc32.ChecksumIEEE([]byte(strings.Join([]string{
			proxy.Name,
			proxy.Host,
			strconv.FormatUint(uint64(proxy.Port), 10),
			proxy.Protocol,
		}, "-")))

		tmp := make([]byte, 4)
		binary.LittleEndian.PutUint32(tmp, checksum)
		hash = hex.EncodeToString(tmp)

	}

	return hash

}
