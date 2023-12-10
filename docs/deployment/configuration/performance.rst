.. _deployment-configuration-performance:

######################################################
Optimizing Performance
######################################################

.. tags:: Infrastructure, Kubernetes, Advanced

.. tip:: Before getting started, it is always important to measure the performance. Nebula project publishes and manages some grafana templates as described in - :ref:`deployment-configuration-monitoring`.

The video below contains an overview of the Nebula architecture, what is meant by "performance", details of one loop in NebulaPropeller, and a demo of the Grafana Labs dashboard.

..  youtube:: FJ-rG9lZDhY 

Scaling up NebulaPropeller
==========================
`NebulaPropeller <https://pkg.go.dev/github.com/nebulaclouds/NebulaPropeller>`_ is the core engine of Nebula that executes the workflows for Nebula.
It is designed as a `Kubernetes Controller <https://kubernetes.io/docs/concepts/architecture/controller/>`_, where the desired state is specified as a NebulaWorkflow `Custom Resource <https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/>`_.

One of the base assumptions of NebulaPropeller is that every workflow is independent and can be executed by a completely distinct process, without a need for communication with other processes. Meanwhile, one workflow tracks the dependencies between tasks using a DAG structure and hence constantly needs synchronization.
Currently, NebulaPropeller executes Workflows by using an event loop to periodically track and amend the execution status. Within each iteration, a single thread requests the state of Workflow nodes and performs operations (i.e., scheduling node executions, handling failures, etc) to gracefully transition a Workflow from the observed state to the desired state (i.e., Success). Consequently, actual node executions are performed by various NebulaPlugins, a diverse collection of operations spanning k8s and other remote services, and NebulaPropeller is only responsible for effectively monitoring and managing these executions.

NebulaPropeller has a lot of knobs that can be tweaked for performance. The default configuration is usually adequate for small to medium sized installations of Nebula, that are running about 500 workflows concurrently with no noticeable overhead. In the case when the number of workflows increases,
NebulaPropeller will automatically slow down, without losing correctness.

Signs of slowdown
------------------

.. list-table:: Important metrics to look out for in Prometheus dashboard
   :widths: 25 25 50 100
   :header-rows: 1

   * - Metric
     - Dashboard
     - Alerting factor
     - Effect
   * - ``nebula:propeller:all:free_workers_count``
     - Nebula Propeller Dashboard
     - Number of free-workers in all NebulaPropeller instances are very low.
     - This will increase overall latency for Each workflow propagation
   * - ``sum(rate(nebula:propeller:all:round:raw_ms[5m])) by (wf)``
     - Nebula Propeller Dashboard
     - Round Latency for Each workflow increases
     - Nebula propeller is taking more time to process each workflow
   * - ``sum(rate(nebula:propeller:all:main_depth[5m]))``
     - Nebula Propeller Dashboard
     - Workflows take longer to start or slow
     - The processing queue depth is long and is taking long to drain

For each of the metrics above you can dig deeper into the cause for this spike in latency. All of them are mostly related to one latency and should be correlated - ``The Round latency!``.
The round-latency is the time it takes for NebulaPropeller to to perform a single iteration of workflow evaluation. To understand this, you have to understand the :ref:`divedeep-execution-timeline`. Once you understand this, continue reading on.

Optimizing round latency
-------------------------

Optimize NebulaPropeller configuration
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

Usually round-latency can be resolved by adjusting NebulaPropeller config specified `here <https://pkg.go.dev/github.com/nebulaclouds/NebulaPropeller@v0.10.3/pkg/controller/config>`_ or sometimes adjusting the global defaults for your deployment or per workflow-execution.
Let us first look at various config properties that can be set and would impact the performance of one round in Nebula and should be tweaked

