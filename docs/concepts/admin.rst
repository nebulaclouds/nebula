.. _divedeep-admin:

##########
NebulaAdmin
##########

.. tags:: Advanced, Design

Admin Structure
===============

NebulaAdmin serves as the main Nebula API to process all client requests to the system. Clients include the NebulaConsole, which calls:

1. NebulaAdmin to list the workflows, get execution details, etc.
2. Nebulakit, which in turn calls NebulaAdmin to register, launch workflows, etc.

Below, we'll dive into each component defined in admin in more detail.

RPC
---

NebulaAdmin uses the `grpc-gateway <https://github.com/grpc-ecosystem/grpc-gateway>`__ library to serve incoming gRPC and HTTP requests with identical handlers.
Refer to the admin service :std:ref:`definition <ref_nebulaidl/service/admin.proto>` for a detailed API overview, including request and response entities.
The RPC handlers are thin shims that enforce request structure validation and call out to the appropriate :ref:`manager <divedeep-admin-manager>` methods to process requests.

You can find a detailed explanation of the service in the :ref:`admin service <divedeep-admin-service>` page.

.. _divedeep-admin-manager:

Managers
--------

The Admin API is broken up into entities:

- Executions
- Launch plans
- Node Executions
- Projects (and their respective domains)
- Task Executions
- Tasks
- Workflows

Each API entity has an entity manager in NebulaAdmin responsible for implementing business logic for the entity.
Entity managers handle full validation of creating, updating and getting requests and
data persistence in the backing store (see the :ref:`divedeep-admin-repository` section).


Additional Components
+++++++++++++++++++++

The managers utilize additional components to process requests. These additional components include:

- :ref:`workflow engine <divedeep-admin-workflowengine>`: compiles workflows and launches workflow executions from launch plans.
- :ref:`data <divedeep-admin-data>` (remote cloud storage): offloads data blobs to the configured cloud provider.
- :ref:`runtime <divedeep-admin-config>`: loads values from a config file to assign task resources, initialization values, execution queues, and more.
- :ref:`async processes <divedeep-admin-async>`: provides functions to schedule and execute the workflows as well as enqueue and trigger notifications.

.. _divedeep-admin-repository:

Repository
----------
Serialized entities (tasks, workflows, launch plans) and executions (workflow-, node- and task-) are stored as protos defined
`here <https://github.com/nebulaclouds/nebulaidl/tree/master/protos/nebulaidl/admin>`__.
We use the excellent `gorm <https://gorm.io/docs/index.html>`__ library to interface with our database, which currently supports a Postgres
implementation.  You can find the actual code for issuing queries with gorm in the
`gormimpl <https://github.com/nebulaclouds/nebulaadmin/blob/master/pkg/repositories/gormimpl>`__ directory.

Models
++++++
Database models are defined in the `models <https://github.com/nebulaclouds/nebulaadmin/blob/master/pkg/repositories/models>`__ directory and correspond 1:1 with the database tables [0]_.

The full set of database tables includes:

- executions
- execution_events
- launch_plans
- node_executions
- node_execution_events
- tasks
- task_executions
- workflows

These database models inherit primary keys and indexes as defined in the corresponding `models <https://github.com/nebulaclouds/nebulaadmin/blob/master/pkg/repositories/models>`__ file.

The repositories code also includes `transformers <https://github.com/nebulaclouds/nebulaadmin/blob/master/pkg/repositories/transformers>`__.
These convert entities from the database format to a response format for the external API.
If you change either of these structures, you must change the corresponding transformers too.


.. _divedeep-admin-async:

Component Details
=================

This section dives into the details of each top-level directory defined in ``pkg/``.

Asynchronous Components
-----------------------

Notifications and schedules are handled by async routines that are responsible for enqueuing and subsequently processing dequeued messages.

NebulaAdmin uses the `gizmo toolkit <https://github.com/nytimes/gizmo>`__ to abstract queueing implementation. Gizmo's
`pubsub <https://github.com/nytimes/gizmo#pubsub>`__ library offers implementations for Amazon SNS/SQS, Google Pubsub, Kafka topics, and publishing over HTTP.

For the sandbox development, no-op implementations of the notifications and schedule handlers are used to remove external cloud dependencies.


