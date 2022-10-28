TEST?=$$(go list ./... | grep -v 'vendor')

test: 
	@go test $(TEST) || exit 1                                                   
	@echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4

coverage:
	@go test ./... -coverprofile cover.out
	@go tool cover -func cover.out | grep total:
	@rm cover.out
	
quality:
	@goreportcard-cli -v ./...
