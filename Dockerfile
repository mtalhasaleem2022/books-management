# Build stage
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/books-api ./cmd/main.go

# Runtime stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/bin/books-api .
EXPOSE 8080
CMD ["/app/books-api"]