.. list-table:: Important Properties
   :widths: 25 25 25 50
   :header-rows: 1

   * - Property
     - Section
     - Rule of thumb
     - Description
   * - ``workers``
     - propeller
     - Larger the number, implies more workflows can be evaluated in parallel. But it should depend on number of CPU cores assigned to NebulaPropeller and evaluated against the cost of context switching. A number usually < 500 - 800 with 4-8 cpu cores works fine.
     - Number of `logical threads` workers, that can work concurrently. Also implies number of workflows that can be executed in parallel. Since NebulaPropeller uses go-routines, it can run way  more than number of physical cores.
   * - ``workflow-reeval-duration``
     - propeller
     - lower the number - lower latency, lower throughput (low throughput is because the same workflow will be evaluated more times)
     - frequency at which, given no external signal, a workflow should be re-evaluated by Nebula propellers reval loop
   * - ``downstream-eval-duration``
     - propeller
     - lower the number - lower latency and lower throughput (low throughput is because the same workflow will be evaluated more times)
     - This indicates how often are external events like pods completion etc recorded.
   * - ``max-streak-length``
     - propeller
     - higher the number lower the latency for end to end workflow, especially for cached workflows
     - number of consecutive rounds to try with one workflow - prioritize a hot workflow over others.
   * - ``kube-client-config``
     - propeller
     - This is how you can control the number of requests ceiling that NebulaPropeller can initiate to KubeAPI. This is usual the #1 bottle neck
     - this configures the kubernetes client used by NebulaPropeller
   * - ``workflowStore.policy``
     - propeller
     - This config uses a trick in etcD to minimize number of redundant loops in NebulaPropeller, thus improving free slots
     - Use this to configure how NebulaPropeller should evaluate workflows, the default is usually a good choice
   * - ``storage.cache``
     - propeller
     - This config is used to configure the write-through cache used by NebulaPropeller on top of the metastore
     - NebulaPropeller uses the configure blob-store (can be changed to something more performant in the future) to optimize read and write latency, for all metadata IO operations. Metadata refers to the input and output pointers
   * - ``admin-launcher.tps``, ``admin-launcher.cacheSize``, ``admin-launcher.workers``
     - propeller
     - This config is used to configure the max rate and launch-plans that NebulaPropeller can launch against NebulaAdmin
     - It is essential to limit the number of writes from NebulaPropeller to nebulaadmin to prevent brown-outs or request throttling at the server. Also the cache reduces number of calls to the server.
   * - ``tasks.backoff.max-duration``
     - propeller
     - This config is used to configure the maximum back-off interval in case of resource-quota errors
     - NebulaPropeller will automatically back-off when k8s or other services request it to slowdown or when desired quotas are met.
   * - ``max-parallelism``
     - admin, per workflow, per execution
     - Refer to examples and documentation below
     - docs below


In the above table the 2 most important configs are ``workers`` and ``kube-client-config``.

The Kube client config controls the request throughput from NebulaPropeller to the Kube API server. These requests may include creating/monitoring Pods or creating/updating NebulaWorkflow CRDs to track workflow execution. The default configuration (provided by k8s) contains very steep rate-limiting, and therefore NebulaPropeller provides a default configuration that offers better performance. However, if your workload involves larger scales (e.g., >5k fanout dynamic or map tasks, >8k concurrent workflows, etc.,) the Kube client config rate limiting may still contribute to a noticeable drop in performance. Increasing the ``qps`` and ``burst`` values may help alleviate back pressure and improve NebulaPropeller performance. An example of Kube-client-config is as follows:

.. code-block:: yaml

    propeller:
        kube-client-config:
            qps: 100 # Refers to max rate of requests to KubeAPI server
            burst: 50 # refers to max burst rate to Kube API server
            timeout: 30s # Refers to timeout when talking with kubeapi server


.. note:: As you increase the number of workers in NebulaPropeller it is important to increase the number of cpu's given to NebulaPropeller pod.

It is worth noting that the Kube API server tends to throttle requests transparently. This means that while tweaking performance by increasing the allowed frequency of Kube API server requests (e.g., increasing NebulaPropeller workers or relaxing Kube client config rate-limiting), there may be steep performance decreases for no apparent reason. Looking at the Kube API server request queue metrics in these cases can assist in identifying whether throttling is to blame. Unfortunately, there is no one-size-fits-all solution here, and customizing these parameters for your workload will require trial and error.

Another area of slowdown could be the size of the input-output cache that NebulaPropeller maintains in-memory. This can be configured, while configuring
the storage for NebulaPropeller. Rule of thumb, for NebulaPropeller with x memory limit, allocate x/2 to the cache

Learn: max-streak-length & ResourceVersionCaching
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
Kubernetes controllers often use Informer caches, rather than reading data directly from KubeAPI. This is to prevent excessive requests to KubeAPI server. The caches are eventually consistent, i.e., every write by the controller is eventually replicated to the cache, but there can be time periods, when the cache lags.
Since NebulaPropeller, runs Workflow evaluations as an event loop, which is triggered by any changes to one of the resources that a workflow spawned.
It is possible that a Workflow will be evaluated, even when the last write has not yet propagated to the Informer cache. EtcD also does not allow stale writes, i.e., writes with an object that is older than the object that was written. This is maintained using a server side vector-clock - called the resource version.
Stale writes are writes when the evaluation resulted in a mutation of an object that is older than the object recorded in etcD.
These stale writes often lead to conflicts and hence increase load on the KubeAPI server and on NebulaPropeller as the workers are busy writing stale objects repeatedly.

