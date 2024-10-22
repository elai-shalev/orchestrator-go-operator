package rhdh

const RHDHConfigTempl = `
app:
  title: Red Hat Developer Hub
  baseUrl: https://backstage-backstage-{{ .TargetNamespace }}.{{ .ClusterDomain }}
backend:
  auth:
	externalAccess:
	  - type: static
		options:
		  token: {{ printf "${%s}" .BackendSecret }}
		  subject: orchestrator
  baseUrl: https://backstage-backstage-{{ .TargetNamespace }}.{{ .ClusterDomain }}
  csp:
	script-src: ["'self'", "'unsafe-inline'", "'unsafe-eval'"]
	script-src-elem: ["'self'", "'unsafe-inline'", "'unsafe-eval'"]
	connect-src: ["'self'", 'http:', 'https:', 'data:']
  cors:
	origin: https://backstage-backstage-{{ .TargetNamespace }}.{{ .ClusterDomain }}
  database:
	client: pg
	connection:
	  password: ${POSTGRESQL_ADMIN_PASSWORD}
	  user: ${POSTGRES_USER}
	  host: ${POSTGRES_HOST}
	  port: ${POSTGRES_PORT}
{{- if .ArgoCDEnabled }}
argocd:
  appLocatorMethods:
  - instances:
	- name: main
	  url: {{ printf "${%s}" .ArgoCDUrl }}
	  username: {{ printf "${%s}" .ArgoCDUsername }}
	  password: {{ printf "${%s}" .ArgoCDPassword }}
	type: config
{{- end }}
`

type RHDHConfig struct {
	TargetNamespace string
	ArgoCDUsername  string
	ArgoCDPassword  string
	ArgoCDUrl       string
	ArgoCDEnabled   bool
	BackendSecret   string
	ClusterDomain   string
}
