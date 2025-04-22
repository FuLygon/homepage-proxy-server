FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o homepage-proxy-server ./cmd/

FROM alpine:3.21
RUN apk --no-cache add ca-certificates curl
WORKDIR /app
COPY --from=builder /app/homepage-proxy-server .

HEALTHCHECK --interval=1m --timeout=3s --start-period=3s --retries=3 \
  CMD curl -f http://127.0.0.1:${PORT:-8080}/health || exit 1
EXPOSE ${PORT:-8080}
ENTRYPOINT ["/app/homepage-proxy-server"]