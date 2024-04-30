# GoTemplate linting for yaml (Helm) files
This tool will list files based on the provided criteria and format them if specified with options.

## How to install
Use go install command, i.e.:
```bash
go install github.com/imunhatep/gotpl-yaml-linter/cmd/gotpl-linter@latest

gotpl-linter --help
```

Project built using go1.22

## How to Run
Linter tool supports to commands: 
 - lint
 - format

```bash
gotpl-yaml-linter help

NAME:
   gotpl-linter - GoLang template for yaml formatting and linting tool

USAGE:
   gotpl-linter [command] [subcommand] [command options]

VERSION:
   v1.2.0

DESCRIPTION:
   https://github.com/imunhatep/gotpl-yaml-linter/README.md

COMMANDS:
   fmt      yaml tpl format
   lint     yaml gotpl linting
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --verbose value, --vv value  Log verbosity (default: 3) [$APP_DEBUG]
   --help, -h                   show help
   --version, -v                print the version
```


### Lint
Running lint command will test files on proper linting.

```bash
gotpl-yaml-linter lint help

NAME:
   yamltpl_linter lint - yaml tpl linting

USAGE:
   Example: bin/yamltpl_{os}-{arch} -vv 10 lint -p ./templates/ -f *.yaml

OPTIONS:
   --path value, -p value    path to go tpl files (default: ./)
   --filter value, -f value  filter files by pattern (default: "*")
   --show, -s                output expected formatting (default: false)
   --help, -h                show help
```


### Format
Executing format command will update file contents in place.

```bash
gotpl-linter fmt help

NAME:
   gotpl-linter fmt - yaml tpl format

USAGE:
   Example: gotpl-linter -vv 10 fmt -p ./templates/ -f *.yaml

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
```gotemplate
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
{{- include  "isControllerTagValid" . -}}
    {{- include "ingress-nginx.labels" . | nindent 4 }}
{{- end }}
```
#### Output
```gotemplate
{{- if or (eq .Values.controller.kind "Deployment") (eq .Values.controller.kind "Both") -}}
  {{- include  "isControllerTagValid" . -}}
  {{- include "ingress-nginx.labels" . | nindent 4 }}
{{- end }}
```

### Example 2
#### Input
```gotemplate
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
```gotemplate
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
```gotemplate
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
```gotemplate
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