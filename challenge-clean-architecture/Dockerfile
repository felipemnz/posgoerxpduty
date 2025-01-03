FROM golang:1.22.5 as builder

WORKDIR /app
COPY . .
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd/ordersystem/main.go ./cmd/ordersystem/wire_gen.go
FROM alpine:3.18

WORKDIR /app

ARG DB_USER
ARG DB_PASS
ARG DB_HOST
ARG DB_PORT
ARG DB_NAME

RUN ls
RUN apk add --no-cache mysql-client curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-arm64.tar.gz | tar xz -C /usr/local/bin

COPY --from=builder /app/app .
COPY --from=builder /app/internal/infra/database/migrations internal/infra/database/migrations
EXPOSE 8000 50051 8080

CMD ["./app/app"]