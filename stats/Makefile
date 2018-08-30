SRC_LOCATION=

default: clean dependencies test build

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/statisticshandler main.go

.PHONY: dependencies
dependencies:
	sudo npm install -g serverless
	@go get github.com/tools/godep
	dep ensure -v

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy 
deploy: 
	sls deploy --verbose -s $(ENVIRONMENT)

.PHONY: delete 
delete: 
	sls delete --verbose -s $(ENVIRONMENT)

test: test-all

.PHONY: test-all
test-all:
	@go test -v -cover ./$(SRC_LOCATION)/...

.PHONY: test-min
test-min:
	@go test ./...