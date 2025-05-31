FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/zipvault

FROM alpine:3.20.0

RUN apk --no-cache add bash

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/wait-for-it.sh .

RUN chmod +x /app/wait-for-it.sh /app/main

ENTRYPOINT ["/app/wait-for-it.sh", "mysql:3306", "--", "./main"]
