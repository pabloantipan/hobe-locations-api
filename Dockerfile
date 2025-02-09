FROM golang:1.23-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN mkdir -p /app/service-accounts
COPY .env /app/.env

COPY datastore_sa.json /app/service-accounts/datastore_sa.json
COPY logging_sa.json /app/service-accounts/logging_sa.json

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go
EXPOSE 8080
CMD ["/app/server"]
