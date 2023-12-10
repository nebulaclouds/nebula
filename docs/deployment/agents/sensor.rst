.. _deployment-agent-setup-sensor:

Sensor Agent
=================

Sensor enables users to continuously check for a file or a condition to be met periodically.

When the condition is met, the sensor will complete.

This guide provides an overview of how to set up Sensor in your Nebula deployment.

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
    * Confirm that you have the correct Nebulactl configuration at ``~/.nebula/config.yaml``.

.. note::

  Add the Nebula chart repo to Helm if you're installing via the Helm charts.

  .. code-block:: bash

    helm repo add nebulaclouds https://nebulaclouds.github.io/nebula

Specify agent configuration
----------------------------

Enable the Sensor agent by adding the following config to the relevant YAML file(s):

.. tabs::

    .. group-tab:: Nebula binary

      Edit the relevant YAML file to specify the agent.

      .. code-block:: bash

        kubectl edit configmap nebula-sandbox-config -n nebula

      .. code-block:: yaml
        :emphasize-lines: 7,11,16
  
        tasks:
          task-plugins:
            enabled-plugins:
              - container
              - sidecar
              - k8s-array
              - agent-service
            default-for-task-types:
              - container: container
              - container_array: k8s-array
              - sensor: agent-service
        
        plugins:
          agent-service:
            supportedTaskTypes:
            - sensor

    .. group-tab:: Nebula core

      Create a file named ``values-override.yaml`` and add the following configuration to it.

      .. code-block:: yaml

        configmap:
          enabled_plugins:
            # -- Tasks specific configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#GetConfig)
            tasks:
              # -- Plugins configuration, [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#TaskPluginConfig)
              task-plugins:
                # -- [Enabled Plugins](https://pkg.go.dev/github.com/nebulaclouds/nebulaplugins/go/tasks/config#Config). Enable sagemaker*, athena if you install the backend
                enabled-plugins:
                  - container
                  - sidecar
                  - k8s-array
                  - agent-service
                default-for-task-types:
                  container: container
                  sidecar: sidecar
                  container_array: k8s-array
                  sensor: agent-service
            plugins:
              agent-service:
                supportedTaskTypes:
                - sensor


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

Wait for the upgrade to complete.

You can check the status of the deployment pods by running the following command:

.. code-block::

  kubectl get pods -n nebula
