### Command-line Flags

- `-path`: Specify the path for listing files (default is current directory).
- `-filter`: Define the pattern for matching in directories (default is "*.yaml").
- `-fmt`: Format files in place if set to true.

The program will list files based on the provided criteria and format them if specified.

## How to Run
Linter tool supports to commands: lint and format

```bash
gotpl-yaml-linter help

NAME:
   yamltpl_linter - GoLang template for yaml formatting and linting tool

USAGE:
   bin/yamltpl_linter_{os}-{arch} [command] [subcommand] [command options]

DESCRIPTION:
   https://github.com/imunhatep/gotpl-yaml-linter/README.md

COMMANDS:
   fmt      yaml tpl format
   lint     yaml tpl linting
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose value, -v value  Log verbosity (default: 3) [$APP_DEBUG]
   --help, -h                 show help
```


### Lint
Running lint command will test files on proper linting.

```bash
gotpl-yaml-linter lint help

NAME:
   yamltpl_linter lint - yaml tpl linting

USAGE:
   Example: bin/yamltpl_{os}-{arch} -v 10 lint --path ./templates/ --filter *.yaml

OPTIONS:
   --path value, -p value    path to go tpl files (default: ./)
   --filter value, -f value  filter files by pattern (default: "*")
   --show, -s                output expected formatting (default: false)
   --help, -h                show help
```


### Format
Executing format command will update file contents in place.

```bash
gotpl-yaml-linter fmt help

NAME:
   yamltpl_linter fmt - yaml tpl format

USAGE:
   Example: bin/yamltpl_{os}-{arch} -v 10 fmt --path ./templates/ --filter *.yaml

OPTIONS:
   --path value, -p value    path to go tpl files (default: ./)
   --filter value, -f value  filter files by pattern (default: "*")
   --show, -s                output expected formatting (default: false)
   --help, -h                show help
```


## Examples
Examples of formatting yaml tpl files

### Example 1
#### Input
```
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
{{- include  "isControllerTagValid" . -}}
    {{- include "ingress-nginx.labels" . | nindent 4 }}
{{- end }}
```
#### Output
```
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
  {{- include  "isControllerTagValid" . -}}
  {{- include "ingress-nginx.labels" . | nindent 4 }}
{{- end }}
```

### Example 2
#### Input
```yaml
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
{{- include  "isControllerTagValid" . -}}
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    {{- include "ingress-nginx.labels" . | nindent 4 }}
    app.kubernetes.io/component: controller
    {{- with .Values.controller.labels }}
    {{- toYaml . | nindent 4 }}
    {{- end }}
  name: {{ include "ingress-nginx.controller.fullname" . }}
  namespace: {{ .Release.Namespace }}
  {{- if .Values.controller.annotations }}
  annotations: {{ toYaml .Values.controller.annotations | nindent 4 }}
  {{- end }}
{{- end }}
```

#### Output
```yaml
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
  {{- include  "isControllerTagValid" . -}}
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    {{- include "ingress-nginx.labels" . | nindent 4 }}
    app.kubernetes.io/component: controller
    {{- with .Values.controller.labels }}
    {{- toYaml . | nindent 4 }}
    {{- end }}
  name: {{ include "ingress-nginx.controller.fullname" . }}
  namespace: {{ .Release.Namespace }}
  {{- if .Values.controller.annotations }}
  annotations: {{ toYaml .Values.controller.annotations | nindent 4 }}
  {{- end }}
{{- end }}
```
### Example 3
#### Input
```yaml
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
{{- include  "isControllerTagValid" . -}}
    {{- include "ingress-nginx.labels" . | nindent 4 }}
    {{- with .Values.controller.labels }}
    {{- toYaml . | nindent 4 }}
    {{- end }}
  {{- if .Values.controller.annotations }}
    {{- with .Values.controller.labels }}
    {{- toYaml . | nindent 8 }}
    {{- end }}
  {{- end }}
      {{- include "ingress-nginx.selectorLabels" . | nindent 6 }}
  {{- if not .Values.controller.autoscaling.enabled }}
  {{- end }}
{{- end }}
```

#### Output
```yaml
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
  {{- include  "isControllerTagValid" . -}}
  {{- include "ingress-nginx.labels" . | nindent 4 }}
  {{- with .Values.controller.labels }}
    {{- toYaml . | nindent 4 }}
  {{- end }}
  {{- if .Values.controller.annotations }}
    {{- with .Values.controller.labels }}
      {{- toYaml . | nindent 8 }}
    {{- end }}
  {{- end }}
  {{- include "ingress-nginx.selectorLabels" . | nindent 6 }}
  {{- if not .Values.controller.autoscaling.enabled }}
  {{- end }}
{{- end }}
```