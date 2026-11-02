package protocol

const (
	MethodChallenge = "rpc.challenge"
	MethodHello     = "rpc.hello"
	MethodWelcome   = "rpc.welcome"
	MethodPing      = "rpc.ping"
	MethodPong      = "rpc.pong"
	MethodCancel    = "rpc.cancel"
)

const (
	MethodSystemReport          = "system.report"
	MethodSystemCollect         = "system.collect"
	MethodIncidentsCollectSince = "incidents.collectSince"
	MethodPackagesSearch        = "packages.search"
	MethodPackagesUpdate        = "packages.update"
	MethodProgramsSearch        = "programs.search"
	MethodServicesSearch        = "services.search"
	MethodUsersSearch           = "users.search"
	MethodUpdatesSearch         = "updates.search"
	MethodNetworksSearch        = "networks.search"
	MethodAntiquesSearch        = "antiques.search"
	MethodDevicesSearch         = "devices.search"
	MethodDrivesSearch          = "drives.search"
	MethodFirewallStatus        = "firewall.status"
	MethodFirewallCheck         = "firewall.check"
	MethodFirewallPermit        = "firewall.permit"
	MethodFirewallForbid        = "firewall.forbid"
)

const (
	MethodSystemsList       = "systems.list"
	MethodSystemsReport     = "systems.report"
	MethodSystemsIncidents  = "systems.incidents"
	MethodSessionCreate     = "session.create"
	MethodSessionClose      = "session.close"
	MethodGrantsRequest     = "grants.request"
	MethodGrantsList        = "grants.list"
	MethodGrantsRevoke      = "grants.revoke"
	MethodEventsSubscribe   = "events.subscribe"
	MethodEventsUnsubscribe = "events.unsubscribe"
	MethodAuditList         = "audit.list"
	MethodCall              = "call"
)

const (
	MethodGuardOnline   = "guard.online"
	MethodGuardOffline  = "guard.offline"
	MethodIncidentNew   = "incident.new"
	MethodReportUpdated = "report.updated"
	MethodGrantChanged  = "grant.changed"
)

const (
	MethodVulnerabilitiesScan = "vulnerabilities.scan"
	MethodEvidenceRead        = "evidence.read"
	MethodFilesRead           = "files.read"
	MethodShellExecute        = "shell.execute"
)
