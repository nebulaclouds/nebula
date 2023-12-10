.. _contribute_Nebula:

#####################
Contributing to Nebula
#####################

.. tags:: Contribute, Basic

Thank you for taking the time to contribute to Nebula!
Please read our `Code of Conduct <https://lfprojects.org/policies/code-of-conduct/>`__ before contributing to Nebula.

Here are some guidelines for you to follow, which will make your first and follow-up contributions easier.

TL;DR: Find the repo-specific contribution guidelines in the `Component Reference <#component-reference>`__ section.

💻 Becoming a contributor
=========================

An issue tagged with `good first issue <https://github.com/nebulaclouds/nebula/labels/good%20first%20issue>`__ is the best place to start for first-time contributors.

**Appetizer for every repo: Fork and clone the concerned repository. Create a new branch on your fork and make the required changes. Create a pull request once your work is ready for review.** 

.. note::
    To open a pull request, refer to `GitHub's guide <https://guides.github.com/activities/forking/>`__ for detailed instructions. 

Example PR for your reference: `GitHub PR <https://github.com/nebulaclouds/nebulapropeller/pull/242>`__.
A couple of checks are introduced to help maintain the robustness of the project. 

#. To get through DCO, sign off on every commit (`Reference <https://github.com/src-d/guide/blob/master/developer-community/fix-DCO.md>`__) 
#. To improve code coverage, write unit tests to test your code
#. Make sure all the tests pass. If you face any issues, please let us know

On a side note, format your Go code with ``golangci-lint`` followed by ``goimports`` (use ``make lint`` and ``make goimports``), and Python code with ``black`` and ``isort`` (use ``make fmt``). 
If make targets are not available, you can manually format the code.
Refer to `Effective Go <https://golang.org/doc/effective_go>`__, `Black <https://github.com/psf/black>`__, and `Isort <https://github.com/PyCQA/isort>`__ for full coding standards.

As you become more involved with the project, you may be able to be added as a contributor to the repos you're working on,
but there is a medium term effort to move all development to forks.

📃 Documentation
================

Nebula uses Sphinx for documentation. ``protoc-gen-doc`` is used to generate the documentation from ``.proto`` files.

Sphinx spans multiple repositories under `nebulaclouds <https://github.com/nebulaclouds>`__. It uses reStructured Text (rst) files to store the documentation content.
For API- and code-related content, it extracts docstrings from the code files. 

To get started, refer to the `reStructuredText reference <https://www.sphinx-doc.org/en/master/usage/restructuredtext/index.html#rst-index>`__. 

For minor edits that don't require a local setup, you can edit the GitHub page in the documentation to propose improvements.

Intersphinx
***********

`Intersphinx <https://www.sphinx-doc.org/en/master/usage/extensions/intersphinx.html>`__ can generate automatic links to the documentation of objects in other projects.

To establish a reference to any other documentation from Nebula or within it, use Intersphinx.

To do so, create an ``intersphinx_mapping`` in the ``conf.py`` file which should be present in the respective ``docs`` repository. 
For example, ``rsts`` is the docs repository for the ``nebula`` repo.

For example:

.. code-block:: python

    intersphinx_mapping = {
        "python": ("https://docs.python.org/3", None),
        "nebulakit": ("https://nebula.readthedocs.io/projects/nebulakit/en/master/", None),
    }

The key refers to the name used to refer to the file (while referencing the documentation), and the URL denotes the precise location. 

Here is an example using ``:std:doc``:
 
* Direct reference

  .. code-block:: text

      Task: :std:doc:`/api/nebulakit/generated/nebulakit.task`

  Output:

  Task: :std:doc:`/api/nebulakit/generated/nebulakit.task`

* Custom name

  .. code-block:: text

      :std:doc:`Using custom words </api/nebulakit/generated/nebulakit.task>`

  Output:

  :std:doc:`Using custom words </api/nebulakit/generated/nebulakit.task>`

|

You can cross-reference multiple Python objects. Check out this `section <https://www.sphinx-doc.org/en/master/usage/restructuredtext/domains.html#cross-referencing-python-objects>`__ to learn more. 

