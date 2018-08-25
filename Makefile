SRC_LOCATION=
IAAS_PATH=iaas/cloudformation/
IAAS_FILE=iaas.go
IAAS_LOCATION=iaas

default: clean dependencies test build

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/dnshandler dnshandler/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/buckethandler buckethandler/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/stackhandler stackhandler/main.go

.PHONY: dependencies
dependencies: bin-data
	npm install -g serverless
	@go get github.com/tools/godep
	dep ensure -v

.PHONY: clean
clean:
	rm -rf ./bin ./vendor Gopkg.lock

.PHONY: deploy
deploy: 
	sls deploy --verbose

.PHONY: bin-data
bin-data:
	./go-bindata -prefix $(IAAS_PATH) -pkg $(IAAS_LOCATION) -o $(IAAS_LOCATION)/$(IAAS_FILE) $(IAAS_PATH)

test: test-all

.PHONY: test-all
test-all:
	@go test -v -cover ./$(SRC_LOCATION)/...

.PHONY: test-min
test-min:
	@go test ./...