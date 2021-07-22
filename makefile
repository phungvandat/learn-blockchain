dev:
	@go run *.go

init: 
	@go mod tidy

install: init
	@go install