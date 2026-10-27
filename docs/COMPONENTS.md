# Tholian® Core Components

This document is the complete catalog of exported components in the current
`tholian-core` module. It covers the value types in `source/types`, the nested
data structures in `source/structs`, and the selector objects in
`source/matchers`.

The JSON forms documented here are part of the API contract. See the
[README](../README.md) for the canonical wire format and general design goals.

- `source/types`: 16 exported types and their constants.
- `source/structs`: 27 exported data structures.
- `source/matchers`: 28 exported selector types.

## Types (`source/types`)

| Type | Go kind | Fields | Canonical form / JSON |
|:-----|:--------|:-------|:----------------------|
| `Architecture` | `string` | — | `"x86_64"` |
| `ASN` | `int` | — | `"AS1337"` |
| `Connection` | `struct` | `Source`, `Target`, `Protocol`, `Type` | `"[tcp] 1.3.3.7:1338 -> 1.3.3.7:443"` |
| `Datetime` | `struct` | `Year`, `Month`, `Day`, `Hour`, `Minute`, `Second` | `"2026-10-19 11:59:59"` |
| `Domain` | `string` | — | `"example.com"` |
| `Geolocation` | `struct` | `Latitude`, `Longitude` | `"geo:52.520008,13.404954"` |
| `Group` | `struct` | `ID`, `Name`, `Password`, `Type` | `"wheel:user:1000:secret"` |
| `IPv4` | `[4]byte` | — | `"1.3.3.7"` |
| `IPv6` | `[16]byte` | — | `"2001:0db8:0000:0000:0000:0000:0000:0001"` |
| `Maintainer` | `struct` | `Name`, `Email` | `"Alice <alice@example.com>"` |
| `Manager` | `string` | — | `"apt"` |
| `Protocol` | `string` | — | `"tcp"` |
| `Socket` | `struct` | `Host`, `Port`, `Scope`, `Type` | `"1.3.3.7:1338"` |
| `Time` | `struct` | `Hour`, `Minute`, `Second` | `"11:59:59"` |
| `User` | `struct` | `ID`, `Name`, `Password`, `Folder`, `Groups`, `Shell`, `Type` | `"alice:user:1000:%2Fhome%2Falice:%2Fbin%2Fbash:secret:wheel:user:1000:secret"` |
| `Version` | `struct` | `Epoche`, `Upstream`, `Revision` | `"0:1.2.3+abc123~1"` |

Canonical string notes:

- `Group` uses `name:type:id:password`, and `User` uses
  `name:type:id:folder:shell:password:groups`.
- `User` groups are a comma-separated list of canonical `Group` strings.
- Textual `Group` and `User` fields are URL query escaped.
- `Version.Upstream` contains `Major`, `Minor`, `Patch`, `Release`, and `Hash`.
  The canonical string is `epoche:major.minor.patch[release][+hash][~revision]`.
- `Version.Upstream.release`, `Version.Upstream.hash`, and `Version.revision`
  are omitted when empty.
- `Time` pads hours to two digits up to `99`; larger hours are not padded.
- `Datetime` is always `YYYY-MM-DD HH:MM:SS`.
- `IPv6` is expanded to eight lowercase groups.
- `Socket` uses `host:port`; an empty host renders as `*:port`.
- `Connection` uses `[protocol] source <direction> target`, where the
  direction is `->` for clients, `<-` for servers, and `<->` for peers.
- `Geolocation` renders coordinates with six decimal places.

## Constants and Wildcards

