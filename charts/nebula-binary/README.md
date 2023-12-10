# nebula-binary

![Version: v0.1.10](https://img.shields.io/badge/Version-v0.1.10-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.16.0](https://img.shields.io/badge/AppVersion-1.16.0-informational?style=flat-square)

Chart for basic single Nebula executable deployment

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://../nebulaagent | nebulaagent(nebulaagent) | v0.1.10 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| clusterResourceTemplates.annotations | object | `{}` |  |
| clusterResourceTemplates.externalConfigMap | string | `""` |  |
| clusterResourceTemplates.inline | object | `{}` |  |
| clusterResourceTemplates.inlineConfigMap | string | `""` |  |
| clusterResourceTemplates.labels | object | `{}` |  |
| commonAnnotations | object | `{}` |  |
| commonLabels | object | `{}` |  |
| configuration.agentService.defaultAgent.defaultTimeout | string | `"10s"` |  |
| configuration.agentService.defaultAgent.endpoint | string | `"dns:///nebulaagent.nebula.svc.cluster.local:8000"` |  |
| configuration.agentService.defaultAgent.insecure | bool | `true` |  |
| configuration.agentService.defaultAgent.timeouts.GetTask | string | `"10s"` |  |
| configuration.agentService.supportedTaskTypes[0] | string | `"default_task"` |  |
| configuration.annotations | object | `{}` |  |
| configuration.auth.authorizedUris | list | `[]` |  |
| configuration.auth.clientSecretsExternalSecretRef | string | `""` |  |
| configuration.auth.enableAuthServer | bool | `true` |  |
| configuration.auth.enabled | bool | `false` |  |
| configuration.auth.nebulaClient.audience | string | `""` |  |
| configuration.auth.nebulaClient.clientId | string | `"nebulactl"` |  |
| configuration.auth.nebulaClient.redirectUri | string | `"http://localhost:53593/callback"` |  |
| configuration.auth.nebulaClient.scopes[0] | string | `"all"` |  |
| configuration.auth.internal.clientId | string | `"nebulapropeller"` |  |
| configuration.auth.internal.clientSecret | string | `""` |  |
| configuration.auth.internal.clientSecretHash | string | `""` |  |
| configuration.auth.oidc.baseUrl | string | `""` |  |
| configuration.auth.oidc.clientId | string | `""` |  |
| configuration.auth.oidc.clientSecret | string | `""` |  |
| configuration.co-pilot.image.repository | string | `"cr.nebula.org/nebulaclouds/nebulacopilot"` |  |
| configuration.co-pilot.image.tag | string | `"v1.10.6"` |  |
| configuration.database.dbname | string | `"nebula"` |  |
| configuration.database.host | string | `"127.0.0.1"` |  |
| configuration.database.options | string | `"sslmode=disable"` |  |
| configuration.database.password | string | `""` |  |
| configuration.database.passwordPath | string | `""` |  |
| configuration.database.port | int | `5432` |  |
| configuration.database.username | string | `"postgres"` |  |
| configuration.externalConfigMap | string | `""` |  |
| configuration.externalSecretRef | string | `""` |  |
| configuration.inline | object | `{}` |  |
| configuration.inlineConfigMap | string | `""` |  |
| configuration.inlineSecretRef | string | `""` |  |
| configuration.labels | object | `{}` |  |
| configuration.logging.level | int | `1` |  |
| configuration.logging.plugins.cloudwatch.enabled | bool | `false` |  |
| configuration.logging.plugins.cloudwatch.templateUri | string | `""` |  |
| configuration.logging.plugins.custom | list | `[]` |  |
| configuration.logging.plugins.kubernetes.enabled | bool | `false` |  |
| configuration.logging.plugins.kubernetes.templateUri | string | `""` |  |
| configuration.logging.plugins.stackdriver.enabled | bool | `false` |  |
| configuration.logging.plugins.stackdriver.templateUri | string | `""` |  |
| configuration.storage.metadataContainer | string | `"my-organization-nebula-container"` |  |
| configuration.storage.provider | string | `"s3"` |  |
| configuration.storage.providerConfig.azure.account | string | `"storage-account-name"` |  |
| configuration.storage.providerConfig.azure.configDomainSuffix | string | `""` |  |
| configuration.storage.providerConfig.azure.configUploadConcurrency | int | `4` |  |
| configuration.storage.providerConfig.azure.key | string | `""` |  |
| configuration.storage.providerConfig.gcs.project | string | `"my-organization-gcp-project"` |  |
| configuration.storage.providerConfig.s3.accessKey | string | `""` |  |
| configuration.storage.providerConfig.s3.authType | string | `"iam"` |  |
| configuration.storage.providerConfig.s3.disableSSL | bool | `false` |  |
| configuration.storage.providerConfig.s3.endpoint | string | `""` |  |
| configuration.storage.providerConfig.s3.region | string | `"us-east-1"` |  |
| configuration.storage.providerConfig.s3.secretKey | string | `""` |  |
| configuration.storage.providerConfig.s3.v2Signing | bool | `false` |  |
| configuration.storage.userDataContainer | string | `"my-organization-nebula-container"` |  |
| deployment.annotations | object | `{}` |  |
| deployment.args | list | `[]` |  |
| deployment.command | list | `[]` |  |
| deployment.extraEnvVars | list | `[]` |  |
| deployment.extraEnvVarsConfigMap | string | `""` |  |
| deployment.extraEnvVarsSecret | string | `""` |  |
| deployment.extraPodSpec | object | `{}` |  |
| deployment.extraVolumeMounts | list | `[]` |  |
| deployment.extraVolumes | list | `[]` |  |
| deployment.genAdminAuthSecret.args | list | `[]` |  |
| deployment.genAdminAuthSecret.command | list | `[]` |  |
| deployment.image.pullPolicy | string | `"IfNotPresent"` |  |
| deployment.image.repository | string | `"cr.nebula.org/nebulaclouds/nebula-binary"` |  |
| deployment.image.tag | string | `"latest"` |  |
| deployment.initContainers | list | `[]` |  |
| deployment.labels | object | `{}` |  |
| deployment.lifecycleHooks | object | `{}` |  |
| deployment.livenessProbe | object | `{}` |  |
| deployment.podAnnotations | object | `{}` |  |
| deployment.podLabels | object | `{}` |  |
| deployment.podSecurityContext.enabled | bool | `false` |  |
| deployment.podSecurityContext.fsGroup | int | `65534` |  |
| deployment.podSecurityContext.runAsGroup | int | `65534` |  |
| deployment.podSecurityContext.runAsUser | int | `65534` |  |
| deployment.readinessProbe | object | `{}` |  |
| deployment.sidecars | list | `[]` |  |
| deployment.startupProbe | object | `{}` |  |
| deployment.waitForDB.args | list | `[]` |  |
| deployment.waitForDB.command | list | `[]` |  |
| deployment.waitForDB.image.pullPolicy | string | `"IfNotPresent"` |  |
| deployment.waitForDB.image.repository | string | `"postgres"` |  |
| deployment.waitForDB.image.tag | string | `"15-alpine"` |  |
| enabled_plugins.tasks | object | `{"task-plugins":{"default-for-task-types":{"container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array","agent-service"]}}` | Tasks specific configuration [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#GetConfig) |
| enabled_plugins.tasks.task-plugins | object | `{"default-for-task-types":{"container":"container","container_array":"k8s-array","sidecar":"sidecar"},"enabled-plugins":["container","sidecar","k8s-array","agent-service"]}` | Plugins configuration, [structure](https://pkg.go.dev/github.com/nebulaclouds/nebulapropeller/pkg/controller/nodes/task/config#TaskPluginConfig) |
| enabled_plugins.tasks.task-plugins.enabled-plugins | list | `["container","sidecar","k8s-array","agent-service"]` | [Enabled Plugins](https://pkg.go.dev/github.com/lyft/nebulaplugins/go/tasks/config#Config). Enable sagemaker*, athena if you install the backend plugins |
| nebula-core-components.admin.disableClusterResourceManager | bool | `false` |  |
| nebula-core-components.admin.disableScheduler | bool | `false` |  |
| nebula-core-components.admin.disabled | bool | `false` |  |
| nebula-core-components.admin.seedProjects[0] | string | `"nebulasnacks"` |  |
| nebula-core-components.dataCatalog.disabled | bool | `false` |  |
| nebula-core-components.propeller.disableWebhook | bool | `false` |  |
| nebula-core-components.propeller.disabled | bool | `false` |  |
| nebulaagent.enabled | bool | `false` |  |
| fullnameOverride | string | `""` |  |
| ingress.commonAnnotations | object | `{}` |  |
| ingress.create | bool | `false` |  |
| ingress.grpcAnnotations | object | `{}` |  |
| ingress.grpcExtraPaths.append | list | `[]` |  |
| ingress.grpcExtraPaths.prepend | list | `[]` |  |
| ingress.grpcIngressClassName | string | `""` |  |
| ingress.grpcTls | list | `[]` |  |
| ingress.host | string | `""` |  |
| ingress.httpAnnotations | object | `{}` |  |
| ingress.httpExtraPaths.append | list | `[]` |  |
| ingress.httpExtraPaths.prepend | list | `[]` |  |
| ingress.httpIngressClassName | string | `""` |  |
| ingress.httpTls | list | `[]` |  |
| ingress.ingressClassName | string | `""` |  |
| ingress.labels | object | `{}` |  |
| ingress.tls | list | `[]` |  |
| nameOverride | string | `""` |  |
| rbac.annotations | object | `{}` |  |
| rbac.create | bool | `true` |  |
| rbac.extraRules | list | `[]` |  |
| rbac.labels | object | `{}` |  |
| service.clusterIP | string | `""` |  |
| service.commonAnnotations | object | `{}` |  |
| service.externalTrafficPolicy | string | `"Cluster"` |  |
| service.extraPorts | list | `[]` |  |
| service.grpcAnnotations | object | `{}` |  |
| service.httpAnnotations | object | `{}` |  |
| service.labels | object | `{}` |  |
| service.loadBalancerIP | string | `""` |  |
| service.loadBalancerSourceRanges | list | `[]` |  |
| service.nodePorts.grpc | string | `""` |  |
| service.nodePorts.http | string | `""` |  |
| service.ports.grpc | string | `""` |  |
| service.ports.http | string | `""` |  |
| service.type | string | `"ClusterIP"` |  |
| serviceAccount.annotations | object | `{}` |  |
| serviceAccount.create | bool | `true` |  |
| serviceAccount.imagePullSecrets | list | `[]` |  |
| serviceAccount.labels | object | `{}` |  |
| serviceAccount.name | string | `""` |  |

