.. _deployment-plugin-setup-k8s:

Configure Kubernetes Plugins
============================

.. tags:: Kubernetes, Integration, Spark, AWS, GCP, Advanced

This guide provides an overview of setting up the Kubernetes Operator backend plugin in your Nebula deployment.

Spin up a cluster
-----------------

.. tabs::

  .. group-tab:: Nebula binary

    .. tabs::

      .. group-tab:: Demo cluster

        .. tabs::

          .. group-tab:: PyTorch

            Enable the PyTorch plugin on the demo cluster by adding the following block to ``~/.nebula/sandbox/config.yaml``:

            .. code-block:: yaml

              tasks:
                task-plugins:
                  default-for-task-types:
                    container: container
                    container_array: k8s-array
                    sidecar: sidecar
                    pytorch: pytorch
                  enabled-plugins:
                  - container
                  - k8s-array
                  - sidecar
                  - pytorch

          .. group-tab:: TensorFlow

            Enable the TensorFlow plugin on the demo cluster by adding the following block to ``~/.nebula/sandbox/config.yaml``:

            .. code-block:: yaml

              tasks:
                task-plugins:
                  default-for-task-types:
                    container: container
                    container_array: k8s-array
                    sidecar: sidecar
                    tensorflow: tensorflow
                  enabled-plugins:
                  - container
                  - k8s-array
                  - sidecar
                  - tensorflow

          .. group-tab:: MPI

            Enable the MPI plugin on the demo cluster by adding the following block to ``~/.nebula/sandbox/config.yaml``:

            .. code-block:: yaml

              tasks:
                task-plugins:
                  default-for-task-types:
                    container: container
                    container_array: k8s-array
                    sidecar: sidecar
                    mpi: mpi
                  enabled-plugins:
                  - container
                  - k8s-array
                  - sidecar
                  - mpi

          .. group-tab:: Ray

            Enable the Ray plugin on the demo cluster by adding the following block to ``~/.nebula/sandbox/config.yaml``:

            .. code-block:: yaml

              tasks:
                task-plugins:
                  default-for-task-types:
                    container: container
                    container_array: k8s-array
                    sidecar: sidecar
                    ray: ray
                  enabled-plugins:
                  - container
                  - k8s-array
                  - sidecar
                  - ray

          .. group-tab:: Spark

            Enable the Spark plugin on the demo cluster by adding the following config to ``~/.nebula/sandbox/config.yaml``:

            .. code-block:: yaml

              tasks:
                task-plugins:
                  default-for-task-types:
                    container: container
                    container_array: k8s-array
                    sidecar: sidecar
                    spark: spark
                  enabled-plugins:
                    - container
                    - sidecar
                    - k8s-array
                    - spark
              plugins:
                spark:
                  spark-config-default:
                    - spark.driver.cores: "1"
                    - spark.hadoop.fs.s3a.aws.credentials.provider: "org.apache.hadoop.fs.s3a.SimpleAWSCredentialsProvider"
                    - spark.hadoop.fs.s3a.endpoint: "http://minio.nebula:9000"
                    - spark.hadoop.fs.s3a.access.key: "minio"
                    - spark.hadoop.fs.s3a.secret.key: "miniostorage"
                    - spark.hadoop.fs.s3a.path.style.access: "true"
                    - spark.kubernetes.allocation.batch.size: "50"
                    - spark.hadoop.fs.s3a.acl.default: "BucketOwnerFullControl"
                    - spark.hadoop.fs.s3n.impl: "org.apache.hadoop.fs.s3a.S3AFileSystem"
                    - spark.hadoop.fs.AbstractFileSystem.s3n.impl: "org.apache.hadoop.fs.s3a.S3A"
                    - spark.hadoop.fs.s3.impl: "org.apache.hadoop.fs.s3a.S3AFileSystem"
                    - spark.hadoop.fs.AbstractFileSystem.s3.impl: "org.apache.hadoop.fs.s3a.S3A"
                    - spark.hadoop.fs.s3a.impl: "org.apache.hadoop.fs.s3a.S3AFileSystem"
                    - spark.hadoop.fs.AbstractFileSystem.s3a.impl: "org.apache.hadoop.fs.s3a.S3A"
              cluster_resources:
                refreshInterval: 5m
                customData:
                  - production:
                      - projectQuotaCpu:
                          value: "5"
                      - projectQuotaMemory:
                          value: "4000Mi"
                  - staging:
                      - projectQuotaCpu:
                          value: "2"
                      - projectQuotaMemory:
                          value: "3000Mi"
                  - development:
                      - projectQuotaCpu:
                          value: "4"
                      - projectQuotaMemory:
                          value: "5000Mi"
                refresh: 5m
              
            Also add the following cluster resource templates to the ``~/.nebula/sandbox/cluster-resource-templates`` directory:

            1. ``serviceaccount.yaml``

            .. code-block:: yaml

              apiVersion: v1
              kind: ServiceAccount
              metadata:
                name: default
                namespace: "{{ namespace }}"
                annotations:
                  eks.amazonaws.com/role-arn: "{{ defaultIamRole }}"
            
            2. ``spark_role.yaml``

            .. code-block:: yaml

              apiVersion: rbac.authorization.k8s.io/v1
              kind: Role
              metadata:
                name: spark-role
                namespace: "{{ namespace }}"
              rules:
                - apiGroups:
                    - ""
                  resources:
                    - pods
                    - services
                    - configmaps
                  verbs:
                    - "*"

            3. ``spark_service_account.yaml``

            .. code-block:: yaml

              apiVersion: v1
              kind: ServiceAccount
              metadata:
                name: spark
                namespace: "{{ namespace }}"
                annotations:
                  eks.amazonaws.com/role-arn: "{{ defaultIamRole }}"
            
            4. ``spark_role_binding.yaml``

            .. code-block:: yaml

              apiVersion: rbac.authorization.k8s.io/v1
              kind: RoleBinding
              metadata:
                name: spark-role-binding
                namespace: "{{ namespace }}"
              roleRef:
                apiGroup: rbac.authorization.k8s.io
                kind: Role
                name: spark-role
              subjects:
                - kind: ServiceAccount
                  name: spark
                  namespace: "{{ namespace }}"

          .. group-tab:: Dask

            Enable the Dask plugin on the demo cluster by adding the following block to ``~/.nebula/sandbox/config.yaml``:

            .. code-block:: yaml

              tasks:
                task-plugins:
                  default-for-task-types:
                    container: container
                    container_array: k8s-array
                    sidecar: sidecar
                    dask: dask
                  enabled-plugins:
                  - container
                  - k8s-array
                  - sidecar
                  - dask

        Start the demo cluster by running the following command:

        .. code-block:: bash
      
          nebulactl demo start

      .. group-tab:: Helm chart

         1. Add the following to your values file under `configmap.inline`:

         .. code-block:: yaml

            tasks:
              task-plugins:
                enabled-plugins:
                  - container
                  - sidecar
                  - K8S-ARRAY
                  - spark
                  - ray
                default-for-task-types:
                  - container: container
                  - container_array: K8S-ARRAY
                  - spark: spark
                  - ray: ray

         2. Install the :ref:`nebula-binary Helm chart <deployment-deployment-cloud-simple>`.
   
  .. group-tab:: Nebula core
   
    If you have installed Nebula using the `nebula-core Helm chart
    <https://github.com/nebulaclouds/nebula/tree/master/charts/nebula-core>`__, please ensure:

    * You have the correct kubeconfig and have selected the correct Kubernetes context.
    * You have configured the correct nebulactl settings in ``~/.nebula/config.yaml``.

