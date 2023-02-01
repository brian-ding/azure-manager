# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

ENV GIN_MODE=release
ENV GOPROXY=https://proxy.golang.com.cn,direct

WORKDIR /app

RUN pwd

COPY . .

RUN ls -l .

RUN go build -o /azure-manager cmd/main.go

EXPOSE 8080

CMD [ "/azure-manager" ]