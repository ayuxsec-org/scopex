default: build

.PHONY: build
build:
	go build -o scopex ./cmd/scopex/.