Common
------

As the name implies, ``common`` houses shared components used across different NebulaAdmin components in a single, top-level directory to avoid cyclic dependencies. These components include execution naming and phase utils, query filter definitions, query sorting definitions, and named constants.

.. _divedeep-admin-data:

Data
-----

Data interfaces are primarily handled by the `storage <https://github.com/nebulaclouds/nebulastdlib>`__ library implemented in ``nebulastdlib``. However, neither this nor the underlying `stow <https://github.com/graymeta/stow>`__ library expose `HEAD <https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/HEAD>`__ support. Hence, the data package in admin exists as the layer responsible for additional, remote data operations.

Errors
------

The errors directory contains centrally defined errors that are designed for compatibility with gRPC statuses.

.. _divedeep-admin-config:

Runtime
-------
Values specific to the NebulaAdmin application, including task, workflow registration, and execution are configured in the `runtime <https://github.com/nebulaclouds/nebulaadmin/tree/master/pkg/runtime>`__ directory. These interfaces expose values configured in the ``nebulaadmin`` top-level key in the application config.

.. _divedeep-admin-workflowengine:

Workflow engine
----------------

This directory contains the interfaces to build and execute workflows leveraging NebulaPropeller compiler and client components.

.. [0] Given the unique naming constraints, some models are redefined in `migration_models <https://github.com/nebulaclouds/nebulaadmin/blob/master/pkg/repositories/config/migration_models.go>`__ to guarantee unique index values.

.. _divedeep-admin-service:


NebulaAdmin Service Background
=============================

Entities
---------

The :std:ref:`admin service definition <ref_nebulaidl/service/admin.proto>` defines REST operations for the entities that
NebulaAdmin administers.

As a refresher, the primary :ref:`entities <divedeep>` across Nebula maps to NebulaAdmin entities.

Static entities
+++++++++++++++

These include:

- Workflows
- Tasks
- Launch Plans

Permitted operations include:

- Create
- Get
- List

The above entities are designated by an :std:ref:`identifier <ref_nebulaidl.core.Identifier>`
that consists of a project, domain, name, and version specification. These entities are, for the most part, immutable. To update one of these entities, the updated
version must be re-registered with a unique and new version identifier attribute.

One caveat is that the launch plan can toggle between :std:ref:`ACTIVE and INACTIVE <ref_nebulaidl.admin.LaunchPlan>` states.
At a given point in time, only one launch plan version across a shared {Project, Domain, Name} specification can be active. The state affects the scheduled launch plans only.
An inactive launch plan can be used to launch individual executions. However, only an active launch plan runs on a schedule (given it has a schedule defined).


Static entities metadata (Named Entities)
+++++++++++++++++++++++++++++++++++++++++

A :std:ref:`named entity <ref_nebulaidl.admin.NamedEntity>` includes metadata for one of the above entities
(workflow, task or launch plan) across versions. It also includes a resource type (workflow, task or launch plan) and an
:std:ref:`id <ref_nebulaidl.admin.NamedEntityIdentifier>` which is composed of project, domain and name.
The named entity also includes metadata, which are mutable attributes about the referenced entity.

This metadata includes:

- Description: a human-readable description for the Named Entity collection.
- State (workflows only): this determines whether the workflow is shown on the overview list of workflows scoped by project and domain.

Permitted operations include:

- Create
- Update
- Get
- List


Execution entities
++++++++++++++++++

These include:

- (Workflow) executions
- Node executions
- Task executions

Permitted operations include:

- Create
- Get
- List

After an execution begins, NebulaPropeller monitors the execution and sends the events which the admin uses to update the above executions.

These :std:ref:`events <ref_nebulaidl/event/event.proto>` include

- WorkflowExecutionEvent
- NodeExecutionEvent
- TaskExecutionEvent

and contain information about respective phase transitions, phase transition time and optional output data if the event concerns a terminal phase change.

These events provide the **only** way to update an execution. No raw update endpoint exists.

To track the lifecycle of an execution, admin and store attributes such as `duration` and `timestamp` at which an execution transitioned to running and end time are used.

For debugging purposes, admin also stores Workflow and Node execution events in its database, but does not currently expose them through an API. Because array tasks can yield many executions, admin does **not** store TaskExecutionEvents.


