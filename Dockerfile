FROM golang:1.23

WORKDIR /app

COPY . .

EXPOSE 8080

RUN go build -o main ./cmd/zipvault
CMD ["./main"]