.. note::

  Add the Nebula chart repo to Helm if you're installing via the Helm charts.

  .. code-block:: bash

    helm repo add nebulaclouds https://nebulaclouds.github.io/nebula

Install the Kubernetes operator
-------------------------------

.. tabs::

  .. group-tab:: PyTorch/TensorFlow/MPI

    First, `install kustomize <https://kubectl.docs.kubernetes.io/installation/kustomize/>`__.

    Build and apply the training-operator.
  
    .. code-block:: bash
  
      export KUBECONFIG=$KUBECONFIG:~/.kube/config:~/.nebula/k3s/k3s.yaml
      kustomize build "https://github.com/kubeflow/training-operator.git/manifests/overlays/standalone?ref=v1.5.0" | kubectl apply -f -

    **Optional: Using a gang scheduler**

    To address potential issues with worker pods of distributed training jobs being scheduled at different times
    due to resource constraints, you can opt for a gang scheduler. This ensures that all worker pods are scheduled
    simultaneously, reducing the likelihood of job failures caused by timeout errors.
    
    To `enable gang scheduling for the Kubeflow training-operator <https://www.kubeflow.org/docs/components/training/job-scheduling/>`__,
    you can install the `Kubernetes scheduler plugins <https://github.com/kubernetes-sigs/scheduler-plugins/tree/master>`__
    or the `Apache YuniKorn scheduler <https://yunikorn.apache.org/>`__.

    1. Install the `scheduler plugin <https://github.com/kubernetes-sigs/scheduler-plugins/tree/master/manifests/install/charts/as-a-second-scheduler>`_ or
       `Apache YuniKorn <https://yunikorn.apache.org/docs/next/#install>`_ as a second scheduler.
    2. Configure the Kubeflow training-operator to use the new scheduler:

        Create a manifest called ``kustomization.yaml`` with the following content:

        .. code-block:: yaml

          apiVersion: kustomize.config.k8s.io/v1beta1
          kind: Kustomization

          resources:
          - github.com/kubeflow/training-operator/manifests/overlays/standalone

          patchesStrategicMerge:
          - patch.yaml

        Create a patch file called ``patch.yaml`` with the following content:

        .. code-block:: yaml

          apiVersion: apps/v1
          kind: Deployment
          metadata:
            name: training-operator
          spec:
            template:
              spec:
                containers:
                - name: training-operator
                  command:
                  - /manager
                  - --gang-scheduler-name=<scheduler-plugins/yunikorn>

        Install the patched kustomization with the following command:

        .. code-block:: bash

          kustomize build path/to/overlay/directory | kubectl apply -f -

       (Only for Apache YuniKorn) To configure gang scheduling with Apache YuniKorn,
       make sure to set the following annotations in Nebula pod templates:

       - ``template.metadata.annotations.yunikorn.apache.org/task-group-name``
       - ``template.metadata.annotations.yunikorn.apache.org/task-groups``
       - ``template.metadata.annotations.yunikorn.apache.org/schedulingPolicyParameters``

       For more configuration details,
       refer to the `Apache YuniKorn Gang-Scheduling documentation 
       <https://yunikorn.apache.org/docs/next/user_guide/gang_scheduling>`__.

    3. Use a Nebula pod template with ``template.spec.schedulerName: scheduler-plugins-scheduler``
       to use the new gang scheduler for your tasks.
      
       See the :ref:`using-k8s-podtemplates` section for more information on pod templates in Nebula.
       You can set the scheduler name in the pod template passed to the ``@task`` decorator. However, to prevent the
       two different schedulers from competing for resources, it is recommended to set the scheduler name in the pod template
       in the ``nebula`` namespace which is applied to all tasks. Non distributed training tasks can be scheduled by the
       gang scheduler as well.


       For more information on pod templates in Nebula, refer to the :ref:`using-k8s-podtemplates` section.
       You can set the scheduler name in the pod template passed to the ``@task`` decorator.
       However, to avoid resource competition between the two different schedulers,
       it is recommended to set the scheduler name in the pod template in the ``nebula`` namespace,
       which is applied to all tasks. This allows non-distributed training tasks to be 
       scheduled by the gang scheduler as well.

  .. group-tab:: Ray
    
    To install the Ray Operator, run the following commands:
  
    .. code-block:: bash
  
        export KUBERAY_VERSION=v0.5.2
        kubectl create -k "github.com/ray-project/kuberay/manifests/cluster-scope-resources?ref=${KUBERAY_VERSION}&timeout=90s"
        kubectl apply -k "github.com/ray-project/kuberay/manifests/base?ref=${KUBERAY_VERSION}&timeout=90s"
  
  .. group-tab:: Spark
  
    To add the Spark repository, run the following commands:
  
    .. code-block:: bash
  
       helm repo add spark-operator https://googlecloudplatform.github.io/spark-on-k8s-operator
  
    To install the Spark operator, run the following command:
  
    .. code-block:: bash
  
       helm install spark-operator spark-operator/spark-operator --namespace spark-operator --create-namespace
  
  .. group-tab:: Dask
  
    To add the Dask repository, run the following command:
  
    .. code-block:: bash
  
       helm repo add dask https://helm.dask.org
  
    To install the Dask operator, run the following command:
  
    .. code-block:: bash
  
       helm install dask-operator dask/dask-kubernetes-operator --namespace dask-operator --create-namespace