To prevent this duplication and redundancy, NebulaPropeller employs a trick. For every write, it records the last known version number in the database and then tries to wait for the change to propagate to the informer cache.

If `max-streaks` are enabled then instead of waiting for the informer cache to be refreshed, NebulaPropeller uses its own inmemory copy to run multiple rounds as long as mutations occur or the max-streak-length configuration is met. This reduces the latency of cache propagation, which can be order of seconds.

Worst case workflows: Poison Pills & max-parallelism
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
The worst case for NebulaPropeller is workflows that have an extremely large fan-out. This is because NebulaPropeller implements a greedy traversal algorithm, that tries to evaluate the entire unblocked nodes within a workflow in every round.
A solution for this is to limit the maximum number of nodes that can be evaluated. This can be done by setting max-parallelism for an execution.
This can done in multiple ways

#. Platform default: This allows to set platform-wide defaults for maximum concurrency within a Workflow execution. This can be overridden per Launch plan or per execution.
   The default `maxParallelism is configured to be 25 <https://github.com/nebulaclouds/nebulaadmin/blob/master/pkg/runtime/application_config_provider.go#L40>`_.
   It can be overridden with this config block in nebulaadmin

   .. code-block:: yaml

       nebulaadmin:
          maxParallelism: 25

#. Default for a specific launch plan. For any launch plan, the maxParallelism value can be changed or altered. This can be done using :py:meth:`nebulakit.LaunchPlan.get_or_create` or the :std:ref:`ref_nebulaidl.admin.LaunchPlanCreateRequest`
   **Nebulakit Example**

   .. code-block:: python

       LaunchPlan.get_or_create(
         name="my_cron_scheduled_lp",
         workflow=date_formatter_wf,
         max_parallelism=30,
       )

#. Specify for an execution. For any specific execution the max-parallelism can be overridden. This can be done using nebulactl (and soon nebulaconsole). Refer to :std:ref:`nebulaCtl docs <nebulactl:nebulactl_create_execution>`




Scaling out NebulaAdmin
=======================
NebulaAdmin is a stateless service. Often time before needing to scale NebulaAdmin, you need to scale the backing database. Check out the NebulaAdmin Dashboard to see signs of latency degradation and increase the size of backing postgres instance.
NebulaAdmin is a stateless service and its replicas (in the kubernetes deployment) can be simply increased to allow higher throughput.

Scaling out Datacatalog
========================
Datacatalog is a stateless service. Often time before needing to scale Datacatalog, you need to scale the backing database. Check out the Datacatalog Dashboard to see signs of latency degradation and increase the size of backing postgres instance.
Datacatalog is a stateless service and its replicas (in the kubernetes deployment) can be simply increased to allow higher throughput.

Scaling out NebulaPropeller
===========================

Manual scale-out
----------------
NebulaPropeller can be run manually per namespace. This is not a recommended solution as it is harder to deploy, but if your organization can deploy and maintain multiple copies of NebulaPropeller, you can use this.

Automatic scale-out
-------------------
NebulaPropeller Manager is a new component introduced as part of `this RFC <https://github.com/nebulaclouds/nebula/blob/master/rfc/system/1483-nebulapropeller-horizontal-scaling.md>`_ to facilitate horizontal scaling of NebulaPropeller through sharding. Effectively, the Manager is responsible for maintaining liveness and proper configuration over a collection of NebulaPropeller instances. This scheme uses k8s label selectors to deterministically assign NebulaWorkflow CRD responsibilities to NebulaPropeller instances, effectively distributing processing load over the shards.

Deployment of NebulaPropeller Manager requires k8s configuration updates including a modified NebulaPropeller Deployment and a new PodTemplate defining managed NebulaPropeller instances. The easiest way to apply these updates is by setting the "nebulapropeller.manager" value to "true" in the `helm deployment <https://docs.nebula.org/en/latest/deployment/overview.html#usage-of-helm>`_ and setting the manager config at "configmap.core.manager".

Nebula provides a variety of Shard Strategies to configure how NebulaWorkflows are sharded among managed NebulaPropeller instances. These include hash, which uses consistent hashing to load-balance evaluation over shards, and project / domain, which map the respective IDs to specific managed NebulaPropeller instances. Below we include examples of helm configurations for each of the existing Shard Strategies.

