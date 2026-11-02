package permissions

import "strings"

// Method is a project-qualified RPC method name.
//
// The wire format is "<go-module>.<ExportedActionName>", for example
// "tholian-firewall.Check" or "tholian-scout.CollectPackages". Action method
// values map 1:1 to exported functions in the owning project's actions package.
type Method string

const (
	MethodRPCChallenge Method = "tholian-rpc.Challenge"
	MethodRPCHello     Method = "tholian-rpc.Hello"
	MethodRPCWelcome   Method = "tholian-rpc.Welcome"
	MethodRPCPing      Method = "tholian-rpc.Ping"
	MethodRPCPong      Method = "tholian-rpc.Pong"
	MethodRPCCancel    Method = "tholian-rpc.Cancel"
)

const (
	MethodFirewallInit              Method = "tholian-firewall.Init"
	MethodFirewallStatus            Method = "tholian-firewall.Status"
	MethodFirewallCheck             Method = "tholian-firewall.Check"
	MethodFirewallSearch            Method = "tholian-firewall.Search"
	MethodFirewallSelfTest          Method = "tholian-firewall.SelfTest"
	MethodFirewallLoad              Method = "tholian-firewall.Load"
	MethodFirewallPermit            Method = "tholian-firewall.Permit"
	MethodFirewallPermitConnection  Method = "tholian-firewall.PermitConnection"
	MethodFirewallPermitConnections Method = "tholian-firewall.PermitConnections"
	MethodFirewallPermitNetwork     Method = "tholian-firewall.PermitNetwork"
	MethodFirewallPermitNetworks    Method = "tholian-firewall.PermitNetworks"
	MethodFirewallForbid            Method = "tholian-firewall.Forbid"
	MethodFirewallForbidConnection  Method = "tholian-firewall.ForbidConnection"
	MethodFirewallForbidConnections Method = "tholian-firewall.ForbidConnections"
	MethodFirewallForbidNetwork     Method = "tholian-firewall.ForbidNetwork"
	MethodFirewallForbidNetworks    Method = "tholian-firewall.ForbidNetworks"
)

const (
	MethodScoutInit                  Method = "tholian-scout.Init"
	MethodScoutCollect               Method = "tholian-scout.Collect"
	MethodScoutCollectAntiques       Method = "tholian-scout.CollectAntiques"
	MethodScoutCollectDevices        Method = "tholian-scout.CollectDevices"
	MethodScoutCollectDrives         Method = "tholian-scout.CollectDrives"
	MethodScoutCollectIncidents      Method = "tholian-scout.CollectIncidents"
	MethodScoutCollectIncidentsSince Method = "tholian-scout.CollectIncidentsSince"
	MethodScoutCollectNetworks       Method = "tholian-scout.CollectNetworks"
	MethodScoutCollectPackages       Method = "tholian-scout.CollectPackages"
	MethodScoutCollectPrograms       Method = "tholian-scout.CollectPrograms"
	MethodScoutCollectServices       Method = "tholian-scout.CollectServices"
	MethodScoutCollectUpdates        Method = "tholian-scout.CollectUpdates"
	MethodScoutCollectUsers          Method = "tholian-scout.CollectUsers"
	MethodScoutInstall               Method = "tholian-scout.Install"
	MethodScoutUninstall             Method = "tholian-scout.Uninstall"
	MethodScoutUpdate                Method = "tholian-scout.Update"
	MethodScoutUpdatePackage         Method = "tholian-scout.UpdatePackage"
	MethodScoutUpdatePackages        Method = "tholian-scout.UpdatePackages"
	MethodScoutSearchAntiques        Method = "tholian-scout.SearchAntiques"
	MethodScoutSearchDevices         Method = "tholian-scout.SearchDevices"
	MethodScoutSearchDrives          Method = "tholian-scout.SearchDrives"
	MethodScoutSearchIncidents       Method = "tholian-scout.SearchIncidents"
	MethodScoutSearchNetworks        Method = "tholian-scout.SearchNetworks"
	MethodScoutSearchPackages        Method = "tholian-scout.SearchPackages"
	MethodScoutSearchPrograms        Method = "tholian-scout.SearchPrograms"
	MethodScoutSearchServices        Method = "tholian-scout.SearchServices"
	MethodScoutSearchUpdates         Method = "tholian-scout.SearchUpdates"
	MethodScoutSearchUsers           Method = "tholian-scout.SearchUsers"
)

const (
	MethodGuardReport      Method = "tholian-guard.Report"
	MethodGuardOnline      Method = "tholian-guard.Online"
	MethodGuardOffline     Method = "tholian-guard.Offline"
	MethodGuardIncidentNew Method = "tholian-guard.IncidentNew"
)

