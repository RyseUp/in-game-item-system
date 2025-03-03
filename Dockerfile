FROM golang:1.23.4 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY .env .env

COPY --from=builder /app/main .
EXPOSE 50051
CMD ["./main"]
