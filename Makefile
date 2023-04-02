gql-gateway:
	go build -o bin/gql-gateway/main cmd/gql-gateway/main.go

gateway-server:
	go run internal/services/gql-gateway/server.go

protoUser:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/shared/proto/user/user.proto