|

For instance, `task` decorator in nebulakit uses the ``func`` role.

.. code-block:: text

    Link to nebulakit code :py:func:`nebulakit:nebulakit.task`

Output:

Link to nebulakit code :py:func:`nebulakit:nebulakit.task`

|

Here are a couple more examples.

.. code-block:: text

    :py:mod:`Module <python:typing>`
    :py:class:`Class <python:typing.Type>`
    :py:data:`Data <python:typing.Callable>`
    :py:func:`Function <python:typing.cast>`
    :py:meth:`Method <python:pprint.PrettyPrinter.format>`

Output:

:py:mod:`Module <python:typing>`

:py:class:`Class <python:typing.Type>`

:py:data:`Data <python:typing.Callable>`

:py:func:`Function <python:typing.cast>`

:py:meth:`Method <python:pprint.PrettyPrinter.format>`

🧱 Component reference
======================

To understand how the below components interact with each other, refer to :ref:`Understand the lifecycle of a workflow <workflow-lifecycle>`.

.. figure:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/contribution_guide/dependency_graph.png
    :alt: Dependency graph between various nebulaclouds repos
    :align: center
    :figclass: align-center

    The dependency graph between various nebula repos

``nebula``
*********

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebula>`__
    * - **Purpose**: Deployment, Documentation, and Issues 
    * - **Languages**: Kustomize & RST
  
.. note::
    For the ``nebula`` repo, run the following command in the repo's root to generate documentation locally.

    .. code-block:: console

        make -C rsts html

``nebulaidl``
************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulaidl>`__
    * - **Purpose**: Nebula workflow specification is in `protocol buffers <https://developers.google.com/protocol-buffers>`__ which forms the core of Nebula
    * - **Language**: Protobuf
    * - **Guidelines**: Refer to the `README <https://github.com/nebulaclouds/nebulaidl#generate-code-from-protobuf>`__
 
``nebulapropeller``
******************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulapropeller>`__ | `Code Reference <https://pkg.go.dev/mod/github.com/nebulaclouds/nebulapropeller>`__
    * - **Purpose**: Kubernetes-native operator
    * - **Language**: Go
    * - **Guidelines:**

        * Check for Makefile in the root repo
        * Run the following commands:
           * ``make generate``
           * ``make test_unit``
           * ``make link``
        * To compile, run ``make compile``

``nebulaadmin``
**************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulaadmin>`__ | `Code Reference <https://pkg.go.dev/mod/github.com/nebulaclouds/nebulaadmin>`__
    * - **Purpose**: Control Plane
    * - **Language**: Go
    * - **Guidelines**:

        * Check for Makefile in the root repo
        * If the service code has to be tested, run it locally:
            * ``make compile``
            * ``make server``
        * To seed data locally:
            * ``make compile``
            * ``make seed_projects``
            * ``make migrate``
        * To run integration tests locally:
            * ``make integration``
            * (or to run in containerized dockernetes): ``make k8s_integration``

``nebulakit``
************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulakit>`__
    * - **Purpose**: Python SDK & Tools
    * - **Language**: Python
    * - **Guidelines**: Refer to the `Nebulakit Contribution Guide <https://docs.nebula.org/projects/nebulakit/en/latest/contributing.html>`__

``nebulaconsole``
****************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulaconsole>`__
    * - **Purpose**: Admin Console
    * - **Language**: Typescript
    * - **Guidelines**: Refer to the `README <https://github.com/nebulaclouds/nebulaconsole/blob/master/README.md>`__

``datacatalog``
***************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/datacatalog>`__ | `Code Reference <https://pkg.go.dev/mod/github.com/nebulaclouds/datacatalog>`__
    * - **Purpose**: Manage Input & Output Artifacts
    * - **Language**: Go

``nebulaplugins``
****************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulaplugins>`__ | `Code Reference <https://pkg.go.dev/mod/github.com/nebulaclouds/nebulaplugins>`__
    * - **Purpose**: Nebula Plugins
    * - **Language**: Go
    * - **Guidelines**:

        * Check for Makefile in the root repo
        * Run the following commands:
            * ``make generate``
            * ``make test_unit``
            * ``make link``

