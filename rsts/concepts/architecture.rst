.. _divedeep-architecture-overview:

######################
Component Architecture
######################

.. tags:: Advanced, Glossary, Design

This document aims to demystify how Nebula's major components ``Nebulaidl``, ``Nebulakit``, ``Nebulactl``, ``NebulaConsole``, ``NebulaAdmin``, ``NebulaPropeller``, and ``NebulaPlugins`` fit together at a high level.

NebulaIDL
========

In Nebula, entities like "Workflows", "Tasks", "Launch Plans", and "Schedules" are recognized by multiple system components. For components to communicate effectively, they need a shared understanding about the structure of these entities.

Nebulaidl (Interface Definition Language) is where shared Nebula entities are defined. It also defines the RPC service definition for the :std:ref:`core Nebula API <protos/docs/service/service:nebulaidl/service/admin.proto>`.

Nebulaidl uses the `protobuf <https://developers.google.com/protocol-buffers/>`_ schema to describe entities. Clients are generated for Python, Golang, and JavaScript and imported by Nebula components.


Planes
======

Nebula components are separated into 3 logical planes. The planes are summarized and explained in detail below. The goal is that these planes can be replaced by alternate implementations.

+-------------------+---------------------------------------------------------------------------------------------------------------+
| **User Plane**    | The User Plane consists of all user tools that assist in interacting with the core Nebula API.                 |
|                   | These tools include the NebulaConsole, Nebulakit, and Nebulactl.                                                 |
+-------------------+---------------------------------------------------------------------------------------------------------------+
| **Control Plane** | The Control Plane implements the core Nebula API.                                                              |
|                   | It serves all client requests coming from the User Plane.                                                     |
|                   | It stores information such as current and past running workflows, and provides that information upon request. |
|                   | It also accepts requests to execute workflows, but offloads the work to the Data Plane.                       |
+-------------------+---------------------------------------------------------------------------------------------------------------+
| **Data Plane**    | The sole responsibility of the the Data Plane is to fulfill workflows.                                        |
|                   | It accepts workflow requests from the Control Plane and guides the workflow to completion,                    |
|                   | launching tasks on a cluster of machines as necessary based on the workflow graph.                            |
|                   | It sends status events back to the control plane so the information can be stored and surfaced to end-users.  |
+-------------------+---------------------------------------------------------------------------------------------------------------+

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/concepts/architecture/nebula-logical-architecture.png

User Plane
----------

In Nebula, workflows are represented as a Directed Acyclic Graph (DAG) of tasks. While this representation is logical for services, managing workflow DAGs in this format is a tedious exercise for humans. The Nebula User Plane provides tools to create, manage, and visualize workflows in a format that is easily digestible to the users.

These tools include: 

Nebulakit
  Nebulakit is an SDK that helps users design new workflows using the Python programming language. It can parse the Python code, compile it into a valid Workflow DAG, and submit it to Nebula for execution.

NebulaConsole
  NebulaConsole provides the Web interface for Nebula. Users and administrators can use the console to view workflows, launch plans, schedules, tasks, and individual task executions. The console provides tools to visualize workflows, and surfaces relevant logs for debugging failed tasks.

Nebulactl
  Nebulactl provides interactive access to Nebula to launch and access workflows via terminal.


Control Plane
-------------

The Control Plane supports the core REST/gRPC API defined in Nebulaidl. User Plane tools like NebulaConsole and Nebulakit contact the control plane on behalf of users to store and retrieve information.

Currently, the entire control plane is handled by a single service called **NebulaAdmin**.

NebulaAdmin is stateless. It processes requests to create entities like tasks, workflows, and schedules by persisting data in a relational database.

While NebulaAdmin serves the Workflow Execution API, it does not itself execute workflows. To launch workflow executions, NebulaAdmin sends the workflow DAG to the DataPlane. For added scalability and fault-tolerance, NebulaAdmin can be configured to load-balance workflows across multiple isolated data-plane clusters.


Data Plane
----------

The Data Plane is the engine that accepts DAGs, and fulfills workflow executions by launching tasks in the order defined by the graph. Requests to the Data Plane generally come via the control plane, and not from end-users.

In order to support compute-intensive workflows at massive scale, the Data Plane needs to launch containers on a cluster of machines. The current implementation leverages `Kubernetes <https://kubernetes.io/>`_ for cluster management.

Unlike the user-facing Control Plane, the Data Plane does not expose a traditional REST/gRPC API. To launch an execution in the Data Plane, you create a “nebulaworkflow” resource in Kubernetes.
A “nebulaworkflow” is a Kubernetes `Custom Resource <https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/>`_ (CRD) created by our team. This custom resource represents the Nebula workflow DAG.

The core state machine that processes nebulaworkflows is the worker known as **NebulaPropeller**.

NebulaPropeller leverages the Kubernetes `operator pattern <https://kubernetes.io/docs/concepts/extend-kubernetes/operator/>`_. It polls the Kubernetes API, looking for newly created nebulaworkflow resources. NebulaPropeller understands the workflow DAG, and launches the appropriate Kubernetes pods as needed to complete tasks. It periodically checks for completed tasks, launching downstream tasks until the workflow is complete.

**Plugins**

Each task in a nebulaworkflow DAG has a specified **type**. The logic for fulfilling a task is determined by its task type.
In the basic case, NebulaPropeller launches a single Kubernetes pod to fulfill a task.
Complex task types require workloads to be distributed across hundreds of pods.

The type-specific task logic is separated into isolated code modules known as **plugins**.
Each task type has an associated plugin that is responsible for handling tasks of its type.
For each task in a workflow, NebulaPropeller activates the appropriate plugin based on the task type in order to fulfill the task.

The Nebula team has pre-built plugins for Hive, Spark, AWS Batch, and :ref:`more <integrations>`.
To support new use-cases, developers can create their own plugins and bundle them in their NebulaPropeller deployment.

Component Code Architecture
===========================

.. panels::
  :container: container-lg pb-4
  :column: col-lg-12 p-2
  :body: text-center

  .. link-button:: nebulapropeller-architecture
     :type: ref
     :text: NebulaPropeller
     :classes: btn-block stretched-link

  ---

  .. link-button:: native-scheduler-architecture
     :type: ref
     :text: Nebula Native Scheduler
     :classes: btn-block stretched-link

Component Code References
=========================

.. panels::
  :container: container-lg pb-4
  :column: col-lg-12 p-2
  :body: text-center

  .. link-button:: https://pkg.go.dev/mod/github.com/nebulaclouds/nebulaadmin
     :type: url
     :text: NebulaAdmin
     :classes: btn-block stretched-link

  ---

  .. link-button:: https://pkg.go.dev/mod/github.com/nebulaclouds/nebulapropeller
     :type: url
     :text: NebulaPropeller
     :classes: btn-block stretched-link

  ---

  .. link-button:: https://pkg.go.dev/mod/github.com/nebulaclouds/datacatalog
     :type: url
     :text: DataCatalog
     :classes: btn-block stretched-link

  ---

  .. link-button:: https://pkg.go.dev/mod/github.com/nebulaclouds/nebulaplugins
     :type: url
     :text: NebulaPlugins
     :classes: btn-block stretched-link

  ---

  .. link-button:: https://pkg.go.dev/github.com/nebulaclouds/nebulaadmin/scheduler
     :type: url
     :text: Nebula Native Scheduler
     :classes: btn-block stretched-link


.. toctree::
    :maxdepth: 1
    :name: component code architecture
    :hidden:

    component_architecture/nebulapropeller_architecture
    component_architecture/native_scheduler_architecture
