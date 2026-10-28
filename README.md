# Tholian® Core Type System

Tholian® Core is the shared type system for the Tholian Network. It defines the
canonical representation of all data types and schemas that are used across
Tholian's cyber defense projects.

The project is not a runtime, database, or transport protocol. It is a small Go
module that gives every consumer the same vocabulary and the same JSON wire
format. That makes it possible to correlate data from all Tholian Agent
implementations and our provided threat intelligence tooling without repeatedly
reinterpreting the same values.

The JSON representations are intentionally documented and stable, so that they
can be consumed by external tools and proprietary cyber security stacks such as
Elastic, Kibana, SIEM, SOAR, or custom dashboards.

## Design Goals

- **Canonical**: All supported values have exactly one normalized representation.
- **Validation**: All parsers check ranges, prefixes, labels, and protocol names.
- **JSON-first**: All types implement `MarshalJSON` and `UnmarshalJSON` directly.
- **Portability**: The implementation is standard-library-first, with only
  `x/term` required for terminal handling.
- **Composability**: Sockets combine with subnets, connections combine with matchers,
  and networks group subnets without duplicating them.
- **Interoperability**: Canonical strings map cleanly to native field types in
  external search and analytics platforms.

## Repository Layout

| Path                                       | Purpose                             |
|:-------------------------------------------|:------------------------------------|
| [source/types](./source/types)             | Value and composite types           |
| [source/structs](./source/structs)         | Nested data structures              |
| [source/matchers](./source/matchers)       | Matcher rules for types and structs |
| [docs/COMPONENTS.md](./docs/COMPONENTS.md) | Complete component catalog          |

The Go module is named `tholian-core` and targets Go `1.27.1`.

## Canonical Wire Format

All core types implement `encoding/json` support. Simple values are encoded as
JSON strings, while structured values such as `Network` and `Internet` use their
declared JSON fields.

See [docs/COMPONENTS.md](./docs/COMPONENTS.md) for the complete catalog of all
types, structs, matchers, constants, and their canonical forms.

A network with one subnet is encoded as:

```json
{
  "name": "AS13335",
  "subnets": [
    {
      "name": "",
      "country": "",
      "address": "1.1.1.0",
      "scope": "public",
      "type": "ipv4",
      "prefix": 24
    }
  ]
}
```

The exact string forms are part of the API contract.

IPv6 addresses are expanded to eight lowercase groups, sockets use `host:port`,
subnets use `address/prefix`, and domains are stored without a trailing dot.

`Group` and `User` values use the canonical forms
`name:type:id:password` and `name:type:id:folder:shell:password:groups`.
Textual fields are URL query escaped, and `groups` is a comma-separated list of
canonical `Group` strings. Passwords are part of the wire format, so payloads
containing users or groups must be handled as sensitive data.

Go's `encoding/json` escapes characters such as `>` as `\u003e` by default;
decoding restores the canonical value.

## Usage

The canonical JSON forms are designed to be published as standalone JSON Schema
definitions under `schemas/` using JSON Schema Draft 2020-12. The schema files
are being prepared for public consumption, so that consumers do not need a Go
toolchain to interpret Tholian data.

The Go implementations and the unit tests in this repository are the reference
for the wire format until the schema files are published. Once available, the
schemas will follow semantic versioning:

- Patch releases clarify descriptions or examples without changing validation.
- Minor releases add optional fields or new type definitions.
- Major releases may change canonical string forms or required fields.

### Usage: Elastic and Kibana

The preferred mapping strategy is to keep the canonical value intact and use a
native field type where one exists:

| Tholian value        | Suggested Elasticsearch / Kibana field type |
|:---------------------|:---------------------------------------------|
| IPv4 / IPv6 address  | `ip`                                         |
| Subnet               | `keyword` or `ip` with a CIDR-aware pipeline |
| ASN                  | `keyword`                                    |
| Domain               | `keyword`                                    |
| Socket               | `keyword`                                    |
| Connection           | `keyword`                                    |
| Protocol             | `keyword`                                    |
| Architecture         | `keyword`                                    |
| Time / Datetime      | `date` or `keyword` depending on precision   |

External consumers should preserve the original canonical string alongside any
normalized field. This keeps the source value available for auditing and avoids
losing information when a downstream index uses a different representation.

A representative payload looks like this:

```json
{
  "connection": "[tcp] 1.3.3.7:1338 -> 1.3.3.7:443",
  "source": "1.3.3.7:1338",
  "target": "1.3.3.7:443",
  "subnet": "1.3.3.0/24",
  "asn": "AS1337",
  "domain": "example.com",
  "datetime": "2026-10-19 11:59:59"
}
```

This payload can be ingested through a standard Elasticsearch ingest pipeline or
a Logstash filter. The mapping above is guidance for integrators. A dedicated
Elastic Common Schema (ECS) mapping will only be documented once it is actually
published and tested.

## Usage: Tholian Core

The module is consumed through Go's module system. In the Tholian Network, all
repositories are checked out next to the consuming project and referenced with
a local replacement:

```text
/path/to/Software/tholian-network/
|- core/         # this module
|- insights/     # Tholian Insights  (OSINT dataset generated by Tholian Recon)
|
|- assembly/     # Tholian Assembly  (Operator interface)
|- authority/    # Tholian Authority
|- oversight/    # Tholian Oversight
|
|- convoy/       # Tholian Convoy    (Offensive Pentest Agent)
|- evac/         # Tholian Evac      (Emergency Backup Agent)
|- firewall/     # Tholian Firewall  (Defensive Firewall Agent)
|- guard/        # Tholian Guard     (EDR Agent)
|- recon/        # Tholian Recon     (Threat Intelligence Agent)
|- scout/        # Tholian Scout     (Inventory Agent)
\- warps/        # Tholian Warps     (Network Router Agent)
```

```go
module example

go 1.27.1

require tholian-core v0.0.0

replace tholian-core => ../../core/source
```

### Naming Conventions

The packages follow a consistent naming scheme:

- `Parse*` returns a pointer or fails with `nil`, delegating to `Parse()`.
- `To*` returns a zero value when parsing fails.
- `Is*` validates strings or reports type predicates.
- `New*` constructs an empty or initialized value.
- `Set*` mutates a value only when the input is valid.
- `Equal` compares values of the same type.
- `Matches` compares matcher or message semantics.
- `IsValid`, `IsAny`, `Has*` describe state.
- `Any` is the shared value for wildcard or unspecified selectors.

## Testing

The module has no build tags. Tests can be run directly from the source folder:

```bash
cd source;
go test ./...;

# optional static checks
go vet ./...;
```

## License

The Tholian Core Schema definitions are published under the
[MIT license](https://opensource.org/licenses/MIT).

The Go source code of our software is licensed separately under the Tholian
Network repository terms. Contact [Tholian](https://tholian.network) for
commercial licensing.

`(c) 2022-2026 by Tholian(r) Network`
