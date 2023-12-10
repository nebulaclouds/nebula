# Nebula 1.2 Release

## Platform
- Support for Ray (https://github.com/nebulaclouds/nebula/issues/2641) - Also see the [blog post](https://blog.nebula.org/ray-and-nebula). 
- Execution names can be longer now, up to 63 characters (https://github.com/nebulaclouds/nebulaadmin/pull/466)
- Offloading NebulaWorkflow CRD static workflow spec (https://github.com/nebulaclouds/nebula/issues/2705)
- Enabled NebulaPropeller subqueue - this means that every time a pod is updated in the workflow it reevals for faster downstream scheduling
- Add container configuration to default pod template (https://github.com/nebulaclouds/nebula/issues/2703)
- Fixed issues with blobstore writes - GCS had duplicate writes and subworkflow inputs were rewritten on every evaluation, this meant slower evaluations
- Support external deletion of non-terminal map task subtasks (as a result of https://github.com/nebulaclouds/nebula/issues/2701)
- Better truncation of nebulapropeller logs (https://github.com/nebulaclouds/nebula/issues/2818)
- Bug: Recovering subworkflows is now supported (https://github.com/nebulaclouds/nebula/issues/2840)
- Fix snowflake plugin (https://github.com/nebulaclouds/nebulaplugins/pull/286)


## Nebulakit
- Hugging Face plugin (https://github.com/nebulaclouds/nebulakit/pull/1116)
- dbt plugin (https://github.com/nebulaclouds/nebula/issues/2202)
- cache overriding behavior is now open to all types (https://github.com/nebulaclouds/nebula/issues/2912)
- Bug: Fallback to pickling in the case of unknown types used Unions (https://github.com/nebulaclouds/nebula/issues/2823)
- [pynebula run](https://docs.nebula.org/projects/nebulakit/en/latest/design/clis.html#pynebula-run) now supports [imperative workflows](https://docs.nebula.org/projects/cookbook/en/latest/auto/core/nebula_basics/imperative_wf_style.html#sphx-glr-auto-core-nebula-basics-imperative-wf-style-py)
- Newlines are now stripped from client secrets (https://github.com/nebulaclouds/nebulakit/pull/1163)
- Ensure repeatability in the generation of cache keys in the case of dictionaries (https://github.com/nebulaclouds/nebulakit/pull/1126) 
- Support for multiple images in the yaml config file (https://github.com/nebulaclouds/nebulakit/pull/1106)

And more. See the full changelog in https://github.com/nebulaclouds/nebulakit/releases/tag/v1.2.0


## Nebulaconsole
- fix: Make sure groups used in graph aren't undefined [#545](https://github.com/nebulaclouds/nebulaconsole/pull/545)
- fix: Graph Center on initial render [#541](https://github.com/nebulaclouds/nebulaconsole/pull/541)
- fix: Graph edge overlaps nodes [#542](https://github.com/nebulaclouds/nebulaconsole/pull/542) 
- Fix searchbar X button [#564](https://github.com/nebulaclouds/nebulaconsole/pull/564)
- fix: Update timeline view to show dynamic wf internals on first render [#562](https://github.com/nebulaclouds/nebulaconsole/pull/562)
- fix: Webmanifest missing crossorigin attribute [#566](https://github.com/nebulaclouds/nebulaconsole/pull/566)
- fix: console showing subworkflows as unknown [#570](https://github.com/nebulaclouds/nebulaconsole/pull/570)
- fix: Dict value loses 1 trailing character on UI Launch. [#561](https://github.com/nebulaclouds/nebulaconsole/pull/561)
- fix: launchform validation [#557](https://github.com/nebulaclouds/nebulaconsole/pull/557)
- fix: integrate timeline and graph tabs wrappers under one component [#572](https://github.com/nebulaclouds/nebulaconsole/pull/572)
- added none type in union type [#577](https://github.com/nebulaclouds/nebulaconsole/pull/577)
- minor: inputHelpers InputProps [#579](https://github.com/nebulaclouds/nebulaconsole/pull/579)
- fix: fix test of launchform [#581](https://github.com/nebulaclouds/nebulaconsole/pull/581)
- Pruning some unused packages [#583](https://github.com/nebulaclouds/nebulaconsole/pull/583)
- fix: floor seconds to int in the edge case moment returns it as float [#582](https://github.com/nebulaclouds/nebulaconsole/pull/582)
- fix: add BASE_URL to dev startup, open deeply nested urls [#589](https://github.com/nebulaclouds/nebulaconsole/pull/589)
- fix: add default disabled state for only mine filter [#585](https://github.com/nebulaclouds/nebulaconsole/pull/585)
- fixed graph/timeline support for launchplanRef [#601](https://github.com/nebulaclouds/nebulaconsole/pull/601)
- fix: enable deeplinks in development [#602](https://github.com/nebulaclouds/nebulaconsole/pull/602)
