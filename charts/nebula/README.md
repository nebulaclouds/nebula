# nebula

![Version: v0.1.10](https://img.shields.io/badge/Version-v0.1.10-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

A Helm chart for Nebula Sandbox

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://../nebula-core | nebula(nebula-core) | v0.1.10 |
| https://charts.bitnami.com/bitnami | contour | 7.10.1 |
| https://googlecloudplatform.github.io/spark-on-k8s-operator | sparkoperator(spark-operator) | 1.1.15 |
| https://helm.dask.org | daskoperator(dask-kubernetes-operator) | 2022.12.0 |
| https://kubernetes.github.io/dashboard/ | kubernetes-dashboard | 4.0.2 |

**NOTE:** Nebula sandbox helm chart is deprecated, From now follow the sandbox [docs](https://docs.nebula.org/en/latest/deployment/sandbox.html) for installing it on cloud

### SANDBOX INSTALLATION:
- [Install helm 3](https://helm.sh/docs/intro/install/)
- Install Nebula sandbox:

```bash
helm repo add nebula https://nebulaclouds.github.io/nebula
helm install -n nebula -f values.yaml --create-namespace nebula nebula/nebula
```

Customize your installation by changing settings in a new file `values-sandbox.yaml`.
You can use the helm diff plugin to review any value changes you've made to your values:

```bash
helm plugin install https://github.com/databus23/helm-diff
helm diff upgrade -f values-sandbox.yaml nebula .
```

Then apply your changes:
```bash
helm upgrade -f values-sandbox.yaml nebula .
```

#### Alternative: Generate raw kubernetes yaml with helm template
- `helm template --name-template=nebula-sandbox . -n nebula -f values-sandbox.yaml > nebula_generated_sandbox.yaml`
- Deploy the manifest `kubectl apply -f nebula_generated_sandbox.yaml`

- When all pods are running - run end2end tests: `kubectl apply -f ../end2end/tests/endtoend.yaml`
- If running on minikube, get nebula host using `minikube service contour -n heptio-contour --url`. And then visit `http://<HOST>/console`

### CONFIGURATION NOTES:
- The docker images, their tags and other default parameters are configured in `values.yaml` file.
- Each Nebula installation type should have separate `values-*.yaml` file: for sandbox, EKS and etc. The configuration in `values.yaml` and the chosen config `values-*.yaml` are merged when generating the deployment manifest.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| contour.affinity | object | `{}` | affinity for Contour deployment |
| contour.contour.resources | object | `{"limits":{"cpu":"100m","memory":"100Mi"},"requests":{"cpu":"10m","memory":"50Mi"}}` | Default resources requests and limits for Contour |
| contour.contour.resources.limits | object | `{"cpu":"100m","memory":"100Mi"}` | Limits are the maximum set of resources needed for this pod |
| contour.contour.resources.requests | object | `{"cpu":"10m","memory":"50Mi"}` | Requests are the minimum set of resources needed for this pod |
| contour.enabled | bool | `true` | - enable or disable Contour deployment installation |
| contour.envoy.resources | object | `{"limits":{"cpu":"100m","memory":"100Mi"},"requests":{"cpu":"10m","memory":"50Mi"}}` | Default resources requests and limits for Envoy |
| contour.envoy.resources.limits | object | `{"cpu":"100m","memory":"100Mi"}` | Limits are the maximum set of resources needed for this pod |
| contour.envoy.resources.requests | object | `{"cpu":"10m","memory":"50Mi"}` | Requests are the minimum set of resources needed for this pod |
| contour.envoy.service.nodePorts.http | int | `30081` |  |
| contour.envoy.service.ports.http | int | `80` |  |
| contour.envoy.service.type | string | `"NodePort"` |  |
| contour.nodeSelector | object | `{}` | nodeSelector for Contour deployment |
| contour.podAnnotations | object | `{}` | Annotations for Contour pods |
| contour.replicaCount | int | `1` | Replicas count for Contour deployment |
| contour.serviceAccountAnnotations | object | `{}` | Annotations for ServiceAccount attached to Contour pods |
| contour.tolerations | list | `[]` | tolerations for Contour deployment |
| daskoperator | object | `{"enabled":false}` | Optional: Dask Plugin using the Dask Operator |
| daskoperator.enabled | bool | `false` | - enable or disable the dask operator deployment installation |
| nebula | object | `{"cluster_resource_manager":{"config":{"cluster_resources":{"customData":[{"production":[{"projectQuotaCpu":{"value":"5"}},{"projectQuotaMemory":{"value":"4000Mi"}}]},{"staging":[{"projectQuotaCpu":{"value":"2"}},{"projectQuotaMemory":{"value":"3000Mi"}}]},{"development":[{"projectQuotaCpu":{"value":"4"}},{"projectQuotaMemory":{"value":"3000Mi"}}]}],"refresh":"5m","refreshInterval":"5m","standaloneDeployment":false,"templatePath":"/etc/nebula/clusterresource/templates"}},"enabled":true,"service_account_name":"nebulaadmin","templates":[{"key":"aa_namespace","value":"apiVersion: v1\nkind: Namespace\nmetadata:\n  name: {{ namespace }}\nspec:\n  finalizers:\n  - kubernetes\n"},{"key":"ab_project_resource_quota","value":"apiVersion: v1\nkind: ResourceQuota\nmetadata:\n  name: project-quota\n  namespace: {{ namespace }}\nspec:\n  hard:\n    limits.cpu: {{ projectQuotaCpu }}\n    limits.memory: {{ projectQuotaMemory }}\n"}]},"common":{"databaseSecret":{"name":"","secretManifest":{}},"nebulaNamespaceTemplate":{"enabled":false},"ingress":{"albSSLRedirect":false,"annotations":{"nginx.ingress.kubernetes.io/app-root":"/console"},"enabled":true,"host":"","separateGrpcIngress":false,"separateGrpcIngressAnnotations":{"nginx.ingress.kubernetes.io/backend-protocol":"GRPC"},"tls":{"enabled":false},"webpackHMR":true}},"configmap":{"adminServer":{"auth":{"appAuth":{"thirdPartyConfig":{"nebulaClient":{"clientId":"nebulactl","redirectUri":"http://localhost:53593/callback","scopes":["offline","all"]}}},"authorizedUris":["https://localhost:30081","http://nebulaadmin:80","http://nebulaadmin.nebula.svc.cluster.local:80"],"userAuth":{"openId":{"baseUrl":"https://accounts.google.com","clientId":"657465813211-6eog7ek7li5k7i7fvgv2921075063hpe.apps.googleusercontent.com","scopes":["profile","openid"]}}},"nebulaadmin":{"eventVersion":2,"metadataStoragePrefix":["metadata","admin"],"metricsScope":"nebula:","profilerPort":10254,"roleNameKey":"iam.amazonaws.com/role","testing":{"host":"http://nebulaadmin"}},"server":{"grpcPort":8089,"httpPort":8088,"security":{"allowCors":true,"allowedHeaders":["Content-Type","nebula-authorization"],"allowedOrigins":["*"],"secure":false,"useAuth":false}}},"catalog":{"catalog-cache":{"endpoint":"datacatalog:89","insecure":true,"type":"datacatalog"}},"console":{"BASE_URL":"/console","CONFIG_DIR":"/etc/nebula/config"},"copilot":{"plugins":{"k8s":{"co-pilot":{"image":"cr.nebula.org/nebulaclouds/nebulacopilot:v1.10.6","name":"nebula-copilot-","start-timeout":"30s"}}}},"core":{"propeller":{"downstream-eval-duration":"30s","enable-admin-launcher":true,"leader-election":{"enabled":true,"lease-duration":"15s","lock-config-map":{"name":"propeller-leader","namespace":"nebula"},"renew-deadline":"10s","retry-period":"2s"},"limit-namespace":"all","max-workflow-retries":30,"metadata-prefix":"metadata/propeller","metrics-prefix":"nebula","prof-port":10254,"queue":{"batch-size":-1,"batching-interval":"2s","queue":{"base-delay":"5s","capacity":1000,"max-delay":"120s","rate":100,"type":"maxof"},"sub-queue":{"capacity":100,"rate":10,"type":"bucket"},"type":"batch"},"rawoutput-prefix":"s3://my-s3-bucket/","workers":4,"workflow-reeval-duration":"30s"},"webhook":{"certDir":"/etc/webhook/certs","serviceName":"nebula-pod-webhook"}},"datacatalogServer":{"application":{"grpcPort":8089,"grpcServerReflection":true,"httpPort":8080},"datacatalog":{"metrics-scope":"datacatalog","profiler-port":10254,"storage-prefix":"metadata/datacatalog"}},"domain":{"domains":[{"id":"development","name":"development"},{"id":"staging","name":"staging"},{"id":"production","name":"production"}]},"enabled_plugins":{"tasks":{"task-plugins":{"default-for-task-types":{"bigquery_query_job_task":"agent-service","container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array","agent-service"]}}},"k8s":{"plugins":{"agent-service":{"defaultAgent":{"endpoint":"dns:///nebulaagent.nebula.svc.cluster.local:8000","insecure":true},"supportedTaskTypes":["bigquery_query_job_task"]},"k8s":{"default-cpus":"100m","default-env-vars":[{"NEBULA_AWS_ENDPOINT":"http://minio.nebula:9000"},{"NEBULA_AWS_ACCESS_KEY_ID":"minio"},{"NEBULA_AWS_SECRET_ACCESS_KEY":"miniostorage"}],"default-memory":"200Mi"}}},"logger":{"logger":{"level":5,"show-source":true}},"remoteData":{"remoteData":{"region":"us-east-1","scheme":"local","signedUrls":{"durationMinutes":3}}},"resource_manager":{"propeller":{"resourcemanager":{"redis":null,"type":"noop"}}},"task_logs":{"plugins":{"logs":{"cloudwatch-enabled":false,"kubernetes-enabled":true,"kubernetes-template-uri":"http://localhost:30082/#/log/{{ \"{{\" }} .namespace {{ \"}}\" }}/{{ \"{{\" }} .podName {{ \"}}\" }}/pod?namespace={{ \"{{\" }} .namespace {{ \"}}\" }}"}}},"task_resource_defaults":{"task_resources":{"defaults":{"cpu":"100m","memory":"200Mi","storage":"5Mi"},"limits":{"cpu":2,"gpu":1,"memory":"1Gi","storage":"20Mi"}}}},"datacatalog":{"affinity":{},"configPath":"/etc/datacatalog/config/*.yaml","image":{"pullPolicy":"IfNotPresent","repository":"cr.nebula.org/nebulaclouds/datacatalog","tag":"v1.10.6"},"nodeSelector":{},"podAnnotations":{},"replicaCount":1,"resources":{"limits":{"cpu":"500m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}},"service":{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"type":"NodePort"},"serviceAccount":{"annotations":{},"create":true,"imagePullSecrets":[]},"tolerations":[]},"db":{"admin":{"database":{"dbname":"nebulaadmin","host":"postgres","port":5432,"username":"postgres"}},"datacatalog":{"database":{"dbname":"datacatalog","host":"postgres","port":5432,"username":"postgres"}}},"deployRedoc":true,"nebulaadmin":{"additionalVolumeMounts":[],"additionalVolumes":[],"affinity":{},"configPath":"/etc/nebula/config/*.yaml","env":[],"image":{"pullPolicy":"IfNotPresent","repository":"cr.nebula.org/nebulaclouds/nebulaadmin","tag":"v1.10.6"},"initialProjects":["nebulasnacks","nebulatester","nebulaexamples"],"nodeSelector":{},"podAnnotations":{},"replicaCount":1,"resources":{"limits":{"cpu":"250m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}},"secrets":{},"service":{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"loadBalancerSourceRanges":[],"type":"ClusterIP"},"serviceAccount":{"annotations":{},"create":true,"imagePullSecrets":[]},"tolerations":[]},"nebulaconsole":{"affinity":{},"ga":{"enabled":true,"tracking_id":"G-0QW4DJWJ20"},"image":{"pullPolicy":"IfNotPresent","repository":"cr.nebula.org/nebulaclouds/nebulaconsole","tag":"v1.10.2"},"nodeSelector":{},"podAnnotations":{},"replicaCount":1,"resources":{"limits":{"cpu":"500m","memory":"275Mi"},"requests":{"cpu":"10m","memory":"250Mi"}},"service":{"annotations":{},"type":"ClusterIP"},"tolerations":[]},"nebulapropeller":{"affinity":{},"cacheSizeMbs":0,"configPath":"/etc/nebula/config/*.yaml","image":{"pullPolicy":"IfNotPresent","repository":"cr.nebula.org/nebulaclouds/nebulapropeller","tag":"v1.10.6"},"manager":false,"nodeSelector":{},"podAnnotations":{},"replicaCount":1,"resources":{"limits":{"cpu":"200m","ephemeral-storage":"100Mi","memory":"200Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}},"serviceAccount":{"annotations":{},"create":true,"imagePullSecrets":[]},"tolerations":[]},"nebulascheduler":{"affinity":{},"configPath":"/etc/nebula/config/*.yaml","image":{"pullPolicy":"IfNotPresent","repository":"cr.nebula.org/nebulaclouds/nebulascheduler","tag":"v1.10.6"},"nodeSelector":{},"podAnnotations":{},"resources":{"limits":{"cpu":"250m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}},"secrets":{},"serviceAccount":{"annotations":{},"create":true,"imagePullSecrets":[]},"tolerations":[]},"storage":{"bucketName":"my-s3-bucket","custom":{},"gcs":null,"s3":{"region":"us-east-1"},"type":"sandbox"},"webhook":{"enabled":true,"service":{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"type":"ClusterIP"},"serviceAccount":{"annotations":{},"create":true,"imagePullSecrets":[]}},"workflow_notifications":{"config":{},"enabled":false},"workflow_scheduler":{"enabled":true,"type":"native"}}` | ------------------------------------------------------------------- Core System settings This section consists of Core components of Nebula and their deployment settings. This includes NebulaAdmin service, Datacatalog, NebulaPropeller and Nebulaconsole |
| nebula.cluster_resource_manager | object | `{"config":{"cluster_resources":{"customData":[{"production":[{"projectQuotaCpu":{"value":"5"}},{"projectQuotaMemory":{"value":"4000Mi"}}]},{"staging":[{"projectQuotaCpu":{"value":"2"}},{"projectQuotaMemory":{"value":"3000Mi"}}]},{"development":[{"projectQuotaCpu":{"value":"4"}},{"projectQuotaMemory":{"value":"3000Mi"}}]}],"refresh":"5m","refreshInterval":"5m","standaloneDeployment":false,"templatePath":"/etc/nebula/clusterresource/templates"}},"enabled":true,"service_account_name":"nebulaadmin","templates":[{"key":"aa_namespace","value":"apiVersion: v1\nkind: Namespace\nmetadata:\n  name: {{ namespace }}\nspec:\n  finalizers:\n  - kubernetes\n"},{"key":"ab_project_resource_quota","value":"apiVersion: v1\nkind: ResourceQuota\nmetadata:\n  name: project-quota\n  namespace: {{ namespace }}\nspec:\n  hard:\n    limits.cpu: {{ projectQuotaCpu }}\n    limits.memory: {{ projectQuotaMemory }}\n"}]}` | Configuration for the Cluster resource manager component. This is an optional component, that enables automatic cluster configuration. This is useful to set default quotas, manage namespaces etc that map to a project/domain |
| nebula.cluster_resource_manager.config.cluster_resources | object | `{"customData":[{"production":[{"projectQuotaCpu":{"value":"5"}},{"projectQuotaMemory":{"value":"4000Mi"}}]},{"staging":[{"projectQuotaCpu":{"value":"2"}},{"projectQuotaMemory":{"value":"3000Mi"}}]},{"development":[{"projectQuotaCpu":{"value":"4"}},{"projectQuotaMemory":{"value":"3000Mi"}}]}],"refresh":"5m","refreshInterval":"5m","standaloneDeployment":false,"templatePath":"/etc/nebula/clusterresource/templates"}` | ClusterResource parameters Refer to the [structure](https://pkg.go.dev/github.com/lyft/nebulaadmin@v0.3.37/pkg/runtime/interfaces#ClusterResourceConfig) to customize. |
| nebula.cluster_resource_manager.config.cluster_resources.standaloneDeployment | bool | `false` | Starts the cluster resource manager in standalone mode with requisite auth credentials to call nebulaadmin service endpoints |
| nebula.cluster_resource_manager.enabled | bool | `true` | Enables the Cluster resource manager component |
| nebula.cluster_resource_manager.service_account_name | string | `"nebulaadmin"` | Service account name to run with |
| nebula.cluster_resource_manager.templates | list | `[{"key":"aa_namespace","value":"apiVersion: v1\nkind: Namespace\nmetadata:\n  name: {{ namespace }}\nspec:\n  finalizers:\n  - kubernetes\n"},{"key":"ab_project_resource_quota","value":"apiVersion: v1\nkind: ResourceQuota\nmetadata:\n  name: project-quota\n  namespace: {{ namespace }}\nspec:\n  hard:\n    limits.cpu: {{ projectQuotaCpu }}\n    limits.memory: {{ projectQuotaMemory }}\n"}]` | Resource templates that should be applied |
| nebula.cluster_resource_manager.templates[0] | object | `{"key":"aa_namespace","value":"apiVersion: v1\nkind: Namespace\nmetadata:\n  name: {{ namespace }}\nspec:\n  finalizers:\n  - kubernetes\n"}` | Template for namespaces resources |
| nebula.common | object | `{"databaseSecret":{"name":"","secretManifest":{}},"nebulaNamespaceTemplate":{"enabled":false},"ingress":{"albSSLRedirect":false,"annotations":{"nginx.ingress.kubernetes.io/app-root":"/console"},"enabled":true,"host":"","separateGrpcIngress":false,"separateGrpcIngressAnnotations":{"nginx.ingress.kubernetes.io/backend-protocol":"GRPC"},"tls":{"enabled":false},"webpackHMR":true}}` | ----------------------------------------------  COMMON SETTINGS  |
| nebula.common.databaseSecret.name | string | `""` | Specify name of K8s Secret which contains Database password. Leave it empty if you don't need this Secret |
| nebula.common.databaseSecret.secretManifest | object | `{}` | Specify your Secret (with sensitive data) or pseudo-manifest (without sensitive data). See https://github.com/godaddy/kubernetes-external-secrets |
| nebula.common.nebulaNamespaceTemplate.enabled | bool | `false` | - Enable or disable creating Nebula namespace in template. Enable when using helm as template-engine only. Disable when using `helm install ...`. |
| nebula.common.ingress.albSSLRedirect | bool | `false` | - albSSLRedirect adds a special route for ssl redirect. Only useful in combination with the AWS LoadBalancer Controller. |
| nebula.common.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/app-root":"/console"}` | - Ingress annotations applied to both HTTP and GRPC ingresses. |
| nebula.common.ingress.enabled | bool | `true` | - Enable or disable creating Ingress for Nebula. Relevant to disable when using e.g. Istio as ingress controller. |
| nebula.common.ingress.host | string | `""` | - Ingress hostname |
| nebula.common.ingress.separateGrpcIngress | bool | `false` | - separateGrpcIngress puts GRPC routes into a separate ingress if true. Required for certain ingress controllers like nginx. |
| nebula.common.ingress.separateGrpcIngressAnnotations | object | `{"nginx.ingress.kubernetes.io/backend-protocol":"GRPC"}` | - Extra Ingress annotations applied only to the GRPC ingress. Only makes sense if `separateGrpcIngress` is enabled. |
| nebula.common.ingress.tls | object | `{"enabled":false}` | - TLS Settings |
| nebula.common.ingress.webpackHMR | bool | `true` | - Enable or disable HMR route to nebulaconsole. This is useful only for frontend development. |
| nebula.configmap | object | `{"adminServer":{"auth":{"appAuth":{"thirdPartyConfig":{"nebulaClient":{"clientId":"nebulactl","redirectUri":"http://localhost:53593/callback","scopes":["offline","all"]}}},"authorizedUris":["https://localhost:30081","http://nebulaadmin:80","http://nebulaadmin.nebula.svc.cluster.local:80"],"userAuth":{"openId":{"baseUrl":"https://accounts.google.com","clientId":"657465813211-6eog7ek7li5k7i7fvgv2921075063hpe.apps.googleusercontent.com","scopes":["profile","openid"]}}},"nebulaadmin":{"eventVersion":2,"metadataStoragePrefix":["metadata","admin"],"metricsScope":"nebula:","profilerPort":10254,"roleNameKey":"iam.amazonaws.com/role","testing":{"host":"http://nebulaadmin"}},"server":{"grpcPort":8089,"httpPort":8088,"security":{"allowCors":true,"allowedHeaders":["Content-Type","nebula-authorization"],"allowedOrigins":["*"],"secure":false,"useAuth":false}}},"catalog":{"catalog-cache":{"endpoint":"datacatalog:89","insecure":true,"type":"datacatalog"}},"console":{"BASE_URL":"/console","CONFIG_DIR":"/etc/nebula/config"},"copilot":{"plugins":{"k8s":{"co-pilot":{"image":"cr.nebula.org/nebulaclouds/nebulacopilot:v1.10.6","name":"nebula-copilot-","start-timeout":"30s"}}}},"core":{"propeller":{"downstream-eval-duration":"30s","enable-admin-launcher":true,"leader-election":{"enabled":true,"lease-duration":"15s","lock-config-map":{"name":"propeller-leader","namespace":"nebula"},"renew-deadline":"10s","retry-period":"2s"},"limit-namespace":"all","max-workflow-retries":30,"metadata-prefix":"metadata/propeller","metrics-prefix":"nebula","prof-port":10254,"queue":{"batch-size":-1,"batching-interval":"2s","queue":{"base-delay":"5s","capacity":1000,"max-delay":"120s","rate":100,"type":"maxof"},"sub-queue":{"capacity":100,"rate":10,"type":"bucket"},"type":"batch"},"rawoutput-prefix":"s3://my-s3-bucket/","workers":4,"workflow-reeval-duration":"30s"},"webhook":{"certDir":"/etc/webhook/certs","serviceName":"nebula-pod-webhook"}},"datacatalogServer":{"application":{"grpcPort":8089,"grpcServerReflection":true,"httpPort":8080},"datacatalog":{"metrics-scope":"datacatalog","profiler-port":10254,"storage-prefix":"metadata/datacatalog"}},"domain":{"domains":[{"id":"development","name":"development"},{"id":"staging","name":"staging"},{"id":"production","name":"production"}]},"enabled_plugins":{"tasks":{"task-plugins":{"default-for-task-types":{"bigquery_query_job_task":"agent-service","container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array","agent-service"]}}},"k8s":{"plugins":{"agent-service":{"defaultAgent":{"endpoint":"dns:///nebulaagent.nebula.svc.cluster.local:8000","insecure":true},"supportedTaskTypes":["bigquery_query_job_task"]},"k8s":{"default-cpus":"100m","default-env-vars":[{"NEBULA_AWS_ENDPOINT":"http://minio.nebula:9000"},{"NEBULA_AWS_ACCESS_KEY_ID":"minio"},{"NEBULA_AWS_SECRET_ACCESS_KEY":"miniostorage"}],"default-memory":"200Mi"}}},"logger":{"logger":{"level":5,"show-source":true}},"remoteData":{"remoteData":{"region":"us-east-1","scheme":"local","signedUrls":{"durationMinutes":3}}},"resource_manager":{"propeller":{"resourcemanager":{"redis":null,"type":"noop"}}},"task_logs":{"plugins":{"logs":{"cloudwatch-enabled":false,"kubernetes-enabled":true,"kubernetes-template-uri":"http://localhost:30082/#/log/{{ \"{{\" }} .namespace {{ \"}}\" }}/{{ \"{{\" }} .podName {{ \"}}\" }}/pod?namespace={{ \"{{\" }} .namespace {{ \"}}\" }}"}}},"task_resource_defaults":{"task_resources":{"defaults":{"cpu":"100m","memory":"200Mi","storage":"5Mi"},"limits":{"cpu":2,"gpu":1,"memory":"1Gi","storage":"20Mi"}}}}` | -----------------------------------------------------------------  CONFIGMAPS SETTINGS  |
| nebula.configmap.adminServer | object | `{"auth":{"appAuth":{"thirdPartyConfig":{"nebulaClient":{"clientId":"nebulactl","redirectUri":"http://localhost:53593/callback","scopes":["offline","all"]}}},"authorizedUris":["https://localhost:30081","http://nebulaadmin:80","http://nebulaadmin.nebula.svc.cluster.local:80"],"userAuth":{"openId":{"baseUrl":"https://accounts.google.com","clientId":"657465813211-6eog7ek7li5k7i7fvgv2921075063hpe.apps.googleusercontent.com","scopes":["profile","openid"]}}},"nebulaadmin":{"eventVersion":2,"metadataStoragePrefix":["metadata","admin"],"metricsScope":"nebula:","profilerPort":10254,"roleNameKey":"iam.amazonaws.com/role","testing":{"host":"http://nebulaadmin"}},"server":{"grpcPort":8089,"httpPort":8088,"security":{"allowCors":true,"allowedHeaders":["Content-Type","nebula-authorization"],"allowedOrigins":["*"],"secure":false,"useAuth":false}}}` | NebulaAdmin server configuration |
| nebula.configmap.adminServer.auth | object | `{"appAuth":{"thirdPartyConfig":{"nebulaClient":{"clientId":"nebulactl","redirectUri":"http://localhost:53593/callback","scopes":["offline","all"]}}},"authorizedUris":["https://localhost:30081","http://nebulaadmin:80","http://nebulaadmin.nebula.svc.cluster.local:80"],"userAuth":{"openId":{"baseUrl":"https://accounts.google.com","clientId":"657465813211-6eog7ek7li5k7i7fvgv2921075063hpe.apps.googleusercontent.com","scopes":["profile","openid"]}}}` | Authentication configuration |
| nebula.configmap.adminServer.server.security.secure | bool | `false` | Controls whether to serve requests over SSL/TLS. |
| nebula.configmap.adminServer.server.security.useAuth | bool | `false` | Controls whether to enforce authentication. Follow the guide in https://docs.nebula.org/ on how to setup authentication. |
| nebula.configmap.catalog | object | `{"catalog-cache":{"endpoint":"datacatalog:89","insecure":true,"type":"datacatalog"}}` | Catalog Client configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/catalog#Config) Additional advanced Catalog configuration [here](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/pluginmachinery/catalog#Config) |
| nebula.configmap.console | object | `{"BASE_URL":"/console","CONFIG_DIR":"/etc/nebula/config"}` | Configuration for Nebula console UI |
| nebula.configmap.copilot | object | `{"plugins":{"k8s":{"co-pilot":{"image":"cr.nebula.org/nebulaclouds/nebulacopilot:v1.10.6","name":"nebula-copilot-","start-timeout":"30s"}}}}` | Copilot configuration |
| nebula.configmap.copilot.plugins.k8s.co-pilot | object | `{"image":"cr.nebula.org/nebulaclouds/nebulacopilot:v1.10.6","name":"nebula-copilot-","start-timeout":"30s"}` | Structure documented [here](https://pkg.go.dev/github.com/lyft/nebulaplugins@v0.5.28/go/tasks/pluginmachinery/nebulak8s/config#NebulaCoPilotConfig) |
| nebula.configmap.core | object | `{"propeller":{"downstream-eval-duration":"30s","enable-admin-launcher":true,"leader-election":{"enabled":true,"lease-duration":"15s","lock-config-map":{"name":"propeller-leader","namespace":"nebula"},"renew-deadline":"10s","retry-period":"2s"},"limit-namespace":"all","max-workflow-retries":30,"metadata-prefix":"metadata/propeller","metrics-prefix":"nebula","prof-port":10254,"queue":{"batch-size":-1,"batching-interval":"2s","queue":{"base-delay":"5s","capacity":1000,"max-delay":"120s","rate":100,"type":"maxof"},"sub-queue":{"capacity":100,"rate":10,"type":"bucket"},"type":"batch"},"rawoutput-prefix":"s3://my-s3-bucket/","workers":4,"workflow-reeval-duration":"30s"},"webhook":{"certDir":"/etc/webhook/certs","serviceName":"nebula-pod-webhook"}}` | Core propeller configuration |
| nebula.configmap.core.propeller | object | `{"downstream-eval-duration":"30s","enable-admin-launcher":true,"leader-election":{"enabled":true,"lease-duration":"15s","lock-config-map":{"name":"propeller-leader","namespace":"nebula"},"renew-deadline":"10s","retry-period":"2s"},"limit-namespace":"all","max-workflow-retries":30,"metadata-prefix":"metadata/propeller","metrics-prefix":"nebula","prof-port":10254,"queue":{"batch-size":-1,"batching-interval":"2s","queue":{"base-delay":"5s","capacity":1000,"max-delay":"120s","rate":100,"type":"maxof"},"sub-queue":{"capacity":100,"rate":10,"type":"bucket"},"type":"batch"},"rawoutput-prefix":"s3://my-s3-bucket/","workers":4,"workflow-reeval-duration":"30s"}` | follows the structure specified [here](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/config). |
| nebula.configmap.datacatalogServer | object | `{"application":{"grpcPort":8089,"grpcServerReflection":true,"httpPort":8080},"datacatalog":{"metrics-scope":"datacatalog","profiler-port":10254,"storage-prefix":"metadata/datacatalog"}}` | Datacatalog server config |
| nebula.configmap.domain | object | `{"domains":[{"id":"development","name":"development"},{"id":"staging","name":"staging"},{"id":"production","name":"production"}]}` | Domains configuration for Nebula projects. This enables the specified number of domains across all projects in Nebula. |
| nebula.configmap.enabled_plugins.tasks | object | `{"task-plugins":{"default-for-task-types":{"bigquery_query_job_task":"agent-service","container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array","agent-service"]}}` | Tasks specific configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#GetConfig) |
| nebula.configmap.enabled_plugins.tasks.task-plugins | object | `{"default-for-task-types":{"bigquery_query_job_task":"agent-service","container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array","agent-service"]}` | Plugins configuration, [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#TaskPluginConfig) |
| nebula.configmap.enabled_plugins.tasks.task-plugins.enabled-plugins | list | `["container","sidecar","k8s-array","agent-service"]` | [Enabled Plugins](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/config#Config). Enable sagemaker*, athena if you install the backend plugins |
| nebula.configmap.k8s | object | `{"plugins":{"agent-service":{"defaultAgent":{"endpoint":"dns:///nebulaagent.nebula.svc.cluster.local:8000","insecure":true},"supportedTaskTypes":["bigquery_query_job_task"]},"k8s":{"default-cpus":"100m","default-env-vars":[{"NEBULA_AWS_ENDPOINT":"http://minio.nebula:9000"},{"NEBULA_AWS_ACCESS_KEY_ID":"minio"},{"NEBULA_AWS_SECRET_ACCESS_KEY":"miniostorage"}],"default-memory":"200Mi"}}}` | Kubernetes specific Nebula configuration |
| nebula.configmap.k8s.plugins.k8s | object | `{"default-cpus":"100m","default-env-vars":[{"NEBULA_AWS_ENDPOINT":"http://minio.nebula:9000"},{"NEBULA_AWS_ACCESS_KEY_ID":"minio"},{"NEBULA_AWS_SECRET_ACCESS_KEY":"miniostorage"}],"default-memory":"200Mi"}` | Configuration section for all K8s specific plugins [Configuration structure](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/pluginmachinery/nebulak8s/config) |
| nebula.configmap.logger | object | `{"logger":{"level":5,"show-source":true}}` | Logger configuration |
| nebula.configmap.resource_manager | object | `{"propeller":{"resourcemanager":{"redis":null,"type":"noop"}}}` | Resource manager configuration |
| nebula.configmap.resource_manager.propeller | object | `{"resourcemanager":{"redis":null,"type":"noop"}}` | resource manager configuration |
| nebula.configmap.task_logs | object | `{"plugins":{"logs":{"cloudwatch-enabled":false,"kubernetes-enabled":true,"kubernetes-template-uri":"http://localhost:30082/#/log/{{ \"{{\" }} .namespace {{ \"}}\" }}/{{ \"{{\" }} .podName {{ \"}}\" }}/pod?namespace={{ \"{{\" }} .namespace {{ \"}}\" }}"}}}` | Section that configures how the Task logs are displayed on the UI. This has to be changed based on your actual logging provider. Refer to [structure](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/logs#LogConfig) to understand how to configure various logging engines |
| nebula.configmap.task_logs.plugins.logs.cloudwatch-enabled | bool | `false` | One option is to enable cloudwatch logging for EKS, update the region and log group accordingly |
| nebula.configmap.task_resource_defaults | object | `{"task_resources":{"defaults":{"cpu":"100m","memory":"200Mi","storage":"5Mi"},"limits":{"cpu":2,"gpu":1,"memory":"1Gi","storage":"20Mi"}}}` | Task default resources configuration Refer to the full [structure](https://pkg.go.dev/github.com/lyft/nebulaadmin@v0.3.37/pkg/runtime/interfaces#TaskResourceConfiguration). |
| nebula.configmap.task_resource_defaults.task_resources | object | `{"defaults":{"cpu":"100m","memory":"200Mi","storage":"5Mi"},"limits":{"cpu":2,"gpu":1,"memory":"1Gi","storage":"20Mi"}}` | Task default resources parameters |
| nebula.datacatalog.affinity | object | `{}` | affinity for Datacatalog deployment |
| nebula.datacatalog.configPath | string | `"/etc/datacatalog/config/*.yaml"` | Default regex string for searching configuration files |
| nebula.datacatalog.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| nebula.datacatalog.image.repository | string | `"cr.nebula.org/nebulaclouds/datacatalog"` | Docker image for Datacatalog deployment |
| nebula.datacatalog.image.tag | string | `"v1.10.6"` | Docker image tag |
| nebula.datacatalog.nodeSelector | object | `{}` | nodeSelector for Datacatalog deployment |
| nebula.datacatalog.podAnnotations | object | `{}` | Annotations for Datacatalog pods |
| nebula.datacatalog.replicaCount | int | `1` | Replicas count for Datacatalog deployment |
| nebula.datacatalog.resources | object | `{"limits":{"cpu":"500m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}}` | Default resources requests and limits for Datacatalog deployment |
| nebula.datacatalog.service | object | `{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"type":"NodePort"}` | Service settings for Datacatalog |
| nebula.datacatalog.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for Datacatalog |
| nebula.datacatalog.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to Datacatalog pods |
| nebula.datacatalog.serviceAccount.create | bool | `true` | Should a service account be created for Datacatalog |
| nebula.datacatalog.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebula.datacatalog.tolerations | list | `[]` | tolerations for Datacatalog deployment |
| nebula.nebulaadmin.affinity | object | `{}` | affinity for Nebulaadmin deployment |
| nebula.nebulaadmin.configPath | string | `"/etc/nebula/config/*.yaml"` | Default regex string for searching configuration files |
| nebula.nebulaadmin.env | list | `[]` | Additional nebulaadmin container environment variables  e.g. SendGrid's API key  - name: SENDGRID_API_KEY    value: "<your sendgrid api key>"  e.g. secret environment variable (you can combine it with .additionalVolumes): - name: SENDGRID_API_KEY   valueFrom:     secretKeyRef:       name: sendgrid-secret       key: api_key |
| nebula.nebulaadmin.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| nebula.nebulaadmin.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulaadmin"` | Docker image for Nebulaadmin deployment |
| nebula.nebulaadmin.image.tag | string | `"v1.10.6"` | Docker image tag |
| nebula.nebulaadmin.initialProjects | list | `["nebulasnacks","nebulatester","nebulaexamples"]` | Initial projects to create |
| nebula.nebulaadmin.nodeSelector | object | `{}` | nodeSelector for Nebulaadmin deployment |
| nebula.nebulaadmin.podAnnotations | object | `{}` | Annotations for Nebulaadmin pods |
| nebula.nebulaadmin.replicaCount | int | `1` | Replicas count for Nebulaadmin deployment |
| nebula.nebulaadmin.resources | object | `{"limits":{"cpu":"250m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}}` | Default resources requests and limits for Nebulaadmin deployment |
| nebula.nebulaadmin.service | object | `{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"loadBalancerSourceRanges":[],"type":"ClusterIP"}` | Service settings for Nebulaadmin |
| nebula.nebulaadmin.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for NebulaAdmin |
| nebula.nebulaadmin.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to Nebulaadmin pods |
| nebula.nebulaadmin.serviceAccount.create | bool | `true` | Should a service account be created for nebulaadmin |
| nebula.nebulaadmin.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebula.nebulaadmin.tolerations | list | `[]` | tolerations for Nebulaadmin deployment |
| nebula.nebulaconsole.affinity | object | `{}` | affinity for Nebulaconsole deployment |
| nebula.nebulaconsole.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| nebula.nebulaconsole.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulaconsole"` | Docker image for Nebulaconsole deployment |
| nebula.nebulaconsole.image.tag | string | `"v1.10.2"` | Docker image tag |
| nebula.nebulaconsole.nodeSelector | object | `{}` | nodeSelector for Nebulaconsole deployment |
| nebula.nebulaconsole.podAnnotations | object | `{}` | Annotations for Nebulaconsole pods |
| nebula.nebulaconsole.replicaCount | int | `1` | Replicas count for Nebulaconsole deployment |
| nebula.nebulaconsole.resources | object | `{"limits":{"cpu":"500m","memory":"275Mi"},"requests":{"cpu":"10m","memory":"250Mi"}}` | Default resources requests and limits for Nebulaconsole deployment |
| nebula.nebulaconsole.service | object | `{"annotations":{},"type":"ClusterIP"}` | Service settings for Nebulaconsole |
| nebula.nebulaconsole.tolerations | list | `[]` | tolerations for Nebulaconsole deployment |
| nebula.nebulapropeller.affinity | object | `{}` | affinity for Nebulapropeller deployment |
| nebula.nebulapropeller.configPath | string | `"/etc/nebula/config/*.yaml"` | Default regex string for searching configuration files |
| nebula.nebulapropeller.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| nebula.nebulapropeller.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulapropeller"` | Docker image for Nebulapropeller deployment |
| nebula.nebulapropeller.image.tag | string | `"v1.10.6"` | Docker image tag |
| nebula.nebulapropeller.nodeSelector | object | `{}` | nodeSelector for Nebulapropeller deployment |
| nebula.nebulapropeller.podAnnotations | object | `{}` | Annotations for Nebulapropeller pods |
| nebula.nebulapropeller.replicaCount | int | `1` | Replicas count for Nebulapropeller deployment |
| nebula.nebulapropeller.resources | object | `{"limits":{"cpu":"200m","ephemeral-storage":"100Mi","memory":"200Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}}` | Default resources requests and limits for Nebulapropeller deployment |
| nebula.nebulapropeller.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for NebulaPropeller |
| nebula.nebulapropeller.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to NebulaPropeller pods |
| nebula.nebulapropeller.serviceAccount.create | bool | `true` | Should a service account be created for NebulaPropeller |
| nebula.nebulapropeller.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebula.nebulapropeller.tolerations | list | `[]` | tolerations for Nebulapropeller deployment |
| nebula.nebulascheduler.affinity | object | `{}` | affinity for Nebulascheduler deployment |
| nebula.nebulascheduler.configPath | string | `"/etc/nebula/config/*.yaml"` | Default regex string for searching configuration files |
| nebula.nebulascheduler.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| nebula.nebulascheduler.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulascheduler"` | Docker image for Nebulascheduler deployment |
| nebula.nebulascheduler.image.tag | string | `"v1.10.6"` | Docker image tag |
| nebula.nebulascheduler.nodeSelector | object | `{}` | nodeSelector for Nebulascheduler deployment |
| nebula.nebulascheduler.podAnnotations | object | `{}` | Annotations for Nebulascheduler pods |
| nebula.nebulascheduler.resources | object | `{"limits":{"cpu":"250m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}}` | Default resources requests and limits for Nebulascheduler deployment |
| nebula.nebulascheduler.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for Nebulascheduler |
| nebula.nebulascheduler.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to Nebulascheduler pods |
| nebula.nebulascheduler.serviceAccount.create | bool | `true` | Should a service account be created for Nebulascheduler |
| nebula.nebulascheduler.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebula.nebulascheduler.tolerations | list | `[]` | tolerations for Nebulascheduler deployment |
| nebula.storage | object | `{"bucketName":"my-s3-bucket","custom":{},"gcs":null,"s3":{"region":"us-east-1"},"type":"sandbox"}` | ----------------------------------------------------  STORAGE SETTINGS  |
| nebula.storage.bucketName | string | `"my-s3-bucket"` | bucketName defines the storage bucket nebula will use. Required for all types except for sandbox. |
| nebula.storage.custom | object | `{}` | Settings for storage type custom. See https://github.com/graymeta/stow for supported storage providers/settings. |
| nebula.storage.gcs | string | `nil` | settings for storage type gcs |
| nebula.storage.s3 | object | `{"region":"us-east-1"}` | settings for storage type s3 |
| nebula.storage.type | string | `"sandbox"` | Sets the storage type. Supported values are sandbox, s3, gcs and custom. |
| nebula.webhook.enabled | bool | `true` | enable or disable secrets webhook |
| nebula.webhook.service | object | `{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"type":"ClusterIP"}` | Service settings for the webhook |
| nebula.webhook.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for the webhook |
| nebula.webhook.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to the webhook |
| nebula.webhook.serviceAccount.create | bool | `true` | Should a service account be created for the webhook |
| nebula.webhook.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebula.workflow_notifications | object | `{"config":{},"enabled":false}` | **Optional Component** Workflow notifications module is an optional dependency. Nebula uses cloud native pub-sub systems to notify users of various events in their workflows |
| nebula.workflow_scheduler | object | `{"enabled":true,"type":"native"}` | **Optional Component** Nebula uses a cloud hosted Cron scheduler to run workflows on a schedule. The following module is optional. Without, this module, you will not have scheduled launchplans / workflows. Docs: https://docs.nebula.org/en/latest/howto/enable_and_use_schedules.html#setting-up-scheduled-workflows |
| nebulaagent.enabled | bool | `true` |  |
| kubernetes-dashboard.enabled | bool | `true` |  |
| kubernetes-dashboard.extraArgs[0] | string | `"--enable-skip-login"` |  |
| kubernetes-dashboard.extraArgs[1] | string | `"--enable-insecure-login"` |  |
| kubernetes-dashboard.extraArgs[2] | string | `"--disable-settings-authorizer"` |  |
| kubernetes-dashboard.protocolHttp | bool | `true` |  |
| kubernetes-dashboard.rbac.clusterReadOnlyRole | bool | `true` |  |
| kubernetes-dashboard.service.externalPort | int | `30082` |  |
| kubernetes-dashboard.service.nodePort | int | `30082` |  |
| kubernetes-dashboard.service.type | string | `"NodePort"` |  |
| minio.affinity | object | `{}` | affinity for Minio deployment |
| minio.enabled | bool | `true` | - enable or disable Minio deployment installation |
| minio.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| minio.image.repository | string | `"ecr.nebula.org/bitnami/minio"` | Docker image for Minio deployment |
| minio.image.tag | string | `"2021.10.13-debian-10-r0"` | Docker image tag |
| minio.nodeSelector | object | `{}` | nodeSelector for Minio deployment |
| minio.podAnnotations | object | `{}` | Annotations for Minio pods |
| minio.replicaCount | int | `1` | Replicas count for Minio deployment |
| minio.resources | object | `{"limits":{"cpu":"200m","memory":"512Mi"},"requests":{"cpu":"10m","memory":"128Mi"}}` | Default resources requests and limits for Minio deployment |
| minio.resources.limits | object | `{"cpu":"200m","memory":"512Mi"}` | Limits are the maximum set of resources needed for this pod |
| minio.resources.requests | object | `{"cpu":"10m","memory":"128Mi"}` | Requests are the minimum set of resources needed for this pod |
| minio.service | object | `{"annotations":{},"type":"NodePort"}` | Service settings for Minio |
| minio.tolerations | list | `[]` | tolerations for Minio deployment |
| postgres.affinity | object | `{}` | affinity for Postgres deployment |
| postgres.enabled | bool | `true` | - enable or disable Postgres deployment installation |
| postgres.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| postgres.image.repository | string | `"ecr.nebula.org/ubuntu/postgres"` | Docker image for Postgres deployment |
| postgres.image.tag | string | `"13-21.04_beta"` | Docker image tag |
| postgres.nodeSelector | object | `{}` | nodeSelector for Postgres deployment |
| postgres.podAnnotations | object | `{}` | Annotations for Postgres pods |
| postgres.replicaCount | int | `1` | Replicas count for Postgres deployment |
| postgres.resources | object | `{"limits":{"cpu":"1000m","memory":"512Mi"},"requests":{"cpu":"10m","memory":"128Mi"}}` | Default resources requests and limits for Postgres deployment |
| postgres.service | object | `{"annotations":{},"type":"NodePort"}` | Service settings for Postgres |
| postgres.tolerations | list | `[]` | tolerations for Postgres deployment |
| redis | object | `{"enabled":false}` | ---------------------------------------------  REDIS SETTINGS  |
| redis.enabled | bool | `false` | - enable or disable Redis Statefulset installation |
| redoc.affinity | object | `{}` | affinity for Minio deployment |
| redoc.enabled | bool | `true` | - enable or disable Minio deployment installation |
| redoc.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| redoc.image.repository | string | `"docker.io/redocly/redoc"` | Docker image for Minio deployment |
| redoc.image.tag | string | `"latest"` | Docker image tag |
| redoc.nodeSelector | object | `{}` | nodeSelector for Minio deployment |
| redoc.podAnnotations | object | `{}` | Annotations for Minio pods |
| redoc.replicaCount | int | `1` | Replicas count for Minio deployment |
| redoc.resources | object | `{"limits":{"cpu":"200m","memory":"512Mi"},"requests":{"cpu":"10m","memory":"128Mi"}}` | Default resources requests and limits for Minio deployment |
| redoc.resources.limits | object | `{"cpu":"200m","memory":"512Mi"}` | Limits are the maximum set of resources needed for this pod |
| redoc.resources.requests | object | `{"cpu":"10m","memory":"128Mi"}` | Requests are the minimum set of resources needed for this pod |
| redoc.service | object | `{"type":"ClusterIP"}` | Service settings for Minio |
| redoc.tolerations | list | `[]` | tolerations for Minio deployment |
| sparkoperator | object | `{"enabled":false}` | Optional: Spark Plugin using the Spark Operator |
| sparkoperator.enabled | bool | `false` | - enable or disable Sparkoperator deployment installation |
