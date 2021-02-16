.PHONY: default
default:
	go install .

.PHONY: docs
docs:
	swagger generate spec -o docs/swagger.yaml --scan-models

.PHONY: build
build:
	go build -v ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: coverage
coverage:
	go test -cover -v ./...
