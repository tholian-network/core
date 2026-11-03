package methods

var all_methods = []Method{

	// Internal RPC Methods
	RPCChallenge,
	RPCHello,
	RPCWelcome,
	RPCPing,
	RPCPong,
	RPCCancel,

	// Firewall Install Capability
	FirewallLoad,
	FirewallInit,
	FirewallInstall,
	FirewallSelfTest,
	FirewallUninstall,

	// Firewall Read Capability
	FirewallCheck,
	FirewallSearch,
	FirewallStatus,

	// Firewall Write Capability
	FirewallForbid,
	FirewallForbidConnection,
	FirewallForbidConnections,
	FirewallForbidNetwork,
	FirewallForbidNetworks,
	FirewallPermit,
	FirewallPermitConnection,
	FirewallPermitConnections,
	FirewallPermitNetwork,
	FirewallPermitNetworks,

	// Scout Install Capability
	ScoutInstall,
	ScoutUninstall,

	// Scout Read Capability
	ScoutCollect,
	ScoutCollectAntiques,
	ScoutCollectDevices,
	ScoutCollectDrives,
	ScoutCollectIncidents,
	ScoutCollectIncidentsSince,
	ScoutCollectNetworks,
	ScoutCollectPackages,
	ScoutCollectPrograms,
	ScoutCollectServices,
	ScoutCollectUpdates,
	ScoutCollectUsers,
	ScoutSearchAntiques,
	ScoutSearchDevices,
	ScoutSearchDrives,
	ScoutSearchIncidents,
	ScoutSearchNetworks,
	ScoutSearchPackages,
	ScoutSearchPrograms,
	ScoutSearchServices,
	ScoutSearchUpdates,
	ScoutSearchUsers,

	// Scout Write Capability
	ScoutUpdate,
	ScoutUpdatePackage,
	ScoutUpdatePackages,

	// Guard Install Capability
	GuardInit,
	GuardInstall,
	GuardUninstall,
	GuardOnline,
	GuardOffline,

	// Guard Read Capability
	GuardCollectSystem,
	GuardCollectVulnerabilities,
	GuardReadEvidence,
	GuardReadFile,
	GuardSearchVulnerabilities,

	// Guard Write Capability
	GuardExecuteShell,
	GuardUpdatePackages,
	GuardUpdatePrograms,
	GuardUpdateServices,
	GuardWriteEvidence,
	GuardWriteFile,

	// Watchtower Methods
	WatchtowerListAudits,
	WatchtowerListFindings,
	WatchtowerListSystems,
	WatchtowerCollectSystem,
	WatchtowerCollectIncidents,
	WatchtowerCollectVulnerabilities,
	WatchtowerConnectSystem,
	WatchtowerDisconnectSystem,
	WatchtowerListGrants,
	WatchtowerRequestGrant,
	WatchtowerRevokeGrant,
	WatchtowerChangeGrant,
	WatchtowerSubscribeEvents,
	WatchtowerUnsubscribeEvents,

	// Operator Methods
	OperatorConnectWatchtower,
	OperatorDisconnectWatchtower,

}

func AllMethods() []Method {

	result := make([]Method, 0)

	result = append(result, all_methods...)

	return result

}

