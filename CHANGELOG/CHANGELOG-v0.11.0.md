# Nebula v0.11.0

## Nebula Platform
* New to nebula? https://start.nebula.org takes you through first run experience. (Thanks to @jeevb)
* [Grafana templates](https://docs.nebula.org/en/latest/howto/monitoring/index.html) for monitoring Nebula System and User Workflows.
* [Extend Nebula](https://docs.nebula.org/en/latest/plugins/index.html) docs.
* [NebulaIdl Docs](https://docs.nebula.org/projects/nebulaidl/en/latest/) are published! You can learn about the core language that makes it all work.
* [Additional knob](https://github.com/nebulaclouds/nebulapropeller/pull/219/files#diff-91657d6448dfbf87f4cecf126ad02bd668ea233edcf74e860ef4f54bdd4cb552R78) for fine tuning nebula propeller performance that speeds up executions drastically.
* OidC support for Google Idp (And other OidC compliant Idps)
* Various stabilization bugs.

## Nebulakit
Since v0.16.0a2, the last nebulakit milestone release, all effort has been towards stabilizing the new API. Please see the individual [releases](https://github.com/nebulaclouds/nebulakit/releases) for detailed information. The highlights are

* Serialization/registration processes have been firmed up and utilities to ease that process introduced (not having to build a container to serialize for instance).
* Plugins structure revamped (eventually we'll move to a separate new repo entirely)
* User-facing imports have been organized into three top-level subpackages (`nebulakit`, `nebulakit.testing`, and `nebulakit.extend`)
* Retries added to read-only Admin calls in client
* Lots of cleanup and additions to the [cookbook](https://nebulacookbook.readthedocs.io/en/latest/) and documentation generally.
* Bug fixes.

