package permissions

import "tholian-core/protocol/capabilities"
import "tholian-core/protocol/methods"

type GrantRequest struct {
	SystemID     string                    `json:"system"`
	Purpose      Purpose                   `json:"purpose"`
	Capabilities []capabilities.Capability `json:"capabilities"`
	Scope        Scope                     `json:"scope"`
	ExpiresIn    int                       `json:"expires-in"`
}

func (request *GrantRequest) IsValid() bool {

	if request == nil {
		return false
	}

	if request.SystemID != "" && request.Purpose.IsValid() && len(request.Capabilities) > 0 {
		return true
	}

	return false

}

func IsSupportedCapability(value capabilities.Capability) bool {

	switch value {
	case capabilities.FirewallInstall,
		capabilities.FirewallRead,
		capabilities.FirewallWrite,
		capabilities.GuardInstall,
		capabilities.GuardRead,
		capabilities.GuardWrite,
		capabilities.ScoutInstall,
		capabilities.ScoutRead,
		capabilities.ScoutWrite:
		return true
	}

	return false

}

func IsReservedCapability(value capabilities.Capability) bool {

	switch value {
	case capabilities.GuardWrite,
		capabilities.ScoutWrite:
		return true
	}

	return false

}

func MethodCapability(method methods.Method) (capabilities.Capability, bool) {

	switch method {

	case methods.GuardInit,
		methods.GuardInstall,
		methods.GuardUninstall,
		methods.GuardOnline,
		methods.GuardOffline:

		return capabilities.GuardInstall, true

	case methods.GuardCollectSystem,
		methods.GuardCollectVulnerabilities,
		methods.GuardReadEvidence,
		methods.GuardReadFile:

		return capabilities.GuardRead, true

	case methods.GuardExecuteShell,
		methods.GuardWriteEvidence,
		methods.GuardWriteFile,
		methods.GuardUpdatePrograms,
		methods.GuardUpdateServices:

		return capabilities.GuardWrite, true

	case methods.ScoutInit,
		methods.ScoutInstall,
		methods.ScoutUninstall:

		return capabilities.ScoutInstall, true

	case methods.ScoutCollect,
		methods.ScoutCollectAntiques,
		methods.ScoutCollectDevices,
		methods.ScoutCollectDrives,
		methods.ScoutCollectIncidents,
		methods.ScoutCollectIncidentsSince,
		methods.ScoutCollectNetworks,
		methods.ScoutCollectPackages,
		methods.ScoutCollectPrograms,
		methods.ScoutCollectServices,
		methods.ScoutCollectUpdates,
		methods.ScoutCollectUsers,
		methods.ScoutSearchAntiques,
		methods.ScoutSearchDevices,
		methods.ScoutSearchDrives,
		methods.ScoutSearchIncidents,
		methods.ScoutSearchNetworks,
		methods.ScoutSearchPackages,
		methods.ScoutSearchPrograms,
		methods.ScoutSearchServices,
		methods.ScoutSearchUpdates,
		methods.ScoutSearchUsers:

		return capabilities.ScoutRead, true

	case methods.ScoutUpdate,
		methods.ScoutUpdatePackage,
		methods.ScoutUpdatePackages:

		return capabilities.ScoutWrite, true

	case methods.FirewallInit,
		methods.FirewallInstall,
		methods.FirewallLoad:

		return capabilities.FirewallInstall, true

	case methods.FirewallStatus,
		methods.FirewallCheck,
		methods.FirewallSearch,
		methods.FirewallSelfTest:

		return capabilities.FirewallRead, true

	case methods.FirewallPermit,
		methods.FirewallPermitConnection,
		methods.FirewallPermitConnections,
		methods.FirewallPermitNetwork,
		methods.FirewallPermitNetworks,
		methods.FirewallForbid,
		methods.FirewallForbidConnection,
		methods.FirewallForbidConnections,
		methods.FirewallForbidNetwork,
		methods.FirewallForbidNetworks:

		return capabilities.FirewallWrite, true

	}

	return capabilities.Capability(""), false

}
