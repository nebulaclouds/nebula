# Nebula v1.0.0

Nebula 1.0 represents the first major release of the platform. Nebula APIs, and tools can now be considered **[stable]**. While it's impossible to assert so, any workflow/task written against nebulakit `v1.0.0` should be expected to run against all `v1.x.x` versions.

Quick Stats:
| Stat   |    Change      |
|--------|----------------|
| 4,125  | PRs Merged     |
| 18,983 | Comments Added |
| 1,511  | Issues Created |

## ChangeLog

### NebulaKit

* File ignore in fast register
Many thanks to @bimtauer for adding this much requested feature! When packaging local files for fast-registration, nebulakit will  now respect `.gitignore` and `.dockerignore`. Please see the [PR](https://github.com/nebulaclouds/nebulakit/pull/967) for more information.
* `pynebula run`
The pynebula run command has been slightly updated to not need a `:` when selecting the workflow.
    ```bash
    $ pynebula run --remote example.py wf --n 500 --mean 42 --sigma 2
    ```
* Script mode: register and run workflows all in one command using a pre-defined base image
* Nebula remote GA: register workflows and interact with Nebula execution artifacts programmmatically
* Configuration overhaul: use the same config across nebulakit and nebulactl
* Fast register without having AWS/GCP or other cloud credentials on your laptop, all you need is Nebula access

### Core Platform

* Improved Garbage Collector
    Garbage collection logic has been revamped to reduce load on KubeAPI and ensure terminated workflows are cleaned up in a timely fashion.
    **[Action Required]** Due to the change in the computation of how to clean up terminated workflows, users are advised to clean up old workflows once by running the following command per namespace:
    ```bash
    kubectl delete fly -l termination-status=terminated --all-namespaces --cascade='background' --wait=false --force --grace-period=0
    ```
* Single binary: deploy the entire Nebula back-end as a single binary. This speeds up sandbox and improves the local contributor experience. Coming soon: faster deployment for small-scale use-cases
* Improved map task subtask handling
    - Cache status reporting
    - Individual log links mapped to subtasks
    - Interruptible failure handling for spot instances
    - Secret injection
* Improved performance for fetching and rendering dynamic nodes
* Set raw output data config at create execution time
* Execution overrides at the project level for
    - Kubernetes service account
    - AssumableIAMRole
    - OutputLocationPrefix

### Console
* Project dashboard page with recent executions overview along with a config and other settings summary
* Dynamic workflow rendering
* Map task UI and UX improvements: see logs, retry attempts and more at the subtask level
