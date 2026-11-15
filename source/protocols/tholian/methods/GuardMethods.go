package methods

const (

	// Install Capability
	GuardInit      Method = "tholian-guard.Init"
	GuardInstall   Method = "tholian-guard.Install"
	GuardUninstall Method = "tholian-guard.Uninstall"

	// Internal States
	GuardOnline  Method = "tholian-guard.Online"
	GuardOffline Method = "tholian-guard.Offline"

	// Read Capability
	GuardCollectFiles           Method = "tholian-guard.CollectFiles"
	GuardCollectSystem          Method = "tholian-guard.CollectSystem"
	GuardCollectVulnerabilities Method = "tholian-guard.CollectVulnerabilities"
	GuardSearchFiles            Method = "tholian-guard.SearchFiles"
	GuardSearchVulnerabilities  Method = "tholian-guard.SearchVulnerabilities"

	// Write Capability
	GuardExecuteProgram Method = "tholian-guard.ExecuteProgram"
	GuardRemoveFiles    Method = "tholian-guard.RemoveFiles"
	GuardUpdateFiles    Method = "tholian-guard.UpdateFiles"
	GuardUpdatePackages Method = "tholian-guard.UpdatePackages"
	GuardUpdatePrograms Method = "tholian-guard.UpdatePrograms"
	GuardUpdateServices Method = "tholian-guard.UpdateServices"

)
