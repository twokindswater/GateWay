PWD := $(shell pwd)
DOCKERS := $(shell docker ps -qa)
docker:
	docker rm -f $(DOCKERS)
	docker run --rm -d --name redis-test -p 6379:6379 -v ~/tmp/redis:/data redis

build:
	@echo "Build homebody binary to './bin/homebody'"
	@go build -mod vendor -o $(PWD)/bin/homebody cmd/homebody/main.go

test: docker build
	@echo "Running unit tests"
	@go clean -testcache
	@go test -cover -mod vendor -parallel 1 -p 1 -failfast -timeout 300s ./...