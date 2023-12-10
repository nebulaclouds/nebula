.. _divedeep-registration:

############
Registration
############

.. tags:: Basic, Glossary, Design

During registration, Nebula validates the workflow structure and saves the workflow. The registration process also updates the workflow graph.

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/concepts/executions/nebula_wf_registration_overview.svg?sanitize=true

Typical Flow 
-------------
The following steps elaborate on the specifics of the registration process:

* Define the tasks using the :py:mod:`Nebulakit <nebulakit:nebulakit>` Task Definition language.
* Define a workflow using the :py:mod:`Nebulakit <nebulakit:nebulakit>` Workflow definition language.
* Use `nebulactl register CLI <https://docs.nebula.org/projects/nebulactl/en/latest/gen/nebulactl_register_files.html>`__ to compile the tasks into their serialized representation as described in :std:ref:`Nebula Specification language <nebulaidl:nebulaidltoc>`. During this, the task representation is bound to a container that constitutes the code for the task. This associated entity is registered with NebulaAdmin using the registerTask API.
* Use nebulactl register CLI to compile the workflow into their serialized representation as described in :std:ref:`Nebula Specification language <nebulaidl:nebulaidltoc>`. The referenced tasks are replaced by their NebulaAdmin registered Identifiers, obtained in the previous step. The associated entity is registered with NebulaAdmin using the registerWorkflow API.
* Launch an execution using the NebulaAdmin launch execution API, which requires the necessary inputs provided. This is automatically done if the user uses nebulactl to launch the execution.
* Use the NebulaAdmin read APIs to get details of the execution, monitor it to completion, or retrieve a historical execution.
* **OR** use the NebulaConsole to visualize the execution in real time as it progresses or visualize any historical execution. The console makes it easy to view debugging information for the execution.
* Set specific rules such as *notification* on **failure** or **success** or publish all events in the execution to a pub-sub system.
* Query the datastore to get a summary of all the executions and the compute resources consumed.

.. note::
    Workflows and tasks are purely specifications and can be provided using tools like ``YAML``, ``JSON``, ``protobuf binary`` or any other programming language, and hence registration is possible using other tools. Contributions welcome!

Registration in the Backend
---------------------------

When NebulaAdmin receives a workflow registration request, it uses the workflow compiler to compile and validate the workflow. It also fetches all the referenced tasks and creates a complete workflow closure, which is stored in the metastore. If the workflow compilation fails, the compiler returns an error to the client.
