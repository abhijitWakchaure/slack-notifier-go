# Slack Notifier Go

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


## How to create your own slack app

1. Go to [https://api.slack.com/apps](https://api.slack.com/apps) and `Create New App` using `From Scratch`

1. Under `Add features and functionality` choose `Incoming Webhooks` and turn on the `Activate Incoming Webhooks` option

1. Create a new webhook URL for your app by clicking the button `Add New Webhook to Workspace`

1. Grant the permission to post in any channel of your choosing and copy the webhook URL. It should look like this:

    ```text
    https://hooks.slack.com/services/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
    ```

1. Export the webhook URL as `SLACK_WEBHOOK_URL` using command:

    ```bash
    export SLACK_WEBHOOK_URL=<your slack webhook URL>
    ```

1. To test if your slack app is working correctly you can use the following command to send `Hello, World!` message to your channel:

    ```bash
    curl -X POST -H 'Content-type: application/json' --data '{"text":"Hello, World!"}' <your slack webhook URL>
    ```

    You should see the `ok` as output and your slack channel should receive the `Hello, World!` message.
