{{- define "ambassador.mapping" -}}
{{- range .Values.ambassador }}
---
apiVersion: ambassador/v1
kind: Mapping
name: {{ $.Values.name }}-ambassador-mapping
prefix: {{ .prefix }}
{{- if .rewrite }}
rewrite: {{ .rewrite }}
{{- end }}
service: "{{ $.Values.name }}.{{ $.Release.Namespace }}:{{ $.Values.servicePort }}"
{{- if .host }}
host: {{ .host }}
{{- end }}
{{- if .prefix_regex }}
prefix_regex: true
{{- end }}
{{- if .tls }}
tls: true
{{- end }}
{{- if .idle_timeout_ms }}
idle_timeout_ms: {{ .idle_timeout_ms }}
{{- end }}
{{- if .rate_limits }}
rate_limits: {{ .rate_limits }}
{{- end }}
{{- if .timeout_ms }}
timeout_ms: {{ .timeout_ms }}
{{- end }}
{{- if .use_websocket }}
use_websocket: true
{{- end }}
{{- if .grpc }}
grpc: true
{{- end }}
{{- if .method }}
method: {{ .method }}
{{- end }}
{{- if .retry_policy }}
retry_policy:
{{- if .retry_policy.retry_on }}
  retry_on: {{ .retry_policy.retry_on }}
{{- end }}
{{- if .retry_policy.num_retries }}
  num_retries: {{ .retry_policy.num_retries }}
{{- end }}
{{- if .retry_policy.per_try_timeout }}
  per_try_timeout: {{ .retry_policy.per_try_timeout }}
{{- end }}
{{- end }}
{{- end }}
{{- end -}}
