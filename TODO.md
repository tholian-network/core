# TODO

## source/protocols/tholian: Ongoing Refactor

- [ ] Make sure that ./source/protocols/tholian/methods/AllMethods.go returns all methods correctly
- [ ] Move ./source/protocols/tholian/methods/Methods.go into GuardMethods.go and WatchtowerMethods.go
- [ ] Verify firewall and scout methods, parameters, and results
- [ ] Continue implementing methods for guard and watchtower

- [ ] Make sure parameters are correctly in-sync with methods
- [ ] Make sure results are correctly in-sync with methods

| Identifier                   | Method | Parameters | Results |
|------------------------------|--------|------------|---------|
| guard.Offline                | [x]    | [x]        |         |
| guard.Online                 | [x]    | [x]        |         |
| watchtower.ListFindings      | [x]    |            |         |
| watchtower.ListGrants        | [x]    |            |         |
| watchtower.ListReports       | [x]    |            |         |
| watchtower.ListSystems       | [x]    |            |         |
| watchtower.CollectReport     | [x]    |            |         |
| watchtower.AccessAccept      | [x]    |            |         |
| watchtower.AccessChange      | [x]    |            |         |
| watchtower.AccessDeny        | [x]    |            |         |
| watchtower.AccessSystem      | [x]    |            |         |
| watchtower.SubscribeSystem   | [x]    |            |         |
| watchtower.UnsubscribeSystem | [x]    |            |         |

