.. _deployment-plugin-setup-webapi-snowflake:

Snowflake Plugin
================

This guide provides an overview of how to set up Snowflake in your Nebula deployment.

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
    `nebula-core helm chart <https://github.com/nebulaclouds/nebula/tree/master/charts/nebula-core>`__,
    please ensure:

    * You have the correct kubeconfig and have selected the correct Kubernetes context.
    * You have configured the correct nebulactl settings in ``~/.nebula/config.yaml``.

.. note::

  Add the Nebula chart repo to Helm if you're installing via the Helm charts.

  .. code-block:: bash

    helm repo add nebulaclouds https://nebulaclouds.github.io/nebula

Specify plugin configuration
----------------------------

.. tabs::

  .. group-tab:: Nebula binary

    .. tabs::
         
      .. group-tab:: Demo cluster

        Enable the Snowflake plugin on the demo cluster by adding the following block to ``~/.nebula/sandbox/config.yaml``:

        .. code-block:: yaml

          tasks:
            task-plugins:
              default-for-task-types:
                container: container
                container_array: k8s-array
                sidecar: sidecar
                snowflake: snowflake
              enabled-plugins:
                - container
                - k8s-array
                - sidecar
                - snowflake

      .. group-tab:: Helm chart

        Edit the relevant YAML file to specify the plugin.

        .. code-block:: yaml
          :emphasize-lines: 7,11

          tasks:
            task-plugins:
              enabled-plugins:
                - container
                - sidecar
                - k8s-array
                - snowflake
              default-for-task-types:
                - container: container
                - container_array: k8s-array
                - snowflake: snowflake

  .. group-tab:: Nebula core
    
    Create a file named ``values-override.yaml`` and add the following config to it:

    .. code-block:: yaml

        configmap:
          enabled_plugins:
            # -- Tasks specific configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#GetConfig)
            tasks:
              # -- Plugins configuration, [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#TaskPluginConfig)
              task-plugins:
                # -- [Enabled Plugins](https://pkg.go.dev/github.com/nebulaclouds/nebulaplugins/go/tasks/config#Config). Enable sagemaker*, athena if you install the backend
                # plugins
                enabled-plugins:
                  - container
                  - sidecar
                  - k8s-array
                  - snowflake
                default-for-task-types:
                  container: container
                  sidecar: sidecar
                  container_array: k8s-array
                  snowflake: snowflake

Obtain and add the Snowflake JWT token
--------------------------------------

Create a Snowflake account, and follow the `Snowflake docs 
<https://docs.snowflake.com/en/developer-guide/sql-api/authenticating#using-key-pair-authentication>`__
to generate a JWT token.
Then, add the Snowflake JWT token to NebulaPropeller.

.. tabs::

  .. group-tab:: Nebula binary

    .. tabs::

      .. group-tab:: Demo cluster

        Add the JWT token as an environment variable to the ``nebula-sandbox`` deployment.

        .. code-block:: bash

          kubectl edit deploy nebula-sandbox -n nebula

        Update the ``env`` configuration:

        .. code-block:: yaml
          :emphasize-lines: 12-13

          env:
          - name: POD_NAME
            valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.name
          - name: POD_NAMESPACE
            valueFrom:
            fieldRef:
              apiVersion: v1
              fieldPath: metadata.namespace
          - name: NEBULA_SECRET_NEBULA_SNOWFLAKE_CLIENT_TOKEN
            value: <JWT_TOKEN>
          image: nebula-binary:sandbox
          ...

      .. group-tab:: Helm chart

        Create an external secret as follows:

        .. code-block:: bash

          cat <<EOF | kubectl apply -f -
          apiVersion: v1
          kind: Secret
          metadata:
            name: nebula-binary-client-secrets-external-secret
            namespace: nebula
          type: Opaque
          stringData:
            NEBULA_SNOWFLAKE_CLIENT_TOKEN: <JWT_TOKEN>
          EOF
        
        Reference the newly created secret in 
        ``.Values.configuration.auth.clientSecretsExternalSecretRef``
        in your YAML file as follows:

        .. code-block:: yaml
          :emphasize-lines: 3

          configuration:
            auth:
              clientSecretsExternalSecretRef: nebula-binary-client-secrets-external-secret
      
    Replace ``<JWT_TOKEN>`` with your JWT token.

  .. group-tab:: Nebula core

    Add the JWT token as a secret to ``nebula-secret-auth``.

    .. code-block:: bash

      kubectl edit secret -n nebula nebula-secret-auth

    .. code-block:: yaml
      :emphasize-lines: 3

      apiVersion: v1
      data:
        NEBULA_SNOWFLAKE_CLIENT_TOKEN: <JWT_TOKEN>
        client_secret: Zm9vYmFy
      kind: Secret
      ...

    Replace ``<JWT_TOKEN>`` with your JWT token.

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