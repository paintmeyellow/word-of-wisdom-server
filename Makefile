server:
	TCP_PORT=8080 go run cmd/server/main.go

client:
	SERVER_ADDR=:8080 go run cmd/client/main.go

test:
	go test ./internal/entity ./internal/usecase -count 1