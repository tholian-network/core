# Tholian® Core Type System

Tholian® Core is the shared type system for the Tholian Network. It defines the
canonical representation of all data types and schemas that are used across
Tholian's cyber defense projects.

The core module is not a general-purpose runtime or transport protocol. It is a
small Go module that gives every consumer the same vocabulary and the same JSON
wire format, and it includes the filesystem database generator used by the
separate Tholian Insights repository. That makes it possible to correlate data
from all Tholian Agent implementations and our provided threat intelligence
tooling without repeatedly reinterpreting the same values.

The JSON representations are intentionally documented and stable, so that they
can be consumed by external tools and proprietary cyber security stacks such as
Elastic, Kibana, SIEM, SOAR, or custom dashboards.

## Design Goals

- **Canonical**: All supported values have exactly one normalized representation.
- **Validation**: All parsers check ranges, prefixes, labels, and protocol names.
- **JSON-first**: Canonical value types marshal to stable JSON strings, and
  structured types use documented JSON fields.
- **Portability**: The implementation is standard-library-first, with only
  `x/term` required for terminal handling.
- **Composability**: Sockets combine with subnets, connections combine with matchers,
  and networks group subnets without duplicating them.
- **Interoperability**: Canonical strings map cleanly to native field types in
  external search and analytics platforms.

## Repository Layout

| Path                                           | Purpose                              |
|:-----------------------------------------------|:-------------------------------------|
| [source/caches](./source/caches)               | Optimized Caches for data structures |
| [source/distributions](./source/distributions) | Distribution detection and parsing   |
| [source/insights](./source/insights)           | Insights Cache and Database          |
| [source/matchers](./source/matchers)           | Matcher rules for types and structs  |
| [source/schemas](./source/schemas)             | External schema mappings (OVAL)      |
| [source/structs](./source/structs)             | Nested data structures               |
| [source/types](./source/types)                 | Value and composite types            |
| [source/utils](./source/utils)                 | Utils and helper methods             |
| [docs/COMPONENTS.md](./docs/COMPONENTS.md)     | Complete component catalog           |

The Go module is named `tholian-core` and targets Go `1.27.1`.

## Canonical Wire Format

All core types implement `encoding/json` support. Simple values are encoded as
JSON strings, while structured values such as `Network` and `Internet` use their
declared JSON fields.

See [docs/COMPONENTS.md](./docs/COMPONENTS.md) for the complete catalog of all
types, structs, matchers, caches, insights helpers, utilities, constants, and
their canonical forms.

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

## Usage: Tholian Insights

`source/insights` contains the filesystem layout and generators used by the
separate Tholian Insights repository. It is split into two layers.

### Insights Cache

`insights.NewCache(folder)` is a small filesystem helper with
`Exists`, `Read`, `Remove`, `Scan`, and `Write` methods. It is independent of the
database layout and only operates on the folder it is given.

### Insights Database

`insights.NewDatabase(folder)` creates and manages the canonical folder layout:

```text
countries/               devices/                distributions/
networks/                providers/              proxies/
timeslots/               weaknesses/             credentials/embedded/
incidents/embedded/      mitigations/embedded/   products/embedded/
responses/embedded/      vulnerabilities/embedded/
```

Cache-backed data types (`Countries`, `Devices`, `Distributions`, `Networks`,
`Providers`, `Proxies`, `Timeslots`, `Weaknesses`) are held in memory and
written by `Flush()`. Flushing writes both `Type.json` and `Type.json.gz` and
generates the corresponding Go cache package.

File-backed data types are persisted one record per file:

| Type | File key | Build tags |
|:-----|:---------|:-----------|
| `Credential` | `Name` | `oversight` |
| `Incident` | `Hash` | `guard`, `oversight` |
| `Mitigation` | `Hash` | `guard` |
| `Product` | `Name` | `scout`, `guard`, `oversight` |
| `Response` | `Hash` | `guard`, `oversight` |
| `Vulnerability` | `Name` | `scout`, `guard`, `oversight` |

Incidents and Responses are distribution-aware for the `guard` tag.
Credentials are oversight-only. Products intentionally emit one all-record cache
per base tag and no distribution-specific copies, so all products remain
available to match distribution-specific vulnerabilities. Vulnerabilities
additionally generate `scout_<distribution>` and `guard_<distribution>` cache
files.

Each file-backed type provides `BorrowType`, `ReadType`, `RemoveType`,
`WriteType`, `WriteTypes`, and `CleanupTypes` methods. `ReadType` falls back to
the compressed `.json.gz` file when the plain `.json` file is absent. Names used
as file keys are checked with `utils/path.IsSafeBasename` before access.

Writes are deterministic: nested collections are sorted by their `Hash()` before
serialization, so regenerating the same data produces line-by-line identical
JSON. Hash-keyed write methods use the sorted hash for both locking and the
filename; name-keyed types use their `Name`.

`Flush()` only handles the cache-backed types. File-backed embeds are generated
explicitly through the corresponding `WriteTypes()` method:

```go
database := insights.NewDatabase("/srv/tholian/insights")

vulnerability := structs.NewVulnerability("CVE-2026-0001")
vulnerability.SetState("published")

if database.WriteVulnerability(&vulnerability) == false {
	log.Fatal("cannot write vulnerability")
}

if database.WriteVulnerabilities() == false {
	log.Fatal("cannot generate vulnerability embeds")
}

if database.Flush() == false {
	log.Fatal("cannot flush cache-backed insights")
}
```

`WriteVulnerability` also maintains the distribution-vendor index used by
`QueryVulnerabilitiesByDistributionVendor`. Generated cache packages import
`tholian-core/caches` and `tholian-core/structs`; when a type has no records, a
valid empty embed package is generated instead of an empty `//go:embed`
directive.

### Naming Conventions

The packages follow a consistent naming scheme:

- `Parse*` returns a pointer or fails with `nil`, delegating to `Parse()`.
- `To*` returns a zero value when parsing fails.
- `Is*` validates strings or reports type predicates.
- `New*` constructs an empty or initialized value.
- `Hash` returns a deterministic content-derived identifier; nameless persisted
  records use it as their filesystem key.
- `Set*` mutates a value only when the input is valid.
- `Equal` compares values of the same type.
- `Matches` compares matcher or message semantics.
- `IsValid`, `IsAny`, `Has*` describe state.
- `Any` is the shared value for wildcard or unspecified selectors.

## Testing

The module source has no build tags. Generated Insights cache files do use
build tags, but they are produced at runtime and are not part of the checked-in
module.

Run the full suite from the repository root:

```bash
make test
```

Or run the suites individually:

```bash
go -C source test ./...
go -C source test -race ./insights/...
go -C source vet ./...
go -C e2e-tests test ./...
```

## License

The Tholian Core Schema definitions are published under the
[MIT license](https://opensource.org/licenses/MIT).

The Go source code of our software is licensed separately under the Tholian
Network repository terms. Contact [Tholian](https://tholian.network) for
commercial licensing.

`(c) 2022-2026 by Tholian(r) Network`
