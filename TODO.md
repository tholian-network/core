# TODO

## Collection helpers

- [ ] Add `Package.AddUnresolved`, `Package.RemoveUnresolved`, `Package.SetUnresolved`, and `Package.HasUnresolved`.
- [ ] Add `Package.HasConflict`, `Package.HasProvide`, `Package.HasReplace`, and `Package.HasMaintainer`.
- [ ] Add `Program.SetPID`.
- [ ] Add `Program.HasConnection`, `Program.HasFilesystem`, and `Program.HasPackage`.
- [ ] Add `Program.RemovePackages`.
- [ ] Add `Country.AddTimezone`, `Country.RemoveTimezone`, and `Country.HasSubnet`.
- [ ] Add `Credential.HasKey` and `Credential.HasPassword`.
- [ ] Add `Device.HasSystem` and `Device.HasSubsystem`.
- [ ] Add `Distribution.AddKernelModule`, `Distribution.RemoveKernelModule`, and `Distribution.HasKernelModule`.
- [ ] Add `Network.String` and `Network.HasSubnet`.
- [ ] Add `Weakness.HasImpact` and `Weakness.HasReference`.

## Matching helpers

- [ ] Add `Incident.MatchesDevice`, `Incident.MatchesProduct`, and `Incident.MatchesWeakness`.
- [ ] Add `Response.MatchesDevice`, `Response.MatchesProduct`, and `Response.MatchesWeakness`.
- [ ] Add `Vulnerability.MatchesAntique` and `Vulnerability.MatchesUpdate`.
- [ ] Consider `Vulnerability.MatchesNetwork` and `Vulnerability.MatchesUser`.

## Internet and time helpers

- [ ] Add `Internet.RemoveNetwork` and `Internet.RemoveSubnet`.
- [ ] Add `Internet.IsValid`.
- [ ] Add `Timezone.IsValid` and `Timezone.String`.
- [ ] Add `DeviceSystem.NewDeviceSystem` and `DeviceSystem.IsValid`.

## JSON

- [ ] Add `MarshalJSON` and `UnmarshalJSON` methods for the remaining schema structs to match the README contract.
