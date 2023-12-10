{{/*
Expand the name of the chart.
*/}}
{{- define "nebula-sandbox.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "nebula-sandbox.fullname" -}}
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
{{- define "nebula-sandbox.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "nebula-sandbox.labels" -}}
helm.sh/chart: {{ include "nebula-sandbox.chart" . }}
{{ include "nebula-sandbox.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "nebula-sandbox.selectorLabels" -}}
app.kubernetes.io/name: {{ include "nebula-sandbox.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "nebula-sandbox.serviceAccountName" -}}
{{- if .Values.serviceAccount.create }}
{{- default (include "nebula-sandbox.fullname" .) .Values.serviceAccount.name }}
{{- else }}
{{- default "default" .Values.serviceAccount.name }}
{{- end }}
{{- end }}

{{/*
Name of inline ConfigMap containing additional configuration or overrides for Nebula
*/}}
{{- define "nebula-sandbox.configuration.inlineConfigMap" -}}
{{- printf "%s-extra-config" .Release.Name -}}
{{- end }}

{{/*
Name of inline ConfigMap containing additional cluster resource templates
*/}}
{{- define "nebula-sandbox.clusterResourceTemplates.inlineConfigMap" -}}
{{- printf "%s-extra-cluster-resource-templates" .Release.Name -}}
{{- end }}

{{/*
Name of PersistentVolume and PersistentVolumeClaim for PostgreSQL database
*/}}
{{- define "nebula-sandbox.persistence.dbVolumeName" -}}
{{- printf "%s-db-storage" .Release.Name -}}
{{- end }}

{{/*
Name of PersistentVolume and PersistentVolumeClaim for Minio
*/}}
{{- define "nebula-sandbox.persistence.minioVolumeName" -}}
{{- printf "%s-minio-storage" .Release.Name -}}
{{- end }}

{{/*
Selector labels for Buildkit
*/}}
{{- define "nebula-sandbox.buildkitSelectorLabels" -}}
{{ include "nebula-sandbox.selectorLabels" . }}
app.kubernetes.io/component: buildkit
{{- end }}

{{/*
Selector labels for Envoy proxy
*/}}
{{- define "nebula-sandbox.proxySelectorLabels" -}}
{{ include "nebula-sandbox.selectorLabels" . }}
app.kubernetes.io/component: proxy
{{- end }}

{{/*
Name of Envoy proxy configmap
*/}}
{{- define "nebula-sandbox.proxyConfigMapName" -}}
{{- printf "%s-proxy-config" .Release.Name -}}
{{- end }}

{{/*
Name of development-mode Nebula headless service
*/}}
{{- define "nebula-sandbox.localHeadlessService" -}}
{{- printf "%s-local" .Release.Name | trunc 63 | trimSuffix "-" -}}
{{- end }}
