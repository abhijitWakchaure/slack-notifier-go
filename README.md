# Slack Notifier

A simple command line utility to send message to your group using slack webhook URL

## Congiguration

Before using this tool make sure you have your slack webhook URL. It should look like this:

```text
https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX
```

Following env variables must be set before using this tool:

| Variable | Where to get it |
---|---
SN_SLACK_WORKSPACE_ID | First ID in your webhook URL, starts with 'T'
SN_SLACK_WEBHOOK_KEY | Second ID in your webhook URL, starts with 'B'
SN_SLACK_WEBHOOK_SECRET | Last ID in your webhook URL
