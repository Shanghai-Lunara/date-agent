FROM golang:1.17-alpine

ENV GO111MODULE=on
ENV GOPROXY=https://mirrors.aliyun.com/goproxy/,https://goproxy.cn,direct

WORKDIR /server

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o date-agent-client cmd/client-env-retry/main.go

CMD ["go", "run", "/server/cmd/hang/main.go"]