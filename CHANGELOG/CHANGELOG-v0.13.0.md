# Nebula v0.13.0

## Platform
- Oauth2 support with SingleSignOn and configuration examples for popular IDP's now available in Nebula.
  Please see the updated [information and description of the feature](https://github.com/nebulaclouds/nebula/blob/master/docs/howto/authentication/index.rst), and the [setup information](https://github.com/nebulaclouds/nebula/blob/master/docs/howto/authentication/setup.rst)
  **Attention: If using Auth already - this is a BREAKING change**. refer to the [migration guide](https://github.com/nebulaclouds/nebula/blob/master/docs/howto/authentication/migration.rst) to update configuration to ensure Admin continues to work. (No migration needed if auth is not turned on.)

* Backend improvements to support dynamic workflow visualization (in future releases).
* Lot of features added to [nebulactl](https://nebulactl.readthedocs.io/en/latest/) .
* Documentation site overhaul and redesign (more in progress)

## Nebulakit
The first two features should be considered beta and subject to change
* First cut of the control plane classes to replace the old `Sdk...` classes. These classes provide programmatic access to a Nebula backend.
* New paradigm for nebulakit-only plugins
* SqlAlchemy/Dolt plugins.

Please see the [nebulakit release](https://github.com/nebulaclouds/nebulakit/releases/tag/v0.18.0) for the full list and more details.
