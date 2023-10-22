GOPATH?=$(shell go env GOPATH)

.PHONY: default
default: build

.PHONY: build
build:
	go build -v

# Exclude some packages from coverage
# * test helpers
# * generated mocks
# * generated code
COVERPKG=go list ./... | grep -v /ent | tr "\n" ","
.PHONY: test
test:
	go test ./... -coverprofile=coverage.out -coverpkg=`$(COVERPKG)`
	go tool cover -func coverage.out

.PHONY: lint
lint:
	golangci-lint run

.PHONY: coverage_report
coverage_report:
	go test ./... -coverprofile=coverage.out -coverpkg=`$(COVERPKG)`
	go tool cover -html coverage.out
