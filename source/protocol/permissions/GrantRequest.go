package permissions

type GrantRequest struct {
	SystemID     string       `json:"system"`
	Purpose      Purpose      `json:"purpose"`
	Capabilities []Capability `json:"capabilities"`
	Scope        Scope        `json:"scope"`
	ExpiresIn    int          `json:"expires-in"`
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

func IsSupportedCapability(value Capability) bool {

	switch value {
	case CapabilitySystemRead,
		CapabilityPackagesRead,
		CapabilityPackagesUpdate,
		CapabilityProgramsRead,
		CapabilityServicesRead,
		CapabilityUsersRead,
		CapabilityUpdatesRead,
		CapabilityIncidentsRead,
		CapabilityFirewallRead,
		CapabilityFirewallWrite:
		return true
	}

	return false

}

func IsReservedCapability(value Capability) bool {

	switch value {
	case CapabilityProgramsUpdate,
		CapabilityServicesUpdate,
		CapabilityUsersUpdate,
		CapabilityEvidenceRead,
		CapabilityEvidenceWrite,
		CapabilityShellExecute:
		return true
	}

	return false

}

func MethodCapability(method Method) (Capability, bool) {

	switch method {
	case MethodScoutInit, MethodScoutCollect, MethodGuardReport:
		return CapabilitySystemRead, true
	case MethodScoutCollectAntiques, MethodScoutCollectPackages, MethodScoutSearchAntiques, MethodScoutSearchPackages:
		return CapabilityPackagesRead, true
	case MethodScoutCollectDevices, MethodScoutCollectDrives, MethodScoutCollectNetworks:
		return CapabilitySystemRead, true
	case MethodScoutCollectIncidents, MethodScoutCollectIncidentsSince, MethodScoutSearchIncidents:
		return CapabilityIncidentsRead, true
	case MethodScoutCollectPrograms, MethodScoutSearchPrograms:
		return CapabilityProgramsRead, true
	case MethodScoutCollectServices, MethodScoutSearchServices:
		return CapabilityServicesRead, true
	case MethodScoutCollectUpdates, MethodScoutSearchUpdates:
		return CapabilityUpdatesRead, true
	case MethodScoutCollectUsers, MethodScoutSearchUsers:
		return CapabilityUsersRead, true
	case MethodScoutInstall, MethodScoutUninstall, MethodScoutUpdate, MethodScoutUpdatePackage, MethodScoutUpdatePackages:
		return CapabilityPackagesUpdate, true
	case MethodScoutSearchDevices, MethodScoutSearchDrives, MethodScoutSearchNetworks:
		return CapabilitySystemRead, true
	case MethodFirewallStatus, MethodFirewallCheck, MethodFirewallSearch, MethodFirewallSelfTest:
		return CapabilityFirewallRead, true
	case MethodFirewallInit, MethodFirewallLoad,
		MethodFirewallPermit, MethodFirewallPermitConnection, MethodFirewallPermitConnections,
		MethodFirewallPermitNetwork, MethodFirewallPermitNetworks,
		MethodFirewallForbid, MethodFirewallForbidConnection, MethodFirewallForbidConnections,
		MethodFirewallForbidNetwork, MethodFirewallForbidNetworks:
		return CapabilityFirewallWrite, true
	}

	return Capability(""), false

}
