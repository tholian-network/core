package capabilities

type Capability string

func (capability Capability) String() string {
	return string(capability)
}

const (

	FirewallInstall Capability = "tholian-firewall.install"
	FirewallRead    Capability = "tholian-firewall.read"
	FirewallWrite   Capability = "tholian-firewall.write"

	GuardInstall    Capability = "tholian-guard.install"
	GuardRead       Capability = "tholian-guard.read"
	GuardWrite      Capability = "tholian-guard.write"

	ScoutInstall    Capability = "tholian-scout.install"
	ScoutRead       Capability = "tholian-scout.read"
	ScoutWrite      Capability = "tholian-scout.write"

)