The Hash Shard Strategy, denoted by "type: hash" in the configuration below, uses consistent hashing to evenly distribute NebulaWorkflows over managed NebulaPropeller instances. This configuration requires a "shard-count" variable which defines the number of managed NebulaPropeller instances.

.. code-block:: yaml

    configmap:
      core:
        # a configuration example using the "hash" shard type
        manager:
          # pod and scanning configuration redacted
          # ...
          shard:
            type: hash     # use the "hash" shard strategy
            shard-count: 4 # the total number of shards
 
The Project and Domain Shard Strategies, denoted by "type: project" and "type: domain" respectively, use the NebulaWorkflow project and domain metadata to shard NebulaWorkflows. These Shard Strategies are configured using a "per-shard-mapping" option, which is a list of ID lists. Each element in the "per-shard-mapping" list defines a new shard and the ID list assigns responsibility for the specified IDs to that shard. A shard configured as a single wildcard ID (i.e. "*") is responsible for all IDs that are not covered by other shards. Only a single shard may be configured with a wildcard ID and on that shard their must be only one ID, namely the wildcard.

.. code-block:: yaml

    configmap:
      core:
        # a configuration example using the "project" shard type
        manager:
          # pod and scanning configuration redacted
          # ...
          shard:
            type: project       # use the "project" shard strategy
            per-shard-mapping:  # a list of per shard mappings - one shard is created for each element
              - ids:            # the list of ids to be managed by the first shard
                - nebulasnacks
              - ids:            # the list of ids to be managed by the second shard
                - nebulaexamples
                - nebulalabs
              - ids:            # the list of ids to be managed by the third shard
                - "*"           # use the wildcard to manage all ids not managed by other shards
    
    configmap:
      core:
        # a configuration example using the "domain" shard type
        manager:
          # pod and scanning configuration redacted
          # ...
          shard:
            type: domain        # use the "domain" shard strategy
            per-shard-mapping:  # a list of per shard mappings - one shard is created for each element
              - ids:            # the list of ids to be managed by the first shard
                - production
              - ids:            # the list of ids to be managed by the second shard
                - "*"           # use the wildcard to manage all ids not managed by other shards
 
Multi-Cluster mode
===================
In our experience at Lyft, we saw that the Kubernetes cluster would have problems before NebulaPropeller or NebulaAdmin would have impact. Thus Nebula supports adding multiple dataplane clusters by default. Each dataplane cluster, has one or more NebulaPropellers running in them, and nebulaadmin manages the routing and assigning of workloads to these clusters.


Improving etcd Performance
===========================

Offloading Static Workflow Information from CRD
-----------------------------------------------

Nebula uses a k8s CRD (Custom Resource Definition) to store and track workflow executions. This resource includes the workflow definition, for example tasks and subworkflows that are involved and the dependencies between nodes, but also includes the execution status of the workflow. The latter information (ie. runtime status) is dynamic, meaning changes during the workflow's execution as nodes transition phases and the workflow execution progresses. However, the former information (ie. workflow definition) remains static, meaning it will never change and is only consulted to retrieve node definitions and workflow dependencies.

CRDs are stored within etcd, a key-value datastore heavily used in kubernetes. Etcd requires a complete rewrite of the value data every time a single field changes. Consequently, the read / write performance of etcd, as with all key-value stores, is strongly correlated with the size of the data. In Nebula's case, to guarantee only-once execution of nodes we need to persist workflow state by updating the CRD at every node phase change. As the size of a workflow increases this means we are frequently rewriting a large CRD. In addition to poor read / write performance in etcd this update may be restricted by a hard limit on the overall CRD size.

To counter the challenges of large NebulaWorkflow CRDs Nebula includes a configuration option to offload the static portions of the CRD (ie. workflow / task / subworkflow definitions and node dependencies) to the blobstore. This functionality can be enabled by setting the ``useOffloadedWorkflowClosure`` option to ``true`` in the `NebulaAdmin configuration <https://docs.nebula.org/en/latest/deployment/cluster_config/nebulaadmin_config.html#useoffloadedworkflowclosure-bool>`_. When set, the NebulaWorkflow CRD will populate a ``WorkflowClosureReference`` field on the CRD with the location of the static data and NebulaPropeller will read this information (through a cache) during each workflow evaluation. One important note is that currently this requires NebulaAdmin and NebulaPropeller to have access to the same blobstore since NebulaAdmin only specifies a blobstore location in the CRD.
