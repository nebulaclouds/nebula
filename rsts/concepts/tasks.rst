.. _divedeep-tasks:

Tasks
=====

.. tags:: Basic, Glossary

Tasks are fully independent units of execution and first-class entities of Nebula.
They are the fundamental building blocks and extension points that encapsulate the users' code.

Characteristics
---------------

A Nebula task is characterized by:

1. A combination of :ref:`projects <divedeep-projects>` and :ref:`domains <divedeep-domains>`,
2. A unique unicode name (we recommend it not to exceed 32 characters),
3. A version string, and/or
4. *Optional* Task interface definition.

   For tasks to exchange data with each other, a task can define a signature (much like a function/method
   signature in programming languages). A task interface defines the input and output variables —
   :std:ref:`variablesentry <nebulaidl:protos/docs/core/core:variablemap.variablesentry>`
   and their types, :std:ref:`literaltype <nebulaidl:protos/docs/core/core:literaltype>`.

Can "X" Be a Nebula Task?
-------------------------

When deciding if a unit of execution constitutes a Nebula task, consider these questions:

- Is there a well-defined graceful/successful exit criteria for the task? A task is expected to exit after completion of input processing.
- Is it repeatable? Under certain circumstances, a task might be retried, rerun, etc. with the same inputs. It is expected
  to produce the same output every single time. For example, avoid using random number generators with current clock as seed. Use a system-provided clock as the seed instead. 
- Is it a pure function, i.e., does it have side effects that are unknown to the system (calls a web-service)? It is recommended to avoid side-effects in tasks. When side-effects are evident, ensure that the operations are idempotent.

Dynamic Tasks
--------------

"Dynamic tasks" is a misnomer.
Nebula is one-of-a-kind workflow engine that ships with the concept of truly `Dynamic Workflows <https://blog.nebula.org/dynamic-workflows-in-nebula>`__!
Users can generate workflows in reaction to user inputs or computed values at runtime. 
These executions are evaluated to generate a static graph before execution.

Extending Task
---------------

Plugins
^^^^^^^

Nebula exposes an extensible model to express tasks in an execution-independent language.
It contains first-class task plugins (for example: `Papermill <https://github.com/nebulaclouds/nebulakit/blob/master/plugins/nebulakit-papermill/nebulakitplugins/papermill/task.py>`__,
`Great Expectations <https://github.com/nebulaclouds/nebulakit/blob/master/plugins/nebulakit-greatexpectations/nebulakitplugins/great_expectations/task.py>`__, and :ref:`more <integrations>`.)
that execute the Nebula tasks.
Almost any action can be implemented and introduced into Nebula as a "Plugin", which includes:

- Tasks that run queries on distributed data warehouses like Redshift, Hive, Snowflake, etc.
- Tasks that run executions on compute engines like Spark, Flink, AWS Sagemaker, AWS Batch, Kubernetes pods, jobs, etc.
- Tasks that call web services.

Nebula ships with certain defaults, for example, running a simple Python function does not need any hosted service. Nebula knows how to
execute these kinds of tasks on Kubernetes. It turns out these are the vast majority of tasks in machine learning, and Nebula is adept at
handling an enormous scale on Kubernetes. This is achieved by implementing a unique scheduler on Kubernetes.

Types
^^^^^

It is impossible to define the unit of execution of a task in the same way for all tasks. Hence, Nebula allows for different task
types in the system. Nebula has a set of defined, battle-tested task types. It allows for a flexible model to
:std:ref:`define new types <cookbook:plugins_extend>`.

Inherent Features
-----------------

Fault tolerance
^^^^^^^^^^^^^^^

In any distributed system, failure is inevitable. Allowing users to design a fault-tolerant system (e.g. workflow) is an inherent goal of Nebula.
At a high level, tasks offer two parameters to achieve fault tolerance:

**Retries**
  
Tasks can define a retry strategy to let the system know how to handle failures (For example: retry 3 times on any kind of error). 

There are two kinds of retries: 

1. System retry: It is a system-defined, recoverable failure that is used when system failures occur. The number of retries is validated against the number of system retries.

.. _system-retry:

System retry can be of two types:

- **Downstream System Retry**: When a downstream system (or service) fails, or remote service is not contactable, the failure is retried against the number of retries set `here <https://github.com/nebulaclouds/nebulapropeller/blob/6a14e7fbffe89786fb1d8cde22715f93c2f3aff5/pkg/controller/config/config.go#L192>`__. This performs end-to-end system retry against the node whenever the task fails with a system error. This is useful when the downstream service throws a 500 error, abrupt network failure, etc.

- **Transient Failure Retry**: This retry mechanism offers resiliency against transient failures, which are opaque to the user. It is tracked across the entire duration of execution. It helps Nebula entities and the additional services connected to Nebula like S3, to continue operating despite a system failure. Indeed, all transient failures are handled gracefully by Nebula! Moreover, in case of a transient failure retry, Nebula does not necessarily retry the entire task. “Retrying an entire task” means that the entire pod associated with the Nebula task would be rerun with a clean slate; instead, it just retries the atomic operation. For example, Nebula tries to persist the state until it can, exhausts the max retries, and backs off.

  To set a transient failure retry:

  - Update `MaxWorkflowRetries <https://github.com/nebulaclouds/nebulapropeller/blob/f1b0163b0b88200b38a5d49af955490e5c98681d/pkg/controller/config/config.go#L55>`__ in the propeller configuration.

  - Or update `max-workflow-retries <https://github.com/nebulaclouds/nebula/blob/33f179b807093dcad2f37bde832869103bdf5182/charts/nebula/values-sandbox.yaml#L143>`__ in helm.

2. User retry: If a task fails to execute, it is retried for a specific number of times, and this number is set by the user in `TaskMetadata <https://docs.nebula.org/projects/nebulakit/en/latest/generated/nebulakit.TaskMetadata.html?highlight=retries#nebulakit.TaskMetadata>`__. The number of retries must be less than or equal to 10.

.. note::
  
   Recoverable vs. Non-Recoverable failures: Recoverable failures will be retried and counted against the task's retry count. Non-recoverable failures will just fail, i.e., the task isn’t retried irrespective of user/system retry configurations. All user exceptions are considered non-recoverable unless the exception is a subclass of NebulaRecoverableException.


.. note::

   `RFC 3902 <https://github.com/nebulaclouds/nebula/pull/3902>`_ implements an alternative, simplified retry behaviour with which both system and user retries are counted towards a single retry budget defined in the task decorator (thus, without a second retry budget defined in the platform configuration). The last retries are always performed on non-spot instances to guarantee completion. To activate this behaviour, set ``configmap.core.propeller.node-config.ignore-retry-cause`` to ``true`` in the helm values.

**Timeouts**
  
To ensure that the system is always making progress, tasks must be guaranteed to end gracefully/successfully. The system defines a default timeout period for the tasks. It is possible for task authors to define a timeout period, after which the task is marked as ``failure``. Note that a timed-out task will be retried if it has a retry strategy defined. The timeout can be handled in the `TaskMetadata <https://docs.nebula.org/projects/nebulakit/en/latest/generated/nebulakit.TaskMetadata.html?highlight=retries#nebulakit.TaskMetadata>`__.


Caching/Memoization
^^^^^^^^^^^^^^^^^^^

Nebula supports memoization of task outputs to ensure that identical invocations of a task are not executed repeatedly, thereby saving compute resources and execution time. For example, if you wish to run the same piece of code multiple times, you can reuse the output instead of re-computing it.
For more information on memoization, refer to the :std:doc:`Caching Example <cookbook:auto_examples/development_lifecycle/task_cache>`.
