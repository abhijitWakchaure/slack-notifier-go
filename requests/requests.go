package requests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/abhijitWakchaure/slack-notifier-go/constants"
	"github.com/abhijitWakchaure/slack-notifier-go/env"
)

// POST ...
func POST(msg string) {
	rawURL := fmt.Sprintf(constants.SlackWebhookBase, env.Read(env.SlackWorkspaceID), env.Read(env.SlackWebhookKey), env.Read(env.SlackWebhookSecret))
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("Failed to parse raw webhook url: [%s] due to %s\n", rawURL, err.Error())
		os.Exit(1)
	}
	payload := `{
		"blocks": [
			{
				"type": "section",
				"text": {
				"type": "mrkdwn",
				"text": "` + msg + `"
				}
		  	}
		]
	}`
	buff := bytes.NewBufferString(payload)
	res, err := http.Post(u.String(), "application/json", buff)
	if err != nil {
		fmt.Printf("Failed to POST webhook url due to %s\n", err.Error())
		os.Exit(1)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Printf("Expected status code 200 but received %d\n", res.StatusCode)
		os.Exit(1)
	}
}
