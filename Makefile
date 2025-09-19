server:
	@go run ./cmd/server/main.go

agent:
	@go run ./cmd/agent/main.go

web:
	@cd ui && npm run dev && cd .. 

build-server:
	@go build -o bin/server/watchtower cmd/server/main.go

build-agent:
	@go build -o bin/agent/watchtower-agent cmd/agent/main.go

build:
	@go build -o bin/server/watchtower cmd/server/main.go
	@go build -o bin/agent/watchtower-agent cmd/agent/main.go


domain:
	@go run ./cmd/cli/main.go new-domain $(n)