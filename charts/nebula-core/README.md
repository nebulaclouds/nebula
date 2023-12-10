# nebula-core

![Version: v0.1.10](https://img.shields.io/badge/Version-v0.1.10-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square)

A Helm chart for Nebula core

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://../nebulaagent | nebulaagent(nebulaagent) | v0.1.10 |

### Nebula INSTALLATION:
- [Install helm 3](https://helm.sh/docs/intro/install/)
- Fetch chart dependencies ``
- Install Nebula:

```bash
helm repo add nebula https://nebulaclouds.github.io/nebula
helm install -n nebula -f values-eks.yaml --create-namespace nebula nebula/nebula-core
```

Customize your installation by changing settings in `values-eks.yaml`.
You can use the helm diff plugin to review any value changes you've made to your values:

```bash
helm plugin install https://github.com/databus23/helm-diff
helm diff upgrade -f values-eks.yaml nebula nebula/nebula-core
```

Then apply your changes:
```bash
helm upgrade -f values-eks.yaml nebula nebula/nebula-core
```

Install ingress controller (By default Nebula helm chart have contour ingress resource)
```bash
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install gateway bitnami/contour -n nebula
```

#### Alternative: Generate raw kubernetes yaml with helm template
- `helm template --name-template=nebula-eks . -n nebula -f values-eks.yaml > nebula_generated_eks.yaml`
- Deploy the manifest `kubectl apply -f nebula_generated_eks.yaml`

- When all pods are running - run end2end tests: `kubectl apply -f ../end2end/tests/endtoend.yaml`
- Get nebula host `minikube service contour -n heptio-contour --url`. And then visit `http://<HOST>/console`

### CONFIGURATION NOTES:
- The docker images, their tags and other default parameters are configured in `values.yaml` file.
- Each Nebula installation type should have separate `values-*.yaml` file: for sandbox, EKS and etc. The configuration in `values.yaml` and the chosen config `values-*.yaml` are merged when generating the deployment manifest.
- The configuration in `values-sandbox.yaml` is ready for installation in minikube. But `values-eks.yaml` should be edited before installation: s3 bucket, RDS hosts, iam roles, secrets and etc need to be modified.

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| cloud_events.aws.region | string | `"us-east-2"` |  |
| cloud_events.enable | bool | `false` |  |
| cloud_events.eventsPublisher.eventTypes[0] | string | `"all"` |  |
| cloud_events.eventsPublisher.topicName | string | `"arn:aws:sns:us-east-2:123456:123-my-topic"` |  |
| cloud_events.type | string | `"aws"` |  |
| cluster_resource_manager | object | `{"config":{"cluster_resources":{"customData":[{"production":[{"projectQuotaCpu":{"value":"5"}},{"projectQuotaMemory":{"value":"4000Mi"}}]},{"staging":[{"projectQuotaCpu":{"value":"2"}},{"projectQuotaMemory":{"value":"3000Mi"}}]},{"development":[{"projectQuotaCpu":{"value":"4"}},{"projectQuotaMemory":{"value":"3000Mi"}}]}],"refreshInterval":"5m","standaloneDeployment":false,"templatePath":"/etc/nebula/clusterresource/templates"}},"enabled":true,"podAnnotations":{},"service_account_name":"nebulaadmin","standaloneDeployment":false,"templates":[{"key":"aa_namespace","value":"apiVersion: v1\nkind: Namespace\nmetadata:\n  name: {{ namespace }}\nspec:\n  finalizers:\n  - kubernetes\n"},{"key":"ab_project_resource_quota","value":"apiVersion: v1\nkind: ResourceQuota\nmetadata:\n  name: project-quota\n  namespace: {{ namespace }}\nspec:\n  hard:\n    limits.cpu: {{ projectQuotaCpu }}\n    limits.memory: {{ projectQuotaMemory }}\n"}]}` | Configuration for the Cluster resource manager component. This is an optional component, that enables automatic cluster configuration. This is useful to set default quotas, manage namespaces etc that map to a project/domain |
| cluster_resource_manager.config | object | `{"cluster_resources":{"customData":[{"production":[{"projectQuotaCpu":{"value":"5"}},{"projectQuotaMemory":{"value":"4000Mi"}}]},{"staging":[{"projectQuotaCpu":{"value":"2"}},{"projectQuotaMemory":{"value":"3000Mi"}}]},{"development":[{"projectQuotaCpu":{"value":"4"}},{"projectQuotaMemory":{"value":"3000Mi"}}]}],"refreshInterval":"5m","standaloneDeployment":false,"templatePath":"/etc/nebula/clusterresource/templates"}}` | Configmap for ClusterResource parameters |
| cluster_resource_manager.config.cluster_resources | object | `{"customData":[{"production":[{"projectQuotaCpu":{"value":"5"}},{"projectQuotaMemory":{"value":"4000Mi"}}]},{"staging":[{"projectQuotaCpu":{"value":"2"}},{"projectQuotaMemory":{"value":"3000Mi"}}]},{"development":[{"projectQuotaCpu":{"value":"4"}},{"projectQuotaMemory":{"value":"3000Mi"}}]}],"refreshInterval":"5m","standaloneDeployment":false,"templatePath":"/etc/nebula/clusterresource/templates"}` | ClusterResource parameters Refer to the [structure](https://pkg.go.dev/github.com/lyft/nebulaadmin@v0.3.37/pkg/runtime/interfaces#ClusterResourceConfig) to customize. |
| cluster_resource_manager.config.cluster_resources.refreshInterval | string | `"5m"` | How frequently to run the sync process |
| cluster_resource_manager.config.cluster_resources.standaloneDeployment | bool | `false` | Starts the cluster resource manager in standalone mode with requisite auth credentials to call nebulaadmin service endpoints |
| cluster_resource_manager.enabled | bool | `true` | Enables the Cluster resource manager component |
| cluster_resource_manager.podAnnotations | object | `{}` | Annotations for ClusterResource pods |
| cluster_resource_manager.service_account_name | string | `"nebulaadmin"` | Service account name to run with |
| cluster_resource_manager.templates | list | `[{"key":"aa_namespace","value":"apiVersion: v1\nkind: Namespace\nmetadata:\n  name: {{ namespace }}\nspec:\n  finalizers:\n  - kubernetes\n"},{"key":"ab_project_resource_quota","value":"apiVersion: v1\nkind: ResourceQuota\nmetadata:\n  name: project-quota\n  namespace: {{ namespace }}\nspec:\n  hard:\n    limits.cpu: {{ projectQuotaCpu }}\n    limits.memory: {{ projectQuotaMemory }}\n"}]` | Resource templates that should be applied |
| cluster_resource_manager.templates[0] | object | `{"key":"aa_namespace","value":"apiVersion: v1\nkind: Namespace\nmetadata:\n  name: {{ namespace }}\nspec:\n  finalizers:\n  - kubernetes\n"}` | Template for namespaces resources |
| common | object | `{"databaseSecret":{"name":"","secretManifest":{}},"nebulaNamespaceTemplate":{"enabled":false},"ingress":{"albSSLRedirect":false,"annotations":{"nginx.ingress.kubernetes.io/app-root":"/console"},"enabled":true,"separateGrpcIngress":false,"separateGrpcIngressAnnotations":{"nginx.ingress.kubernetes.io/backend-protocol":"GRPC"},"tls":{"enabled":false},"webpackHMR":false}}` | ----------------------------------------------  COMMON SETTINGS  |
| common.databaseSecret.name | string | `""` | Specify name of K8s Secret which contains Database password. Leave it empty if you don't need this Secret |
| common.databaseSecret.secretManifest | object | `{}` | Specify your Secret (with sensitive data) or pseudo-manifest (without sensitive data). See https://github.com/godaddy/kubernetes-external-secrets |
| common.nebulaNamespaceTemplate.enabled | bool | `false` | - Enable or disable creating Nebula namespace in template. Enable when using helm as template-engine only. Disable when using `helm install ...`. |
| common.ingress.albSSLRedirect | bool | `false` | - albSSLRedirect adds a special route for ssl redirect. Only useful in combination with the AWS LoadBalancer Controller. |
| common.ingress.annotations | object | `{"nginx.ingress.kubernetes.io/app-root":"/console"}` | - Ingress annotations applied to both HTTP and GRPC ingresses. |
| common.ingress.enabled | bool | `true` | - Enable or disable creating Ingress for Nebula. Relevant to disable when using e.g. Istio as ingress controller. |
| common.ingress.separateGrpcIngress | bool | `false` | - separateGrpcIngress puts GRPC routes into a separate ingress if true. Required for certain ingress controllers like nginx. |
| common.ingress.separateGrpcIngressAnnotations | object | `{"nginx.ingress.kubernetes.io/backend-protocol":"GRPC"}` | - Extra Ingress annotations applied only to the GRPC ingress. Only makes sense if `separateGrpcIngress` is enabled. |
| common.ingress.tls | object | `{"enabled":false}` | - Ingress hostname host: |
| common.ingress.webpackHMR | bool | `false` | - Enable or disable HMR route to nebulaconsole. This is useful only for frontend development. |
| configmap.admin | object | `{"admin":{"clientId":"{{ .Values.secrets.adminOauthClientCredentials.clientId }}","clientSecretLocation":"/etc/secrets/client_secret","endpoint":"nebulaadmin:81","insecure":true},"event":{"capacity":1000,"rate":500,"type":"admin"}}` | Admin Client configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/subworkflow/launchplan#AdminConfig) |
| configmap.adminServer | object | `{"auth":{"appAuth":{"thirdPartyConfig":{"nebulaClient":{"clientId":"nebulactl","redirectUri":"http://localhost:53593/callback","scopes":["offline","all"]}}},"authorizedUris":["https://localhost:30081","http://nebulaadmin:80","http://nebulaadmin.nebula.svc.cluster.local:80"],"userAuth":{"openId":{"baseUrl":"https://accounts.google.com","clientId":"657465813211-6eog7ek7li5k7i7fvgv2921075063hpe.apps.googleusercontent.com","scopes":["profile","openid"]}}},"nebulaadmin":{"eventVersion":2,"metadataStoragePrefix":["metadata","admin"],"metricsScope":"nebula:","profilerPort":10254,"roleNameKey":"iam.amazonaws.com/role","testing":{"host":"http://nebulaadmin"}},"server":{"grpcPort":8089,"httpPort":8088,"security":{"allowCors":true,"allowedHeaders":["Content-Type","nebula-authorization"],"allowedOrigins":["*"],"secure":false,"useAuth":false}}}` | NebulaAdmin server configuration |
| configmap.adminServer.auth | object | `{"appAuth":{"thirdPartyConfig":{"nebulaClient":{"clientId":"nebulactl","redirectUri":"http://localhost:53593/callback","scopes":["offline","all"]}}},"authorizedUris":["https://localhost:30081","http://nebulaadmin:80","http://nebulaadmin.nebula.svc.cluster.local:80"],"userAuth":{"openId":{"baseUrl":"https://accounts.google.com","clientId":"657465813211-6eog7ek7li5k7i7fvgv2921075063hpe.apps.googleusercontent.com","scopes":["profile","openid"]}}}` | Authentication configuration |
| configmap.adminServer.server.security.secure | bool | `false` | Controls whether to serve requests over SSL/TLS. |
| configmap.adminServer.server.security.useAuth | bool | `false` | Controls whether to enforce authentication. Follow the guide in https://docs.nebula.org/ on how to setup authentication. |
| configmap.catalog | object | `{"catalog-cache":{"endpoint":"datacatalog:89","insecure":true,"type":"datacatalog"}}` | Catalog Client configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/catalog#Config) Additional advanced Catalog configuration [here](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/pluginmachinery/catalog#Config) |
| configmap.clusters.clusterConfigs | list | `[]` |  |
| configmap.clusters.labelClusterMap | object | `{}` |  |
| configmap.console | object | `{"BASE_URL":"/console","CONFIG_DIR":"/etc/nebula/config"}` | Configuration for Nebula console UI |
| configmap.copilot | object | `{"plugins":{"k8s":{"co-pilot":{"image":"cr.nebula.org/nebulaclouds/nebulacopilot:v1.10.6","name":"nebula-copilot-","start-timeout":"30s"}}}}` | Copilot configuration |
| configmap.copilot.plugins.k8s.co-pilot | object | `{"image":"cr.nebula.org/nebulaclouds/nebulacopilot:v1.10.6","name":"nebula-copilot-","start-timeout":"30s"}` | Structure documented [here](https://pkg.go.dev/github.com/lyft/nebulaplugins@v0.5.28/go/tasks/pluginmachinery/nebulak8s/config#NebulaCoPilotConfig) |
| configmap.core | object | `{"manager":{"pod-application":"nebulapropeller","pod-template-container-name":"nebulapropeller","pod-template-name":"nebulapropeller-template"},"propeller":{"downstream-eval-duration":"30s","enable-admin-launcher":true,"leader-election":{"enabled":true,"lease-duration":"15s","lock-config-map":{"name":"propeller-leader","namespace":"nebula"},"renew-deadline":"10s","retry-period":"2s"},"limit-namespace":"all","max-workflow-retries":30,"metadata-prefix":"metadata/propeller","metrics-prefix":"nebula","prof-port":10254,"queue":{"batch-size":-1,"batching-interval":"2s","queue":{"base-delay":"5s","capacity":1000,"max-delay":"120s","rate":100,"type":"maxof"},"sub-queue":{"capacity":100,"rate":10,"type":"bucket"},"type":"batch"},"rawoutput-prefix":"s3://my-s3-bucket/","workers":4,"workflow-reeval-duration":"30s"},"webhook":{"certDir":"/etc/webhook/certs","serviceName":"nebula-pod-webhook"}}` | Core propeller configuration |
| configmap.core.manager | object | `{"pod-application":"nebulapropeller","pod-template-container-name":"nebulapropeller","pod-template-name":"nebulapropeller-template"}` | follows the structure specified [here](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/manager/config#Config). |
| configmap.core.propeller | object | `{"downstream-eval-duration":"30s","enable-admin-launcher":true,"leader-election":{"enabled":true,"lease-duration":"15s","lock-config-map":{"name":"propeller-leader","namespace":"nebula"},"renew-deadline":"10s","retry-period":"2s"},"limit-namespace":"all","max-workflow-retries":30,"metadata-prefix":"metadata/propeller","metrics-prefix":"nebula","prof-port":10254,"queue":{"batch-size":-1,"batching-interval":"2s","queue":{"base-delay":"5s","capacity":1000,"max-delay":"120s","rate":100,"type":"maxof"},"sub-queue":{"capacity":100,"rate":10,"type":"bucket"},"type":"batch"},"rawoutput-prefix":"s3://my-s3-bucket/","workers":4,"workflow-reeval-duration":"30s"}` | follows the structure specified [here](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/config). |
| configmap.datacatalogServer | object | `{"application":{"grpcPort":8089,"grpcServerReflection":true,"httpPort":8080},"datacatalog":{"heartbeat-grace-period-multiplier":3,"max-reservation-heartbeat":"30s","metrics-scope":"datacatalog","profiler-port":10254,"storage-prefix":"metadata/datacatalog"}}` | Datacatalog server config |
| configmap.domain | object | `{"domains":[{"id":"development","name":"development"},{"id":"staging","name":"staging"},{"id":"production","name":"production"}]}` | Domains configuration for Nebula projects. This enables the specified number of domains across all projects in Nebula. |
| configmap.enabled_plugins.tasks | object | `{"task-plugins":{"default-for-task-types":{"container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array"]}}` | Tasks specific configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#GetConfig) |
| configmap.enabled_plugins.tasks.task-plugins | object | `{"default-for-task-types":{"container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array"]}` | Plugins configuration, [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#TaskPluginConfig) |
| configmap.enabled_plugins.tasks.task-plugins.enabled-plugins | list | `["container","sidecar","k8s-array"]` | [Enabled Plugins](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/config#Config). Enable sagemaker*, athena if you install the backend plugins |
| configmap.k8s | object | `{"plugins":{"k8s":{"default-cpus":"100m","default-env-vars":[],"default-memory":"100Mi"}}}` | Kubernetes specific Nebula configuration |
| configmap.k8s.plugins.k8s | object | `{"default-cpus":"100m","default-env-vars":[],"default-memory":"100Mi"}` | Configuration section for all K8s specific plugins [Configuration structure](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/pluginmachinery/nebulak8s/config) |
| configmap.remoteData.remoteData.region | string | `"us-east-1"` |  |
| configmap.remoteData.remoteData.scheme | string | `"local"` |  |
| configmap.remoteData.remoteData.signedUrls.durationMinutes | int | `3` |  |
| configmap.resource_manager | object | `{"propeller":{"resourcemanager":{"type":"noop"}}}` | Resource manager configuration |
| configmap.resource_manager.propeller | object | `{"resourcemanager":{"type":"noop"}}` | resource manager configuration |
| configmap.schedulerConfig.scheduler.metricsScope | string | `"nebula:"` |  |
| configmap.schedulerConfig.scheduler.profilerPort | int | `10254` |  |
| configmap.task_logs | object | `{"plugins":{"logs":{"cloudwatch-enabled":false,"kubernetes-enabled":false}}}` | Section that configures how the Task logs are displayed on the UI. This has to be changed based on your actual logging provider. Refer to [structure](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/logs#LogConfig) to understand how to configure various logging engines |
| configmap.task_logs.plugins.logs.cloudwatch-enabled | bool | `false` | One option is to enable cloudwatch logging for EKS, update the region and log group accordingly |
| configmap.task_resource_defaults | object | `{"task_resources":{"defaults":{"cpu":"100m","memory":"500Mi","storage":"500Mi"},"limits":{"cpu":2,"gpu":1,"memory":"1Gi","storage":"20Mi"}}}` | Task default resources configuration Refer to the full [structure](https://pkg.go.dev/github.com/lyft/nebulaadmin@v0.3.37/pkg/runtime/interfaces#TaskResourceConfiguration). |
| configmap.task_resource_defaults.task_resources | object | `{"defaults":{"cpu":"100m","memory":"500Mi","storage":"500Mi"},"limits":{"cpu":2,"gpu":1,"memory":"1Gi","storage":"20Mi"}}` | Task default resources parameters |
| daskoperator | object | `{"enabled":false}` | Optional: Dask Plugin using the Dask Operator |
| daskoperator.enabled | bool | `false` | - enable or disable the dask operator deployment installation |
| databricks | object | `{"enabled":false,"plugin_config":{"plugins":{"databricks":{"databricksInstance":"dbc-a53b7a3c-614c","entrypointFile":"dbfs:///FileStore/tables/entrypoint.py"}}}}` | Optional: Databricks Plugin allows us to run the spark job on the Databricks platform. |
| datacatalog.additionalContainers | list | `[]` | Appends additional containers to the deployment spec. May include template values. |
| datacatalog.additionalVolumeMounts | list | `[]` | Appends additional volume mounts to the main container's spec. May include template values. |
| datacatalog.additionalVolumes | list | `[]` | Appends additional volumes to the deployment spec. May include template values. |
| datacatalog.affinity | object | `{}` | affinity for Datacatalog deployment |
| datacatalog.configPath | string | `"/etc/datacatalog/config/*.yaml"` | Default regex string for searching configuration files |
| datacatalog.enabled | bool | `true` |  |
| datacatalog.extraArgs | object | `{}` | Appends extra command line arguments to the main command |
| datacatalog.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| datacatalog.image.repository | string | `"cr.nebula.org/nebulaclouds/datacatalog"` | Docker image for Datacatalog deployment |
| datacatalog.image.tag | string | `"v1.10.6"` | Docker image tag |
| datacatalog.nodeSelector | object | `{}` | nodeSelector for Datacatalog deployment |
| datacatalog.podAnnotations | object | `{}` | Annotations for Datacatalog pods |
| datacatalog.priorityClassName | string | `""` | Sets priorityClassName for datacatalog pod(s). |
| datacatalog.replicaCount | int | `1` | Replicas count for Datacatalog deployment |
| datacatalog.resources | object | `{"limits":{"cpu":"500m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}}` | Default resources requests and limits for Datacatalog deployment |
| datacatalog.service | object | `{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"type":"NodePort"}` | Service settings for Datacatalog |
| datacatalog.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for Datacatalog |
| datacatalog.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to Datacatalog pods |
| datacatalog.serviceAccount.create | bool | `true` | Should a service account be created for Datacatalog |
| datacatalog.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| datacatalog.tolerations | list | `[]` | tolerations for Datacatalog deployment |
| db.admin.database.dbname | string | `"nebulaadmin"` |  |
| db.admin.database.host | string | `"postgres"` |  |
| db.admin.database.port | int | `5432` |  |
| db.admin.database.username | string | `"postgres"` |  |
| db.datacatalog.database.dbname | string | `"datacatalog"` |  |
| db.datacatalog.database.host | string | `"postgres"` |  |
| db.datacatalog.database.port | int | `5432` |  |
| db.datacatalog.database.username | string | `"postgres"` |  |
| deployRedoc | bool | `false` |  |
| external_events | object | `{"aws":{"region":"us-east-2"},"enable":false,"eventsPublisher":{"eventTypes":["all"],"topicName":"arn:aws:sns:us-east-2:123456:123-my-topic"},"type":"aws"}` | **Optional Component** External events are used to send events (unprocessed, as Admin see them) to an SNS topic (or gcp equivalent) The config is here as an example only - if not enabled, it won't be used. |
| nebulaadmin.additionalContainers | list | `[]` | Appends additional containers to the deployment spec. May include template values. |
| nebulaadmin.additionalVolumeMounts | list | `[]` | Appends additional volume mounts to the main container's spec. May include template values. |
| nebulaadmin.additionalVolumes | list | `[]` | Appends additional volumes to the deployment spec. May include template values. |
| nebulaadmin.affinity | object | `{}` | affinity for Nebulaadmin deployment |
| nebulaadmin.configPath | string | `"/etc/nebula/config/*.yaml"` | Default regex string for searching configuration files |
| nebulaadmin.enabled | bool | `true` |  |
| nebulaadmin.env | list | `[]` | Additional nebulaadmin container environment variables  e.g. SendGrid's API key  - name: SENDGRID_API_KEY    value: "<your sendgrid api key>"  e.g. secret environment variable (you can combine it with .additionalVolumes): - name: SENDGRID_API_KEY   valueFrom:     secretKeyRef:       name: sendgrid-secret       key: api_key |
| nebulaadmin.extraArgs | object | `{}` | Appends extra command line arguments to the serve command |
| nebulaadmin.image.pullPolicy | string | `"IfNotPresent"` |  |
| nebulaadmin.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulaadmin"` | Docker image for Nebulaadmin deployment |
| nebulaadmin.image.tag | string | `"v1.10.6"` |  |
| nebulaadmin.initialProjects | list | `["nebulasnacks","nebulatester","nebulaexamples"]` | Initial projects to create |
| nebulaadmin.nodeSelector | object | `{}` | nodeSelector for Nebulaadmin deployment |
| nebulaadmin.podAnnotations | object | `{}` | Annotations for Nebulaadmin pods |
| nebulaadmin.priorityClassName | string | `""` | Sets priorityClassName for nebulaadmin pod(s). |
| nebulaadmin.replicaCount | int | `1` | Replicas count for Nebulaadmin deployment |
| nebulaadmin.resources | object | `{"limits":{"cpu":"250m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}}` | Default resources requests and limits for Nebulaadmin deployment |
| nebulaadmin.secrets | object | `{}` |  |
| nebulaadmin.service | object | `{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"loadBalancerSourceRanges":[],"type":"ClusterIP"}` | Service settings for Nebulaadmin |
| nebulaadmin.serviceAccount | object | `{"alwaysCreate":false,"annotations":{},"clusterRole":{"apiGroups":["","nebula.lyft.com","rbac.authorization.k8s.io"],"resources":["configmaps","nebulaworkflows","namespaces","pods","resourcequotas","roles","rolebindings","secrets","services","serviceaccounts","spark-role","limitranges"],"verbs":["*"]},"create":true,"createClusterRole":true,"imagePullSecrets":[]}` | Configuration for service accounts for NebulaAdmin |
| nebulaadmin.serviceAccount.alwaysCreate | bool | `false` | Should a service account always be created for nebulaadmin even without an actual nebulaadmin deployment running (e.g. for multi-cluster setups) |
| nebulaadmin.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to Nebulaadmin pods |
| nebulaadmin.serviceAccount.clusterRole | object | `{"apiGroups":["","nebula.lyft.com","rbac.authorization.k8s.io"],"resources":["configmaps","nebulaworkflows","namespaces","pods","resourcequotas","roles","rolebindings","secrets","services","serviceaccounts","spark-role","limitranges"],"verbs":["*"]}` | Configuration for ClusterRole created for Nebulaadmin |
| nebulaadmin.serviceAccount.clusterRole.apiGroups | list | `["","nebula.lyft.com","rbac.authorization.k8s.io"]` | Specifies the API groups that this ClusterRole can access |
| nebulaadmin.serviceAccount.clusterRole.resources | list | `["configmaps","nebulaworkflows","namespaces","pods","resourcequotas","roles","rolebindings","secrets","services","serviceaccounts","spark-role","limitranges"]` | Specifies the resources that this ClusterRole can access |
| nebulaadmin.serviceAccount.clusterRole.verbs | list | `["*"]` | Specifies the verbs (actions) that this ClusterRole can perform on the specified resources |
| nebulaadmin.serviceAccount.create | bool | `true` | Should a service account be created for nebulaadmin |
| nebulaadmin.serviceAccount.createClusterRole | bool | `true` | Should a ClusterRole be created for Nebulaadmin |
| nebulaadmin.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebulaadmin.serviceMonitor | object | `{"enabled":false,"interval":"60s","labels":{},"scrapeTimeout":"30s"}` | Settings for nebulaadmin service monitor |
| nebulaadmin.serviceMonitor.enabled | bool | `false` | If enabled create the nebulaadmin service monitor |
| nebulaadmin.serviceMonitor.interval | string | `"60s"` | Sets the interval at which metrics will be scraped by prometheus |
| nebulaadmin.serviceMonitor.labels | object | `{}` | Sets the labels for the service monitor which are required by the prometheus to auto-detect the service monitor and start scrapping the metrics |
| nebulaadmin.serviceMonitor.scrapeTimeout | string | `"30s"` | Sets the timeout after which request to scrape metrics will time out |
| nebulaadmin.tolerations | list | `[]` | tolerations for Nebulaadmin deployment |
| nebulaagent.enabled | bool | `false` |  |
| nebulaconsole.affinity | object | `{}` | affinity for Nebulaconsole deployment |
| nebulaconsole.enabled | bool | `true` |  |
| nebulaconsole.ga.enabled | bool | `false` |  |
| nebulaconsole.ga.tracking_id | string | `"G-0QW4DJWJ20"` |  |
| nebulaconsole.image.pullPolicy | string | `"IfNotPresent"` |  |
| nebulaconsole.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulaconsole"` | Docker image for Nebulaconsole deployment |
| nebulaconsole.image.tag | string | `"v1.10.2"` |  |
| nebulaconsole.nodeSelector | object | `{}` | nodeSelector for Nebulaconsole deployment |
| nebulaconsole.podAnnotations | object | `{}` | Annotations for Nebulaconsole pods |
| nebulaconsole.priorityClassName | string | `""` | Sets priorityClassName for nebula console pod(s). |
| nebulaconsole.replicaCount | int | `1` | Replicas count for Nebulaconsole deployment |
| nebulaconsole.resources | object | `{"limits":{"cpu":"500m","memory":"250Mi"},"requests":{"cpu":"10m","memory":"50Mi"}}` | Default resources requests and limits for Nebulaconsole deployment |
| nebulaconsole.service | object | `{"annotations":{},"type":"ClusterIP"}` | Service settings for Nebulaconsole |
| nebulaconsole.tolerations | list | `[]` | tolerations for Nebulaconsole deployment |
| nebulapropeller.additionalContainers | list | `[]` | Appends additional containers to the deployment spec. May include template values. |
| nebulapropeller.additionalVolumeMounts | list | `[]` | Appends additional volume mounts to the main container's spec. May include template values. |
| nebulapropeller.additionalVolumes | list | `[]` | Appends additional volumes to the deployment spec. May include template values. |
| nebulapropeller.affinity | object | `{}` | affinity for Nebulapropeller deployment |
| nebulapropeller.cacheSizeMbs | int | `0` |  |
| nebulapropeller.clusterName | string | `""` | Defines the cluster name used in events sent to Admin |
| nebulapropeller.configPath | string | `"/etc/nebula/config/*.yaml"` | Default regex string for searching configuration files |
| nebulapropeller.createCRDs | bool | `true` | Whether to install the nebulaworkflows CRD with helm |
| nebulapropeller.enabled | bool | `true` |  |
| nebulapropeller.extraArgs | object | `{}` | Appends extra command line arguments to the main command |
| nebulapropeller.image.pullPolicy | string | `"IfNotPresent"` |  |
| nebulapropeller.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulapropeller"` | Docker image for Nebulapropeller deployment |
| nebulapropeller.image.tag | string | `"v1.10.6"` |  |
| nebulapropeller.manager | bool | `false` |  |
| nebulapropeller.nodeSelector | object | `{}` | nodeSelector for Nebulapropeller deployment |
| nebulapropeller.podAnnotations | object | `{}` | Annotations for Nebulapropeller pods |
| nebulapropeller.priorityClassName | string | `""` | Sets priorityClassName for propeller pod(s). |
| nebulapropeller.replicaCount | int | `1` | Replicas count for Nebulapropeller deployment |
| nebulapropeller.resources | object | `{"limits":{"cpu":"200m","ephemeral-storage":"100Mi","memory":"200Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"100Mi"}}` | Default resources requests and limits for Nebulapropeller deployment |
| nebulapropeller.service | object | `{"enabled":false}` | Settings for nebulapropeller service |
| nebulapropeller.service.enabled | bool | `false` | If enabled create the nebulapropeller service |
| nebulapropeller.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for NebulaPropeller |
| nebulapropeller.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to NebulaPropeller pods |
| nebulapropeller.serviceAccount.create | bool | `true` | Should a service account be created for NebulaPropeller |
| nebulapropeller.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebulapropeller.serviceMonitor | object | `{"enabled":false,"interval":"60s","labels":{},"scrapeTimeout":"30s"}` | Settings for nebulapropeller service monitor |
| nebulapropeller.serviceMonitor.enabled | bool | `false` | If enabled create the flyetepropeller service monitor |
| nebulapropeller.serviceMonitor.interval | string | `"60s"` | Sets the interval at which metrics will be scraped by prometheus |
| nebulapropeller.serviceMonitor.labels | object | `{}` | Sets the labels for the service monitor which are required by the prometheus to auto-detect the service monitor and start scrapping the metrics |
| nebulapropeller.serviceMonitor.scrapeTimeout | string | `"30s"` | Sets the timeout after which request to scrape metrics will time out |
| nebulapropeller.terminationMessagePolicy | string | `"FallbackToLogsOnError"` | Error reporting |
| nebulapropeller.tolerations | list | `[]` | tolerations for Nebulapropeller deployment |
| nebulascheduler.additionalContainers | list | `[]` | Appends additional containers to the deployment spec. May include template values. |
| nebulascheduler.additionalVolumeMounts | list | `[]` | Appends additional volume mounts to the main container's spec. May include template values. |
| nebulascheduler.additionalVolumes | list | `[]` | Appends additional volumes to the deployment spec. May include template values. |
| nebulascheduler.affinity | object | `{}` | affinity for Nebulascheduler deployment |
| nebulascheduler.configPath | string | `"/etc/nebula/config/*.yaml"` | Default regex string for searching configuration files |
| nebulascheduler.image.pullPolicy | string | `"IfNotPresent"` | Docker image pull policy |
| nebulascheduler.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulascheduler"` | Docker image for Nebulascheduler deployment |
| nebulascheduler.image.tag | string | `"v1.10.6"` | Docker image tag |
| nebulascheduler.nodeSelector | object | `{}` | nodeSelector for Nebulascheduler deployment |
| nebulascheduler.podAnnotations | object | `{}` | Annotations for Nebulascheduler pods |
| nebulascheduler.priorityClassName | string | `""` | Sets priorityClassName for nebula scheduler pod(s). |
| nebulascheduler.resources | object | `{"limits":{"cpu":"250m","ephemeral-storage":"100Mi","memory":"500Mi"},"requests":{"cpu":"10m","ephemeral-storage":"50Mi","memory":"50Mi"}}` | Default resources requests and limits for Nebulascheduler deployment |
| nebulascheduler.runPrecheck | bool | `true` | Whether to inject an init container which waits on nebulaadmin |
| nebulascheduler.secrets | object | `{}` |  |
| nebulascheduler.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for Nebulascheduler |
| nebulascheduler.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to Nebulascheduler pods |
| nebulascheduler.serviceAccount.create | bool | `true` | Should a service account be created for Nebulascheduler |
| nebulascheduler.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| nebulascheduler.tolerations | list | `[]` | tolerations for Nebulascheduler deployment |
| secrets.adminOauthClientCredentials.clientId | string | `"nebulapropeller"` |  |
| secrets.adminOauthClientCredentials.clientSecret | string | `"foobar"` |  |
| secrets.adminOauthClientCredentials.enabled | bool | `true` | If enabled is true, helm will create and manage `nebula-secret-auth` and populate it with `clientSecret`. If enabled is false, it's up to the user to create `nebula-secret-auth` as described in https://docs.nebula.org/en/latest/deployment/cluster_config/auth_setup.html#oauth2-authorization-server |
| sparkoperator | object | `{"enabled":false,"plugin_config":{"plugins":{"spark":{"spark-config-default":[{"spark.hadoop.fs.s3a.aws.credentials.provider":"com.amazonaws.auth.DefaultAWSCredentialsProviderChain"},{"spark.hadoop.mapreduce.fileoutputcommitter.algorithm.version":"2"},{"spark.kubernetes.allocation.batch.size":"50"},{"spark.hadoop.fs.s3a.acl.default":"BucketOwnerFullControl"},{"spark.hadoop.fs.s3n.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3n.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3a.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3a.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3a.multipart.threshold":"536870912"},{"spark.blacklist.enabled":"true"},{"spark.blacklist.timeout":"5m"},{"spark.task.maxfailures":"8"}]}}}}` | Optional: Spark Plugin using the Spark Operator |
| sparkoperator.enabled | bool | `false` | - enable or disable Sparkoperator deployment installation |
| sparkoperator.plugin_config | object | `{"plugins":{"spark":{"spark-config-default":[{"spark.hadoop.fs.s3a.aws.credentials.provider":"com.amazonaws.auth.DefaultAWSCredentialsProviderChain"},{"spark.hadoop.mapreduce.fileoutputcommitter.algorithm.version":"2"},{"spark.kubernetes.allocation.batch.size":"50"},{"spark.hadoop.fs.s3a.acl.default":"BucketOwnerFullControl"},{"spark.hadoop.fs.s3n.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3n.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3a.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3a.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3a.multipart.threshold":"536870912"},{"spark.blacklist.enabled":"true"},{"spark.blacklist.timeout":"5m"},{"spark.task.maxfailures":"8"}]}}}` | Spark plugin configuration |
| sparkoperator.plugin_config.plugins.spark.spark-config-default | list | `[{"spark.hadoop.fs.s3a.aws.credentials.provider":"com.amazonaws.auth.DefaultAWSCredentialsProviderChain"},{"spark.hadoop.mapreduce.fileoutputcommitter.algorithm.version":"2"},{"spark.kubernetes.allocation.batch.size":"50"},{"spark.hadoop.fs.s3a.acl.default":"BucketOwnerFullControl"},{"spark.hadoop.fs.s3n.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3n.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3a.impl":"org.apache.hadoop.fs.s3a.S3AFileSystem"},{"spark.hadoop.fs.AbstractFileSystem.s3a.impl":"org.apache.hadoop.fs.s3a.S3A"},{"spark.hadoop.fs.s3a.multipart.threshold":"536870912"},{"spark.blacklist.enabled":"true"},{"spark.blacklist.timeout":"5m"},{"spark.task.maxfailures":"8"}]` | Spark default configuration |
| storage | object | `{"bucketName":"my-s3-bucket","custom":{},"enableMultiContainer":false,"gcs":null,"limits":{"maxDownloadMBs":10},"s3":{"accessKey":"","authType":"iam","region":"us-east-1","secretKey":""},"type":"sandbox"}` | ----------------------------------------------------  STORAGE SETTINGS  |
| storage.bucketName | string | `"my-s3-bucket"` | bucketName defines the storage bucket nebula will use. Required for all types except for sandbox. |
| storage.custom | object | `{}` | Settings for storage type custom. See https://github.com/graymeta/stow for supported storage providers/settings. |
| storage.enableMultiContainer | bool | `false` | toggles multi-container storage config |
| storage.gcs | string | `nil` | settings for storage type gcs |
| storage.limits | object | `{"maxDownloadMBs":10}` | default limits being applied to storage config |
| storage.s3 | object | `{"accessKey":"","authType":"iam","region":"us-east-1","secretKey":""}` | settings for storage type s3 |
| storage.s3.accessKey | string | `""` | AWS IAM user access key ID to use for S3 bucket auth, only used if authType is set to accesskey |
| storage.s3.authType | string | `"iam"` | type of authentication to use for S3 buckets, can either be iam or accesskey |
| storage.s3.secretKey | string | `""` | AWS IAM user secret access key to use for S3 bucket auth, only used if authType is set to accesskey |
| storage.type | string | `"sandbox"` | Sets the storage type. Supported values are sandbox, s3, gcs and custom. |
| webhook.enabled | bool | `true` | enable or disable secrets webhook |
| webhook.service | object | `{"annotations":{"projectcontour.io/upstream-protocol.h2c":"grpc"},"type":"ClusterIP"}` | Service settings for the webhook |
| webhook.serviceAccount | object | `{"annotations":{},"create":true,"imagePullSecrets":[]}` | Configuration for service accounts for the webhook |
| webhook.serviceAccount.annotations | object | `{}` | Annotations for ServiceAccount attached to the webhook |
| webhook.serviceAccount.create | bool | `true` | Should a service account be created for the webhook |
| webhook.serviceAccount.imagePullSecrets | list | `[]` | ImagePullSecrets to automatically assign to the service account |
| workflow_notifications | object | `{"config":{},"enabled":false}` | **Optional Component** Workflow notifications module is an optional dependency. Nebula uses cloud native pub-sub systems to notify users of various events in their workflows |
| workflow_scheduler | object | `{"config":{},"enabled":false,"type":""}` | **Optional Component** Nebula uses a cloud hosted Cron scheduler to run workflows on a schedule. The following module is optional. Without, this module, you will not have scheduled launchplans / workflows. Docs: https://docs.nebula.org/en/latest/howto/enable_and_use_schedules.html#setting-up-scheduled-workflows |
