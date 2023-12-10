{{/* vim: set filetype=mustache: */}}

{{- define "nebula.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "nebula.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" -}}
{{- end -}}

{{- define "nebula.namespace" -}}
{{- default .Release.Namespace .Values.forceNamespace | trunc 63 | trimSuffix "-" -}}
{{- end -}}


{{- define "nebulaadmin.name" -}}
nebulaadmin
{{- end -}}

{{- define "nebulaadmin.selectorLabels" -}}
app.kubernetes.io/name: {{ template "nebulaadmin.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "nebulaadmin.labels" -}}
{{ include "nebulaadmin.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "nebulascheduler.name" -}}
nebulascheduler
{{- end -}}

{{- define "nebulascheduler.selectorLabels" -}}
app.kubernetes.io/name: {{ template "nebulascheduler.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}


{{- define "nebulascheduler.labels" -}}
{{ include "nebulascheduler.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "nebulaclusterresourcesync.name" -}}
nebulaclusterresourcesync
{{- end -}}

{{- define "nebulaclusterresourcesync.selectorLabels" -}}
app.kubernetes.io/name: {{ template "nebulaclusterresourcesync.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "nebulaclusterresourcesync.labels" -}}
{{ include "nebulaclusterresourcesync.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "datacatalog.name" -}}
datacatalog
{{- end -}}

{{- define "datacatalog.selectorLabels" -}}
app.kubernetes.io/name: {{ template "datacatalog.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "datacatalog.labels" -}}
{{ include "datacatalog.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "nebulaagent.name" -}}
nebulaagent
{{- end -}}

{{- define "nebulaagent.selectorLabels" -}}
app.kubernetes.io/name: {{ template "nebulaagent.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "nebulaagent.labels" -}}
{{ include "nebulaagent.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "nebulapropeller.name" -}}
nebulapropeller
{{- end -}}

{{- define "nebulapropeller.selectorLabels" -}}
app.kubernetes.io/name: {{ template "nebulapropeller.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "nebulapropeller.labels" -}}
{{ include "nebulapropeller.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "nebulapropeller-manager.name" -}}
nebulapropeller-manager
{{- end -}}

{{- define "nebulapropeller-manager.selectorLabels" -}}
app.kubernetes.io/name: {{ template "nebulapropeller-manager.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "nebulapropeller-manager.labels" -}}
{{ include "nebulapropeller-manager.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

{{- define "nebula-pod-webhook.name" -}}
nebula-pod-webhook
{{- end -}}


{{- define "nebulaconsole.name" -}}
nebulaconsole
{{- end -}}

{{- define "nebulaconsole.selectorLabels" -}}
app.kubernetes.io/name: {{ template "nebulaconsole.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end -}}

{{- define "nebulaconsole.labels" -}}
{{ include "nebulaconsole.selectorLabels" . }}
helm.sh/chart: {{ include "nebula.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end -}}

# Optional blocks for secret mount

{{- define "databaseSecret.volume" -}}
{{- with .Values.common.databaseSecret.name -}}
- name: {{ . }}
  secret:
    secretName: {{ . }}
{{- end }}
{{- end }}

{{- define "databaseSecret.volumeMount" -}}
{{- with .Values.common.databaseSecret.name -}}
- mountPath: /etc/db
  name: {{ . }}
{{- end }}
{{- end }}

{{- define "storage.base" -}}
storage:
{{- if eq .Values.storage.type "s3" }}
  type: s3
  container: {{ .Values.storage.bucketName | quote }}
  connection:
    auth-type: {{ .Values.storage.s3.authType }}
    region: {{ .Values.storage.s3.region }}
    {{- if eq .Values.storage.s3.authType "accesskey" }}
    access-key: {{ .Values.storage.s3.accessKey }}
    secret-key: {{ .Values.storage.s3.secretKey }}
    {{- end }}
{{- else if eq .Values.storage.type "gcs" }}
  type: stow
  stow:
    kind: google
    config:
      json: ""
      project_id: {{ .Values.storage.gcs.projectId }}
      scopes: https://www.googleapis.com/auth/cloud-platform
  container: {{ .Values.storage.bucketName | quote }}
{{- else if eq .Values.storage.type "sandbox" }}
  type: minio
  container: {{ .Values.storage.bucketName | quote }}
  stow:
    kind: s3
    config:
      access_key_id: minio
      auth_type: accesskey
      secret_key: miniostorage
      disable_ssl: true
      endpoint: http://minio.{{ .Release.Namespace }}.svc.cluster.local:9000
      region: us-east-1
  signedUrl:
    stowConfigOverride:
      endpoint: http://localhost:30084
{{- else if eq .Values.storage.type "custom" }}
{{- with .Values.storage.custom -}}
  {{ tpl (toYaml .) $ | nindent 2 }}
{{- end }}
{{- end }}
{{- end }}

{{- define "storage" -}}
{{ include "storage.base" .}}
  enable-multicontainer: {{ .Values.storage.enableMultiContainer }}
  limits:
    maxDownloadMBs: {{ .Values.storage.limits.maxDownloadMBs }}
{{- end }}