``nebulastdlib``
***************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulastdlib>`__
    * - **Purpose**: Standard Library for Shared Components
    * - **Language**: Go

``nebulasnacks``
***************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulasnacks>`__
    * - **Purpose**: Examples, Tips, and Tricks to use Nebulakit SDKs
    * - **Language**: Python (In the future, Java examples will be added)
    * - **Guidelines**: Refer to the `Nebulasnacks Contribution Guide <https://docs.nebula.org/projects/cookbook/en/latest/contribute.html>`__

``nebulactl``
************

.. list-table::

    * - `Repo <https://github.com/nebulaclouds/nebulactl>`__
    * - **Purpose**: A standalone Nebula CLI
    * - **Language**: Go
    * - **Guidelines**: Refer to the `NebulaCTL Contribution Guide <https://docs.nebula.org/projects/nebulactl/en/stable/contribute.html>`__


🔮 Development Environment Setup Guide
======================================

This guide provides a step-by-step approach to setting up a local development environment for 
`nebulaidl <https://github.com/nebulaclouds/nebulaidl>`_, `nebulaadmin <https://github.com/nebulaclouds/nebulaadmin>`_,
`nebulaplugins <https://github.com/nebulaclouds/nebulaplugins>`_, `nebulapropeller <https://github.com/nebulaclouds/nebulapropeller>`_,
`nebulakit <https://github.com/nebulaclouds/nebulakit>`_ , `nebulaconsole <https://github.com/nebulaclouds/nebulaconsole>`_,
`datacatalog <https://github.com/nebulaclouds/datacatalog>`_, and `nebulastdlib <https://github.com/nebulaclouds/nebulastdlib>`_.

The video below is a tutorial on how to set up a local development environment for Nebula.

..  youtube:: V-KlVQmQAjE 

Requirements
************

This guide has been tested and used on AWS EC2 with an Ubuntu 22.04
image. The following tools are required:

- `Docker <https://docs.docker.com/install/>`__
- `Kubectl <https://kubernetes.io/docs/tasks/tools/install-kubectl/>`__
- `Go <https://golang.org/doc/install>`__

Content
*******

-  `How to setup dev environment for nebulaidl, nebulaadmin, nebulaplugins,
   nebulapropeller, datacatalog and nebulastdlib? <#how-to-setup-dev-environment-for-nebulaidl-nebulaadmin-nebulaplugins-nebulapropeller-datacatalog-and-nebulastdlib>`__

-  `How to setup dev environment for
   nebulakit? <#how-to-setup-dev-environment-for-nebulakit>`__

-  `How to setup dev environment for
   nebulaconsole? <#how-to-setup-dev-environment-for-nebulaconsole>`__

-  `How to access Nebula UI, minio, postgres, k3s, and endpoints?
   <#how-to-access-nebula-ui-minio-postgres-k3s-and-endpoints>`__

How to setup dev environment for nebulaidl, nebulaadmin, nebulaplugins, nebulapropeller, datacatalog and nebulastdlib?
******************************************************************************************************************************

**1. Install nebulactl**


`Nebulactl <https://github.com/nebulaclouds/nebulactl>`__ is a portable and lightweight command-line interface to work with Nebula.

.. code:: shell

   # Step1: Install the latest version of nebulactl
   curl -sL https://ctl.nebula.org/install | bash
   # nebulaclouds/nebulactl info checking GitHub for latest tag
   # nebulaclouds/nebulactl info found version: 0.6.39 for v0.6.39/Linux/x86_64
   # nebulaclouds/nebulactl info installed ./bin/nebulactl

   # Step2: Export nebulactl path based on the previous log "nebulaclouds/nebulactl info installed ./bin/nebulactl"
   export PATH=$PATH:/home/ubuntu/bin # replace with your path

**2. Build a k3s cluster that runs minio and postgres Pods.**


