.PHONY: default
default:
	go install .

.PHONY: docs
docs:
	swagger generate spec -o docs/swagger.yaml --scan-models
