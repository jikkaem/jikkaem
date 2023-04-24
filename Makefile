gql-gateway:
	go build -o bin/gql-gateway/main cmd/gql-gateway/main.go

protocUser:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/user/user.proto

protocFancam:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/fancam/fancam.proto

protocSearch:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/search/search.proto

scrape:
	python3 services/scraper/main.py

pyProtocFancam:
	python3 -m grpc_tools.protoc -Iproto/fancam --python_out=services/scraper --pyi_out=services/scraper --grpc_python_out=services/scraper proto/fancam/fancam.proto 

devUser:
	air --build.cmd "go build -o bin/user cmd/user/main.go" --build.bin "./bin/user"

devFancam:
	air --build.cmd "go build -o bin/fancam cmd/fancam/main.go" --build.bin "./bin/fancam"

devGw:
	air --build.cmd "go build -o bin/gateway services/gql-gateway/server.go" --build.bin "./bin/gateway"