| Package | Constant | Value | Meaning |
|:--------|:---------|:------|:--------|
| `types`, `matchers` | `Any` | `"any"` | Wildcard or unspecified selector |
| `types` | `HostWildcard` | `"*"` | Wildcard host |
| `types` | `AddressTypeIPv4` | `"ipv4"` | IPv4 address type |
| `types` | `AddressTypeIPv6` | `"ipv6"` | IPv6 address type |
| `types` | `AddressTypeDomain` | `"domain"` | Domain address type |
| `types`, `matchers` | `ConnectionTypeClient` | `"client"` | Client connection |
| `types`, `matchers` | `ConnectionTypeServer` | `"server"` | Server connection |
| `types`, `matchers` | `ConnectionTypePeer` | `"peer"` | Peer connection |
| `types` | `UnspecifiedIPv4` | `"0.0.0.0"` | Unspecified IPv4 address |
| `types` | `UnspecifiedIPv6` | `"0000:0000:0000:0000:0000:0000:0000:0000"` | Unspecified IPv6 address |
| `types` | `ArchitectureX86` | `"x86"` | x86 architecture |
| `types` | `ArchitectureX86_64` | `"x86_64"` | x86-64 architecture |
| `types` | `ArchitectureARMv6` | `"armv6"` | ARMv6 architecture |
| `types` | `ArchitectureARMv7` | `"armv7"` | ARMv7 architecture |
| `types` | `ArchitectureARMv8` | `"armv8"` | ARMv8 architecture |
| `types` | `ArchitectureRISCV32` | `"riscv32"` | RISC-V 32-bit architecture |
| `types` | `ArchitectureRISCV64` | `"riscv64"` | RISC-V 64-bit architecture |
| `types` | `ArchitectureSPARC32` | `"sparc"` | SPARC 32-bit architecture |
| `types` | `ArchitectureSPARC64` | `"sparc64"` | SPARC 64-bit architecture |
| `types` | `ProtocolDNS` | `"dns"` | DNS protocol |
| `types` | `ProtocolDNSoverTLS` | `"dns-over-tls"` | DNS over TLS protocol |
| `types` | `ProtocolHTTPS` | `"https"` | HTTPS protocol |
| `types` | `ProtocolHTTP` | `"http"` | HTTP protocol |
| `types` | `ProtocolICMP` | `"icmp"` | ICMP protocol |
| `types` | `ProtocolSSH` | `"ssh"` | SSH protocol |
| `types` | `ProtocolSOCKS` | `"socks"` | SOCKS protocol |
| `types` | `ProtocolTCP` | `"tcp"` | TCP protocol |
| `types` | `ProtocolUDP` | `"udp"` | UDP protocol |
| `types` | `ProtocolWHOIS` | `"whois"` | WHOIS protocol |
| `matchers` | `ProtocolTCP` | `"tcp"` | TCP matcher value |
| `matchers` | `ProtocolUDP` | `"udp"` | UDP matcher value |

Supported `types.Manager` values:

- Distribution managers: `apk`, `apt`, `dnf`, `pacman`, `rpm`, `tdnf`,
  `zypper`, `pkg`, `pkgsrc`, `msi`.
- Language managers: `cargo`, `chocolatey`, `cocoapods`, `composer`, `conan`,
  `conda`, `cran`, `gem`, `go`, `gradle`, `hackage`, `hex`, `maven`, `npm`,
  `nuget`, `pear`, `phar`, `pip`.

## Structs (`source/structs`)

| Struct | Purpose | JSON fields |
|:-------|:--------|:------------|
| `Antique` | Legacy or vendor-specific package metadata | `name`, `version`, `architecture`, `manager`, `vendor`, `service`, `url` |
| `Console` | Terminal console with messages and output streams | `messages`, `stdout`, `stderr` |
| `ConsoleMessage` | Single console log entry with caller information | `datetime`, `method`, `value`, `caller` (`file`, `line`) |
| `Country` | Country metadata with geolocation, subnets, and timezones | `iso`, `name`, `continent`, `geolocation`, `population`, `allegiances`, `subnets`, `timezones`, `registry` |
| `Credential` | Credential with keys, passwords, and type | `name`, `keys`, `passwords`, `type` |
| `Device` | Hardware device with optional system and subsystem data | `name`, `bus`, `system`, `subsystem` |
| `DeviceSystem` | Vendor and device descriptor for a device system | `name`, `vendor`, `device` |
| `Distribution` | Operating system and kernel metadata | `name`, `version`, `kernel`, `kernel_architecture`, `kernel_modules`, `kernel_version`, `manager`, `vendor`, `keywords` |
| `Drive` | Drive or volume capacity and mountpoint | `name`, `size`, `free`, `used`, `mountpoint`, `type` |
| `Incident` | Incident with affected matcher collections | `system`, `type`, `datetime`, `connections`, `distributions`, `networks`, `packages`, `products`, `programs`, `users`, `vulnerabilities`, `weaknesses` |
| `Internet` | Network and subnet index | `networks` (`Mapv4` and `Mapv6` are excluded from JSON) |
| `Mitigation` | Incident and response pairing | `incident`, `response` |
| `Network` | Named network containing subnets | `name`, `subnets` |
| `Package` | Installed package metadata and dependency graph | `name`, `version`, `architecture`, `manager`, `vendor`, `url`, `datetime`, `maintainers`, `filesystem`, `conflicts`, `dependencies`, `provides`, `replaces`, `unresolved` |
| `Product` | Vendor product identification and state | `name`, `vendor`, `product`, `version`, `type`, `state` |
| `Program` | Process or service metadata | `pid`, `name`, `command`, `arguments`, `folder`, `environment`, `user`, `manager`, `filesystem`, `connections`, `dependencies`, `packages` |
| `Provider` | Hosting or network provider metadata | `name`, `address`, `country`, `subnets` |
| `Proxy` | Proxy configuration | `name`, `addresses`, `port`, `protocol` |
| `Response` | Response with affected matcher collections | `system`, `type`, `datetime`, `connections`, `distributions`, `networks`, `packages`, `products`, `programs`, `users`, `vulnerabilities`, `weaknesses` |
| `Subnet` | Network subnet with address, prefix, scope, and location | `name`, `country`, `address`, `scope`, `type`, `prefix` |
| `System` | Complete system inventory | `name`, `hostname`, `datetime`, `distribution`, `fingerprint`, `bios`, `board`, `devices`, `drives`, `incidents`, `mitigations`, `networks`, `packages`, `programs`, `responses`, `services`, `antiques`, `updates`, `users` |
| `SystemFingerprint` | System fingerprint based on country, locale, timezone, and token | `country`, `locale`, `timezone`, `token` |
| `Timeslot` | Time offset with population, weight, countries, and timezones | `offset`, `population`, `weight`, `countries`, `timezones` |
| `Timezone` | Timezone name and offset | `name`, `offset` |
| `Update` | Available update metadata | `name`, `version`, `architecture`, `manager`, `vendor`, `url` |
| `Vulnerability` | Vulnerability metadata with affected matcher collections | `name`, `description`, `type`, `severity`, `state`, `references`, `datetime`, `connections`, `devices`, `distributions`, `packages`, `products`, `programs`, `weaknesses` |
| `Weakness` | Weakness description with impacts, references, and scopes | `name`, `description`, `impacts`, `references`, `scopes` |

