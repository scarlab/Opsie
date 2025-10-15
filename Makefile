PROJECT_NAME := $(shell grep '^name:' app.config.yaml | awk '{print $$2}')
VERSION := $(shell grep '^version:' app.config.yaml | awk '{print $$2}')

v:
	@echo ${PROJECT_NAME} $(VERSION)

dev:
	@docker compose up
	
log-server:
	@docker logs opsie-server-1 -f

log-ui:
	@docker logs opsie-ui-1 -f


agent:
	@go run ./cmd/agent/main.go

build-agent:
	@echo "Building agent $(VERSION)"
	@go build -o bin/agent/opsie-agent cmd/agent/main.go



domain:
	@go run cmd/cli/main.go create-domain $(filter-out $@,$(MAKECMDGOALS))

domain-ws:
	@go run cmd/cli/main.go create-domain $(filter-out $@,$(MAKECMDGOALS)) --ws


migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

mg-force:
	@go run cmd/cli/main.go migrate force $(filter-out $@,$(MAKECMDGOALS))

mg-up:
	@go run cmd/cli/main.go migrate up

mg-down:
	@go run cmd/cli/main.go migrate down