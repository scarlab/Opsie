PROJECT_NAME := $(shell grep '^name:' app.config.yaml | awk '{print $$2}')
VERSION := $(shell grep '^version:' app.config.yaml | awk '{print $$2}')

v:
	@echo ${PROJECT_NAME} $(VERSION)

serve:
	OPSIE_ENV=production
	@./bin/server/opsie

dev:
	@docker compose -f docker-compose.dev.yml up -d
	
dev-down:
	@docker stop opsie-server-1 opsie-ui-1
	
slog:
	@docker logs opsie-server-1 -f

ulog:
	@docker logs opsie-ui-1 -f


agent:
	@go run ./cmd/agent/main.go


build-agent:
	@echo "Building Agent $(VERSION)"
	@go build -o bin/agent/opsie-agent cmd/agent/main.go


build-server:
	@echo "Building Server $(VERSION)"
	@cd ui && VITE_APP_VERSION=$(VERSION) VITE_APP_ENV=production npm run build && cd ..
	@go build -o bin/server/opsie cmd/server/main.go


build:
	@echo "Building Opsie $(VERSION)"
	@cd ui && VITE_APP_VERSION=$(VERSION) VITE_APP_ENV=production npm run build && cd ..
	@go build -o bin/server/opsie cmd/server/main.go
	@go build -o bin/agent/opsie-agent cmd/agent/main.go



api:
	@go run cmd/cli/main.go api create $(filter-out $@,$(MAKECMDGOALS))

api-ws:
	@go run cmd/cli/main.go api create $(filter-out $@,$(MAKECMDGOALS)) --ws

api-delete:
	@go run cmd/cli/main.go api delete $(filter-out $@,$(MAKECMDGOALS))

