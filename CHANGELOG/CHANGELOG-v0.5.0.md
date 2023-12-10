# Nebula v0.5.0

## Infrastructure
- Moved CI/CD to Github Actions
- Added end-to-end tests as part of the PR & master merges. 
- Enable CI system to run on forks.

## Core Platform
- [Single Task Execution](https://nebula.readthedocs.io/en/latest/user/features/single_task_execution.html) to enable registering and launching tasks outside the scope of a workflow to enable faster iteration and a more intuitive development workflow.
- [Run to completion](https://nebula.readthedocs.io/en/latest/user/features/on_failure_policy.html) to enable workflows to continue executing even if one or more branches fail.
- Fixed retries for dynamically yielded nodes.
- PreAlpha Support for Raw container with NebulaCoPilot. (docs coming soon). [Sample Notebooks](https://github.com/lyft/nebulakit/blob/master/sample-notebooks/raw-container-shell.ipynb). This makes it possible to run workflows with arbitrary containers

## Plugins
- Cleaned up Presto executor to send namespace as user to improve queueing