const (
	MethodWatchtowerSystemsList       Method = "tholian-watchtower.SystemsList"
	MethodWatchtowerSystemsReport     Method = "tholian-watchtower.SystemsReport"
	MethodWatchtowerSystemsIncidents  Method = "tholian-watchtower.SystemsIncidents"
	MethodWatchtowerSessionCreate     Method = "tholian-watchtower.SessionCreate"
	MethodWatchtowerSessionClose      Method = "tholian-watchtower.SessionClose"
	MethodWatchtowerGrantsRequest     Method = "tholian-watchtower.GrantsRequest"
	MethodWatchtowerGrantsList        Method = "tholian-watchtower.GrantsList"
	MethodWatchtowerGrantsRevoke      Method = "tholian-watchtower.GrantsRevoke"
	MethodWatchtowerEventsSubscribe   Method = "tholian-watchtower.EventsSubscribe"
	MethodWatchtowerEventsUnsubscribe Method = "tholian-watchtower.EventsUnsubscribe"
	MethodWatchtowerAuditList         Method = "tholian-watchtower.AuditList"
	MethodWatchtowerCall              Method = "tholian-watchtower.Call"
	MethodWatchtowerReportUpdated     Method = "tholian-watchtower.ReportUpdated"
	MethodWatchtowerGrantChanged      Method = "tholian-watchtower.GrantChanged"
)

const (
	MethodOperatorConnectionClosed Method = "tholian-operator.ConnectionClosed"
)

const (
	MethodGuardCollectVulnerabilities Method = "tholian-guard.CollectVulnerabilities"
	MethodGuardReadEvidence           Method = "tholian-guard.ReadEvidence"
	MethodGuardReadFile               Method = "tholian-guard.ReadFile"
	MethodGuardWriteFile              Method = "tholian-guard.WriteFile"
	MethodGuardExecuteShell           Method = "tholian-guard.ExecuteShell"
)

var methods = []Method{
	MethodRPCChallenge,
	MethodRPCHello,
	MethodRPCWelcome,
	MethodRPCPing,
	MethodRPCPong,
	MethodRPCCancel,
	MethodFirewallInit,
	MethodFirewallStatus,
	MethodFirewallCheck,
	MethodFirewallSearch,
	MethodFirewallSelfTest,
	MethodFirewallLoad,
	MethodFirewallPermit,
	MethodFirewallPermitConnection,
	MethodFirewallPermitConnections,
	MethodFirewallPermitNetwork,
	MethodFirewallPermitNetworks,
	MethodFirewallForbid,
	MethodFirewallForbidConnection,
	MethodFirewallForbidConnections,
	MethodFirewallForbidNetwork,
	MethodFirewallForbidNetworks,
	MethodScoutInit,
	MethodScoutCollect,
	MethodScoutCollectAntiques,
	MethodScoutCollectDevices,
	MethodScoutCollectDrives,
	MethodScoutCollectIncidents,
	MethodScoutCollectIncidentsSince,
	MethodScoutCollectNetworks,
	MethodScoutCollectPackages,
	MethodScoutCollectPrograms,
	MethodScoutCollectServices,
	MethodScoutCollectUpdates,
	MethodScoutCollectUsers,
	MethodScoutInstall,
	MethodScoutUninstall,
	MethodScoutUpdate,
	MethodScoutUpdatePackage,
	MethodScoutUpdatePackages,
	MethodScoutSearchAntiques,
	MethodScoutSearchDevices,
	MethodScoutSearchDrives,
	MethodScoutSearchIncidents,
	MethodScoutSearchNetworks,
	MethodScoutSearchPackages,
	MethodScoutSearchPrograms,
	MethodScoutSearchServices,
	MethodScoutSearchUpdates,
	MethodScoutSearchUsers,
	MethodGuardReport,
	MethodGuardOnline,
	MethodGuardOffline,
	MethodGuardIncidentNew,
	MethodWatchtowerSystemsList,
	MethodWatchtowerSystemsReport,
	MethodWatchtowerSystemsIncidents,
	MethodWatchtowerSessionCreate,
	MethodWatchtowerSessionClose,
	MethodWatchtowerGrantsRequest,
	MethodWatchtowerGrantsList,
	MethodWatchtowerGrantsRevoke,
	MethodWatchtowerEventsSubscribe,
	MethodWatchtowerEventsUnsubscribe,
	MethodWatchtowerAuditList,
	MethodWatchtowerCall,
	MethodWatchtowerReportUpdated,
	MethodWatchtowerGrantChanged,
	MethodOperatorConnectionClosed,
	MethodGuardCollectVulnerabilities,
	MethodGuardReadEvidence,
	MethodGuardReadFile,
	MethodGuardWriteFile,
	MethodGuardExecuteShell,
}

func AllMethods() []Method {
	return append([]Method(nil), methods...)
}

func (method Method) String() string {
	return string(method)
}

func (method Method) IsValid() bool {

	for _, candidate := range methods {
		if candidate == method {
			return true
		}
	}

	return false

}

func (method Method) Project() string {

	project, _, ok := strings.Cut(string(method), ".")

	if ok == true {
		return project
	}

	return ""

}

func (method Method) Action() string {

	_, action, ok := strings.Cut(string(method), ".")

	if ok == true {
		return action
	}

	return ""

}

func ParseMethod(value string) (Method, bool) {

	method := Method(value)

	if method.IsValid() == true {
		return method, true
	}

	return Method(""), false

}
