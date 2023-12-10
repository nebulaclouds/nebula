# Nebula 1.1 Changelog

## Platform
### User Improvements
Support for [Optional types](https://github.com/nebulaclouds/nebula/issues/2426). With the inclusion of Union types in nebulakit, we can now support optional types.

[Nebula Deck](https://github.com/nebulaclouds/nebula/issues/2175) is now available. Please take a look at the [documentation](https://docs.nebula.org/projects/cookbook/en/latest/auto/core/nebula_basics/deck.html#sphx-glr-auto-core-nebula-basics-deck-py) and also the [OSS presentation](https://www.youtube.com/watch?v=KqyBYIaAZ7c) that was done a few weeks back.


### Backend Improvements
* Allow different [cookie settings](https://github.com/nebulaclouds/nebula/issues/2596) in Admin.
* Pluggable [middleware](https://github.com/nebulaclouds/nebula/issues/2507) in Admin authentication.
* [Server-side compiler](https://github.com/nebulaclouds/nebula/issues/2516) strips Type Metadata.
* [Nebulactl now uses latest](https://github.com/nebulaclouds/nebula/issues/2329) launch plans for scheduled workflows
* [Allow spot instances](https://github.com/nebulaclouds/nebula/issues/2284) at workflow start time


### Bug Fixes
* [Propeller](https://github.com/nebulaclouds/nebula/issues/2298) calling finalize rather than abort
* [Propeller](https://github.com/nebulaclouds/nebula/issues/2404) correctly identify error when requesting a launch plan that does not exist.
* Better handle [execution CRDs](https://github.com/nebulaclouds/nebula/issues/2275) that don't exist in Admin.
* [Fix panic](https://github.com/nebulaclouds/nebula/issues/2597) when creating additional label options.
* Check [validity](https://github.com/nebulaclouds/nebula/issues/2601) of notifications.
* [Revert Spark tasks](https://github.com/nebulaclouds/nebulaadmin/pull/450) to use the `spark` role if unspecified by the user.

...and more!

## Nebulakit
* Work for Nebula Decks and Optional types.
* `pynebula run` now [supports](https://github.com/nebulaclouds/nebula/issues/2471) executing tasks.
* `pynebula register` combines the UX of `run` with the functionality of `package`. Please see the [video](https://www.youtube.com/watch?v=Z_KLl0qhp0Y) posted along with the forthcoming documentation.

### Bug Fixes
* Pynebula fixes - dot separated python packages [fix](https://github.com/nebulaclouds/nebula/issues/2476),  [`pynebula run`](https://github.com/nebulaclouds/nebula/issues/2474) to respect the `--config` flag and [read packages](https://github.com/nebulaclouds/nebulakit/pull/1002) from environment variables
* Set authorization key in the case of [external command](https://github.com/nebulaclouds/nebulakit/pull/1065)



## UI
* fix: support mapped tasks [#494](https://github.com/nebulaclouds/nebulaconsole/pull/494)
* feat: support nebula decks [#504](https://github.com/nebulaclouds/nebulaconsole/issues/504)
* feat: launch plans list & detail page [#507](https://github.com/nebulaclouds/nebulaconsole/issues/507)
* fix(bug-508): executions can not be filtered by start time [#509](https://github.com/nebulaclouds/nebulaconsole/issues/509)
* feat: navbar navigation dropdown [#511](https://github.com/nebulaclouds/nebulaconsole/issues/511)
* chore: support internal/external navigsation better [#513](https://github.com/nebulaclouds/nebulaconsole/issues/513)
* chore: Update Contributing.md [#515](https://github.com/nebulaclouds/nebulaconsole/issues/515)
* chore: update navigationDropdown usage [#517](https://github.com/nebulaclouds/nebulaconsole/issues/517)
* fix: cache icon for map task [#519](https://github.com/nebulaclouds/nebulaconsole/issues/519)
* fix: Relaunch form does not persist security context values when changed [#527](https://github.com/nebulaclouds/nebulaconsole/pull/527)
* fix: release process [#529](https://github.com/nebulaclouds/nebulaconsole/pull/529)
* fix: fix semantic-release config [#532](https://github.com/nebulaclouds/nebulaconsole/pull/532)
* test: fix time sensitive test [#533](https://github.com/nebulaclouds/nebulaconsole/pull/533)
* feat: Rename upgrade idl workflow [#534](https://github.com/nebulaclouds/nebulaconsole/pull/534)
* fix(491): remove favicon package + use favicon.svg by default [#537](https://github.com/nebulaclouds/nebulaconsole/pull/537)

*The lists above contains changes from the 1.0.x release as well.*

