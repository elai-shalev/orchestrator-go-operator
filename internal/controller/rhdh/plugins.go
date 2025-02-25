package rhdh

type Plugin struct {
	Package   string
	Integrity string
}

const Orchestrator string = "orchestrator"
const OrchestratorBackend string = "orchestratorBackend"

func getPlugins() map[string]Plugin {
	return map[string]Plugin{
		Orchestrator: {
			Package:   "backstage-plugin-orchestrator-1.5.0-rc.2.tgz",
			Integrity: "sha512-k+oXawNBQa0TFskAoYvExWZ/EOJ9H4s2+y4ujE+RFzsu7rkm4YmElDIrVYMZhJLRqBhSoHgCdGyn7nSPW20rcg==",
		},
		OrchestratorBackend: {
			Package:   "backstage-plugin-orchestrator-backend-dynamic-1.5.0-rc.2.tgz",
			Integrity: "sha512-TmG54OazZLSuzPFmqQSi11koChBE+T8q0ZA7zVkSZZHZjkxvXy2fjqi4Vozz/2hYDUuXRXMJFJ806ijlsiwUsw==",
		},
	}

}
