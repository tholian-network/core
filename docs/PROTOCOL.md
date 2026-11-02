# Tholian RPC Protocol

The `tholian-core/protocol` package is the shared wire contract for
`tholian-operator`, `tholian-watchtower` and `tholian-guard`.

- Transport: WebSocket version 13 via `github.com/cookiengineer/gows`.
- Security: TLS 1.2+ with Ed25519 challenge-response authentication.
- Encoding: JSON text frames with a versioned envelope.
- Compression: `permessage-deflate` for large system reports.
- Version: `1`.
- Subprotocol name: `tholian-rpc.v2`. `gows.Client` v0.0.2 does not send
  `Sec-WebSocket-Protocol`, so the server does not require it and the value is
  informational for now.

## Envelope

```json
{
    "version": "1",
    "id": "request-id",
    "kind": "request|response|event|error",
    "method": "tholian-scout.SearchPackages",
    "params": {},
    "result": {},
    "error": {
        "code": "forbidden",
        "message": "human-readable message"
    },
    "metadata": {
        "capability_token": "..."
    }
}
```

Requests use `params`; responses use `result`; server errors use `error` with a
stable code. Events use `params`.

## Handshake

1. The client connects with TLS.
2. Watchtower sends `tholian-rpc.Challenge` with a random nonce.
3. The guard or operator sends `tholian-rpc.Hello` with its role, ID, Ed25519
   public key, a signature over `tholian-hello-v1:<nonce>:<id>`, an optional
   one-time enrollment token and the protocol version.
4. Watchtower verifies the token or the registered public key and replies with
   `tholian-rpc.Welcome`. The challenge nonce is carried in request `params`;
   the welcome payload is carried in response `result` and contains the server
   public key and a signature over `tholian-welcome-v1:<nonce>:<id>`.
5. `gows.Client.Connect()` does not start the read loop; clients install
   `OnMessage`/`OnClose` and then call `go socket.Init()`.

## Methods

Method names are project-qualified: `<module>.<ExportedActionName>`. The
authoritative catalog lives in `source/protocol/permissions`.

Handshake:

- `tholian-rpc.Challenge`, `tholian-rpc.Hello`, `tholian-rpc.Welcome`,
  `tholian-rpc.Ping`, `tholian-rpc.Pong`, `tholian-rpc.Cancel`

Firewall actions:

- `tholian-firewall.Init`, `tholian-firewall.Status`, `tholian-firewall.Check`,
  `tholian-firewall.Search`, `tholian-firewall.SelfTest`, `tholian-firewall.Load`
- `tholian-firewall.Permit`, `tholian-firewall.PermitConnection`,
  `tholian-firewall.PermitConnections`, `tholian-firewall.PermitNetwork`,
  `tholian-firewall.PermitNetworks`
- `tholian-firewall.Forbid`, `tholian-firewall.ForbidConnection`,
  `tholian-firewall.ForbidConnections`, `tholian-firewall.ForbidNetwork`,
  `tholian-firewall.ForbidNetworks`

Scout actions:

- `tholian-scout.Init`, `tholian-scout.Collect`
- `tholian-scout.CollectAntiques`, `tholian-scout.CollectDevices`,
  `tholian-scout.CollectDrives`, `tholian-scout.CollectIncidents`,
  `tholian-scout.CollectIncidentsSince`, `tholian-scout.CollectNetworks`,
  `tholian-scout.CollectPackages`, `tholian-scout.CollectPrograms`,
  `tholian-scout.CollectServices`, `tholian-scout.CollectUpdates`,
  `tholian-scout.CollectUsers`
- `tholian-scout.Install`, `tholian-scout.Uninstall`, `tholian-scout.Update`,
  `tholian-scout.UpdatePackage`, `tholian-scout.UpdatePackages`
- `tholian-scout.SearchAntiques`, `tholian-scout.SearchDevices`,
  `tholian-scout.SearchDrives`, `tholian-scout.SearchIncidents`,
  `tholian-scout.SearchNetworks`, `tholian-scout.SearchPackages`,
  `tholian-scout.SearchPrograms`, `tholian-scout.SearchServices`,
  `tholian-scout.SearchUpdates`, `tholian-scout.SearchUsers`

Guard:

- `tholian-guard.Report`

Watchtower:

- `tholian-watchtower.SystemsList`, `tholian-watchtower.SystemsReport`,
  `tholian-watchtower.SystemsIncidents`
- `tholian-watchtower.SessionCreate`, `tholian-watchtower.SessionClose`
- `tholian-watchtower.GrantsRequest`, `tholian-watchtower.GrantsList`,
  `tholian-watchtower.GrantsRevoke`
- `tholian-watchtower.EventsSubscribe`, `tholian-watchtower.EventsUnsubscribe`
- `tholian-watchtower.AuditList`, `tholian-watchtower.Call`

Events:

- `tholian-guard.Online`, `tholian-guard.Offline`,
  `tholian-guard.IncidentNew`
- `tholian-watchtower.ReportUpdated`, `tholian-watchtower.GrantChanged`
- `tholian-operator.ConnectionClosed`

Reserved: `tholian-guard.VulnerabilitiesScan`,
`tholian-guard.EvidenceRead`, `tholian-guard.FilesRead`,
`tholian-guard.ShellExecute`.

## Permissions

- Capabilities are project-qualified strings. Scout grants use
  `tholian-scout.system.read`, `tholian-scout.packages.read`,
  `tholian-scout.packages.update`, `tholian-scout.programs.read`,
  `tholian-scout.services.read`, `tholian-scout.users.read`,
  `tholian-scout.updates.read` and `tholian-scout.incidents.read`. Firewall
  grants use `tholian-firewall.read` and `tholian-firewall.write`.
- Grants are per system, per purpose, time-bounded, use-limited and optionally
  scope-limited with core matchers. Watchtower rejects unsupported or reserved
  capabilities and caps grant lifetimes at 24 hours.
- Watchtower issues short-lived Ed25519-signed capability tokens for each
  forwarded call.
- Tokens bind the issuer, subject, system, method, params hash, purpose, grant,
  expiry and a replay nonce.
- Guards verify the signature, system binding, method, params hash, scope,
  expiry and replay cache before executing an action.

## Results

Guard actions return a uniform `protocol.ActionResult`:

```json
{
    "success": true,
    "data": {},
    "console": [],
    "error": ""
}
```

`data` contains the typed payload for the method, for example
`protocol.PackagesSearchResult`, `protocol.SystemReportResult` or
`protocol.FirewallStatusResult`. Audit entries record the project-qualified
method string, so a result can be traced to the exact downstream action.

## Size Limits

`protocol.DefaultMaxMessageBytes` is 32 MiB. Components enforce a configurable
limit and negotiate `permessage-deflate` for system reports.
