FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/zipvault

RUN chmod +x /app/wait-for-it.sh

ENTRYPOINT ["/app/wait-for-it.sh", "mysql:3306", "--", "./main"]
