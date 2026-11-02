package permissions

type Capability string

func (capability Capability) String() string {
	return string(capability)
}

const (
	CapabilitySystemRead     Capability = "tholian-scout.system.read"
	CapabilityPackagesRead   Capability = "tholian-scout.packages.read"
	CapabilityPackagesUpdate Capability = "tholian-scout.packages.update"
	CapabilityProgramsRead   Capability = "tholian-scout.programs.read"
	CapabilityServicesRead   Capability = "tholian-scout.services.read"
	CapabilityUsersRead      Capability = "tholian-scout.users.read"
	CapabilityUpdatesRead    Capability = "tholian-scout.updates.read"
	CapabilityIncidentsRead  Capability = "tholian-scout.incidents.read"
	CapabilityFirewallRead   Capability = "tholian-firewall.read"
	CapabilityFirewallWrite  Capability = "tholian-firewall.write"
)

const (
	CapabilityProgramsUpdate Capability = "tholian-scout.programs.update"
	CapabilityServicesUpdate Capability = "tholian-scout.services.update"
	CapabilityUsersUpdate    Capability = "tholian-scout.users.update"
	CapabilityEvidenceRead   Capability = "tholian-guard.evidence.read"
	CapabilityEvidenceWrite  Capability = "tholian-guard.evidence.write"
	CapabilityShellExecute   Capability = "tholian-guard.shell.execute"
)
