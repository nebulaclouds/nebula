# Nebula v1.6.0-b1

## What's Changed
### Console
* feat: show launchplan in execution table by @pradithya in https://github.com/nebulaclouds/nebulaconsole/pull/738
* feat: show launch plan information in workflow's schedules by @pradithya in https://github.com/nebulaclouds/nebulaconsole/pull/739
* fix: passthrough runtime env vars by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/741
* chore: add fallback to task execution link  by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/743
* chore: allow custom subnav by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/734
* fix: force node executions to pull their status  by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/737
* chore: fix details panel card padding by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/745
* chore: fix crash by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/746
* [UI Feature] Add full-list log output to execution detail panel by @james-union in https://github.com/nebulaclouds/nebulaconsole/pull/744
* TLM add log-message window to left panel by @james-union in https://github.com/nebulaclouds/nebulaconsole/pull/748
* [Snyk] Upgrade eslint from 8.31.0 to 8.33.0 by @EngHabu in https://github.com/nebulaclouds/nebulaconsole/pull/695
* chore: [tlm] comprehensive node execution query by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/749
* chore: guard against /tasks failing by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/750
* chore: propagate dynamic parent id  by @ursucarina in https://github.com/nebulaclouds/nebulaconsole/pull/751
* Add support fetching description entity by @pingsutw in https://github.com/nebulaclouds/nebulaconsole/pull/735

### Admin
* Infer GOOS and GOARCH from environment by @jeevb in https://github.com/nebulaclouds/nebulaadmin/pull/550
* Enrich TerminateExecution error to tell propeller the execution already terminated by @EngHabu in https://github.com/nebulaclouds/nebulaadmin/pull/551
* Address resolution by @wild-endeavor in https://github.com/nebulaclouds/nebulaadmin/pull/546
* Add migration to turn `parent_id` column into `bigint` only if necessary by @eapolinario in https://github.com/nebulaclouds/nebulaadmin/pull/554

### Propeller
* Moved controller-runtime start out of webhook Run function by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/546
* Fixing recovering of SKIPPED nodes by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/551
* Remove resource injection on the node for container task by @ByronHsu in https://github.com/nebulaclouds/nebulapropeller/pull/544
* Infer GOOS and GOARCH from environment by @jeevb in https://github.com/nebulaclouds/nebulapropeller/pull/552
* fix makefile to read variables from environment and overrides by @jeevb in https://github.com/nebulaclouds/nebulapropeller/pull/554
* Remove BarrierTick by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/545
* Check for TerminateExecution error and eat Precondition status by @EngHabu in https://github.com/nebulaclouds/nebulapropeller/pull/553
* Setting primaryContainerName by default on Pod plugin by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/555
* Implement ability to specify additional/override annotations when using Vault Secret Manager by @pradithya in https://github.com/nebulaclouds/nebulapropeller/pull/556
* Maintaining Interruptible and OverwriteCache for reference launchplans by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/557
* Added support for aborting task nodes reported as failures by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/541
* Added support for EnvironmentVariables on ExecutionConfig by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/558
* Fast fail if task resource requests exceed k8s resource limits by @hamersaw in https://github.com/nebulaclouds/nebulapropeller/pull/488

* @ByronHsu made their first contribution in https://github.com/nebulaclouds/nebulapropeller/pull/544

