.PHONY: proto register client

DOMAIN := $(shell echo ${DOCKER_REGISTRY_DOMAIN})
PROJECT := $(shell echo ${DOCKER_REGISTRY_PROJECT})
SERVER_IMAGE := "$(DOMAIN)/$(PROJECT)/date-agent:latest"

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o date-agent cmd/register/main.go
	cp cmd/register/Dockerfile . && docker build -t $(SERVER_IMAGE) .
	rm -f Dockerfile && rm -f date-agent
	docker push $(SERVER_IMAGE)
	bash hack/guldan.sh cli $(PROJECT) date-agent latest

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
	#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o date-agent-client cmd/client-env-retry/main.go
	cp cmd/hang/Dockerfile . &
	docker build -t local-date-agent-build:latest .
	rm -f Dockerfile


