# 0.19.2 Change Log

## UX
* Added support for Archive/Unarchive Workflow Executions
* Added Support for viewing node metadata on unexecuted nodes
* Minor bug fixes

## NebulaKit
* Introducing the new StructuredDataset type. Turn it on by setting the env var `NEBULA_SDK_USE_STRUCTURED_DATASET=TRUE`. It'll be turned on by default in v1.0. [Docs]()
* This release also removes the legacy API.

Review the full changelog [here](https://github.com/nebulaclouds/nebulakit/releases/tag/v0.30.0).

## System
* Archive ([issue](https://github.com/nebulaclouds/nebula/issues/2045)).
  Expose the ability to archive/unarchive executions to remove clutter from the UI.
* Security Updates:
    * Access token validation supports multiple-audience tokens ([issue](https://github.com/nebulaclouds/nebula/issues/1809))
    * nebulactl supports caching access tokens for multiple endpoints at the same time ([issue](https://github.com/nebulaclouds/nebula/issues/1962))
* You can use nebulactl to update projects' display names and descriptions ([issue](https://github.com/nebulaclouds/nebula/issues/1620))
    ```bash
    nebulactl update project  --id data --name datadata --description "Team that manage Data Platform" --labels "team=data,managedby=nebula"
    ```
