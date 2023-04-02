gql-gateway:
	go build -o bin/gql-gateway/main cmd/gql-gateway/main.go

gateway-server:
	go run internal/services/gql-gateway/server.go
