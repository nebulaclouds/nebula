.. _deployment-agent-setup-mmcloud:

MMCloud Agent
=================

MemVerge Memory Machine Cloud (MMCloud) empowers users to continuously optimize cloud resources during runtime,
safely execute stateful tasks on spot instances,
and monitor resource usage in real time.
These capabilities make it an excellent fit for long-running batch workloads.

This guide provides an overview of how to set up MMCloud in your Nebula deployment.

Set up MMCloud
--------------

To run a Nebula workflow with Memory Machine Cloud, you will need to deploy Memory Machine Cloud.
Check out the `MMCloud User Guide <https://docs.memverge.com/mmce/current/userguide/olh/index.html>`_ to get started!

By the end of this step, you should have deployed an MMCloud OpCenter.

Spin up a cluster
-----------------

.. tabs::

  .. group-tab:: Nebula binary

    You can spin up a demo cluster using the following command:

    .. code-block:: bash

      nebulactl demo start

    Or install Nebula using the :ref:`nebula-binary helm chart <deployment-deployment-cloud-simple>`.

  .. group-tab:: Nebula core

    If you've installed Nebula using the
    `nebula-core helm chart <https://github.com/nebulaclouds/nebula/tree/master/charts/nebula-core>`__, please ensure:

    * You have the correct kubeconfig and have selected the correct Kubernetes context.
    * You have configured the correct nebulactl settings in ``~/.nebula/config.yaml``.

.. note::

  Add the Nebula chart repo to Helm if you're installing via the Helm charts.

  .. code-block:: bash

    helm repo add nebulaclouds https://nebulaclouds.github.io/nebula

Specify agent configuration
----------------------------

Enable the MMCloud agent by adding the following config to the relevant YAML file(s):

.. code-block:: yaml

  tasks:
    task-plugins:
      enabled-plugins:
        - agent-service
      default-for-task-types:
        - mmcloud_task: agent-service

.. code-block:: yaml

  plugins:
    agent-service:
      agents:
        mmcloud-agent:
          endpoint: <AGENT_ENDPOINT>
          insecure: true
      supportedTaskTypes:
      - mmcloud_task
      agentForTaskTypes:
      - mmcloud_task: mmcloud-agent

Substitute ``<AGENT_ENDPOINT>`` with the endpoint of your MMCloud agent.

Upgrade the deployment
----------------------

.. tabs::

  .. group-tab:: Nebula binary

    .. tabs::

      .. group-tab:: Demo cluster

        .. code-block:: bash

          kubectl rollout restart deployment nebula-sandbox -n nebula

      .. group-tab:: Helm chart

        .. code-block:: bash

          helm upgrade <RELEASE_NAME> nebulaclouds/nebula-binary -n <YOUR_NAMESPACE> --values <YOUR_YAML_FILE>

        Replace ``<RELEASE_NAME>`` with the name of your release (e.g., ``nebula-backend``),
        ``<YOUR_NAMESPACE>`` with the name of your namespace (e.g., ``nebula``),
        and ``<YOUR_YAML_FILE>`` with the name of your YAML file.

  .. group-tab:: Nebula core

    .. code-block::

      helm upgrade <RELEASE_NAME> nebula/nebula-core -n <YOUR_NAMESPACE> --values values-override.yaml

    Replace ``<RELEASE_NAME>`` with the name of your release (e.g., ``nebula``)
    and ``<YOUR_NAMESPACE>`` with the name of your namespace (e.g., ``nebula``).

Wait for the upgrade to complete. You can check the status of the deployment pods by running the following command:

.. code-block::

  kubectl get pods -n nebula

For mmcloud plugin on the Nebula cluster, please refer to `Memory Machine Cloud Plugin Example <https://docs.nebula.org/projects/cookbook/en/latest/auto_examples/mmcloud_plugin/index.html>`_
