run:
	@go build && ./microservice-assignment

PACKAGES=$(shell go list ./...|grep -v -e mocks -e config -e app|tr '\n' ',')

test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test ./... -cover -coverprofile=coverage.txt -covermode=count -coverpkg=$(PACKAGES)
	@go tool cover -func=coverage.txt