# 使用官方 Golang 镜像作为构建环境
FROM golang:1.23-alpine AS builder

FROM golang:1.20

# Enviroment variable
WORKDIR /usr/src/some-api

RUN go install github.com/cosmtrek/air@latest

#Copying files to work directory
COPY go.mod ./
RUN go mod download && go mod verify
COPY . .

# Run and expose the server on port 3000
EXPOSE 3000
# 容器启动时运行 Go 程序
CMD ["./myapp"]