Specify plugin configuration
----------------------------

.. tabs::

  .. group-tab:: PyTorch

    .. tabs::

      .. group-tab:: Nebula binary

        To specify the plugin when using the Helm chart, edit the relevant YAML file.

        .. code-block:: yaml
          :emphasize-lines: 7,11

          tasks:
            task-plugins:
              enabled-plugins:
                - container
                - sidecar
                - k8s-array
                - pytorch
              default-for-task-types:
                - container: container
                - container_array: k8s-array
                - pytorch: pytorch

      .. group-tab:: Nebula core
    
        Create a file named ``values-override.yaml`` and add the following config to it:
    
        .. code-block:: yaml
    
          configmap:
            enabled_plugins:
              tasks:
                task-plugins:
                  enabled-plugins:
                    - container
                    - sidecar
                    - k8s-array
                    - pytorch
                  default-for-task-types:
                    container: container
                    sidecar: sidecar
                    container_array: k8s-array
                    pytorch: pytorch
   
  .. group-tab:: TensorFlow
   
    .. tabs::

      .. group-tab:: Nebula binary

        To specify the plugin when using the Helm chart, edit the relevant YAML file.

        .. code-block:: yaml
          :emphasize-lines: 7,11

          tasks:
            task-plugins:
              enabled-plugins:
                - container
                - sidecar
                - k8s-array
                - tensorflow
              default-for-task-types:
                - container: container
                - container_array: k8s-array
                - tensorflow: tensorflow

      .. group-tab:: Nebula core
    
        Create a file named ``values-override.yaml`` and add the following config to it:
    
        .. code-block:: yaml
    
          configmap:
            enabled_plugins:
              tasks:
                task-plugins:
                  enabled-plugins:
                    - container
                    - sidecar
                    - k8s-array
                    - tensorflow
                  default-for-task-types:
                    container: container
                    sidecar: sidecar
                    container_array: k8s-array
                    tensorflow: tensorflow
   
  .. group-tab:: MPI
   
    .. tabs::

      .. group-tab:: Nebula binary

        To specify the plugin when using the Helm chart, edit the relevant YAML file.

        .. code-block:: yaml
          :emphasize-lines: 7,11

          tasks:
            task-plugins:
              enabled-plugins:
                - container
                - sidecar
                - k8s-array
                - mpi
              default-for-task-types:
                - container: container
                - container_array: k8s-array
                - mpi: mpi

      .. group-tab:: Nebula core
    
        Create a file named ``values-override.yaml`` and add the following config to it:
    
        .. code-block:: yaml
    
          configmap:
            enabled_plugins:
              tasks:
                task-plugins:
                  enabled-plugins:
                    - container
                    - sidecar
                    - k8s-array
                    - mpi
                  default-for-task-types:
                    container: container
                    sidecar: sidecar
                    container_array: k8s-array
                    mpi: mpi
  
  .. group-tab:: Ray

    .. tabs::

      .. group-tab:: Nebula binary

        1. Make sure that your Helm values file includes the following configuration:

        .. code-block:: yaml

           configuration:
             inline:
              tasks:
                task-plugins:
                  enabled-plugins:
                    - container
                    - sidecar
                    - k8s-array
                    - ray
                  default-for-task-types:
                    - container: container
                    - container_array: k8s-array
                    - ray: ray

           rbac:
             extraRules:
               - apiGroups:
               - "ray.io"
               resources:
               - rayjob
               verbs:
               - create
               - get
               - list
               - patch
               - update

        2. Run a ``helm upgrade`` operation

      .. group-tab:: Nebula core
    
        Create a file named ``values-override.yaml`` and add the following config to it:
    
        .. code-block:: yaml
    
          configmap:
            enabled_plugins:
              tasks:
                task-plugins:
                  enabled-plugins:
                    - container
                    - sidecar
                    - k8s-array
                    - ray
                  default-for-task-types:
                    container: container
                    sidecar: sidecar
                    container_array: k8s-array
                    ray: ray
   
  .. group-tab:: Spark
   
      .. tabs:: 

        .. group-tab:: Nebula binary

          To specify the plugin when using the Helm chart, edit the relevant YAML file.

        .. group-tab:: Nebula core

          Create a file named ``values-override.yaml`` and add the following config to it:
   
          .. code-block:: yaml
   
            cluster_resource_manager:
              enabled: true
              config:
                cluster_resources:
                  refreshInterval: 5m
                  templatePath: "/etc/nebula/clusterresource/templates"
                  customData:
                    - production:
                        - projectQuotaCpu:
                            value: "5"
                        - projectQuotaMemory:
                            value: "4000Mi"
                    - staging:
                        - projectQuotaCpu:
                            value: "2"
                        - projectQuotaMemory:
                            value: "3000Mi"
                    - development:
                        - projectQuotaCpu:
                            value: "4"
                        - projectQuotaMemory:
                            value: "3000Mi"
                  refresh: 5m
      
              # -- Resource templates that should be applied
              templates:
                # -- Template for namespaces resources
                - key: aa_namespace
                  value: |
                    apiVersion: v1
                    kind: Namespace
                    metadata:
                      name: {{ namespace }}
                    spec:
                      finalizers:
                      - kubernetes
      
                - key: ab_project_resource_quota
                  value: |
                    apiVersion: v1
                    kind: ResourceQuota
                    metadata:
                      name: project-quota
                      namespace: {{ namespace }}
                    spec:
                      hard:
                        limits.cpu: {{ projectQuotaCpu }}
                        limits.memory: {{ projectQuotaMemory }}
      
                - key: ac_spark_role
                  value: |
                    apiVersion: rbac.authorization.k8s.io/v1beta1
                    kind: Role
                    metadata:
                      name: spark-role
                      namespace: {{ namespace }}
                    rules:
                    - apiGroups: ["*"]
                      resources:
                      - pods
                      verbs:
                      - '*'
                    - apiGroups: ["*"]
                      resources:
                      - services
                      verbs:
                      - '*'
                    - apiGroups: ["*"]
                      resources:
                      - configmaps
                      verbs:
                      - '*'
      
                - key: ad_spark_service_account
                  value: |
                    apiVersion: v1
                    kind: ServiceAccount
                    metadata:
                      name: spark
                      namespace: {{ namespace }}
      
                - key: ae_spark_role_binding
                  value: |
                    apiVersion: rbac.authorization.k8s.io/v1beta1
                    kind: RoleBinding
                    metadata:
                      name: spark-role-binding
                      namespace: {{ namespace }}
                    roleRef:
                      apiGroup: rbac.authorization.k8s.io
                      kind: Role
                      name: spark-role
                    subjects:
                    - kind: ServiceAccount
                      name: spark
                      namespace: {{ namespace }}
      
            sparkoperator:
              enabled: true
              plugin_config:
                plugins:
                  spark:
                    # Edit the Spark configuration as you see fit
                    spark-config-default:
                      - spark.driver.cores: "1"
                      - spark.hadoop.fs.s3a.aws.credentials.provider: "com.amazonaws.auth.DefaultAWSCredentialsProviderChain"
                      - spark.kubernetes.allocation.batch.size: "50"
                      - spark.hadoop.fs.s3a.acl.default: "BucketOwnerFullControl"
                      - spark.hadoop.fs.s3n.impl: "org.apache.hadoop.fs.s3a.S3AFileSystem"
                      - spark.hadoop.fs.AbstractFileSystem.s3n.impl: "org.apache.hadoop.fs.s3a.S3A"
                      - spark.hadoop.fs.s3.impl: "org.apache.hadoop.fs.s3a.S3AFileSystem"
                      - spark.hadoop.fs.AbstractFileSystem.s3.impl: "org.apache.hadoop.fs.s3a.S3A"
                      - spark.hadoop.fs.s3a.impl: "org.apache.hadoop.fs.s3a.S3AFileSystem"
                      - spark.hadoop.fs.AbstractFileSystem.s3a.impl: "org.apache.hadoop.fs.s3a.S3A"
                      - spark.network.timeout: 600s
                      - spark.executorEnv.KUBERNETES_REQUEST_TIMEOUT: 100000
                      - spark.executor.heartbeatInterval: 60s
            configmap:
              enabled_plugins:
                tasks:
                  task-plugins:
                    enabled-plugins:
                      - container
                      - sidecar
                      - k8s-array
                      - spark
                    default-for-task-types:
                      container: container
                      sidecar: sidecar
                      container_array: k8s-array
                      spark: spark
   
  .. group-tab:: Dask
   
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
                - dask
              default-for-task-types:
                - container: container
                - container_array: k8s-array
                - dask: dask

      .. group-tab:: Nebula core
    
        Create a file named ``values-override.yaml`` and add the following config to it:
    
        .. code-block:: yaml
    
          configmap:
            enabled_plugins:
              tasks:
                task-plugins:
                  enabled-plugins:
                    - container
                    - sidecar
                    - k8s-array
                    - dask
                  default-for-task-types:
                    container: container
                    sidecar: sidecar
                    container_array: k8s-array
                    dask: dask

Upgrade the deployment
----------------------

.. tabs::

  .. group-tab:: Nebula binary

    If you are installing Nebula via the Helm chart, run the following command:

    .. note::

      There is no need to run ``helm upgrade`` for Spark.

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

Wait for the upgrade to complete. You can check the status of the deployment pods by running the following command:

.. code-block:: bash

  kubectl get pods -n --all-namespaces
