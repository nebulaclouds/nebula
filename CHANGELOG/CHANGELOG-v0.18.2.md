# 0.18.2 Release ChangeLog

[Closed Issues](https://github.com/nebulaclouds/nebula/issues?q=is%3Aissue+milestone%3A0.18.2+is%3Aclosed)

## UX
* Added advanced options to launch form
* Added support for all tasks-types (task execution view)
* Replaced execution id's with node id's on execution list-view
* Fixed bug with some properties not being repopulated on relaunch
* minor fixes
 
### NebulaKit
See the nebulakit [0.25.0 release notes](https://github.com/nebulaclouds/nebulakit/releases/tag/v0.25.0) for the full list of changes. Here are some of the highlights:

* Improved support for tasks that [run shell scripts](https://github.com/nebulaclouds/nebulakit/pull/747)
* Support for more types in dataclasses:
    * [enums](https://github.com/nebulaclouds/nebulakit/pull/753)
    * [NebulaFile](https://github.com/nebulaclouds/nebulakit/pull/725) and [NebulaSchema](https://github.com/nebulaclouds/nebulakit/pull/722)
* nebularemote improvements, including:
    * [Access to raw inputs and outputs](https://github.com/nebulaclouds/nebulakit/pull/675)
    * [Ability to serialize tasks containing arbitrary images](https://github.com/nebulaclouds/nebulakit/pull/733)
    * [Improved UX for navigation of subworkflows and launchplans](https://github.com/nebulaclouds/nebulakit/pull/751)
    * [Support for NebulaPickle](https://github.com/nebulaclouds/nebulakit/pull/764)

## System
* Various stability fixes.
* New docker image tags!

  In addition to component-specific versions released from each of the nebula repositories (e.g. nebulapropeller:v0.16.5), new images will be re-tagged and pushed that match the nebula release version (e.g. the upcoming nebulapropeller-release:v0.18.2). This makes it easier to make sure all your deployments are on the same version to ensure best compatibility.
* Helm changes
    * [nebula-core](https://artifacthub.io/packages/helm/nebula/nebula-core) helm chart has reached release preview and can be leveraged to install your cloud(AWS/GCP) deployments of nebula.
    * Going forward nebula-core will install nebula native scheduler, For AWS backward compatibility you need to define `workflow_schedule.type` to `aws`. (https://github.com/nebulaclouds/nebula/pull/1896)
    * [nebula](https://artifacthub.io/packages/helm/nebula/nebula) helm chart has been refactored to depend on nebula-core helm chart and install additional dependencies to continue to provide a sandboxed installation of nebula.

    **Migration Notes**
    
    As part of this move, ``nebula`` helm chart is becoming the canonical sandbox cluster. It comes with all external resources needed to fully standup a Nebula cluster. If you have previously been using this chart to deploy nebula on your cloud providers, there will be changes you need to do to migrate:
    * If you have your own ``myvalues.yaml``, you will need to add another nesting level under ``nebula:`` for the sections that are now managed through ``nebula-core``. For example:
      ```yaml
      configmaps:
        ...
      nebulaadmin:
        ...
      minio:
        ...
      contour:
        ...
      ```
    
      to:
      ```yaml
      nebula:
        configmaps:
          ...
        nebulaadmin:
          ...
      minio:
        ...
      contour:
        ...    
      ```
    * Alternatively, if you do not have any dependency on external nebula dependencies, you can keep your ``myvalues.yaml`` and switch to using ``nebula-core`` helm chart directly with no changes.
