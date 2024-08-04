FROM golang:1.22-alpine AS build

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]