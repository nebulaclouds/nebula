.. _divedeep:

#############
Concepts
#############

.. panels::
    :header: text-center
    :column: col-lg-12 p-2

    .. link-button:: divedeep-tasks
       :type: ref
       :text: Tasks
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    A **Task** is any independent unit of processing. Tasks can be pure functions or functions with side-effects.
    Each definition of a task also has associated configurations and requirements specifications.

    ---

    .. link-button:: divedeep-workflows
       :type: ref
       :text: Workflows
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    **Workflows** are programs that are guaranteed to eventually reach a terminal state and are represented as
    Directed Acyclic Graphs (DAGs) expressed in protobuf.

    ---

    .. link-button:: divedeep-nodes
       :type: ref
       :text: Nodes
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    A **Node** is an encapsulation of an instance of a Task. Nodes represent the unit of work, where multiple Nodes are
    interconnected via workflows.

    ---

    .. link-button:: divedeep-launchplans
       :type: ref
       :text: Launch Plans
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    **Launch Plans** provide a mechanism to specialize input parameters for workflows associated with different schedules.

    ---

    .. link-button:: concepts-schedules
       :type: ref
       :text: Scheduling Launch Plans
       :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    **Scheduling** is critical to data and ML jobs. Nebula provides a native Cron-style scheduler.

    ---

    .. link-button:: divedeep-registration
        :type: ref
        :text: Registration
        :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    **Registration** is the process of uploading a workflow and its task definitions to the NebulaAdmin service.
    Registration creates an inventory of available tasks, workflows and launch plans, declared per project and domain.

    ---

    .. link-button:: divedeep-executions
        :type: ref
        :text: Executions
        :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    **Executions** are instances of workflows, nodes or tasks created in the system as a result of a user-requested
    execution or a scheduled execution.

    ---

    .. link-button:: divedeep-state-machine
        :type: ref
        :text: State Machine for an Execution
        :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    The various states an **Execution** passes through.

    ---

    .. link-button:: divedeep-execution-timeline
        :type: ref
        :text: Life of an Execution
        :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    How an **Execution** progresses through the NebulaPropeller execution engine and the timeline.

    ---

    .. link-button:: divedeep-data-management
        :type: ref
        :text: How Nebula Manages Data
        :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    What is **metadata**? How are large amounts of **raw data** handled? How does data flow between tasks?

    ---

    .. link-button:: ui
        :type: ref
        :text: Nebula UI Walkthrough
        :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    A quick overview of the **NebulaConsole**.

    ---

    .. link-button:: divedeep-catalog
        :type: ref
        :text: Platform-wide Memoization/Caching
        :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    A deeper dive into **memoization** and the mechanics of memoization in Nebula.

    ---

    .. link-button:: divedeep-versioning
         :type: ref
         :text: Workflow & Task Versioning
         :classes: btn-block stretched-link
    ^^^^^^^^^^^^
    A deeper dive into one of Nebula's most important features: versioning of workflows and tasks.


The diagram below shows how inputs flow through tasks and workflows to produce outputs.

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/concepts/basics/nebula_wf_tasks_high_level.png


.. toctree::
    :maxdepth: 1
    :name: Core Concepts
    :hidden:

    tasks
    workflows
    nodes
    launchplans
    schedules
    registration
    executions
    state_machine
    execution_timeline
    data_management
    nebula_console
    catalog
    versioning
    workflow_lifecycle
