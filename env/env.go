package env

import (
	"log"
	"os"
)

// Env configs
const (
	SlackWorkspaceID   = "SN_SLACK_WORKSPACE_ID"
	SlackWebhookKey    = "SN_SLACK_WEBHOOK_KEY"
	SlackWebhookSecret = "SN_SLACK_WEBHOOK_SECRET"
)

var defaults map[string]string = map[string]string{}

// Read ...
func Read(name string, defaultVal ...string) string {
	// Lookup env override
	v, ok := os.LookupEnv(name)
	if ok {
		log.Printf("Environment override detected for [%v]\n", name)
		return v
	}
	// Check user default
	if len(defaultVal) > 0 {
		log.Printf("Inline override detected for [%v]\n", name)
		return defaultVal[0]
	}
	v, ok = defaults[name]
	if !ok {
		log.Fatalf("Failed to lookup value for env var: %s\n", name)
		return ""
	}
	log.Printf("Using default value for [%v]\n", name)
	return v
}
