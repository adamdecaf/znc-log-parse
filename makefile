.PHONY: build test

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/znc-log-parse-linux-amd64 .
	GOOS=darwin GOARCH=amd64 go build -o bin/znc-log-parse-osx-amd64 .

test:
	go test -v .
