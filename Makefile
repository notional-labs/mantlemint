all: install

build:
	@echo "Building mantlemint"
	@go build -mod readonly $(BUILD_FLAGS) -o build/mantlemint cmd/mantlemint/main.go

install:
	@echo "Installing mantlemint"
	@go install -mod readonly $(BUILD_FLAGS) ./...

.PHONY: all lint test race msan tools clean build