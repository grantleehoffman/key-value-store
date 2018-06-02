all: clean deps vet build
deps:
	@dep ensure

vet:
	@go vet ./...

build:
	@rm -rf vendor
	@dep ensure
	@go build -o bin/key-value

clean:
	@rm -rf vendor/* bin/*
