# Nebula v1.7.0 release

In this release we're announcing support for Nebula Agents, a new way of writing backend plugins, only now with a much more tightly integrated developer experience. Also lots of bug fixes all around in a buch of first-time contributors.

## Nebulakit

* NEBULACTL_CONFIG env var should take highest precedence in https://github.com/nebulaclouds/nebulakit/pull/1662  by @wild-endeavor 
* Change nebulakit Pytorch, TFJob and MPI plugins to use new kubeflow config in https://github.com/nebulaclouds/nebulakit/pull/1627  by @yubofredwang 
* Root cert should be byte string when loading from caCertFilePath in https://github.com/nebulaclouds/nebulakit/pull/1669 by @wild-endeavor 
* Explicitly set the content type for nebula deck by in https://github.com/nebulaclouds/nebulakit/pull/1658 @pingsutw 
* Set a less strict deadline for hypothesis tests in https://github.com/nebulaclouds/nebulakit/pull/1682 by @eapolinario 
* Dep: Use protos of new `kubeflow.pytorch` plugin instead of legacy `pytorch` plugin in https://github.com/nebulaclouds/nebulakit/pull/1678 by @fg91 
* More time info for time line deck in https://github.com/nebulaclouds/nebulakit/pull/1680 by @Yicheng-Lu-llll

## Nebulaadmin

* Add logs and stats to sync loop in https://github.com/nebulaclouds/nebulaadmin/pull/573 by Haytham Abuelfutuh
* Fix utf-8 encoding issues with trimmed error messages in https://github.com/nebulaclouds/nebulaadmin/pull/569 by Katrina Rogan
* Tiny url improvements in https://github.com/nebulaclouds/nebulaadmin/pull/565 by Yee Hing Tong
* Update startedAt timestamp only if not set in https://github.com/nebulaclouds/nebulaadmin/pull/567 pmahindrakar-oss
* Save execution namespace in system metadata in https://github.com/nebulaclouds/nebulaadmin/pull/568 by Katrina Rogan

## Nebulaconsole

* fix: navlink in https://github.com/nebulaclouds/nebulaconsole/pull/772 by 4nalog
* fix: bump version for console in https://github.com/nebulaclouds/nebulaconsole/pull/769 by 4nalog
* fix: preserve domain when navigating using sidebar in https://github.com/nebulaclouds/nebulaconsole/pull/768 by 4nalog
* fix: dynamic-node-tasks in https://github.com/nebulaclouds/nebulaconsole/pull/765 by 4nalog
* chore: hide map task runtime info in https://github.com/nebulaclouds/nebulaconsole7/66 by Carina Ursu
* Bug: union val missing for LP relaunch in  https://github.com/nebulaclouds/nebulaconsole/pull/762 by Frank Flitton
* Feature: Fullview Nebula Deck modal in https://github.com/nebulaclouds/nebulaconsole/pull/764 by Frank Flitton
* chore: add item when mapped task in https://github.com/nebulaclouds/nebulaconsole/pull/761 by Jason Porter
* Bug: Execution Page's back button returns Workflows route from Launch Plan route https://github.com/nebulaclouds/nebulaconsolepatch in https://github.com/nebulaclouds/nebulaconsole/pull/760 by Frank Flitton

## Nebulapropeller

* Register gRPC plugin after reading configmap in https://github.com/nebulaclouds/nebulapropeller/pull/564 by Kevin Su
* Not stripping structure from literal types in https://github.com/nebulaclouds/nebulapropeller/pull/571 by Dan Rammer
* Bump nebulaplugins to v1.0.63 in https://github.com/nebulaclouds/nebulapropeller/pull/568 by bstadlbauer
* bumped nebulaplugins in https://github.com/nebulaclouds/nebulapropeller/pull/566 by Dan Rammer
* Use correct k8 client in https://github.com/nebulaclouds/nebulapropeller/pull/563 by sonjaer

