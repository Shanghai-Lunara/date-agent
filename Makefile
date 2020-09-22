.PHONY: proto register client

proto:
	cd proto && protoc --go_out=plugins=grpc:. *.proto

mod:
	go mod download