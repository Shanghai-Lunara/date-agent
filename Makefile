.PHONY: proto register client

HARBOR_DOMAIN := $(shell echo ${HARBOR})
PROJECT := lunara-common
SERVER_IMAGE := "$(HARBOR_DOMAIN)/$(PROJECT)/date-agent:latest"

build:
	-i docker image rm $(SERVER_IMAGE)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o date-agent cmd/register/main.go
	cp cmd/register/Dockerfile . && docker build -t $(SERVER_IMAGE) .
	rm -f Dockerfile && rm -f date-agent
	docker push $(SERVER_IMAGE)

proto:
	cd proto && protoc --go_out=plugins=grpc:. *.proto

mod:
	go mod download
	go mod tidy

server:
	go run cmd/register/main.go --grpcservice=0.0.0.0:10000 --httpservice=0.0.0.0:10001 -v=4

client:
	go run cmd/client/main.go --grpcservice=127.0.0.1:10000 -v=4

os:
	go run cmd/client-env-retry/main.go -v=4

build-client:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o date-agent-client cmd/client-env-retry/main.go