| `Minio <https://min.io/>`__ is an S3-compatible object store that will be used later to store task output, input, etc.
| `Postgres <https://www.postgresql.org/>`__ is an open-source object-relational database that will later be used by nebulaadmin/dataCatalog to
  store all Nebula information.

.. code:: shell

   # Step1: Start k3s cluster, create Pods for postgres and minio. Note: We cannot access Nebula UI yet! but we can access the minio console now.
   nebulactl demo start --dev
   # 👨‍💻 Nebula is ready! Nebula UI is available at http://localhost:30080/console 🚀 🚀 🎉
   # ❇️ Run the following command to export demo environment variables for accessing nebulactl
   #         export NEBULACTL_CONFIG=/home/ubuntu/.nebula/config-sandbox.yaml
   # 🐋 Nebula sandbox ships with a Docker registry. Tag and push custom workflow images to localhost:30000
   # 📂 The Minio API is hosted on localhost:30002. Use http://localhost:30080/minio/login for Minio console

   # Step2: Export NEBULACTL_CONFIG as the previous log indicated.
   NEBULACTL_CONFIG=/home/ubuntu/.nebula/config-sandbox.yaml

   # Step3: The kubeconfig will be automatically copied to the user's main kubeconfig (default is `/.kube/config`) with "nebula-sandbox" as the context name.
   # Check that we can access the K3s cluster. Verify that postgres and minio are running.
   kubectl get pod -n nebula
   # NAME                                                  READY   STATUS    RESTARTS   AGE
   # nebula-sandbox-docker-registry-85745c899d-dns8q        1/1     Running   0          5m
   # nebula-sandbox-kubernetes-dashboard-6757db879c-wl4wd   1/1     Running   0          5m
   # nebula-sandbox-proxy-d95874857-2wc5n                   1/1     Running   0          5m
   # nebula-sandbox-minio-645c8ddf7c-sp6cc                  1/1     Running   0          5m
   # nebula-sandbox-postgresql-0                            1/1     Running   0          5m


**3. Run all Nebula components (nebulaadmin, nebulapropeller, datacatalog, nebulaconsole, etc) in a single binary.**

The `Nebula repository <https://github.com/nebulaclouds/nebula>`__ includes Go code
that integrates all Nebula components into a single binary.

.. code:: shell

   # Step1: Clone nebula repo
   git clone https://github.com/nebulaclouds/nebula.git
   cd nebula

   # Step2: Build a single binary that bundles all the Nebula components.
   # The version of each component/library used to build the single binary are defined in `go.mod`.
   sudo apt-get -y install jq # You may need to install jq
   go mod tidy
   make compile

   # Step3: Edit the config file: ./nebula-single-binary-local.yaml.
   # Replace occurrences of $HOME with the actual path of your home directory.
   sedi=(-i)
   case "$(uname)" in
     Darwin*) sedi=(-i "")
   esac
   sed "${sedi[@]}" -e "s|\$HOME|${HOME}|g" nebula-single-binary-local.yaml

   # Step 4: Prepare a namespace template for the cluster resource controller.
   # The configuration file "nebula-single-binary-local.yaml" has an entry named cluster_resources.templatePath.
   # This entry needs to direct to a directory containing the templates for the cluster resource controller to use.
   # We will now create a simple template that allows the automatic creation of required namespaces for projects.
   # For example, with Nebula's default project "nebulasnacks", the controller will auto-create the following namespaces:
   # nebulasnacks-staging, nebulasnacks-development, and nebulasnacks-production.
   mkdir $HOME/.nebula/cluster-resource-templates/
   echo "apiVersion: v1
   kind: Namespace
   metadata:
     name: '{{ namespace }}'" > $HOME/.nebula/cluster-resource-templates/namespace.yaml

   # Step5: Running the single binary.
   # The POD_NAMESPACE environment variable is necessary for the webhook to function correctly. 
   # You may encounter an error due to `ERROR: duplicate key value violates unique constraint`. Running the command again will solve the problem.
   POD_NAMESPACE=nebula ./nebula start --config nebula-single-binary-local.yaml
   # All logs from nebulaadmin, nebulaplugins, nebulapropeller, etc. will appear in the terminal.


