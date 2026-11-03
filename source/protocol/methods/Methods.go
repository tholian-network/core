package methods

const (
	RPCChallenge Method = "tholian-rpc.Challenge"
	RPCHello     Method = "tholian-rpc.Hello"
	RPCWelcome   Method = "tholian-rpc.Welcome"
	RPCPing      Method = "tholian-rpc.Ping"
	RPCPong      Method = "tholian-rpc.Pong"
	RPCCancel    Method = "tholian-rpc.Cancel"
)

const (

	// Install Capability
	FirewallLoad      Method = "tholian-firewall.Load"
	FirewallInit      Method = "tholian-firewall.Init"
	FirewallInstall   Method = "tholian-firewall.Install"
	FirewallSelfTest  Method = "tholian-firewall.SelfTest"
	FirewallUninstall Method = "tholian-firewall.Uninstall"

	// Read Capability
	FirewallCheck  Method = "tholian-firewall.Check"
	FirewallSearch Method = "tholian-firewall.Search"
	FirewallStatus Method = "tholian-firewall.Status"

	// Write Capability
	FirewallForbid            Method = "tholian-firewall.Forbid"
	FirewallForbidConnection  Method = "tholian-firewall.ForbidConnection"
	FirewallForbidConnections Method = "tholian-firewall.ForbidConnections"
	FirewallForbidNetwork     Method = "tholian-firewall.ForbidNetwork"
	FirewallForbidNetworks    Method = "tholian-firewall.ForbidNetworks"
	FirewallPermit            Method = "tholian-firewall.Permit"
	FirewallPermitConnection  Method = "tholian-firewall.PermitConnection"
	FirewallPermitConnections Method = "tholian-firewall.PermitConnections"
	FirewallPermitNetwork     Method = "tholian-firewall.PermitNetwork"
	FirewallPermitNetworks    Method = "tholian-firewall.PermitNetworks"

)

const (

	// Install Capability
	ScoutInit      Method = "tholian-scout.Init"
	ScoutInstall   Method = "tholian-scout.Install"
	ScoutUninstall Method = "tholian-scout.Uninstall"

	// Read Capability
	ScoutCollect               Method = "tholian-scout.Collect"
	ScoutCollectAntiques       Method = "tholian-scout.CollectAntiques"
	ScoutCollectDevices        Method = "tholian-scout.CollectDevices"
	ScoutCollectDrives         Method = "tholian-scout.CollectDrives"
	ScoutCollectIncidents      Method = "tholian-scout.CollectIncidents"
	ScoutCollectIncidentsSince Method = "tholian-scout.CollectIncidentsSince"
	ScoutCollectNetworks       Method = "tholian-scout.CollectNetworks"
	ScoutCollectPackages       Method = "tholian-scout.CollectPackages"
	ScoutCollectPrograms       Method = "tholian-scout.CollectPrograms"
	ScoutCollectServices       Method = "tholian-scout.CollectServices"
	ScoutCollectUpdates        Method = "tholian-scout.CollectUpdates"
	ScoutCollectUsers          Method = "tholian-scout.CollectUsers"
	ScoutSearchAntiques        Method = "tholian-scout.SearchAntiques"
	ScoutSearchDevices         Method = "tholian-scout.SearchDevices"
	ScoutSearchDrives          Method = "tholian-scout.SearchDrives"
	ScoutSearchIncidents       Method = "tholian-scout.SearchIncidents"
	ScoutSearchNetworks        Method = "tholian-scout.SearchNetworks"
	ScoutSearchPackages        Method = "tholian-scout.SearchPackages"
	ScoutSearchPrograms        Method = "tholian-scout.SearchPrograms"
	ScoutSearchServices        Method = "tholian-scout.SearchServices"
	ScoutSearchUpdates         Method = "tholian-scout.SearchUpdates"
	ScoutSearchUsers           Method = "tholian-scout.SearchUsers"

	// Write Capability
	ScoutUpdate         Method = "tholian-scout.Update"
	ScoutUpdatePackage  Method = "tholian-scout.UpdatePackage"
	ScoutUpdatePackages Method = "tholian-scout.UpdatePackages"

)

const (

	// Install Capability
	GuardInit      Method = "tholian-guard.Init"
	GuardInstall   Method = "tholian-guard.Install"
	GuardUninstall Method = "tholian-guard.Uninstall"
	GuardOnline    Method = "tholian-guard.Online"
	GuardOffline   Method = "tholian-guard.Offline"

	// Read Capability
	GuardCollectSystem          Method = "tholian-guard.CollectSystem"
	GuardCollectVulnerabilities Method = "tholian-guard.CollectVulnerabilities"
	GuardReadEvidence           Method = "tholian-guard.ReadEvidence"
	GuardReadFile               Method = "tholian-guard.ReadFile"

	// Write Capability
	GuardExecuteShell   Method = "tholian-guard.ExecuteShell"
	GuardUpdatePackages Method = "tholian-guard.UpdatePackages"
	GuardUpdatePrograms Method = "tholian-guard.UpdatePrograms"
	GuardUpdateServices Method = "tholian-guard.UpdateServices"
	GuardWriteEvidence  Method = "tholian-guard.WriteEvidence"
	GuardWriteFile      Method = "tholian-guard.WriteFile"

)

const (

	WatchtowerListAudits          Method = "tholian-watchtower.ListAudits"
	WatchtowerListFindings        Method = "tholian-watchtower.ListFindings"
	WatchtowerListSystems         Method = "tholian-watchtower.ListSystems"

	WatchtowerCollectSystem          Method = "tholian-watchtower.CollectSystem"
	WatchtowerCollectIncidents       Method = "tholian-watchtower.CollectIncidents"
	WatchtowerCollectVulnerabilities Method = "tholian-watchtower.CollectVulnerabilities"

	WatchtowerConnectSystem    Method = "tholian-watchtower.ConnectSystem"
	WatchtowerDisconnectSystem Method = "tholian-watchtower.DisconnectSystem"

	WatchtowerListGrants   Method = "tholian-watchtower.ListGrants"
	WatchtowerRequestGrant Method = "tholian-watchtower.RequestGrant"
	WatchtowerRevokeGrant  Method = "tholian-watchtower.RevokeGrant"
	WatchtowerChangeGrant  Method = "tholian-watchtower.ChangeGrant"

	WatchtowerSubscribeEvents   Method = "tholian-watchtower.SubscribeEvents"
	WatchtowerUnsubscribeEvents Method = "tholian-watchtower.UnsubscribeEvents"

)

const (
	OperatorConnectWatchtower    Method = "tholian-operator.ConnectWatchtower"
	OperatorDisconnectWatchtower Method = "tholian-operator.DisconnectWatchtower"
)

