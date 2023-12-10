# Nebula v0.17.3

## Platform
1.  [Nebula Native scheduler](https://www.youtube.com/watch?v=YljIIJx1_q8)
2.  [Support for Snowflake](https://github.com/nebulaclouds/nebulasnacks/blob/master/cookbook/integrations/external_services/snowflake/README.rst) including backend plugin and nebulakit tasks
3.  Expose default MaxParallelism in nebulaadmin [configuration](https://github.com/nebulaclouds/nebulaadmin/pull/262)
4.  Support [custom resource cleanup policy](https://github.com/nebulaclouds/nebula/issues/1345) in backend plugins
5.  Improved error message in the case of [images with invalid names](https://github.com/nebulaclouds/nebula/issues/306)


## nebulakit

1.  Continued changes to the NebulaRemote interface, including:
    - `remote.sync` now operates on the object in-place.
    - A `sync_nodes` argument has been added to the `remote.sync` call - by default it's True but set it to False if you want to only render inputs/outputs at the top level (rather than at every node within the execution).
    - The `sync` call on `NebulaWorkflowExecution` objects was removed (it was already deprecated and a noop but if you were calling it, you'll need to remove it.)
    - gRPC credentials to NebulaRemote
2.  Improved typing errors when scanning user code. Added better type inference to some of the transformers.
3.  Plugin tests have been moved into each individual plugin's `test` folder.
4.  Snowflake task has been added
5.  Fixes to SQLAlchemy task secrets handling
6.  Modin schema transformer

Please see the [nebulakit release](https://github.com/nebulaclouds/nebulakit/releases/tag/v0.23.0) for the full list and more details.


## UI
1.  Additional information when a Task in a non-terminal state
2.  Support for workflow versions


## nebulactl
1.  Sandbox docker images can now be provided as a parameter
2.  Bug fixes:
    -   panics in calls to get execution details and launchplans.
    -   datetime format generated in execFile are now valid
