package methods

var all_methods = []Method{

	// from CoreMethods.go
	CoreChallenge,
	CoreHello,
	CoreWelcome,
	CorePing,
	CorePong,

	// from FirewallMethods.go
	FirewallLoad,
	FirewallInit,
	FirewallInstall,
	FirewallSelfTest,
	FirewallUninstall,
	FirewallCheck,
	FirewallSearch,
	FirewallStatus,
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

	// from GuardMethods.go
	GuardInit,
	GuardInstall,
	GuardUninstall,
	GuardOnline,
	GuardOffline,
	GuardCollectFiles,
	GuardCollectSystem,
	GuardCollectVulnerabilities,
	GuardSearchFiles,
	GuardSearchVulnerabilities,
	GuardExecuteProgram,
	GuardRemoveFiles,
	GuardUpdateFiles,
	GuardUpdatePackages,
	GuardUpdatePrograms,
	GuardUpdateServices,

	// from ScoutMethods.go
	ScoutInstall,
	ScoutUninstall,
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
	ScoutUpdate,
	ScoutUpdatePackage,
	ScoutUpdatePackages,

	// from WatchtowerMethods.go
	WatchtowerListFindings,
	WatchtowerListGrants,
	WatchtowerListReports,
	WatchtowerListSystems,
	WatchtowerCollectReport,
	WatchtowerAccessAccept,
	WatchtowerAccessChange,
	WatchtowerAccessDeny,
	WatchtowerAccessSystem,
	WatchtowerSubscribeSystem,
	WatchtowerUnsubscribeSystem,

}

func AllMethods() []Method {

	result := make([]Method, 0)

	result = append(result, all_methods...)

	return result

}
