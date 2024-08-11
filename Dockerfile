# Build Stage
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /server


FROM alpine:latest

COPY --from=builder /server /server
COPY --from=builder /app/.env .env
EXPOSE 8080

ENTRYPOINT ["/server"]
