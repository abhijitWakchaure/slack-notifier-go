#!/usr/bin/env bash

APP_NAME="slack-notifier-go"
APP_VERSION=$(cat VERSION)
LDFLAGS="-s -w -X 'github.com/abhijitWakchaure/slack-notifier-go/cmd.VERSION=${APP_VERSION}'"

echo "Building binary for ${APP_NAME}-v${APP_VERSION}"

export CGO_ENABLED=0

go build -ldflags "${LDFLAGS}" -o dist/${APP_NAME} || {
    echo "Failed to build app binary"
    exit 1
}

echo "You can now execute the app binary by running command ./dist/${APP_NAME} or make run m=\"Hello world\""