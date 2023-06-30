dev: build run

build: build.sh
	chmod +x build.sh
	./build.sh

run: ./dist/slack-notifier-go
	./dist/slack-notifier-go -m "$(m)"

install:
	go install github.com/abhijitWakchaure/slack-notifier-go@latest

clean:
	rm -rf dist