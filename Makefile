build:
	@echo "====> Build server"
	@go build -o ./bin/server server/server.go
.PHONY: build

build.tools:
	@echo "====> Build tools cli"
	@go build -o ./bin/tools ./cmd/mock/main.go
.PHONY: build.tools

build.all: build build.tools
.PHONY: build.all
