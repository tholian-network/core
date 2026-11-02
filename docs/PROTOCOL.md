# Tholian RPC Protocol

The `tholian-core/protocol` package is the shared wire contract for
`tholian-operator`, `tholian-watchtower` and `tholian-guard`.

- Transport: WebSocket version 13 via `github.com/cookiengineer/gows`.
- Security: TLS 1.2+ with Ed25519 challenge-response authentication.
- Encoding: JSON text frames with a versioned envelope.
- Compression: `permessage-deflate` for large system reports.
- Version: `1`.
- Subprotocol name: `tholian-rpc.v1`. `gows.Client` v0.0.2 does not send
  `Sec-WebSocket-Protocol`, so the server does not require it and the value is
  informational for now.

## Envelope

```json
{
    "version": "1",
    "id": "request-id",
    "kind": "request|response|event|error",
    "method": "packages.search",
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
2. Watchtower sends `rpc.challenge` with a random nonce.
3. The guard or operator sends `rpc.hello` with its role, ID, Ed25519 public
   key, a signature over `tholian-hello-v1:<nonce>:<id>`, an optional
   one-time enrollment token and the protocol version.
4. Watchtower verifies the token or the registered public key and replies with
   `rpc.welcome`. The challenge nonce is carried in request `params`; the
   welcome payload is carried in response `result` and contains the server
   public key and a signature over `tholian-welcome-v1:<nonce>:<id>`.
5. `gows.Client.Connect()` does not start the read loop; clients install
   `OnMessage`/`OnClose` and then call `go socket.Init()`.

## Methods

Handshake:

- `rpc.challenge`, `rpc.hello`, `rpc.welcome`, `rpc.ping`, `rpc.pong`,
  `rpc.cancel`

Guard:

- `system.report`, `system.collect`
- `incidents.collectSince`
- `packages.search`, `packages.update`
- `programs.search`, `services.search`, `users.search`, `updates.search`
- `networks.search`, `antiques.search`, `devices.search`, `drives.search`
- `firewall.status`, `firewall.check`, `firewall.permit`, `firewall.forbid`

Watchtower:

- `systems.list`, `systems.report`, `systems.incidents`
- `session.create`, `session.close`
- `grants.request`, `grants.list`, `grants.revoke`
- `events.subscribe`, `events.unsubscribe`
- `audit.list`
- `call`

Reserved: `vulnerabilities.*`, `evidence.*`, `files.*`, `shell.*`.

## Permissions

- Capabilities are strings such as `packages.read`, `packages.update`,
  `firewall.read` and `firewall.write`.
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
`protocol.FirewallStatusResult`.

## Size Limits

`protocol.DefaultMaxMessageBytes` is 32 MiB. Components enforce a configurable
limit and negotiate `permessage-deflate` for system reports.
