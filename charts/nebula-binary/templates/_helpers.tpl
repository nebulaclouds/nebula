{{/*
Expand the name of the chart.
*/}}
{{- define "nebula-binary.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "nebula-binary.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "nebula-binary.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Base labels
*/}}
{{- define "nebula-binary.baseLabels" -}}
app.kubernetes.io/name: {{ include "nebula-binary.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "nebula-binary.labels" -}}
helm.sh/chart: {{ include "nebula-binary.chart" . }}
{{ include "nebula-binary.baseLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "nebula-binary.selectorLabels" -}}
{{ include "nebula-binary.baseLabels" . }}
app.kubernetes.io/component: nebula-binary
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "nebula-binary.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "nebula-binary.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Flag to use external configuration.
*/}}
{{- define "nebula-binary.configuration.externalConfiguration" -}}
{{- or .Values.configuration.externalConfigMap .Values.configuration.externalSecretRef -}}
{{- end -}}

{{/*
Get the Nebula configuration ConfigMap name.
*/}}
{{- define "nebula-binary.configuration.configMapName" -}}
{{- printf "%s-config" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula configuration Secret name.
*/}}
{{- define "nebula-binary.configuration.configSecretName" -}}
{{- printf "%s-config-secret" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula user data prefix.
*/}}
{{- define "nebula-binary.configuration.storage.userDataPrefix" -}}
{{- $userDataContainer := required "User data container required" .Values.configuration.storage.userDataContainer -}}
{{- if eq "s3" .Values.configuration.storage.provider -}}
{{- printf "s3://%s/data" $userDataContainer -}}
{{- else if eq "gcs" .Values.configuration.storage.provider -}}
{{- printf "gs://%s/data" $userDataContainer -}}
{{- else if eq "azure" .Values.configuration.storage.provider -}}
{{- printf "abfs://%s/data" $userDataContainer -}}
{{- end -}}
{{- end -}}

{{/*
Get the Nebula logging configuration.
*/}}
{{- define "nebula-binary.configuration.logging.plugins" -}}
{{- with .Values.configuration.logging.plugins -}}
kubernetes-enabled: {{ .kubernetes.enabled }}
{{- if .kubernetes.enabled }}
kubernetes-template-uri: {{ required "Template URI required for Kubernetes logging plugin" .kubernetes.templateUri }}
{{- end }}
cloudwatch-enabled: {{ .cloudwatch.enabled }}
{{- if .cloudwatch.enabled }}
cloudwatch-template-uri: {{ required "Template URI required for CloudWatch logging plugin" .cloudwatch.templateUri }}
{{- end }}
stackdriver-enabled: {{ .stackdriver.enabled }}
{{- if .stackdriver.enabled }}
stackdriver-template-uri: {{ required "Template URI required for stackdriver logging plugin" .stackdriver.templateUri }}
{{- end }}
{{- if .custom }}
templates: {{- toYaml .custom | nindent 2 -}}
{{- end -}}
{{- end -}}
{{- end -}}

{{/*
Get the Secret name for Nebula admin authentication secrets.
*/}}
{{- define "nebula-binary.configuration.auth.adminAuthSecretName" -}}
{{- printf "%s-admin-auth" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Secret name for Nebula authentication client secrets.
*/}}
{{- define "nebula-binary.configuration.auth.clientSecretName" -}}
{{- printf "%s-client-secrets" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula cluster resource templates ConfigMap name.
*/}}
{{- define "nebula-binary.clusterResourceTemplates.configMapName" -}}
{{- printf "%s-cluster-resource-templates" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula HTTP service name
*/}}
{{- define "nebula-binary.service.http.name" -}}
{{- printf "%s-http" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula GRPC service name
*/}}
{{- define "nebula-binary.service.grpc.name" -}}
{{- printf "%s-grpc" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula service HTTP port.
*/}}
{{- define "nebula-binary.service.http.port" -}}
{{- default 8088 .Values.service.ports.http -}}
{{- end -}}

{{/*
Get the Nebula service GRPC port.
*/}}
{{- define "nebula-binary.service.grpc.port" -}}
{{- default 8089 .Values.service.ports.grpc -}}
{{- end -}}

{{/*
Get the Nebula agent service GRPC port.
*/}}
{{- define "nebula-binary.nebulaagent.grpc.port" -}}
{{- default 8000 .Values.service.ports.grpc -}}
{{- end -}}


{{/*
Get the Nebula webhook service name.
*/}}
{{- define "nebula-binary.webhook.serviceName" -}}
{{- printf "%s-webhook" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula webhook secret name.
*/}}
{{- define "nebula-binary.webhook.secretName" -}}
{{- printf "%s-webhook-secret" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the Nebula ClusterRole name.
*/}}
{{- define "nebula-binary.rbac.clusterRoleName" -}}
{{- printf "%s-cluster-role" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Get the name of the Nebula Agent Deployment.
*/}}
{{- define "nebula-binary.agent.name" -}}
{{- printf "%s-agent" (include "nebula-binary.fullname" .) -}}
{{- end -}}

{{/*
Nebula Agent selector labels
*/}}
{{- define "nebula-binary.agent.selectorLabels" -}}
{{ include "nebula-binary.baseLabels" . }}
app.kubernetes.io/component: agent
{{- end }}