Platform entities
+++++++++++++++++
Projects: Like named entities, projects have mutable metadata such as human-readable names and descriptions, in addition to their unique string ids.

Permitted project operations include:

- Register
- List

.. _divedeep-admin-matchable-resources:

Matchable resources
+++++++++++++++++++

A thorough background on :ref:`matchable resources <deployment-configuration-general>` explains
their purpose and application logic. As a summary, these are used to override system level defaults for Kubernetes cluster
resource management, default execution values, and more across different levels of specificity.

These entities consist of:

- ProjectDomainAttributes
- WorkflowAttributes

``ProjectDomainAttributes`` configure customizable overrides at the project and domain level, and ``WorkflowAttributes`` configure customizable overrides at the project, domain and workflow level.

Permitted attribute operations include:

- Update (implicitly creates if there is no existing override)
- Get
- Delete


Defaults
--------

Task resource defaults
++++++++++++++++++++++

User-facing documentation on configuring task resource requests and limits can be found in :std:ref:`cookbook:customizing task resources`.

As a system administrator you may want to define default task resource requests and limits across your Nebula deployment.
This can be done through the nebulaadmin config.

**Default** values get injected as the task requests and limits when a task definition omits a specific resource.
**Limit** values are only used as validation. Neither a task request nor limit can exceed the limit for a resource type.


Using the Admin Service
-----------------------

Adding request filters  
++++++++++++++++++++++  

We use `gRPC Gateway <https://github.com/grpc-ecosystem/grpc-gateway>`_ to reverse proxy HTTP requests into gRPC. 
While this allows for a single implementation for both HTTP and gRPC, an important limitation is that fields mapped to the path pattern cannot be 
repeated and must have a primitive (non-message) type. Unfortunately this means that repeated string filters cannot use a proper protobuf message. Instead, they use  
the internal syntax shown below:: 

 func(field,value) or func(field, value)  

For example, multiple filters would be appended to an http request like:: 

 ?filters=ne(version, TheWorst)+eq(workflow.name, workflow) 

Timestamp fields use the ``RFC3339Nano`` spec (For example: "2006-01-02T15:04:05.999999999Z07:00") 

The fully supported set of filter functions are 

- contains  
- gt (greater than) 
- gte (greter than or equal to) 
- lt (less than)  
- lte (less than or equal to) 
- eq (equal)  
- ne (not equal)  
- value_in (for repeated sets of values)  

"value_in" is a special case where multiple values are passed to the filter expression. For example:: 

 value_in(phase, RUNNING;SUCCEEDED;FAILED)  

.. note::
   If you're issuing your requests over http(s), be sure to URL encode the ";" semicolon using ``%3B`` like so: ``value_in(phase, RUNNING%3BSUCCEEDED%3BFAILED)``

Filterable fields vary based on entity types: 

- Task  

  - project 
  - domain  
  - name  
  - version 
  - created_at  
  
- Workflow  

  - project 
  - domain  
  - name  
  - version 
  - created_at
  
- Launch plans  

  - project 
  - domain  
  - name  
  - version 
  - created_at  
  - updated_at  
  - workflows.{any workflow field above} (for example: workflow.domain) 
  - state (you must use the integer enum, for example: 1)  
     - States are defined in :std:ref:`launchplanstate <ref_nebulaidl.admin.LaunchPlanState>`.
     
- Named Entity Metadata

  - state (you must use the integer enum, for example: 1)  
     - States are defined in :std:ref:`namedentitystate <ref_nebulaidl.admin.NamedEntityState>`.
     
- Executions (Workflow executions)  

  - project 
  - domain  
  - name  
  - workflow.{any workflow field above} (for example: workflow.domain)  
  - launch_plan.{any launch plan field above} (for example: launch_plan.name) 
  - phase (you must use the upper-cased string name, for example: ``RUNNING``) 
     - Phases are defined in :std:ref:`workflowexecution.phase <ref_nebulaidl.core.WorkflowExecution.Phase>`.
  - execution_created_at  
  - execution_updated_at  
  - duration (in seconds) 
  - mode (you must use the integer enum, for example: 1)  
     - Modes are defined in :std:ref:`executionmode <ref_nebulaidl.admin.ExecutionMetadata.ExecutionMode>`.
  - user (authenticated user or role from nebulakit config)

