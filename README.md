# date-agent
An agent which was deployed for managing or syncing time between different containers or k8s pods

## Run register server
```sh
go run cmd/register/main.go --grpcservice=0.0.0.0:10000 --httpservice=0.0.0.0:10001 -v=4
```

## Run client
```sh
go run cmd/client/main.go --grpcservice=127.0.0.1:10000 -v=4
```