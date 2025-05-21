package rhdh

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/rhdhorchestrator/orchestrator-operator/api/v1alpha3"
)

func ConfigMapTemplateFactory(
	cmTemplateType, clusterDomain, serverlessWorkflowNamespace string,
	argoCDEnabled, tektonEnabled bool,
	rhdhConfig v1alpha3.RHDHConfig) (string, error) {
	switch cmTemplateType {
	case AppConfigRHDHName:
		configData := RHDHConfig{
			RHDHNamespace:  rhdhConfig.Namespace,
			RHDHName:       rhdhConfig.Name,
			ArgoCDUsername: ArgoCDUsername,
			ArgoCDPassword: ArgoCDPassword,
			ArgoCDUrl:      ArgoCDUrl,
			ArgoCDEnabled:  argoCDEnabled,
			BackendSecret:  BackendSecretKey,
			ClusterDomain:  clusterDomain,
		}
		formattedConfig, err := parseConfigTemplate(RHDHConfigTempl, configData)
		if err != nil {
			return "", err
		}
		return formattedConfig, nil
	case AppConfigRHDHAuthName:
		configData := RHDHConfigAuth{
			GitHubToken:         GitHubToken,
			Environment:         "development",
			GitHubClientId:      GitHubClientID,
			GitHubClientSecret:  GitHubClientSecret,
			GitLabToken:         GitLabToken,
			GitLabHost:          GitLabHost,
			EnableGuestProvider: rhdhConfig.DevMode,
		}
		formattedConfig, err := parseConfigTemplate(RHDHAuthTempl, configData)
		if err != nil {
			return "", err
		}
		return formattedConfig, nil
	case AppConfigRHDHCatalogName:
		configData := RHDHConfigCatalog{
			EnableGuestProvider: rhdhConfig.DevMode,
			CatalogBranch:       CatalogBranch,
		}
		formattedConfig, err := parseConfigTemplate(RHDHCatalogTempl, configData)
		if err != nil {
			return "", err
		}
		return formattedConfig, nil
	case AppConfigRHDHDynamicPluginName:
		pluginsMap := getPlugins()
		configData := RHDHDynamicPluginConfig{
			K8ClusterToken:                         ClusterUrl,
			K8ClusterUrl:                           ClusterToken,
			TektonEnabled:                          tektonEnabled,
			ArgoCDEnabled:                          argoCDEnabled,
			ArgoCDUrl:                              ArgoCDUrl,
			ArgoCDUsername:                         ArgoCDUsername,
			ArgoCDPassword:                         ArgoCDPassword,
			OrchestratorBackendPackage:             pluginsMap[OrchestratorBackend].Package,
			OrchestratorBackendIntegrity:           pluginsMap[OrchestratorBackend].Integrity,
			OrchestratorPackage:                    pluginsMap[Orchestrator].Package,
			OrchestratorIntegrity:                  pluginsMap[Orchestrator].Integrity,
			Scope:                                  Scope,
			NotificationEmailEnabled:               rhdhConfig.RHDHPlugins.NotificationsConfig.Enabled,
			NotificationEmailHostname:              NotificationHostname,
			NotificationEmailUsername:              NotificationUsername,
			NotificationEmailPassword:              NotificationPassword,
			NotificationEmailSender:                rhdhConfig.RHDHPlugins.NotificationsConfig.Sender,
			NotificationEmailReplyTo:               rhdhConfig.RHDHPlugins.NotificationsConfig.Recipient,
			NotificationEmailPort:                  rhdhConfig.RHDHPlugins.NotificationsConfig.Port,
			WorkflowNamespace:                      serverlessWorkflowNamespace,
			ScaffolderBackendOrchestratorPackage:   pluginsMap[ScaffolderBackendOrchestrator].Package,
			ScaffolderBackendOrchestratorIntegrity: pluginsMap[ScaffolderBackendOrchestrator].Integrity,
			OrchestratorFormWidgetsPackage:         pluginsMap[OrchestratorFormWidgets].Package,
			OrchestratorFormWidgetsIntegrity:       pluginsMap[OrchestratorFormWidgets].Integrity,
		}
		formattedConfig, err := parseConfigTemplate(RHDHDynamicPluginTempl, configData)
		if err != nil {
			return "", err
		}
		return formattedConfig, nil
	default:
		return "", nil
	}
}

func parseConfigTemplate(templateString string, configData any) (string, error) {
	// parse the template
	templ, err := template.New("config").Parse(templateString)
	if err != nil {
		fmt.Printf("Error occurred when parsing template: %v\n", err)
		return "", err
	}

	// execute template with the dynamic data
	var output bytes.Buffer
	err = templ.Execute(&output, configData)
	if err != nil {
		fmt.Printf("Error occurred when executing template: %v\n", err)
		return "", err
	}
	return output.String(), nil
}
