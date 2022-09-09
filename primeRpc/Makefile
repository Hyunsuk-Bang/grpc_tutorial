
generate:
	protoc -Iproto --go_opt=module=primeRpc --go_out=. \
	--go-grpc_opt=module=primeRpc --go-grpc_out=. proto/*.proto

build:
	go build -o ./bin/server ./server
	go build -o ./bin/client ./client

PHONY: generate, build