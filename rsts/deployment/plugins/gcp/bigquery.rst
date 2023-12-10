.. _deployment-plugin-setup-gcp-bigquery:

Google BigQuery Plugin
======================

This guide provides an overview of setting up BigQuery in your Nebula deployment.
Please note that the BigQuery plugin requires Nebula deployment in the GCP cloud;
it is not compatible with demo/AWS/Azure.

Set up the GCP Nebula cluster
----------------------------

* Ensure you have a functional Nebula cluster running in `GCP <https://docs.nebula.org/en/latest/deployment/gcp/index.html#deployment-gcp>`__.
* Create a service account for BigQuery. For more details, refer to: https://cloud.google.com/bigquery/docs/quickstarts/quickstart-client-libraries.
* Verify that you have the correct kubeconfig and have selected the appropriate Kubernetes context.
* Confirm that you have the correct Nebulactl configuration at ``~/.nebula/config.yaml``.

Specify plugin configuration
----------------------------

.. tabs::

  .. group-tab:: Nebula binary

    Edit the relevant YAML file to specify the plugin.

    .. code-block:: yaml
      :emphasize-lines: 7,11

      tasks:
        task-plugins:
          enabled-plugins:
            - container
            - sidecar
            - k8s-array
            - bigquery
          default-for-task-types:
            - container: container
            - container_array: k8s-array
            - bigquery_query_job_task: bigquery

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
                  - bigquery
                default-for-task-types:
                  container: container
                  sidecar: sidecar
                  container_array: k8s-array
                  bigquery_query_job_task: bigquery

Ensure that the propeller has the correct service account for BigQuery.

Upgrade the Nebula Helm release
------------------------------

.. tabs::

  .. group-tab:: Nebula binary

    .. code-block:: bash

      helm upgrade <RELEASE_NAME> nebulaclouds/nebula-binary -n <YOUR_NAMESPACE> --values <YOUR_YAML_FILE>

    Replace ``<RELEASE_NAME>`` with the name of your release (e.g., ``nebula-backend``),
    ``<YOUR_NAMESPACE>`` with the name of your namespace (e.g., ``nebula``),
    and ``<YOUR_YAML_FILE>`` with the name of your YAML file.

  .. group-tab:: Nebula core

    .. code-block:: bash
    
      helm upgrade <RELEASE_NAME> nebula/nebula-core -n <YOUR_NAMESPACE> --values values-override.yaml

    Replace ``<RELEASE_NAME>`` with the name of your release (e.g., ``nebula``)
    and ``<YOUR_NAMESPACE>`` with the name of your namespace (e.g., ``nebula``).
