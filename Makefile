gql-gateway:
	go build -o bin/gql-gateway/main cmd/gql-gateway/main.go

protocUser:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    internal/shared/proto/user/user.proto

devUser:
	air --build.cmd "go build -o bin/user cmd/user/main.go" --build.bin "./bin/user"

devGw:
	air --build.cmd "go build -o bin/gateway internal/services/gql-gateway/server.go" --build.bin "./bin/gateway"
