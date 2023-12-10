.. _deployment-agent-setup-databricks:

Databricks Agent
=================

This guide provides an overview of how to set up Databricks agent in your Nebula deployment.

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

Databricks workspace
--------------------

To set up your Databricks account, follow these steps:

1. Create a `Databricks account <https://www.databricks.com/>`__.

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/deployment/plugins/databricks/databricks_workspace.png
    :alt: A screenshot of Databricks workspace creation.

2. Ensure that you have a Databricks workspace up and running.

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/deployment/plugins/databricks/open_workspace.png
    :alt: A screenshot of Databricks workspace.

3. Generate a `personal access token 
   <https://docs.databricks.com/dev-tools/auth.html#databricks-personal-ACCESS_TOKEN-authentication>`__ to be used in the Nebula configuration.
   You can find the personal access token in the user settings within the workspace. ``User settings`` -> ``Developer`` -> ``Access tokens``

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/deployment/plugins/databricks/databricks_access_token.png
    :alt: A screenshot of access token.

4. Enable custom containers on your Databricks cluster before you trigger the workflow.

.. code-block:: bash

   curl -X PATCH -n -H "Authorization: Bearer <your-personal-access-token>" \
   https://<databricks-instance>/api/2.0/workspace-conf \
   -d '{"enableDcs": "true"}'

For more detail, check `custom containers <https://docs.databricks.com/administration-guide/clusters/container-services.html>`__.

5. Create an `instance profile 
<https://docs.databricks.com/administration-guide/cloud-configurations/aws/instance-profiles.html>`__ 
for the Spark cluster. This profile enables the Spark job to access your data in the S3 bucket.

Create an instance profile using the AWS console (For AWS Users)
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^

1. In the AWS console, go to the IAM service.
2. Click the Roles tab in the sidebar.
3. Click Create role.

   a. Under Trusted entity type, select AWS service.
   b. Under Use case, select **EC2**.
   c. Click Next.
   d. At the bottom of the page, click Next.
   e. In the Role name field, type a role name.
   f. Click Create role.

4. In the role list, click the **AmazonS3FullAccess** role.
5. Click Create role button.

In the role summary, copy the Role ARN.

.. image:: https://raw.githubusercontent.com/nebulaclouds/static-resources/main/nebula/deployment/plugins/databricks/s3_arn.png
    :alt: A screenshot of s3 arn.

Locate the IAM role that created the Databricks deployment
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
If you don’t know which IAM role created the Databricks deployment, do the following:

1. As an account admin, log in to the account console.
2. Go to ``Workspaces`` and click your workspace name.
3. In the Credentials box, note the role name at the end of the Role ARN

For example, in the Role ARN ``arn:aws:iam::123456789123:role/finance-prod``, the role name is finance-prod

Edit the IAM role that created the Databricks deployment
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
1. In the AWS console, go to the IAM service.
2. Click the Roles tab in the sidebar.
3. Click the role that created the Databricks deployment.
4. On the Permissions tab, click the policy.
5. Click Edit Policy.
6. Append the following block to the end of the Statement array. Ensure that you don’t overwrite any of the existing policy. Replace <iam-role-for-s3-access> with the role you created in Configure S3 access with instance profiles.

.. code-block:: bash

    {
      "Effect": "Allow",
      "Action": "iam:PassRole",
      "Resource": "arn:aws:iam::<aws-account-id-databricks>:role/<iam-role-for-s3-access>"
    }

Specify agent configuration
----------------------------

.. tabs::

  .. group-tab:: Nebula binary

    .. tabs::
      
      .. group-tab:: Demo cluster

        Enable the Databricks agent on the demo cluster by updating the ConfigMap:

        .. code-block:: bash

          kubectl edit configmap nebula-sandbox-config -n nebula

        .. code-block:: yaml
          :emphasize-lines: 7,12,16

          tasks:
            task-plugins:
              default-for-task-types:
                container: container
                container_array: k8s-array
                sidecar: sidecar
                spark: agent-service
              enabled-plugins:
                - container
                - sidecar
                - k8s-array
                - agent-service
          plugins:
            agent-service:
              supportedTaskTypes:
              - spark

      .. group-tab:: Helm chart

        Edit the relevant YAML file to specify the plugin.

        .. code-block:: yaml
          :emphasize-lines: 7,11,15

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
                - spark: agent-service
            plugins:
              agent-service:
                supportedTaskTypes:
                - spark

  .. group-tab:: Nebula core

    Create a file named ``values-override.yaml`` and add the following config to it:

    .. code-block:: yaml
      :emphasize-lines: 9,14-17

        enabled_plugins:
          tasks:
            task-plugins:
              enabled-plugins:
                - container
                - sidecar
                - k8s-array
                - agent-service
              default-for-task-types:
                container: container
                sidecar: sidecar
                container_array: k8s-array
                spark: agent-service
          plugins:
            agent-service:
              supportedTaskTypes:
              - spark

Add the Databricks access token
-------------------------------

You have to set the Databricks token to the Nebula configuration.

1. Install nebulaagent pod using helm
  
.. code-block::
  
  helm repo add nebulaclouds https://nebulaclouds.github.io/nebula
  helm install nebulaagent nebulaclouds/nebulaagent --namespace nebula

2. Get the base64 value of your Databricks token.

.. code-block::

  echo -n "<DATABRICKS_TOKEN>" | base64

3. Edit the nebulaagent secret
  
      .. code-block:: bash
    
        kubectl edit secret nebulaagent -n nebula
    
      .. code-block:: yaml
        :emphasize-lines: 3

        apiVersion: v1
        data:
          nebula_databricks_access_token: <BASE64_ENCODED_DATABRICKS_TOKEN>
        kind: Secret
        metadata:
          annotations:
            meta.helm.sh/release-name: nebulaagent
            meta.helm.sh/release-namespace: nebula
          creationTimestamp: "2023-10-04T04:09:03Z"
          labels:
            app.kubernetes.io/managed-by: Helm
          name: nebulaagent
          namespace: nebula
          resourceVersion: "753"
          uid: 5ac1e1b6-2a4c-4e26-9001-d4ba72c39e54
        type: Opaque


Upgrade the deployment
----------------------

.. tabs::

  .. group-tab:: Nebula binary

    .. tabs::

      .. group-tab:: Demo cluster

        .. code-block::

          kubectl rollout restart deployment nebula-sandbox -n nebula

      .. group-tab:: Helm chart

        .. code-block::

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

For databricks plugin on the Nebula cluster, please refer to `Databricks Plugin Example <https://docs.nebula.org/projects/cookbook/en/latest/auto_examples/databricks_plugin/index.html>`_