- Node Executions 

  - node_id 
  - execution.{any execution field above} (for example: execution.domain) 
  - phase (you must use the upper-cased string name, for example: ``QUEUED``) 
     - Phases are defined in :std:ref:`nodeexecution.phase <ref_nebulaidl.core.NodeExecution.Phase>`.
  - started_at  
  - node_execution_created_at 
  - node_execution_updated_at 
  - duration (in seconds)
  
- Task Executions 

  - retry_attempt 
  - task.{any task field above} (for example: task.version) 
  - execution.{any execution field above} (for example: execution.domain) 
  - node_execution.{any node execution field above} (for example: node_execution.phase) 
  - phase (you must use the upper-cased string name, for example: ``SUCCEEDED``)  
     - Phases are defined in :std:ref:`taskexecution.phase <ref_nebulaidl.core.TaskExecution.Phase>`.
  - started_at  
  - task_execution_created_at 
  - task_execution_updated_at 
  - duration (in seconds) 

Putting It All Together 
----------------------- 

If you wish to query specific executions that were launched using a specific launch plan for a workflow with specific attributes, use:

::  

   gte(duration, 100)+value_in(phase,RUNNING;SUCCEEDED;FAILED)+eq(lauch_plan.project, foo)  
   +eq(launch_plan.domain, bar)+eq(launch_plan.name, baz) 
   +eq(launch_plan.version, 1234) 
   +lte(workflow.created_at,2018-11-29T17:34:05.000000000Z07:00)  
    
    

Adding sorting to requests  
++++++++++++++++++++++++++  

Only a subset of fields are supported for sorting list queries. The explicit list is shown below: 

- ListTasks 

  - project 
  - domain  
  - name  
  - version 
  - created_at
  
- ListTaskIds 

  - project 
  - domain  
  
- ListWorkflows 

  - project 
  - domain  
  - name  
  - version 
  - created_at  
  
- ListWorkflowIds 

  - project 
  - domain  
  
- ListLaunchPlans 

  - project 
  - domain  
  - name  
  - version 
  - created_at  
  - updated_at  
  - state (you must use the integer enum, for example: 1) 
     - States are defined in :std:ref:`launchplanstate <ref_nebulaidl.admin.LaunchPlanState>`.
     
- ListWorkflowIds 

  - project 
  - domain  
  
- ListExecutions  

  - project 
  - domain  
  - name  
  - phase (you must use the upper-cased string name, for example: ``RUNNING``)  
     - Phases are defined in :std:ref:`workflowexecution.phase <ref_nebulaidl.core.WorkflowExecution.Phase>`.
  - execution_created_at  
  - execution_updated_at  
  - duration (in seconds) 
  - mode (you must use the integer enum, for example: 1)  
     - Modes are defined :std:ref:`execution.proto <ref_nebulaidl.admin.ExecutionMetadata.ExecutionMode>`.
     
- ListNodeExecutions  

  - node_id 
  - retry_attempt 
  - phase (you must use the upper-cased string name, for example: ``QUEUED``) 
     - Phases are defined in :std:ref:`nodeexecution.phase <ref_nebulaidl.core.NodeExecution.Phase>`.
  - started_at  
  - node_execution_created_at 
  - node_execution_updated_at 
  - duration (in seconds) 
  
- ListTaskExecutions  

  - retry_attempt 
  - phase (you must use the upper-cased string name, for example: ``SUCCEEDED``)  
     - Phases are defined in :std:ref:`taskexecution.phase <ref_nebulaidl.core.TaskExecution.Phase>`.
  - started_at  
  - task_execution_created_at 
  - task_execution_updated_at 
  - duration (in seconds) 

Sorting syntax  
--------------  

Adding sorting to a request requires specifying the ``key``. For example: The attribute you wish to sort on. Sorting can also optionally specify the direction (one of ``ASCENDING`` or ``DESCENDING``) where ``DESCENDING`` is the default. 

Example sorting HTTP parameter: 

::  

   sort_by.key=created_at&sort_by.direction=DESCENDING  
    
Alternatively, since ``DESCENDING`` is the default sorting direction, the above could be written as 

::  

   sort_by.key=created_at