**4. Build single binary with your own code.**


The following instructions provide guidance on how to build single binary with your customized code under the ``nebulaadmin`` as an example.


- **Note** Although we'll use ``nebulaadmin`` as an example, these steps can be applied to other Nebula components or libraries as well.
  ``{nebulaadmin}`` below can be substituted with other Nebula components/libraries: ``nebulaidl``, ``nebulaplugins``, ``nebulapropeller``, ``datacatalog``, or ``nebulastdlib``.
- **Note** If you want to learn how nebula compiles those components and replace the repositories, you can study how ``go mod edit`` works.

.. code:: shell

   # Step1: Install Go. Nebula uses Go 1.19, so make sure to switch to Go 1.19.
   export PATH=$PATH:$(go env GOPATH)/bin
   go install golang.org/dl/go1.19@latest
   go1.19 download
   export GOROOT=$(go1.19 env GOROOT)
   export PATH="$GOROOT/bin:$PATH"

   # You may need to install goimports to fix lint errors.
   # Refer to https://pkg.go.dev/golang.org/x/tools/cmd/goimports
   go install golang.org/x/tools/cmd/goimports@latest
   export PATH=$(go env GOPATH)/bin:$PATH

   # Step2: Go to the {nebulaadmin} repository, modify the source code accordingly.
   cd nebula/nebulaadmin

   # Step3: Now, you can build the single binary. Go back to Nebula directory.
   go mod tidy
   make compile
   POD_NAMESPACE=nebula ./nebula start --config nebula-single-binary-local.yaml

**5. Test by running a hello world workflow.**


.. code:: shell

   # Step1: Install nebulakit
   pip install nebulakit && export PATH=$PATH:/home/ubuntu/.local/bin

   # Step2: Run a hello world example
   pynebula run --remote https://raw.githubusercontent.com/nebulaclouds/nebulasnacks/master/examples/basics/basics/hello_world.py  hello_world_wf
   # Go to http://localhost:30080/console/projects/nebulasnacks/domains/development/executions/fd63f88a55fed4bba846 to see execution in the console.
   # You can go to the [nebulasnacks repository](https://github.com/nebulaclouds/nebulasnacks) to see more useful examples.

**6. Tear down the k3s cluster after finishing developing.**


.. code:: shell

   nebulactl demo teardown
   # context removed for "nebula-sandbox".
   # 🧹 🧹 Sandbox cluster is removed successfully.
   # ❇️ Run the following command to unset sandbox environment variables for accessing nebulactl
   #        unset NEBULACTL_CONFIG

How to setup dev environment for nebulakit?
*******************************************

**1. Set up local Nebula Cluster.**


If you are also modifying the code for nebulaidl, nebulaadmin, nebulaplugins, nebulapropeller datacatalog, or nebulastdlib,
refer to the instructions in the  `previous section <#how-to-setup-dev-environment-for-nebulaidl-nebulaadmin-nebulaplugins-nebulapropeller-datacatalog-and-nebulastdlib>`__ to set up a local Nebula cluster.

If not, we can start backends with a single command.

.. code:: shell

   # Step1: Install the latest version of nebulactl, a portable and lightweight command-line interface to work with Nebula.
   curl -sL https://ctl.nebula.org/install | bash
   # nebulaclouds/nebulactl info checking GitHub for latest tag
   # nebulaclouds/nebulactl info found version: 0.6.39 for v0.6.39/Linux/x86_64
   # nebulaclouds/nebulactl info installed ./bin/nebulactl

   # Step2: Export nebulactl path based on the previous log "nebulaclouds/nebulactl info installed ./bin/nebulactl"
   export PATH=$PATH:/home/ubuntu/bin # replace with your path

   # Step3: Starts the Nebula demo cluster. This will setup a k3s cluster running minio, postgres Pods, and all Nebula components: nebulaadmin, nebulaplugins, nebulapropeller, etc.
   # See https://docs.nebula.org/projects/nebulactl/en/latest/gen/nebulactl_demo_start.html for more details.
   nebulactl demo start
   # 👨‍💻 Nebula is ready! Nebula UI is available at http://localhost:30080/console 🚀 🚀 🎉
   # ❇️ Run the following command to export demo environment variables for accessing nebulactl
   #         export NEBULACTL_CONFIG=/home/ubuntu/.nebula/config-sandbox.yaml
   # 🐋 Nebula sandbox ships with a Docker registry. Tag and push custom workflow images to localhost:30000
   # 📂 The Minio API is hosted on localhost:30002. Use http://localhost:30080/minio/login for Minio console

