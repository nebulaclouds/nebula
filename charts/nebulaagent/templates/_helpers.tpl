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


{{- define "nebulaagent.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" -}}
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

# Optional blocks for secret mount

{{- define "agentSecret.volume" -}}
- name: {{ include "nebula.name" . }}
  secret:
    secretName: {{ include "nebula.name" . }}
{{- end }}

{{- define "agentSecret.volumeMount" -}}
- mountPath: /etc/secrets
  name: {{ include "nebula.name" . }}
{{- end }}

{{- define "nebulaagent.servicePort" -}}
{{ include .Values.ports.containerPort}}
{{- end }}