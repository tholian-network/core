package permissions

import "tholian-core/protocols/tholian/capabilities"
import "tholian-core/protocols/tholian/methods"

type GrantRequest struct {
	System       string                    `json:"system"`
	Purpose      Purpose                   `json:"purpose"`
	Capabilities []capabilities.Capability `json:"capabilities"`
	Scope        Scope                     `json:"scope"`
	ExpiresIn    int                       `json:"expires-in"`
}

func (request *GrantRequest) IsValid() bool {

	if request == nil {
		return false
	}

	if request.System != "" && request.Purpose.IsValid() && len(request.Capabilities) > 0 {
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
	case capabilities.WatchtowerRead,
		capabilities.WatchtowerWrite:
		return true
	}

	return false

}

func MethodCapability(method methods.Method) (capabilities.Capability, bool) {

	switch method {

	case methods.FirewallLoad,
		methods.FirewallInit,
		methods.FirewallInstall,
		methods.FirewallSelfTest,
		methods.FirewallUninstall:

		return capabilities.FirewallInstall, true

	case methods.FirewallCheck,
		methods.FirewallSearch,
		methods.FirewallStatus:

		return capabilities.FirewallRead, true

	case methods.FirewallForbid,
		methods.FirewallForbidConnection,
		methods.FirewallForbidConnections,
		methods.FirewallForbidNetwork,
		methods.FirewallForbidNetworks,
		methods.FirewallPermit,
		methods.FirewallPermitConnection,
		methods.FirewallPermitConnections,
		methods.FirewallPermitNetwork,
		methods.FirewallPermitNetworks:

		return capabilities.FirewallWrite, true

	case methods.GuardInit,
		methods.GuardInstall,
		methods.GuardUninstall,
		methods.GuardOnline,
		methods.GuardOffline:

		return capabilities.GuardInstall, true

	case methods.GuardCollectFiles,
		methods.GuardCollectSystem,
		methods.GuardCollectVulnerabilities,
		methods.GuardSearchFiles,
		methods.GuardSearchVulnerabilities:

		return capabilities.GuardRead, true

	case methods.GuardExecuteProgram,
		methods.GuardProtectService,
		methods.GuardProtectServices,
		methods.GuardRemoveFiles,
		methods.GuardUpdateFiles,
		methods.GuardUpdatePrograms,
		methods.GuardUpdateServices:

		return capabilities.GuardWrite, true

	case methods.ScoutInstall,
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

	}

	return capabilities.Capability(""), false

}
