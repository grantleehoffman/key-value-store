all: clean fmt deps vet build
deps:
	@dep ensure
fmt:
	@go fmt ./...
vet:
	@go vet ./...

build:
	@rm -rf vendor
	@dep ensure
	@env CGO_ENABLED=0 GOOS=darwin go build -v -o bin/darwin/key-value
	@env CGO_ENABLED=0 GOOS=linux go build -v -o bin/linux/key-value

clean:
	@rm -rf vendor/* bin/*
