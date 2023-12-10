# Nebula v0.17.0

## Platform
1. Recovery Mode: Executions that fail due to external system failures (e.g. external system being down) can now be rerun in recovery mode ([nebulactl --recover docs](https://docs.nebula.org/projects/nebulactl/en/latest/gen/nebulactl_create_execution.html)). It's also available in the UI:

   <img src="https://i.imgur.com/hYYzkLK.png" alt="Recovery mode in the UI" width="800"/>


## Nebulakit
1. Great Expectations Integration ([docs](https://docs.nebula.org/projects/cookbook/en/latest/auto/integrations/nebulakit_plugins/greatexpectations/index.html#great-expectations)).
1. Access to durable blob stores (AWS/GCS/etc) are now pluggable.
1. Local task execution has been updated to also trigger the type engine.
1. Tasks that have `cache=True` should now be cached when running locally as well ([docs](https://docs.nebula.org/projects/cookbook/en/latest/auto/core/nebula_basics/task_cache.html#how-local-caching-works)).

Please see the [nebulakit release](https://github.com/nebulaclouds/nebulakit/releases/tag/v0.22.0) for the full list and more details.

## UI
1. Shiny new Graph UX. The graph rendering has been revamped to be more functional and accessible. More updates are coming for better visualization for nested executions and branches.

   <img src="https://i.imgur.com/HTfuios.png" alt="New Graph UX" width="800"/>
1. JSON Validation for json-based types in the UI.

   

| Before | After |
| -------- | -------- |
| ![](https://i.imgur.com/OKi4rEu.png) | ![](https://i.imgur.com/LX8yQ1x.png) |

1. Enum support in UI

   ![](https://i.imgur.com/9bFZlei.png)


## NebulaCtl
1. `nebulactl upgrade` to automatically upgrade itself ([docs](https://docs.nebula.org/projects/nebulactl/en/latest/gen/nebulactl_upgrade.html)).
1. `--dryRun` is available in most commands with server-side-effects to simulate the operations before committing any changes.

And various stabilization [fixes](https://github.com/nebulaclouds/nebula/milestone/17?closed=1)!
