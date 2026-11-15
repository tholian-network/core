package capabilities

type Capability string

func (capability Capability) String() string {
	return string(capability)
}

const (

	FirewallInstall Capability = "tholian-firewall.Install"
	FirewallRead    Capability = "tholian-firewall.Read"
	FirewallWrite   Capability = "tholian-firewall.Write"

	GuardInstall Capability = "tholian-guard.Install"
	GuardRead    Capability = "tholian-guard.Read"
	GuardWrite   Capability = "tholian-guard.Write"

	ScoutInstall Capability = "tholian-scout.Install"
	ScoutRead    Capability = "tholian-scout.Read"
	ScoutWrite   Capability = "tholian-scout.Write"

	WatchtowerRead  Capability = "tholian-watchtower.Read"
	WatchtowerWrite Capability = "tholian-watchtower.Write"

)
