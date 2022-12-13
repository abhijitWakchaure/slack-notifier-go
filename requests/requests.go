package requests

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/abhijitWakchaure/slack-notifier/constants"
	"github.com/abhijitWakchaure/slack-notifier/env"
)

// POST ...
func POST(msg string) {
	rawURL := fmt.Sprintf(constants.SlackWebhookBase, env.Read(env.SlackWorkspaceID), env.Read(env.SlackWebhookKey), env.Read(env.SlackWebhookSecret))
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
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
		panic(err)
	}
	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("Expected status code 200 but received %d", res.StatusCode)
		panic(err)
	}
}
