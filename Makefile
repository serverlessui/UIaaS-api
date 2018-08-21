SRC_LOCATION=

default: clean dependencies test build

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/dnshandler dnshandler/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/buckethandler buckethandler/main.go

.PHONY: dependencies
dependencies:
	npm install -g serverless
	@go get github.com/tools/godep
	dep ensure -v

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: 
	sls deploy --verbose
	
test: test-all

.PHONY: test-all
test-all:
	@go test -v -cover ./$(SRC_LOCATION)/...

.PHONY: test-min
test-min:
	@go test ./...