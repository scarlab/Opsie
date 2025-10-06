server:
	@go run ./cmd/server/main.go

agent:
	@go run ./cmd/agent/main.go

web:
	@cd ui && npm run dev && cd .. 

build-server:
	@go build -o bin/server/opsie cmd/server/main.go

build-agent:
	@go build -o bin/agent/opsie-agent cmd/agent/main.go

build:
	@go build -o bin/server/opsie cmd/server/main.go
	@go build -o bin/agent/opsie-agent cmd/agent/main.go



domain:
	@go run cmd/cli/main.go create-domain $(filter-out $@,$(MAKECMDGOALS))

migration:
	@migrate create -ext sql -dir db/migrations $(filter-out $@,$(MAKECMDGOALS))

mg-force:
	@go run cmd/cli/main.go migrate force $(filter-out $@,$(MAKECMDGOALS))

mg-up:
	@go run cmd/cli/main.go migrate up

mg-down:
	@go run cmd/cli/main.go migrate down