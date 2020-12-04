.PHONY: proto register client

proto:
	cd proto && protoc --go_out=plugins=grpc:. *.proto

mod:
	go mod download
	go mod tidy

server:
	go run cmd/register/main.go --grpcservice=0.0.0.0:10000 --httpservice=0.0.0.0:10001 -v=4

client:
	go run cmd/client/main.go --grpcservice=127.0.0.1:10000 -v=4