**2. Run workflow locally.**


.. code:: shell

   # Step1: Build a virtual environment for developing Nebulakit. This will allow your local changes to take effect when the same Python interpreter runs `import nebulakit`.
   git clone https://github.com/nebulaclouds/nebulakit.git # replace with your own repo
   cd nebulakit
   virtualenv ~/.virtualenvs/nebulakit
   source ~/.virtualenvs/nebulakit/bin/activate
   make setup
   pip install -e .
   
   # If you are also developing the plugins, consider the following:

   # Installing Specific Plugins:
   # If you wish to only use few plugins, you can install them individually.
   # Take [Nebulakit BigQuery Plugin](https://github.com/nebulaclouds/nebulakit/tree/master/plugins/nebulakit-bigquery#nebulakit-bigquery-plugin) for example:
   # You have to go to the bigquery plugin folder and install it.
   cd plugins/nebulakit-bigquery/
   pip install -e .
   # Now you can use the bigquery plugin, and the performance is fast.

   # (Optional) Installing All Plugins:
   # If you wish to install all available plugins, you can execute the command below.
   # However, it's not typically recommended because the current version of plugins does not support
   # lazy loading. This can lead to a slowdown in the performance of your Python engine.
   cd plugins
   pip install -e .
   # Now you can use all plugins, but the performance is slow.

   # Step2: Modify the source code for nebulakit, then run unit tests and lint.
   make lint
   make test

   # Step3: Run a hello world sample to test locally
   pynebula run https://raw.githubusercontent.com/nebulaclouds/nebulasnacks/master/examples/basics/basics/hello_world.py hello_world_wf
   # Running hello_world_wf() hello world

**3. Run workflow in sandbox.**


Before running your workflow in the sandbox, make sure you're able to successfully run it locally. 
To deploy the workflow in the sandbox, you'll need to build a Nebulakit image.
Create a Dockerfile in your Nebulakit directory with the minimum required configuration to run a task, as shown below.
If your task requires additional components, such as plugins, you may find it useful to refer to the construction of the `officail flitekit image <https://github.com/nebulaclouds/nebulakit/blob/master/Dockerfile>`__

.. code:: Dockerfile

   FROM python:3.9-slim-buster
   USER root
   WORKDIR /root
   ENV PYTHONPATH /root
   RUN apt-get update && apt-get install build-essential -y
   RUN apt-get install git -y
   # The following line is an example of how to install your modified plugins. In this case, it demonstrates how to install the 'deck' plugin.
   # RUN pip install -U git+https://github.com/Yicheng-Lu-llll/nebulakit.git@"demo#egg=nebulakitplugins-deck-standard&subdirectory=plugins/nebulakit-deck-standard" # replace with your own repo and branch
   RUN pip install -U git+https://github.com/Yicheng-Lu-llll/nebulakit.git@demo # replace with your own repo and branch
   ENV NEBULA_INTERNAL_IMAGE "localhost:30000/nebulakit:demo" # replace with your own image name and tag

The instructions below explain how to build the image, push the image to
the Nebula cluster, and finally submit the workflow.

