# 0.19.0 - Eagle Release ChangeLog

v0.19.0 codenamed Eagle, for Nebula marks the first quarterly release for Nebula in 2022. To remind, starting October 2021, the Nebula community has decided to use patch version for monthly releases and minor versions for Quarterly releases. 

The release is the last in set of the 3 releases. The prior 2 can be found at
- [v0.18.1 changelog](https://github.com/nebulaclouds/nebula/blob/master/CHANGELOG/CHANGELOG-v0.18.1.md)
- [v0.18.2 changelog](https://github.com/nebulaclouds/nebula/blob/master/CHANGELOG/CHANGELOG-v0.18.2.md)

The focus on this release and the previous 2 can be divided into 2 parts
 - UX (UI and Nebulakit)
 - System

As part of the UX, we have been focusing a lot on getting feature breadth in the UI for Nebula (nebulaconsole). This is following our current UX design philosophy. We plan to complete the full design in the next 3 months and then create a fully re-designed UX (Stay tuned for more on this.)

The design for Nebulakit has been wildly appreciated and we have been heavily working on improving it further and allowing users to express their various requirements much more easily. We are pretty certain on the design of nebulakit and we promise that all our users will not experience a breaking change when we release a v1.0.0. 

The focus on the system always is improving reliability and performance. This quarter we were able to drastically improve the size of workflows that can be executed and improvements in transition performance between nodes.

The Eagle release sets a great stage for v1.0.0 - Phoenix release slated for April 2022. We think, the Phoenix release will be a major milestone for the platform and we will be focusing on UX improvements.


# Updates specific to v0.19.0

[Closed Issues](https://github.com/nebulaclouds/nebula/issues?q=is%3Aissue+milestone%3A0.19.0+is%3Aclosed)

## UX
* New search UX for workflows
* Support for "security context" in launch workflow form
* Support for Google Analytics
* Minor fixes:
  - Removed DISABLE_AUTH env var
  - Issue with launch plan metadata not showing on execution view   

## NebulaKit

Please find the [full changelog here](https://github.com/nebulaclouds/nebulakit/releases/tag/v0.26.0).

### Changes
* Support for delayed annotations by @bethebunny in https://github.com/nebulaclouds/nebulakit/pull/760
  Users can now add `from __future__ import annotations` to the top of their code to stop using `""` in type hints.
* Add cache_serialize parameter to tasks by @hamersaw in https://github.com/nebulaclouds/nebulakit/pull/673
  Turn this on to avoid two cached tasks from running at the same time. Make sure you're at least at propeller `v0.16.1` or later and `0.3.17` or later in datacatalog.
* Use `functools.wraps` basically within the `task` and `workflow` decorators, @bstadlbauer in https://github.com/nebulaclouds/nebulakit/pull/780

## System
* Auto-horizontal-scale for NebulaPropeller [Docs](https://docs.nebula.org/en/latest/deployment/cluster_config/performance.html#automatic-scale-out)
* `nebulactl sandbox start` will add sandbox cluster context to the local kubeconfig. No more fiddling with env vars to get your KUBECONFIG right!
* Egress Events [Docs](https://docs.nebula.org/en/latest/deployment/cluster_config/eventing.html)
* Full reference of [Admin](https://docs.nebula.org/en/latest/deployment/cluster_config/nebulaadmin_config.html#nebulaadmin-config-specification) and [Propeller](https://docs.nebula.org/en/latest/deployment/cluster_config/nebulapropeller_config.html#nebulapropeller-config-specification) configs.
