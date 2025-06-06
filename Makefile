fmt:
	@go fmt ./...

run:
	@go run main.go

tidy:
	@go mod tidy -v

test:
	@go test -v ./...
	
test_with_coverage:
	@rm -f ${pwd}/coverage.out
	@go test -v -coverprofile=coverage.out ./...

test_with_coverage_html: test_with_coverage
	@go tool cover -html=coverage.out