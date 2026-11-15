# TODO

## source/protocols/tholian: Ongoing Refactor

- [ ] Make sure that ./source/protocols/tholian/methods/AllMethods.go returns all methods correctly
- [ ] Move ./source/protocols/tholian/methods/Methods.go into GuardMethods.go and WatchtowerMethods.go
- [ ] Verify firewall and scout methods, parameters, and results
- [ ] Continue implementing methods for guard and watchtower

- [ ] Make sure parameters are correctly in-sync with methods
- [ ] Make sure results are correctly in-sync with methods

| Identifier                        | Method | Parameters | Results |
|-----------------------------------|--------|------------|---------|
| guard.CollectFiles                | [x]    | [x]        |         |
| guard.CollectSystem               | [x]    | [x]        |         |
| guard.CollectVulnerabilities      | [x]    | [x]        |         |
| guard.ExecuteProgram              | [x]    | [x]        |         |
| guard.Init                        | [x]    | [x]        |         |
| guard.Install                     | [x]    | [x]        |         |
| guard.Offline                     | [x]    | [x]        |         |
| guard.Online                      | [x]    | [x]        |         |
| guard.RemoveFiles                 | [x]    | [x]        |         |
| guard.SearchFiles                 | [x]    | [x]        |         |
| guard.SearchVulnerabilities       | [x]    | [x]        |         |
| guard.Uninstall                   | [x]    | [x]        |         |
| guard.UpdateFiles                 | [x]    | [x]        |         |
| guard.UpdatePackages              | [x]    | [x]        |         |
| guard.UpdatePrograms              | [x]    | [x]        |         |
| guard.UpdateServices              | [x]    | [x]        |         |
| watchtower.ListFindings           | [x]    |            |         |
| watchtower.ListGrants             | [x]    |            |         |
| watchtower.ListReports            | [x]    |            |         |
| watchtower.ListSystems            | [x]    |            |         |
| watchtower.CollectReport          | [x]    |            |         |
| watchtower.AccessAccept           | [x]    |            |         |
| watchtower.AccessChange           | [x]    |            |         |
| watchtower.AccessDeny             | [x]    |            |         |
| watchtower.AccessSystem           | [x]    |            |         |
| watchtower.SubscribeSystem        | [x]    |            |         |
| watchtower.UnsubscribeSystem      | [x]    |            |         |

