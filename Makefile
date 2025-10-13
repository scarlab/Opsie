PROJECT_NAME := $(shell grep '^name:' app.config.yaml | awk '{print $$2}')
VERSION := $(shell grep '^version:' app.config.yaml | awk '{print $$2}')

v:
	@echo ${PROJECT_NAME} $(VERSION)
	
server:
	@go run ./cmd/server/main.go

agent:
	@go run ./cmd/agent/main.go

web:
	@cd ui && VITE_APP_VERSION=$(VERSION) VITE_APP_ENV=development npm run dev && cd .. 

build-server:
	@echo "Building server $(VERSION)"
	@go build -o bin/server/opsie cmd/server/main.go

build-agent:
	@echo "Building agent $(VERSION)"
	@go build -o bin/agent/opsie-agent cmd/agent/main.go

build:
	@echo "Building Opsie $(VERSION)"
	@cd ui && VITE_APP_VERSION=$(VERSION) VITE_APP_ENV=production npm run build && cd ..
	@go build -o bin/server/opsie cmd/server/main.go
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