.. code:: shell

   # Step1: Ensure you have pushed your changes to the remote repo
   # In the nebulakit folder
   git add . && git commit -s -m "develop" && git push

   # Step2: Build the image
   # In the nebulakit folder
   export NEBULA_INTERNAL_IMAGE="localhost:30000/nebulakit:demo" # replace with your own image name and tag
   docker build --no-cache -t  "${NEBULA_INTERNAL_IMAGE}" -f ./Dockerfile .

   # Step3: Push the image to the Nebula cluster
   docker push ${NEBULA_INTERNAL_IMAGE}

   # Step4: Submit a hello world workflow to the Nebula cluster
   cd nebulasnacks
   pynebula run --image ${NEBULA_INTERNAL_IMAGE} --remote https://raw.githubusercontent.com/nebulaclouds/nebulasnacks/master/examples/basics/basics/hello_world.py hello_world_wf
   # Go to http://localhost:30080/console/projects/nebulasnacks/domains/development/executions/f5c17e1b5640c4336bf8 to see execution in the console.

How to setup dev environment for nebulaconsole?
**********************************************

**1. Set up local Nebula cluster.**

Depending on your needs, refer to one of the following guides to setup up the Nebula cluster:

- If you do not need to change the backend code, refer to the section on `How to Set Up a Dev Environment for Nebulakit? <#how-to-setup-dev-environment-for-nebulakit>`__
- If you need to change the backend code, refer to the section on `How to setup dev environment for nebulaidl, nebulaadmin, nebulaplugins, nebulapropeller, datacatalog and nebulastdlib? <#how-to-setup-dev-environment-for-nebulaidl-nebulaadmin-nebulaplugins-nebulapropeller-datacatalog-and-nebulastdlib>`__


**2. Start nebulaconsole.**


.. code:: shell

   # Step1: Clone the repo and navigate to the Nebulaconsole folder
   git clone https://github.com/nebulaclouds/nebulaconsole.git
   cd nebulaconsole

   # Step2: Install Node.js 18. Refer to https://github.com/nodesource/distributions/blob/master/README.md#using-ubuntu-2.
   curl -fsSL https://deb.nodesource.com/setup_18.x | sudo -E bash - &&\
   sudo apt-get install -y nodejs

   # Step3: Install yarn. Refer to https://classic.yarnpkg.com/lang/en/docs/install/#debian-stable.
   curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
   echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
   sudo apt update && sudo apt install yarn

   # Step4: Add environment variables
   export BASE_URL=/console
   export ADMIN_API_URL=http://localhost:30080
   export DISABLE_AUTH=1
   export ADMIN_API_USE_SSL="http"

   # Step5: Generate SSL certificate
   # Note, since we will use HTTP, SSL is not required. However, missing an SSL certificate will cause an error when starting Nebulaconsole.
   make generate_ssl

   # Step6: Install node packages
   yarn install
   yarn build:types # It is fine if seeing error `Property 'at' does not exist on type 'string[]'`
   yarn run build:prod

   # Step7: Start nebulaconsole
   yarn start

**3. Install the Chrome plugin:** `Moesif Origin & CORS Changer <https://chrome.google.com/webstore/detail/moesif-origin-cors-change/digfbfaphojjndkpccljibejjbppifbc>`__.


We need to disable `CORS <https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS>`__ to load resources.

::

   1. Activate plugin (toggle to "on")
   2. Open 'Advanced Settings':
   3. set Access-Control-Allow-Credentials: true

**4. Go to** http://localhost:3000/console/.


How to access Nebula UI, minio, postgres, k3s, and endpoints?
*************************************************************************


This section presumes a local Nebula cluster is already setup. If it isn't, refer to either:

- `How to setup dev environment for nebulakit? <#how-to-setup-dev-environment-for-nebulakit>`__
- `How to setup dev environment for nebulaidl, nebulaadmin, nebulaplugins, nebulapropeller, datacatalog and nebulastdlib? <#how-to-setup-dev-environment-for-nebulaidl-nebulaadmin-nebulaplugins-nebulapropeller-datacatalog-and-nebulastdlib>`__


**1. Access the Nebula UI.**


`Nebula UI <https://docs.nebula.org/en/latest/concepts/nebula_console.html>`__ is a web-based user interface for Nebula that lets you interact with Nebula objects and build directed acyclic graphs (DAGs) for your workflows.

You can access it via http://localhost:30080/console.

**2. Access the minio console.**


