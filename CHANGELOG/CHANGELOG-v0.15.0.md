# Nebula v0.15.0

## Platform
- Enum type support throughout the system
- Stabilization improvements for conditionals
- Support AWS Secrets Manager as a source for secrets injector
- Support Map tasks over Pod tasks
- Support max parallelism to limit how many nodes can be allowed to run in parallel.
- Add Athena nebulakit plugin and examples
- Add BigQuery plugin

## Nebulakit
 - Support Schema of Dataclasses
 - Support node resource overrides

Please see the [nebulakit release](https://github.com/nebulaclouds/nebulakit/releases/tag/v0.20.0) for the full list and more details.

## nebulactl
 - `nebulactl sandbox start` to start a sandbox cluster locally.
 - `nebulactl get workflow .... -o dot` to visualize a workflow graph locally.
 - Add Bash completion support
