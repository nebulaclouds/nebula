.. _nebulapropeller-architecture:

###########################
NebulaPropeller Architecture
###########################

.. tags:: Advanced, Design

.. note::
   In the frame of this document, we use the term “workflow” to describe the single execution of a workflow definition.

Introduction
============

A Nebula :ref:`workflow <divedeep-workflows>` is represented as a Directed Acyclic Graph (DAG) of interconnected Nodes. Nebula supports a robust collection of Node types to ensure diverse functionality.

- ``TaskNodes`` support a plugin system to externally add system integrations.
- Control flow can be altered during runtime using ``BranchNodes``, which prune downstream evaluation paths based on input. 
- ``DynamicNodes`` add nodes to the DAG.
- ``WorkflowNodes`` allow embedding workflows within each other.

NebulaPropeller is responsible for scheduling and tracking execution of Nebula workflows. It is implemented using a K8s controller and adheres to the established K8s design principles. In this scheme, resources are periodically evaluated and the goal is to transition from the observed state to a requested state.

In our case, workflows are the resources and they are iteratively evaluated to transition from the current state to success. During each loop, the current workflow state is established as the phase of workflow nodes and subsequent tasks, and NebulaPropeller performs operations to transition this state to success. The operations may include scheduling (or rescheduling) node executions, evaluating dynamic or branch nodes, etc. These design decisions ensure that NebulaPropeller can scale to manage a large number of concurrent workflows without performance degradation.

This document attempts to break down the NebulaPropeller architecture by tracking workflow life cycle through each internal component. Below is a high-level illustration of the NebulaPropeller architecture and a flow chart of each component's responsibilities during NebulaWorkflow execution.

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/concepts/architecture/nebulapropeller_architecture.png

Components
==========

NebulaWorkflow CRD / K8s Integration
-----------------------------------

Workflows in Nebula are maintained as Custom Resource Definitions (CRDs) in Kubernetes, which are stored in the backing etcd cluster. Each execution of a workflow definition results in the creation of a new NebulaWorkflow CR (Custom Resource) which maintains a state for the entirety of processing. CRDs provide variable definitions to describe both resource specifications (spec) and status' (status). The NebulaWorkflow CRD uses the spec subsection to detail the workflow DAG, embodying node dependencies, etc. The status subsection tracks workflow metadata including overall workflow status, node/task phases, status/phase transition timestamps, etc.

K8s exposes a powerful controller/operator API that enables entities to track creation/updates over a specific resource type. NebulaPropeller uses this API to track NebulaWorkflows, meaning every time an instance of the NebulaWorkflow CR is created/updated, the NebulaPropeller instance is notified. NebulaAdmin is the common entry point, where initialization of NebulaWorkflow CRs may be triggered by user workflow definition executions, automatic relaunches, or periodically scheduled workflow definition executions. However, it is conceivable to manually create NebulaWorkflow CRs, but this will have limited visibility and usability.

WorkQueue/WorkerPool
----------------------

NebulaPropeller supports concurrent execution of multiple, unique workflows using a WorkQueue and WorkerPool.

The WorkQueue is a FIFO queue storing workflow ID strings that require a lookup to retrieve the NebulaWorkflow CR to ensure up-to-date status. A workflow may be added to the queue in a variety of circumstances:

#. A new NebulaWorkflow CR is created or an existing instance is updated
#. The K8s Informer resyncs the NebulaWorkflow periodically (necessary to detect workflow timeouts and ensure liveness)
#. A NebulaPropeller worker experiences an error during a processing loop
#. The WorkflowExecutor observes a completed downstream node
#. A NodeHandler observes state change and explicitly enqueues its owner (For example, K8s pod informer observes completion of a task)

The WorkerPool is implemented as a collection of goroutines, one for each worker. Using this lightweight construct, NebulaPropeller can scale to 1000s of workers on a single CPU. Workers continually poll the WorkQueue for workflows. On success, the workflow is executed (passed to WorkflowExecutor).

WorkflowExecutor
----------------

The WorkflowExecutor is responsible for handling high-level workflow operations. This includes maintaining the workflow phase (for example: running, failing, succeeded, etc.) according to the underlying node phases and administering pending cleanup operations. For example, aborting existing node evaluations during workflow failures or removing NebulaWorkflow CRD finalizers on completion to ensure the CR is deleted. Additionally, at the conclusion of each evaluation round, the WorkflowExecutor updates the NebulaWorkflow CR with updated metadata fields to track the status between evaluation iterations.

NodeExecutor
------------

The NodeExecutor is executed on a single node, beginning with the workflow's start node. It traverses the workflow using a visitor pattern with a modified depth-first search (DFS), evaluating each node along the path. A few examples of node evaluation based on phase: successful nodes are skipped, unevaluated nodes are queued for processing, and failed nodes may be reattempted up to a configurable threshold. There are many configurable parameters to tune evaluation criteria including max parallelism which restricts the number of nodes which may be scheduled concurrently. Additionally, nodes may be retried to ensure recoverability on failure.  

The NodeExecutor is also responsible for linking data readers/writers to facilitate data transfer between node executions. The data transfer process occurs automatically within Nebula, using efficient K8s events rather than a polling listener pattern which incurs more overhead. Relatively small amounts of data may be passed between nodes inline, but it is more common to pass data URLs to backing storage. A component of this is writing to and checking the data cache, which facilitates the reuse of previously completed evaluations.

NodeHandlers
------------

NebulaPropeller includes a robust collection of NodeHandlers to support diverse evaluation of the workflow DAG:

* **TaskHandler (Plugins)**: These are responsible for executing plugin specific tasks. This may include contacting NebulaAdmin to schedule K8s pod to perform work, calling a web API to begin/track evaluation, and much more. The plugin paradigm exposes an extensible interface for adding functionality to Nebula workflows.
* **DynamicHandler**: Nebula workflow CRs are initialized using a DAG compiled during the registration process. The numerous benefits of this approach are beyond the scope of this document. However, there are situations where the complete DAG is unknown at compile time. For example, when executing a task on each value of an input list. Using Dynamic nodes, a new DAG subgraph may be dynamically compiled during runtime and linked to the existing NebulaWorkflow CR.
* **WorkflowHandler**: This handler allows embedding workflows within another workflow definition. The API exposes this functionality using either (1) an inline execution, where the workflow function is invoked directly resulting in a single NebulaWorkflow CR with an appended sub-workflow, or (2) a launch plan, which uses a TODO to create a separate sub-NebulaWorkflow CR whose execution state is linked to the parent NebulaWorkflow CR.
* **BranchHandler**: The branch handler allows the DAG to follow a specific control path based on input (or computed) values.
* **Start / End Handlers**: These are dummy handlers which process input and output data and in turn transition start and end nodes to success.

NebulaAdmin Events
-----------------

It should be noted that the WorkflowExecutor, NodeExecutor, and TaskHandlers send events to NebulaAdmin, enabling it to track workflows in near real-time.
