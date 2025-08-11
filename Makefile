BINARY_NAME=cli-deploy-tool

build:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) main.go
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME).exe main.go
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME)_darwin main.go

run: build
	$(BINARY_NAME).exe