`Device.system` and `Device.subsystem` are omitted when empty, and
`Distribution.keywords` is omitted when nil.

## Matchers (`source/matchers`)

String selectors initialize to `Any` (`"any"`); numeric selectors such as
`Port` default to `0`. Most `Set*` methods normalize `"all"`, `"*"`, and
`"any"` to `Any` where the matcher supports wildcard values. Matchers expose `Matches(...)` and/or individual
`Matches*` methods, plus `Parse` and `Set*` helpers where applicable.

| Matcher | Matches | JSON fields |
|:--------|:--------|:------------|
| `Antique` | Package name, version, architecture, manager, vendor, service | `name`, `version`, `architecture`, `manager`, `vendor`, `service` |
| `Connection` | Host, port, protocol, connection type | `host`, `port`, `protocol`, `type` |
| `Country` | Name, continent, allegiance, subnet, timezone | `name`, `continent`, `allegiance`, `subnet`, `timezone` |
| `Credential` | Name, password, credential type | `name`, `password`, `type` |
| `Datetime` | Datetime range (`from` to `until`) | `from`, `until` |
| `Device` | Name, vendor, device, bus | `name`, `vendor`, `device`, `bus` |
| `Distribution` | Name, version, manager, vendor | `name`, `version`, `manager`, `vendor` |
| `Drive` | Name, mountpoint, drive type | `name`, `mountpoint`, `type` |
| `Incident` | System, incident type, datetime | `system`, `type`, `datetime` |
| `Manager` | Package manager name | `name` |
| `Mitigation` | System, incident type, response type, datetime | `system`, `incident_type`, `response_type`, `datetime` |
| `Network` | Network name or subnet | `name`, `subnet` |
| `Package` | Name, version, architecture, manager, vendor | `name`, `version`, `architecture`, `manager`, `vendor` |
| `Product` | Vendor, product, version, product type | `vendor`, `product`, `version`, `type` |
| `Program` | Name, command, arguments | `name`, `command`, `arguments` |
| `Provider` | Name, country, subnet | `name`, `country`, `subnet` |
| `Proxy` | Name, host, port, protocol | `name`, `host`, `port`, `protocol` |
| `Response` | System, response type, datetime | `system`, `type`, `datetime` |
| `Subnet` | Subnet name or CIDR value | `name`, `subnet` |
| `System` | Name, hostname, datetime, country, timezone | `name`, `hostname`, `datetime`, `country`, `timezone` |
| `Time` | Time range (`from` to `until`) | `from`, `until` |
| `Timeslot` | Offset, country, timezone | `offset`, `country`, `timezone` |
| `Timezone` | Name, offset | `name`, `offset` |
| `Unresolved` | Candidate packages for unresolved dependencies | `Candidates` (no JSON tag) |
| `Update` | Name, version, architecture, manager, vendor | `name`, `version`, `architecture`, `manager`, `vendor` |
| `User` | Name, password, user type | `name`, `password`, `type` |
| `Vulnerability` | Name, type, severity, state, datetime | `name`, `type`, `severity`, `state`, `datetime` |
| `Weakness` | Name, impact, scope | `name`, `impact`, `scope` |

## Hash Support

All 16 types implement `Hash() string`. Most implementations delegate to the
shared CRC32-based `toHash(...)` helper.

Current gaps:

- Structs without `Hash()`: `Console`, `ConsoleMessage`.
- Matchers without `Hash()`: `Unresolved`.

Special cases:

- Matchers that derive their hash directly from `Name`: `Network`, `Manager`, `Subnet`.
- Structs that derive their hash directly from `Name`: `Network`, `Provider`, `Proxy`.

