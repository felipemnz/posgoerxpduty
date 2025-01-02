FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o challenge-weather-by-cep main.go

FROM scratch
WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/challenge-weather-by-cep .
ENTRYPOINT ["./challenge-weather-by-cep"]