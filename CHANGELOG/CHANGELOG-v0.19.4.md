# Nebula v0.19.4 Changelog

## Core Platform
- Single binary: deploy the entire Nebula back-end as a single binary. This speeds up sandbox and improves the local contributor experience. Coming soon: faster deployment for small-scale use-cases
- Improved map task subtask handling
    - Cache status reporting
    - Individual log links mapped to subtasks
    - Interruptible failure handling for spot instances
    - Secret injection
- Improved performance for fetching and rendering dynamic nodes
- Set raw output data config at create execution time
- Execution overrides at the project level for
    - Kubernetes service account
    - AssumableIAMRole
    - OutputLocationPrefix


## SDK
- Script mode: register and run workflows all in one command using a pre-defined base image
- Nebula remote GA: register workflows and interact with Nebula execution artifacts programmmatically
- Configuration overhaul: use the same config across nebulakit and nebulactl
- Fast register without having AWS or other cloud credentials on your laptop, all you need is Nebula access

## Console
- Project dashboard page with recent executions overview along with a config and other settings summary
- Dynamic workflow rendering
- Map task UI and UX improvements: see logs, retry attempts and more at the subtask level