Core Nebula components, such as admin, propeller, and datacatalog, as well as user runtime containers rely on an object store (in this case, minio) to hold files.
During development, you might need to examine files such as `input.pb/output.pb <https://docs.nebula.org/en/latest/concepts/data_management.html#serialization-time>`__, or `deck.html <https://docs.nebula.org/projects/cookbook/en/latest/auto_examples/basics/deck.html#nebula-decks>`__ stored in minio.

Access the minio console at: http://localhost:30080/minio/login. The default credentials are:

- Username: ``minio``
- Password: ``miniostorage``


**3. Access the postgres.**


NebulaAdmin and datacatalog use postgres to store persistent records, and you can interact with postgres on port ``30001``. Here is an example of using `psql` to connect:

.. code:: shell
    
    # Step1: Install the PostgreSQL client.
    sudo apt-get update
    sudo apt-get install postgresql-client

    # Step2: Connect to the PostgreSQL server. The password is "postgres".
    psql -h localhost -p 30001 -U postgres -d nebula


**4. Access the k3s dashboard.** 


Access the k3s dashboard at: http://localhost:30080/kubernetes-dashboard.

**5. Access the endpoints.**


Service endpoints are defined in the `nebulaidl` repository under the `service` directory. You can browse them at `here <https://github.com/nebulaclouds/nebulaidl/tree/master/protos/nebulaidl/service>`__.

For example, the endpoint for the `ListTaskExecutions <https://github.com/nebulaclouds/nebulaidl/blob/b219c2ab37886801039fda67d913760ac6fc4c8b/protos/nebulaidl/service/admin.proto#L442>`__ API is:

.. code:: shell

   /api/v1/task_executions/{node_execution_id.execution_id.project}/{node_execution_id.execution_id.domain}/{node_execution_id.execution_id.name}/{node_execution_id.node_id}

You can access this endpoint at:

.. code:: shell

   # replace with your specific task execution parameters
   http://localhost:30080/api/v1/task_executions/nebulasnacks/development/fe92c0a8cbf684ad19a8/n0?limit=10000






🐞 File an issue
================

We use `GitHub Issues <https://github.com/nebulaclouds/nebula/issues>`__ for issue tracking. The following issue types are available for filing an issue:

* `Plugin Request <https://github.com/nebulaclouds/nebula/issues/new?assignees=&labels=untriaged%2Cplugins&template=backend-plugin-request.md&title=%5BPlugin%5D>`__
* `Bug Report <https://github.com/nebulaclouds/nebula/issues/new?assignees=&labels=bug%2C+untriaged&template=bug_report.md&title=%5BBUG%5D+>`__
* `Documentation Bug/Update Request <https://github.com/nebulaclouds/nebula/issues/new?assignees=&labels=documentation%2C+untriaged&template=docs_issue.md&title=%5BDocs%5D>`__
* `Core Feature Request <https://github.com/nebulaclouds/nebula/issues/new?assignees=&labels=enhancement%2C+untriaged&template=feature_request.md&title=%5BCore+Feature%5D>`__
* `Nebulactl Feature Request <https://github.com/nebulaclouds/nebula/issues/new?assignees=&labels=enhancement%2C+untriaged%2C+nebulactl&template=nebulactl_issue.md&title=%5BNebulactl+Feature%5D>`__
* `Housekeeping <https://github.com/nebulaclouds/nebula/issues/new?assignees=&labels=housekeeping&template=housekeeping_template.md&title=%5BHousekeeping%5D+>`__
* `UI Feature Request <https://github.com/nebulaclouds/nebula/issues/new?assignees=&labels=enhancement%2C+untriaged%2C+ui&template=ui_feature_request.md&title=%5BUI+Feature%5D>`__

If none of the above fit your requirements, file a `blank <https://github.com/nebulaclouds/nebula/issues/new>`__ issue.
Also, add relevant labels to your issue. For example, if you are filing a Nebulakit plugin request, add the ``nebulakit`` label.

For feedback at any point in the contribution process, feel free to reach out to us on `Slack <https://slack.nebula.org/>`__.
