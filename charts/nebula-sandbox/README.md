# nebula-sandbox

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 1.16.0](https://img.shields.io/badge/AppVersion-1.16.0-informational?style=flat-square)

A Helm chart for the Nebula local sandbox

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://../nebula-binary | nebula-binary | v0.1.10 |
| https://charts.bitnami.com/bitnami | minio | 12.1.1 |
| https://charts.bitnami.com/bitnami | postgresql | 12.1.9 |
| https://helm.twun.io/ | docker-registry | 2.2.2 |
| https://kubernetes.github.io/dashboard/ | kubernetes-dashboard | 6.0.0 |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| docker-registry.enabled | bool | `true` |  |
| docker-registry.image.pullPolicy | string | `"Never"` |  |
| docker-registry.image.tag | string | `"sandbox"` |  |
| docker-registry.persistence.enabled | bool | `false` |  |
| docker-registry.service.nodePort | int | `30000` |  |
| docker-registry.service.type | string | `"NodePort"` |  |
| nebula-binary.clusterResourceTemplates.inlineConfigMap | string | `"{{ include \"nebula-sandbox.clusterResourceTemplates.inlineConfigMap\" . }}"` |  |
| nebula-binary.configuration.database.host | string | `"{{ printf \"%s-postgresql\" .Release.Name | trunc 63 | trimSuffix \"-\" }}"` |  |
| nebula-binary.configuration.database.password | string | `"postgres"` |  |
| nebula-binary.configuration.inline.plugins.k8s.default-env-vars[0].NEBULA_AWS_ENDPOINT | string | `"http://{{ printf \"%s-minio\" .Release.Name | trunc 63 | trimSuffix \"-\" }}.{{ .Release.Namespace }}:9000"` |  |
| nebula-binary.configuration.inline.plugins.k8s.default-env-vars[1].NEBULA_AWS_ACCESS_KEY_ID | string | `"minio"` |  |
| nebula-binary.configuration.inline.plugins.k8s.default-env-vars[2].NEBULA_AWS_SECRET_ACCESS_KEY | string | `"miniostorage"` |  |
| nebula-binary.configuration.inline.storage.signedURL.stowConfigOverride.endpoint | string | `"http://localhost:30002"` |  |
| nebula-binary.configuration.inline.task_resources.defaults.cpu | string | `"500m"` |  |
| nebula-binary.configuration.inline.task_resources.defaults.ephemeralStorage | int | `0` |  |
| nebula-binary.configuration.inline.task_resources.defaults.gpu | int | `0` |  |
| nebula-binary.configuration.inline.task_resources.defaults.memory | string | `"1Gi"` |  |
| nebula-binary.configuration.inline.task_resources.limits.cpu | int | `0` |  |
| nebula-binary.configuration.inline.task_resources.limits.ephemeralStorage | int | `0` |  |
| nebula-binary.configuration.inline.task_resources.limits.gpu | int | `0` |  |
| nebula-binary.configuration.inline.task_resources.limits.memory | int | `0` |  |
| nebula-binary.configuration.inlineConfigMap | string | `"{{ include \"nebula-sandbox.configuration.inlineConfigMap\" . }}"` |  |
| nebula-binary.configuration.logging.level | int | `6` |  |
| nebula-binary.configuration.logging.plugins.kubernetes.enabled | bool | `true` |  |
| nebula-binary.configuration.logging.plugins.kubernetes.templateUri | string | `"http://localhost:30080/kubernetes-dashboard/#/log/{{.namespace }}/{{ .podName }}/pod?namespace={{ .namespace }}"` |  |
| nebula-binary.configuration.storage.metadataContainer | string | `"my-s3-bucket"` |  |
| nebula-binary.configuration.storage.provider | string | `"s3"` |  |
| nebula-binary.configuration.storage.providerConfig.s3.accessKey | string | `"minio"` |  |
| nebula-binary.configuration.storage.providerConfig.s3.authType | string | `"accesskey"` |  |
| nebula-binary.configuration.storage.providerConfig.s3.disableSSL | bool | `true` |  |
| nebula-binary.configuration.storage.providerConfig.s3.endpoint | string | `"http://{{ printf \"%s-minio\" .Release.Name | trunc 63 | trimSuffix \"-\" }}.{{ .Release.Namespace }}:9000"` |  |
| nebula-binary.configuration.storage.providerConfig.s3.secretKey | string | `"miniostorage"` |  |
| nebula-binary.configuration.storage.providerConfig.s3.v2Signing | bool | `true` |  |
| nebula-binary.configuration.storage.userDataContainer | string | `"my-s3-bucket"` |  |
| nebula-binary.deployment.image.pullPolicy | string | `"Never"` |  |
| nebula-binary.deployment.image.repository | string | `"nebula-binary"` |  |
| nebula-binary.deployment.image.tag | string | `"sandbox"` |  |
| nebula-binary.deployment.waitForDB.image.pullPolicy | string | `"Never"` |  |
| nebula-binary.deployment.waitForDB.image.repository | string | `"bitnami/postgresql"` |  |
| nebula-binary.deployment.waitForDB.image.tag | string | `"sandbox"` |  |
| nebula-binary.enabled | bool | `true` |  |
| nebula-binary.nameOverride | string | `"nebula-sandbox"` |  |
| nebula-binary.rbac.extraRules[0].apiGroups[0] | string | `"*"` |  |
| nebula-binary.rbac.extraRules[0].resources[0] | string | `"*"` |  |
| nebula-binary.rbac.extraRules[0].verbs[0] | string | `"*"` |  |
| kubernetes-dashboard.enabled | bool | `true` |  |
| kubernetes-dashboard.extraArgs[0] | string | `"--enable-insecure-login"` |  |
| kubernetes-dashboard.extraArgs[1] | string | `"--enable-skip-login"` |  |
| kubernetes-dashboard.image.pullPolicy | string | `"Never"` |  |
| kubernetes-dashboard.image.tag | string | `"sandbox"` |  |
| kubernetes-dashboard.protocolHttp | bool | `true` |  |
| kubernetes-dashboard.rbac.clusterReadOnlyRole | bool | `true` |  |
| kubernetes-dashboard.rbac.clusterRoleMetrics | bool | `false` |  |
| kubernetes-dashboard.rbac.create | bool | `true` |  |
| kubernetes-dashboard.service.externalPort | int | `80` |  |
| minio.auth.rootPassword | string | `"miniostorage"` |  |
| minio.auth.rootUser | string | `"minio"` |  |
| minio.defaultBuckets | string | `"my-s3-bucket"` |  |
| minio.enabled | bool | `true` |  |
| minio.extraEnvVars[0].name | string | `"MINIO_BROWSER_REDIRECT_URL"` |  |
| minio.extraEnvVars[0].value | string | `"http://localhost:30080/minio"` |  |
| minio.image.pullPolicy | string | `"Never"` |  |
| minio.image.tag | string | `"sandbox"` |  |
| minio.persistence.enabled | bool | `true` |  |
| minio.persistence.existingClaim | string | `"{{ include \"nebula-sandbox.persistence.minioVolumeName\" . }}"` |  |
| minio.service.nodePorts.api | int | `30002` |  |
| minio.service.type | string | `"NodePort"` |  |
| minio.volumePermissions.enabled | bool | `true` |  |
| minio.volumePermissions.image.pullPolicy | string | `"Never"` |  |
| minio.volumePermissions.image.tag | string | `"sandbox"` |  |
| postgresql.auth.postgresPassword | string | `"postgres"` |  |
| postgresql.enabled | bool | `true` |  |
| postgresql.image.pullPolicy | string | `"Never"` |  |
| postgresql.image.tag | string | `"sandbox"` |  |
| postgresql.primary.persistence.enabled | bool | `true` |  |
| postgresql.primary.persistence.existingClaim | string | `"{{ include \"nebula-sandbox.persistence.dbVolumeName\" . }}"` |  |
| postgresql.primary.service.nodePorts.postgresql | int | `30001` |  |
| postgresql.primary.service.type | string | `"NodePort"` |  |
| postgresql.shmVolume.enabled | bool | `false` |  |
| postgresql.volumePermissions.enabled | bool | `true` |  |
| postgresql.volumePermissions.image.pullPolicy | string | `"Never"` |  |
| postgresql.volumePermissions.image.tag | string | `"sandbox"` |  |
| sandbox.buildkit.enabled | bool | `true` |  |
| sandbox.buildkit.image.pullPolicy | string | `"Never"` |  |
| sandbox.buildkit.image.repository | string | `"moby/buildkit"` |  |
| sandbox.buildkit.image.tag | string | `"sandbox"` |  |
| sandbox.dev | bool | `false` |  |
| sandbox.proxy.enabled | bool | `true` |  |
| sandbox.proxy.image.pullPolicy | string | `"Never"` |  |
| sandbox.proxy.image.repository | string | `"envoyproxy/envoy"` |  |
| sandbox.proxy.image.tag | string | `"sandbox"` |  |

