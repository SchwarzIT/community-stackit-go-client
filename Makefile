TEST?=$$(go list ./... | grep -v -E 'services|vendor|config|examples|internal')

test: 
	@GOPRIVATE=dev.azure.com go test $(TEST) || exit 1                                                   
	@echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=1

coverage:
	@go test $(TEST) -coverprofile cover.out
	@go tool cover -func cover.out | grep total:
	@rm cover.out
	
quality:
	@goreportcard-cli -v ./...

generate:
	GOPRIVATE=dev.azure.com go generate ./internal/config/...