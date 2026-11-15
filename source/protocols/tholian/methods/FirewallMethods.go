package methods

const (

	// Install Capability
	FirewallLoad      Method = "tholian-firewall.Load"
	FirewallInit      Method = "tholian-firewall.Init"
	FirewallInstall   Method = "tholian-firewall.Install"
	FirewallSelfTest  Method = "tholian-firewall.SelfTest"
	FirewallUninstall Method = "tholian-firewall.Uninstall"

	//// Read Capability
	FirewallCheck  Method = "tholian-firewall.Check"
	FirewallSearch Method = "tholian-firewall.Search"
	FirewallStatus Method = "tholian-firewall.Status"

	//// Write Capability
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
