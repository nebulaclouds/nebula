(integrations)=

# Integrations

Nebula is designed to be highly extensible and can be customized in multiple ways.

```{note}
Want to contribute an example? Check out the {doc}`Example Contribution Guide <nebulasnacks/contribute>`.
```

## Nebulakit Plugins

Nebulakit plugins are simple plugins that can be implemented purely in python, unit tested locally and allow extending
Nebulakit functionality. These plugins can be anything and for comparison can be thought of like
[Airflow Operators](https://airflow.apache.org/docs/apache-airflow/stable/howto/operator/index.html).

```{list-table}
:header-rows: 0
:widths: 20 30

* - {doc}`SQL <nebulasnacks/examples/sql_plugin/index>`
  - Execute SQL queries as tasks.
* - {doc}`Great Expectations <nebulasnacks/examples/greatexpectations_plugin/index>`
  - Validate data with `great_expectations`.
* - {doc}`Papermill <nebulasnacks/examples/papermill_plugin/index>`
  - Execute Jupyter Notebooks with `papermill`.
* - {doc}`Pandera <nebulasnacks/examples/pandera_plugin/index>`
  - Validate pandas dataframes with `pandera`.
* - {doc}`Modin <nebulasnacks/examples/modin_plugin/index>`
  - Scale pandas workflows with `modin`.
* - {doc}`Dolt <nebulasnacks/examples/dolt_plugin/index>`
  - Version your SQL database with `dolt`.
* - {doc}`DBT <nebulasnacks/examples/dbt_plugin/index>`
  - Run and test your `dbt` pipelines in Nebula.
* - {doc}`WhyLogs <nebulasnacks/examples/whylogs_plugin/index>`
  - `whylogs`: the open standard for data logging.
* - {doc}`MLFlow <nebulasnacks/examples/mlflow_plugin/index>`
  - `mlflow`: the open standard for model tracking.
* - {doc}`ONNX <nebulasnacks/examples/onnx_plugin/index>`
  - Convert ML models to ONNX models seamlessly.
* - {doc}`DuckDB <nebulasnacks/examples/duckdb_plugin/index>`
  - Run analytical queries using DuckDB.
```

:::{dropdown} {fa}`info-circle` Using nebulakit plugins
:animate: fade-in-slide-down

Data is automatically marshalled and unmarshalled in and out of the plugin. Users should mostly implement the
{py:class}`~nebulakit.core.base_task.PythonTask` API defined in Nebulakit.

Nebulakit Plugins are lazily loaded and can be released independently like libraries. We follow a convention to name the
plugin like `nebulakitplugins-*`, where \* indicates the package to be integrated into Nebulakit. For example
`nebulakitplugins-papermill` enables users to author Nebulakit tasks using [Papermill](https://papermill.readthedocs.io/en/latest/).

You can find the plugins maintained by the core Nebula team [here](https://github.com/nebulaclouds/nebulakit/tree/master/plugins).
:::

## Native Backend Plugins

Native Backend Plugins are the plugins that can be executed without any external service dependencies because the compute is
orchestrated by Nebula itself, within its provisioned Kubernetes clusters.

```{list-table}
:header-rows: 0
:widths: 20 30

* - {doc}`K8s Pods <nebulasnacks/examples/k8s_pod_plugin/index>`
  - Execute K8s pods for arbitrary workloads.
* - {doc}`K8s Cluster Dask Jobs <nebulasnacks/examples/k8s_dask_plugin/index>`
  - Run Dask jobs on a K8s Cluster.
* - {doc}`K8s Cluster Spark Jobs <nebulasnacks/examples/k8s_spark_plugin/index>`
  - Run Spark jobs on a K8s Cluster.
* - {doc}`Kubeflow PyTorch <nebulasnacks/examples/kfpytorch_plugin/index>`
  - Run distributed PyTorch training jobs using `Kubeflow`.
* - {doc}`Kubeflow TensorFlow <nebulasnacks/examples/kftensorflow_plugin/index>`
  - Run distributed TensorFlow training jobs using `Kubeflow`.
* - {doc}`MPI Operator <nebulasnacks/examples/kfmpi_plugin/index>`
  - Run distributed deep learning training jobs using Horovod and MPI.
* - {doc}`Ray Task <nebulasnacks/examples/ray_plugin/index>`
  - Run Ray jobs on a K8s Cluster.
```

(external_service_backend_plugins)=

## External Service Backend Plugins

As the term suggests, external service backend plugins relies on external services like
[AWS Sagemaker](https://aws.amazon.com/sagemaker),
[Hive](https://docs.qubole.com/en/latest/user-guide/engines/hive/index.html) or
[Snowflake](https://www.snowflake.com/) for handling the workload defined in
the Nebula task that use the respective plugin.

```{list-table}
:header-rows: 0
:widths: 20 30

* - {doc}`AWS Sagemaker: Model Training <nebulasnacks/examples/sagemaker_training_plugin/index>`
  - Train models with built-in or define your own custom algorithms.
* - {doc}`AWS Sagemaker: Pytorch Training <nebulasnacks/examples/sagemaker_pytorch_plugin/index>`
  - Train Pytorch models using Sagemaker, with support for distributed training.
* - {doc}`AWS Athena <nebulasnacks/examples/athena_plugin/index>`
  - Execute queries using AWS Athena
* - {doc}`AWS Batch <nebulasnacks/examples/aws_batch_plugin/index>`
  - Running tasks and workflows on AWS batch service
* - {doc}`Hive <nebulasnacks/examples/hive_plugin/index>`
  - Run Hive jobs in your workflows.
* - {doc}`MMCloud <nebulasnacks/examples/mmcloud_plugin/index>`
  - Execute tasks using MemVerge Memory Machine Cloud
* - {doc}`Snowflake <nebulasnacks/examples/snowflake_plugin/index>`
  - Run Snowflake jobs in your workflows.
* - {doc}`Databricks <nebulasnacks/examples/databricks_plugin/index>`
  - Run Databricks jobs in your workflows.
* - {doc}`BigQuery <nebulasnacks/examples/bigquery_plugin/index>`
  - Run BigQuery jobs in your workflows.
```

(enable-backend-plugins)=

::::{dropdown} {fa}`info-circle` Enabling Backend Plugins
:animate: fade-in-slide-down

To enable a backend plugin you have to add the `ID` of the plugin to the enabled plugins list. The `enabled-plugins` is available under the `tasks > task-plugins` section of NebulaPropeller's configuration.
The plugin configuration structure is defined [here](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller@v0.6.1/pkg/controller/nodes/task/config#TaskPluginConfig). An example of the config follows,

:::{rli} https://raw.githubusercontent.com/nebulaclouds/nebula/master/kustomize/overlays/sandbox/nebula/config/propeller/enabled_plugins.yaml
:language: yaml
:::

**Finding the `ID` of the Backend Plugin**

This is a little tricky since you have to look at the source code of the plugin to figure out the `ID`. In the case of Spark, for example, the value of `ID` is used [here](https://github.com/nebulaclouds/nebulaplugins/blob/v0.5.25/go/tasks/plugins/k8s/spark/spark.go#L424) here, defined as [spark](https://github.com/nebulaclouds/nebulaplugins/blob/v0.5.25/go/tasks/plugins/k8s/spark/spark.go#L41).

**Enabling a Specific Backend Plugin in Your Own Kustomize Generator**

Nebula uses Kustomize to generate the the deployment configuration which can be leveraged to [kustomize your own deployment](https://github.com/nebulaclouds/nebula/tree/master/kustomize).

::::

## Custom Container Tasks

Because Nebula uses executable docker containers as the smallest unit of compute, you can write custom tasks with the
{py:class}`nebulakit.ContainerTask` via the [nebulakit](https://github.com/nebulaclouds/nebulakit) SDK.

```{list-table}
:header-rows: 0
:widths: 20 30

* - {doc}`Raw Container Tasks <nebulasnacks/examples/customizing_dependencies/raw_container>`
  - Execute arbitrary containers: You can write C++ code, bash scripts and any containerized program.
```

## SDKs for Writing Tasks and Workflows

The {ref}`community <community>` would love to help you with your own ideas of building a new SDK. Currently the available SDKs are:

```{list-table}
:header-rows: 0
:widths: 20 30

* - [nebulakit](https://nebulakit.readthedocs.io)
  - The Python SDK for Nebula.
* - [nebulakit-java](https://github.com/spotify/nebulakit-java)
  - The Java/Scala SDK for Nebula.
```

## Nebula Operators

Nebula can be integrated with other orchestrators to help you leverage Nebula's
constructs natively within other orchestration tools.

```{list-table}
:header-rows: 0
:widths: 20 30

* - {doc}`Airflow <nebulasnacks/examples/airflow_plugin/index>`
  - Trigger Nebula executions from Airflow.
```

```{toctree}
:maxdepth: -1
:caption: Integrations
:hidden:

nebulasnacks/examples/sql_plugin/index
nebulasnacks/examples/greatexpectations_plugin/index
nebulasnacks/examples/papermill_plugin/index
nebulasnacks/examples/pandera_plugin/index
nebulasnacks/examples/modin_plugin/index
nebulasnacks/examples/dolt_plugin/index
nebulasnacks/examples/dbt_plugin/index
nebulasnacks/examples/whylogs_plugin/index
nebulasnacks/examples/mlflow_plugin/index
nebulasnacks/examples/onnx_plugin/index
nebulasnacks/examples/duckdb_plugin/index
nebulasnacks/examples/k8s_pod_plugin/index
nebulasnacks/examples/k8s_dask_plugin/index
nebulasnacks/examples/k8s_spark_plugin/index
nebulasnacks/examples/kfpytorch_plugin/index
nebulasnacks/examples/kftensorflow_plugin/index
nebulasnacks/examples/kfmpi_plugin/index
nebulasnacks/examples/ray_plugin/index
nebulasnacks/examples/sagemaker_training_plugin/index
nebulasnacks/examples/sagemaker_pytorch_plugin/index
nebulasnacks/examples/athena_plugin/index
nebulasnacks/examples/aws_batch_plugin/index
nebulasnacks/examples/hive_plugin/index
nebulasnacks/examples/mmcloud_plugin/index
nebulasnacks/examples/sensor/index
nebulasnacks/examples/snowflake_plugin/index
nebulasnacks/examples/databricks_plugin/index
nebulasnacks/examples/bigquery_plugin/index
nebulasnacks/examples/airflow_plugin/index
```
